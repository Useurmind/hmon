package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type LogTypeService interface {
	StoreLog(logSource *LogSource, path string)
}

type LogChecker struct {
	stop                    chan bool
	checkInterval           time.Duration
	logConfigurationService *LogConfigurationService
	fileWatcher             *FileWatcher
	logServices             map[string]LogTypeService
}

func NewLogChecker(logConfigurationService *LogConfigurationService, checkInterval time.Duration, logServices map[string]LogTypeService) *LogChecker {
	return &LogChecker{
		logConfigurationService: logConfigurationService,
		stop:                    make(chan bool),
		fileWatcher:             NewFileWatcher(),
		checkInterval:           checkInterval,
		logServices:             logServices,
	}
}

func (lc *LogChecker) GoStart() {
	go lc.Start()
}

func (lc *LogChecker) Start() {
	log.Println("Starting log checker")
	for {
		select {
		case _ = <-lc.stop:
			log.Println("Stopping log checker...")
		case <-time.After(lc.checkInterval):
			err := lc.checkChangedLogFiles()
			if err != nil {
				log.Printf("ERROR: While checking changed log files: %s", err)
			}
		}
	}
}

func (lc *LogChecker) Stop() {
	lc.stop <- true
}

func (lc *LogChecker) checkChangedLogFiles() error {
	logSources, err := lc.logConfigurationService.GetLogSources()
	if err != nil {
		return err
	}

	for _, logSource := range logSources {
		lc.fileWatcher.AddWatchedFolder(logSource.SourceFolder)
	}

	err = lc.fileWatcher.ForEachChangedFile(func(path string, isNew bool) error {

		// check for matching log sources
		for _, logSource := range logSources {
			sourceFolder := filepath.FromSlash(logSource.SourceFolder)
			absolutePath, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			matchesSourceRegex := false
			isInSourceFolder := strings.HasPrefix(absolutePath, sourceFolder)
			if isInSourceFolder {
				_, fileName := filepath.Split(path)
				matchesSourceRegex, err = regexp.Match(logSource.FileRegex, []byte(fileName))
				if err != nil {
					return err
				}
			}

			belongsToSource := isInSourceFolder && matchesSourceRegex
			if belongsToSource {
				err := lc.storeLogfile(logSource, path)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (lc *LogChecker) storeLogfile(logSource *LogSource, path string) error {

	logService, ok := lc.logServices[logSource.Type]
	if !ok {
		return fmt.Errorf("Could not find log service for log type %s in source %s, log file in question is %s", logSource.Type, logSource.Name, path)
	}

	logService.StoreLog(logSource, path)

	return nil
}
