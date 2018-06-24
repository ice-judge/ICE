package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) NewHandler() http.Handler {
	r := chi.NewRouter()
	r.Route("/assets", s.registerAssets)

	return r
}
