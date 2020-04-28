package main

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"github.com/boltdb/bolt"
)

func TestDBServiceEnsureAndDeleteWorks(t *testing.T) {
	dbFilePath := filepath.FromSlash("testing/dbservice.db")
	dbService := NewDBService(dbFilePath)

	err := os.Mkdir("testing", 0644)
	assert.Nil(t, err)

	defer func() {
		os.Remove(dbFilePath)	
		os.Remove("testing")
	}()

	err = dbService.ensureDB()
	assert.Nil(t, err)

	_, err = os.Stat(dbFilePath)
	assert.Nil(t, err)

	err = dbService.DeleteDB()
	assert.Nil(t, err)

	_, err = os.Stat(dbFilePath)
	assert.True(t, os.IsNotExist(err))
}

type TestObject struct {
	Id int
	Name string
}

func (t *TestObject) GetId() int {
	return t.Id
}

func TestDBServiceWriteReadObjectWorks(t *testing.T) {
	dbFilePath := filepath.FromSlash("testing/dbservice.db")
	dbService := NewDBService(dbFilePath)

	err := os.Mkdir("testing", 0644)
	assert.Nil(t, err)

	defer func() {
		dbService.Close()
		os.Remove(dbFilePath)	
		os.Remove("testing")
	}()

	err = dbService.ensureDB()
	assert.Nil(t, err)

	db, err := dbService.getDB()
	assert.Nil(t, err)

	toSave := TestObject{
		Id: 1,
		Name: "test",
	}

	loaded := TestObject{
		Id: 1,
	}

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("bucket"))
		assert.Nil(t, err)

		err = saveObject(b, &toSave)
		assert.Nil(t, err)

		err = readObject(b, &loaded)
		assert.Nil(t, err)

		return nil
	})

	assert.Equal(t, toSave.Id, loaded.Id)
	assert.Equal(t, toSave.Name, loaded.Name)
}