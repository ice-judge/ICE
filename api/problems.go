package api

import (
	"encoding/json"
	"net/http"

	"github.com/ice-judge/ICE-api/ice"
)

type problemAllRespose struct {
	Problems []ice.ProblemHeader `json:"problems"`
}

// ProblemsAllHandler is for /problems
func (api *API) problemsAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, header, _ := api.ice.GetProblemHeaders(ice.Pagination{})
	response := problemAllRespose{
		Problems: header,
	}
	json.NewEncoder(w).Encode(response)
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
