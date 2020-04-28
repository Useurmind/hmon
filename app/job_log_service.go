package main

import (
	"log"
	"github.com/boltdb/bolt"
)

type JobLog struct {
	ID int
}

func (jl *JobLog) GetId() int {
	return jl.ID
}

type JobLogService struct {
	*DBService
}

func (s *JobLogService) StoreLog(logSource *LogSource, path string) error {
	return nil
}

func (s *JobLogService) StoreJobLog(jl *JobLog) (*JobLog, error) {
	log.Printf("Call to StoreJobLog(%+v)", jl)

	db, err := s.getDB()
	if err != nil {
		return nil, err
	}

	return jl, db.Update(func (tx *bolt.Tx) error {
		b, err := ensureJobLogBucket(tx)
		if err != nil {
			return err
		}

		isNew := jl.ID <= 0
		if isNew {
			nextId, _ := b.NextSequence()

			jl.ID = int(nextId)
		}

		err = saveObject(b, jl)
		if err != nil {
			return err
		}

		return nil
	})
}


const jobLogBucket = "log_jobs"

func ensureJobLogBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return tx.CreateBucketIfNotExists([]byte(jobLogBucket))
}


