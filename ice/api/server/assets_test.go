package server_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/ice-judge/ICE/ice/api/server"
	"github.com/ice-judge/ICE/ice/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestLocalAssets(t *testing.T) {
	a := assert.New(t)

	// prepare test local data
	dir := utils.RandString(20)
	err := os.Mkdir(dir, os.ModePerm)
	a.Nil(err)
	err = ioutil.WriteFile(dir+"/test", []byte("asdf"), os.ModePerm)
	a.Nil(err)
	defer os.RemoveAll(dir)

	// prepare server.Server
	conf := server.Config{
		AssetsPath: dir,
	}
	s, err := server.New(conf)
	a.Nil(err)
	handler := s.NewHandler()

	// tests
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/assets", nil)
	handler.ServeHTTP(w, req)
	a.Equal(404, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/assets/", nil)
	handler.ServeHTTP(w, req)
	a.Equal(404, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/assets/asdf", nil)
	handler.ServeHTTP(w, req)
	a.Equal(404, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/assets/test", nil)
	handler.ServeHTTP(w, req)
	a.Equal(200, w.Code)
	a.Equal("asdf", w.Body.String())
}

func TestProxyAssets(t *testing.T) {
	a := assert.New(t)

	// prepare test target server
	r := chi.NewRouter()
	r.Get("/asdf.js", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "asdf")
	})
	ts := httptest.NewServer(r)
	defer ts.Close()

	// prepare server.Server
	conf := server.Config{
		AssetsPath: ts.URL,
	}
	s, err := server.New(conf)
	a.Nil(err)
	handler := s.NewHandler()

	// tests
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/assets/", nil)
	handler.ServeHTTP(w, req)
	a.Equal(404, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/assets/asdf.js", nil)
	handler.ServeHTTP(w, req)
	a.Equal(200, w.Code)
	a.Equal("asdf", w.Body.String())
}
