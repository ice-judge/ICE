package server_test

import (
	"net/http/httptest"

	"github.com/ice-judge/ICE/ice/api/server"
)

func NewTestServer() *httptest.Server {
	conf := server.Config{}
	s, err := server.New(conf)
	if err != nil {
		panic(err)
	}
	return httptest.NewServer(s.NewHandler())
}
