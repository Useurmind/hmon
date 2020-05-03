package main 

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestFolder struct {
	folder string 
}

func (tf *TestFolder) EnsureDeleted(t *testing.T) {
	err := os.RemoveAll(tf.folder)
	if err != nil && !os.IsNotExist(err) {
		t.Logf("Error while deleting test folder: %s", err)
		assert.Nil(t, err)
	}	
}

func ensureTestingFolder(t *testing.T) *TestFolder {
	testFolder := "testing"
	err := os.Mkdir(testFolder, 0644)
	assert.Nil(t, err)

	return &TestFolder{
		folder: testFolder,
	}
}