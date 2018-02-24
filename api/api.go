package api

import (
	"github.com/gorilla/mux"
)

// API provides a RESTful API
type API struct {
}

// NewAPI creates a new API instance
func NewAPI() *API {
	return &API{}
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
