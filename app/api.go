package main

import (
	"net/http"
	"fmt"
)

type APIOptions struct {
}

type API struct {
	options *APIOptions
}

func NewAPIHandler(options *APIOptions) (http.Handler, error) {
	api := API{
		options: options,
	}

	return api, nil
}

func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Heres the api")
}