package server

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/ice-judge/ICE/ice/pkg/utils"
)

// registerAssets will register asset handlers to r
func (s *Server) registerAssets(r chi.Router) {
	r.Get("/*", s.newAssetsHandler())
}

// newAssetsHandler creates a GET /* handler that
// will serve static asset files based on the config
func (s *Server) newAssetsHandler() func(http.ResponseWriter, *http.Request) {
	if utils.IsHTTPURL(s.config.AssetsPath) {
		return s.newGetProxyHandler()
	}
	return s.newLocalFileHandler()
}

// newFileProxyHandler creates a GET /* handler that
// will proxy the GET requests to url
func (s *Server) newGetProxyHandler() func(http.ResponseWriter, *http.Request) {
	url_ := s.config.AssetsPath
	if !strings.HasSuffix(url_, "/") {
		url_ += "/"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// wildcard parameter
		file := chi.URLParam(r, "*")

		cli := utils.NewHTTPClient()
		resp, err := cli.Get(url_ + file)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		defer resp.Body.Close()

		utils.WriteHTTPResponse(w, resp)
	}
}

// newLocalFileHandler creates a GET /* handler that
// will serve local files from path
func (s *Server) newLocalFileHandler() func(http.ResponseWriter, *http.Request) {
	fs := http.FileServer(http.Dir(s.config.AssetsPath))
	return func(w http.ResponseWriter, r *http.Request) {
		// wildcard parameter
		file := chi.URLParam(r, "*")

		// to disable default root page of http.FileServer
		if file == "" {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		url_, err := url.Parse(file)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		r.URL = url_

		fs.ServeHTTP(w, r)
	}
}
