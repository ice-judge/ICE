package api

import (
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

	return r
}
