package api

import (
	"net/http"
)

// ProblemsAllHandler is for /problems
func (api *API) problemsAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("Content-Type", "application/json")
}

// ProblemsOfDifficultyHandler is for /problems/difficulty/{difficulty}
func (api *API) problemsOfDifficultyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("Content-Type", "application/json")
}

// ProblemsSearchHandler is for /problems/search
func (api *API) problemsSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("Content-Type", "application/json")
}
