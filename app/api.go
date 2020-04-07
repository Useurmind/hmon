package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type APIOptions struct {
	LogConfigurationService *LogConfigurationService
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
	log.Printf("API Request %s: %s", r.Method, r.URL.String())

	mux := NewHTTPMux()
	mux.PathHandlers["/jobsource"] = jobSourceAPI{
		LogConfigurationService: a.options.LogConfigurationService,
	}
	mux.PathHandlers["/sources"] = sourcesAPI{
		LogConfigurationService: a.options.LogConfigurationService,
	}

	mux.ServeHTTP(w, r)
}

type sourcesAPI struct {
	LogConfigurationService *LogConfigurationService
}

func (a sourcesAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	sources, err := a.LogConfigurationService.GetLogSources()
	if err != nil {
		handleError(w, "Unkown error", http.StatusInternalServerError, err)
		return
	}

	json, err := json.Marshal(sources)
	if err != nil {
		handleError(w, "Unkown error", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))
}

type jobSourceAPI struct {
	LogConfigurationService *LogConfigurationService
}

func (a jobSourceAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		jlsId, err := getNextPathPartAsId(r)
		if err != nil {
			handleError(w, "Could not find id of JobLogSource in path", http.StatusBadRequest, err)
			return
		}

		jls, err := a.LogConfigurationService.GetJobLogSource(jlsId)
		if err != nil {
			handleError(w, "Could not get JobLogSource", http.StatusInternalServerError, err)
			return
		}

		if jls == nil {
			handleError(w, "Could not find JobLogSource", http.StatusNotFound, fmt.Errorf("JobLogSource returned by service was nil"))
			return
		}

		json, err := json.Marshal(jls)
		if err != nil {
			handleError(w, "Unkown error", http.StatusInternalServerError, err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(json))

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handleError(w, "Request body could not be read", http.StatusBadRequest, err)
			return
		}

		var jls = JobLogSource{}
		err = json.Unmarshal(body, &jls)
		if err != nil {
			handleError(w, "Body could not be parsed as JobLogSource", http.StatusBadRequest, err)
			return
		}

		err = a.LogConfigurationService.CreateOrUpdateJobLogSource(&jls)
		if err != nil {
			handleError(w, "Unkown error", http.StatusInternalServerError, err)
			return
		}

		result, err := json.Marshal(jls)
		if err != nil {
			handleError(w, "Unkown error", http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(result))

	case "DELETE":
		jlsId, err := getNextPathPartAsId(r)
		if err != nil {
			handleError(w, "Could not find id of JobLogSource in path", http.StatusBadRequest, err)
			return
		}

		err = a.LogConfigurationService.DeleteJobLogSource(jlsId)
		if err != nil {
			handleError(w, "Could not delete JobLogSource", http.StatusInternalServerError, err)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleError(w http.ResponseWriter, errorMsg string, statusCode int, err error) {
	http.Error(w, errorMsg, statusCode)
	log.Printf("ERROR: %s (%s, %d)\r\n", err, errorMsg, statusCode)
}

func getNextPathPartAsId(r *http.Request) (int, error) {
	pathParts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 1)
	if len(pathParts) < 1 {
		return -1, fmt.Errorf("Path parts slice length only %d", len(pathParts))
	}

	id, err := strconv.Atoi(pathParts[0])
	if err != nil {
		return -1, err
	}

	return id, nil
}
