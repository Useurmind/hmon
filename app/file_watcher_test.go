package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
)

func TestNewFileIsRegistered(t *testing.T) {
	filePath := filepath.FromSlash("testing/file1")

	watcher := NewFileWatcher()
	watcher.AddWatchedFolder("testing")	

	testFolder := ensureTestingFolder(t)

	err := ioutil.WriteFile(filePath, []byte("content 1"), 0644)
	assert.Nil(t, err)

	defer func() {
		testFolder.EnsureDeleted(t)
	}()

	done := false
	watcher.ForEachChangedFile(func(path string, isNew bool) error{
		assert.Equal(t, false, done)
		assert.Equal(t, filePath, path)
		done = true

		return nil
	})
}