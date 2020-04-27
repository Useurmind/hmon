package main


import (
	"net/http"
	"strings"
)

type HTTPMuxHandler struct {
	Path string
	HTTPHandler http.Handler
}

type HTTPMux struct {
	PathHandlers []HTTPMuxHandler
}

func NewHTTPMux(handlers []HTTPMuxHandler) HTTPMux {
	return HTTPMux{
		PathHandlers: handlers,
	}
}

func (m HTTPMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, muxHandler := range m.PathHandlers {
		path := muxHandler.Path
		handler := muxHandler.HTTPHandler
		isMatchPath := len(strings.TrimPrefix(req.URL.Path, path)) != len(req.URL.Path)
		if isMatchPath {
			jobHandler := http.StripPrefix(path, handler)

			jobHandler.ServeHTTP(w, req)
			return
		}
	}

	http.NotFound(w, req)
}