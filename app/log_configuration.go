package main

import (
	"log"
	"github.com/boltdb/bolt"
	"encoding/json"
)

// LogSource describes the source of log files.
type LogSource struct {
	Id int
	Name string
	Type string
	SourceFolder string
	FileRegex string
}

func (ls *LogSource) GetId() int {
	return ls.Id
}

const JobLogSourceType string = "job"

// JobLogSource is a source of log files that contain the results of a single job.
type JobLogSource struct {
	Id int
	Name string
	Type string

	// The source folder is the absolute path to scan for matching log files.
	SourceFolder string

	// A regex to match all files in the folder that should be scanned
	// To retrieve a time from the file name integrate a group named 'time' into the regex.
	// If no time group is found or matched the file time will be used.
	FileRegex string

	// A regex that applied to a line in the file states that the job was successful.
	// Parsing will start from the end.
	SuccessRegex string

	// A regex that applied to a line in the file indicates an error.
	// Parsing will start from the end.
	ErrorRegex string
}

func (ls *JobLogSource) GetId() int {
	return ls.Id
}

func (ls *JobLogSource) ToLogSource() *LogSource {
	return &LogSource{
		Id: ls.Id,
		Name: ls.Name,
		Type: ls.Type,
		SourceFolder: ls.SourceFolder,
		FileRegex: ls.FileRegex,
	}
}

func (ls *JobLogSource) ApplyType() {
	ls.Type = JobLogSourceType
}

type LogConfigurationService struct {
	DBService
}

func (s *LogConfigurationService) GetLogSources() ([]*LogSource, error) {
	log.Println("Call to GetLogSources()")
	db, err := s.getDB()
	if err != nil {
		return nil, err
	}

	var lss []*LogSource = make([]*LogSource, 0)

	err = db.View(func(tx *bolt.Tx) error {
		bLS := tx.Bucket([]byte(logSourcesBucket))
		if bLS == nil {
			return nil
		}

		err = bLS.ForEach(func(k, v []byte) error {			
			var ls *LogSource = &LogSource{}
			err = json.Unmarshal(v, ls)
			if err != nil {
				return err
			}
			lss = append(lss, ls)

			return nil
		})

		return err
	})

	return lss, err
}

func (s *LogConfigurationService) GetJobLogSource(id int) (*JobLogSource, error) {
	log.Printf("Call to GetJobLogSource(%d)", id)
	db, err := s.getDB()
	if err != nil {
		return nil, err
	}

	var jls *JobLogSource = nil

	err = db.View(func(tx *bolt.Tx) error {
		bJLS := tx.Bucket([]byte(jobLogSourcesBucket))
		if bJLS == nil {
			return nil
		}

		val := bJLS.Get(itob(id))
		if val != nil {
			jls = &JobLogSource{}
			err = json.Unmarshal(val, jls)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return jls, err
}

func (s *LogConfigurationService) CreateOrUpdateJobLogSource(jls *JobLogSource) (*JobLogSource, error) {
	log.Printf("Call to CreateOrUpdateJobLogSource(%+v)", jls)
	jls.ApplyType()

	db, err := s.getDB()
	if err != nil {
		return nil, err
	}

	return jls, db.Update(func (tx *bolt.Tx) error {
		bLS, err := ensureLogSourceBucket(tx)
		if err != nil {
			return err
		}
		bJLS, err := ensureJobLogSourceBucket(tx)
		if err != nil {
			return err
		}

		lsIsNew := jls.Id <= 0
		if lsIsNew {
			// this is a new one
			id, _ := bLS.NextSequence()

			jls.Id = int(id)
		}

		err = saveObject(bJLS, jls)
		if err != nil {
			return err
		}

		err = saveObject(bLS, jls.ToLogSource())
		if err != nil {
			return err
		}

		return nil
	}) 
}

func (s *LogConfigurationService) DeleteJobLogSource(id int) error {
	log.Printf("Call to DeleteJobLogSource(%d)", id)
	db, err := s.getDB()
	if err != nil {
		return err
	}

	return db.Update(func (tx *bolt.Tx) error {
		bLS, err := ensureLogSourceBucket(tx)
		if err != nil {
			return err
		}
		bJLS, err := ensureJobLogSourceBucket(tx)
		if err != nil {
			return err
		}

		err = bJLS.Delete(itob(id))
		if err != nil {
			return err
		}

		err = bLS.Delete(itob(id))
		if err != nil {
			return err
		}

		return nil
	})
}

const logSourcesBucket = "log_sources"
const jobLogSourcesBucket = "log_sources_jobs"

func ensureLogSourceBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return tx.CreateBucketIfNotExists([]byte(logSourcesBucket))
}

func ensureJobLogSourceBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return tx.CreateBucketIfNotExists([]byte(jobLogSourcesBucket))
}