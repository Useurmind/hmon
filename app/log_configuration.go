package main

import (
	"log"
	"github.com/boltdb/bolt"
	"encoding/json"
	"encoding/binary"
)

type WithId interface {
	GetId() int
}

// LogSource describes the source of log files.
type LogSource struct {
	Id int
	Name string
	Type string
}

func (ls *LogSource) GetId() int {
	return ls.Id
}

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
	}
}

func (ls *JobLogSource) ApplyType() {
	ls.Type = "job"
}

type LogConfigurationService struct {
	DBFilePath string
	DB *bolt.DB
}

func (s *LogConfigurationService) GetLogSources() ([]*LogSource, error) {
	err := s.ensureDB()
	if err != nil {
		return nil, err
	}

	var lss []*LogSource = make([]*LogSource, 0)

	err = s.DB.View(func(tx *bolt.Tx) error {
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
	err := s.ensureDB()
	if err != nil {
		return nil, err
	}

	var jls *JobLogSource = nil

	err = s.DB.View(func(tx *bolt.Tx) error {
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

func (s *LogConfigurationService) CreateOrUpdateJobLogSource(jls *JobLogSource) error {
	jls.ApplyType()

	err := s.ensureDB()
	if err != nil {
		return err
	}

	return s.DB.Update(func (tx *bolt.Tx) error {
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
	err := s.ensureDB()
	if err != nil {
		return err
	}

	return s.DB.Update(func (tx *bolt.Tx) error {
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

func (s *LogConfigurationService) PingDB() error {
	err := s.ensureDB()

	return err
}

func (s *LogConfigurationService) ensureDB() error {
	if s.DB != nil {
		return nil
	}

	log.Printf("Using database in %s", s.DBFilePath)

	db, err := bolt.Open(s.DBFilePath, 0600, nil)
	if err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s *LogConfigurationService) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}

func saveObject(b *bolt.Bucket, obj WithId) error {
	json, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	err = b.Put(itob(obj.GetId()), json)
	if err != nil {
		return err
	}

	return nil
}

const logSourcesBucket = "log_sources"
const jobLogSourcesBucket = "log_sources_jobs"

func ensureLogSourceBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return tx.CreateBucketIfNotExists([]byte(logSourcesBucket))
}

func ensureJobLogSourceBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return tx.CreateBucketIfNotExists([]byte(jobLogSourcesBucket))
}

func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}