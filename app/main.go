package main

import (
	"os"
)

func main() {
	dbFilePath := os.Getenv("HMON_DBFILEPATH")
	if dbFilePath == "" {
		panic("Set HMON_DBFILEPATH to configure db location")
	}
	address := os.Getenv("HMON_ADDRESS")
	port := os.Getenv("HMON_PORT")
	if port == "" {
		port = "8080"
	}

	server := Server{
		Address:    address,
		Port:       port,
		DBFilePath: dbFilePath,
	}

	err := server.Run()
	if err != nil {
		panic(err)
	}
}
