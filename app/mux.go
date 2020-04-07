package main


import (
	"net/http"
	"strings"
)

type HTTPMux struct {
	PathHandlers map[string]http.Handler
}

func NewHTTPMux() HTTPMux {
	return HTTPMux{
		PathHandlers: make(map[string]http.Handler),
	}
}

func (m HTTPMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for path, handler := range m.PathHandlers {
		isMatchPath := len(strings.TrimPrefix(req.URL.Path, path)) != len(req.URL.Path)
		if isMatchPath {
			jobHandler := http.StripPrefix(path, handler)

			jobHandler.ServeHTTP(w, req)
			return
		}
	}

	http.NotFound(w, req)
}