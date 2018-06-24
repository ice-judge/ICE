package utils

import (
	"io"
	"net/http"
	"time"
)

// NewHTTPClient creates new a http.client with 10 seconds timeout
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

// WriteHTTPResponse writes http.Reponse to http.ResponseWriter
func WriteHTTPResponse(w http.ResponseWriter, resp *http.Response) {
	for name, values := range resp.Header {
		w.Header()[name] = values
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
