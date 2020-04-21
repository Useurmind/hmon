package main

import (
	"log"
	"github.com/boltdb/bolt"
	"encoding/json"
)

type DBService struct {
	dbFilePath string
	db *bolt.DB
}

func NewDBService(dbFilePath string) *DBService {
	return &DBService{
		dbFilePath: dbFilePath,
		db: nil,
	}
}

func (s *DBService) getDB() (*bolt.DB, error) {
	err := s.ensureDB()
	if err != nil {
		return nil, err
	}

	return s.db, err
}

func (s *DBService) PingDB() error {
	err := s.ensureDB()

	return err
}

func (s *DBService) ensureDB() error {
	if s.db != nil {
		return nil
	}

	log.Printf("Using database in %s", s.dbFilePath)

	db, err := bolt.Open(s.dbFilePath, 0600, nil)
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *DBService) Close() {
	if s.db != nil {
		s.db.Close()
		s.db = nil
	}
}


type WithId interface {
	GetId() int
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
