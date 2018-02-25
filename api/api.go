package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ice-judge/ICE-api/ice"
)

// API provides a RESTful API
type API struct {
	ice *ice.ICE
}

// NewAPI creates a new API instance
func NewAPI(ice *ice.ICE) *API {
	return &API{
		ice: ice,
	}
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, ice.Version)
}

// GetRouter will return api router
func (api *API) GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.Path("/problems").
		HandlerFunc(api.problemsAllHandler).
		Methods("GET")

	r.Path("/problems/difficulty/{difficulty}").
		HandlerFunc(api.problemsOfDifficultyHandler).
		Methods("GET")

	r.Path("/problems/search").
		HandlerFunc(api.problemsSearchHandler).
		Methods("GET")

	r.Path("/version").
		HandlerFunc(versionHandler).
		Methods("GET")
	return r
}
