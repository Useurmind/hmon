package main

import (
	"log"
)

type FileWatcher struct {
	FoundFiles chan string
	watchedFolders []string
	stop chan bool
}

func NewFileWatcher(watchedFolders []string) *FileWatcher {
	return &FileWatcher{
		FoundFiles: make(chan string),
		watchedFolders: watchedFolders,
		stop: make(chan bool),
	}
}

func (w *FileWatcher) GoStart() {
	go w.Start()
}

func (w *FileWatcher) Start() {
	select {
    case _ = <-w.stop:
        log.Println("Stopping file watcher...")
    }
}

func (w *FileWatcher) Stop() {
	w.stop <- true
}