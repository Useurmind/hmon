package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	// The IP address on which to listen, empty for all.
	Address string
	// The port on which to listen, default 8080.
	Port string
}

func (s *Server) Run() error {
	apiOptions := APIOptions{}
	apiHandler, err := NewAPIHandler(&apiOptions)
	if err != nil {
		return err
	}

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", s.Address, s.Port), apiHandler)
	if err != nil {
		return err
	}

	return nil
}