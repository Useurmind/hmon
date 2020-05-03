package main

import (
	"context"
	"log"
	"fmt"
	"net/http"
	"sync"

	"github.com/Useurmind/spas/handler"
)

type Server struct {
	// The IP address on which to listen, empty for all.
	Address string

	// The port on which to listen, default 8080.
	Port string

	// The path to the file where the db is stored
	DBFilePath string

	// The underlying http server that listens for requests.
	server *http.Server

	// Wait group used when running the server async.
	wg *sync.WaitGroup

	// Responsible for closing all dependencies
	closer *Closer
}

func (s *Server) Shutdown() error {
	if s.wg == nil {
		return fmt.Errorf("Cannot use shutdown when RunAsync was not called before")
	}

	log.Printf("Shutting down server\r\n")
	err := s.server.Shutdown(context.Background())
	if err != nil {
		return err
	}

	s.wg.Wait()

	s.closer.Close()

	return nil
}

func (s *Server) Run() error {
	server, err := s.createServer()
	if err != nil {
		return err
	}

	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) RunAsync() error {
	s.wg = &sync.WaitGroup{}

	server, err := s.createServer()
	if err != nil {
		return err
	}
	s.server = server

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe crashed: %v", err)
		}

		log.Printf("Shutdown of server complete\r\n")
	}()

	return nil
}

func (s *Server) createServer() (*http.Server, error) {
	mux, err := s.createHandler()
	if err != nil {
		return nil, err
	}

	listenOn := fmt.Sprintf("%s:%s", s.Address, s.Port)
	log.Printf("Server listening on: %s\r\n", listenOn)

	server := &http.Server{
		Addr: listenOn,
		Handler: mux,
	}

	return server, nil
}

func (s *Server) createHandler() (http.Handler, error) {
	closer := NewCloser()
	s.closer = closer

	logConfigService, err := NewLogConfigurationService(NewDBService(s.DBFilePath))
	if err != nil {
		return nil, err
	}

	s.closer.AddClosable(logConfigService)

	apiOptions := APIOptions{
		LogConfigurationService: logConfigService,
	}
	apiHandler, err := NewAPIHandler(&apiOptions)
	if err != nil {
		return nil, err
	}

	spasOptions := handler.Options{
		ServeFolder: "www",
		HTMLIndexFile: "index.html",
	}
	spasHandler := handler.NewSPASHandler(&spasOptions)

	mux := NewHTTPMux([]HTTPMuxHandler{
		HTTPMuxHandler{ 
			Path: "/api",
			HTTPHandler: apiHandler,
		},
		HTTPMuxHandler{ 
			Path: "/ui",
			HTTPHandler: spasHandler,
		},
		HTTPMuxHandler{ 
			Path: "/",
			HTTPHandler: http.RedirectHandler("/ui", 301),
		},
	})

	err = logConfigService.PingDB()
	if err != nil {
		return nil, err
	}

	return mux, nil
}