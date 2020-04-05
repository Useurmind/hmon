package main

import (
	"log"
	"fmt"
	"net/http"
)

type Server struct {
	// The IP address on which to listen, empty for all.
	Address string

	// The port on which to listen, default 8080.
	Port string

	// The path to the file where the db is stored
	DBFilePath string
}

func (s *Server) Run() error {
	logConfigService := LogConfigurationService{
		DBFilePath: s.DBFilePath,
	}
	defer logConfigService.Close()

	apiOptions := APIOptions{
		LogConfigurationService: &logConfigService,
	}
	apiHandler, err := NewAPIHandler(&apiOptions)
	if err != nil {
		return err
	}

	err = logConfigService.PingDB()
	if err != nil {
		return err
	}

	listenOn := fmt.Sprintf("%s:%s", s.Address, s.Port)
	log.Printf("Listening on: %s\r\n", listenOn)
	err = http.ListenAndServe(listenOn, apiHandler)
	if err != nil {
		return err
	}

	return nil
}