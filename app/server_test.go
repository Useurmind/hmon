package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"path/filepath"
	"testing"
	"net"
	"time"

	"github.com/stretchr/testify/assert"
)

func createTestServer(t *testing.T) *Server {
	address := "127.0.0.1"
	port := "59856"
	dbFilePath := filepath.FromSlash("testing/db")

	server := Server{
		Address:    address,
		Port:       port,
		DBFilePath: dbFilePath,
	}

	return &server
}

func TestServerListenOnCorrectAddress(t *testing.T) {
	server := createTestServer(t)

	testFolder := ensureTestingFolder(t)

	defer func() {
		server.Shutdown()
		testFolder.EnsureDeleted(t)
	}()

	err := server.RunAsync()
	assert.Nil(t, err)

	timeout, err := time.ParseDuration("3s")
	assert.Nil(t, err)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server.Address, server.Port), timeout)
	assert.Nil(t, err)
	defer conn.Close()
}

func TestServerReturnsEmptyListOfLogSources(t *testing.T) {
	server := createTestServer(t)

	testFolder := ensureTestingFolder(t)

	defer func() {
		server.Shutdown()
		testFolder.EnsureDeleted(t)
	}()

	err := server.RunAsync()
	assert.Nil(t, err)

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/api/sources", server.Address, server.Port))
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "[]", string(bodyBytes))
}
