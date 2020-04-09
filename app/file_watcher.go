package main

import (
	"log"
	"time"
	"path/filepath"
	"os"
)

type FileWatcher struct {
	watchedFolders map[string]bool
	fileTimes map[string]time.Time
}

func NewFileWatcher() *FileWatcher {
	return &FileWatcher{
		watchedFolders: make(map[string]bool),
		fileTimes: make(map[string]time.Time),
	}
}

func (w *FileWatcher) AddWatchedFolder(folder string) {
	w.watchedFolders[folder] = true
}

func (w *FileWatcher) RemoveWatchedFolder(folder string) {
	delete(w.watchedFolders, folder)
}

func (w *FileWatcher) ForEachChangedFile(action func(filePath string, isNew bool) error) error {
	log.Println("Check for changed files...")
	for folder, _ := range w.watchedFolders {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				// we are only interested in files
				return nil
			}
			
			log.Printf("Checking if file %s changed...\r\n", path)
			lastFileTime, ok := w.fileTimes[path]
			if !ok {
				log.Printf("File is new...\r\n")
				w.fileTimes[path] = info.ModTime()

				err := action(path, true)
				if err != nil {
					return err
				}
			} else if lastFileTime.Before(info.ModTime()) {
				log.Printf("File changed...\r\n")
				w.fileTimes[path] = info.ModTime()

				err := action(path, false)
				if err != nil {
					return err
				}
			} else {
				log.Printf("File did not change...\r\n")
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}