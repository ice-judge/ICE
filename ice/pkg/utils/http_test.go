package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ice-judge/ICE/ice/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPClient(t *testing.T) {
	a := assert.New(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	}))
	defer ts.Close()

	cli := utils.NewHTTPClient()
	a.Equal(10*time.Second, cli.Timeout)

	// to prevent test timeout
	cli.Timeout = time.Second

	c := make(chan bool)
	go func() {
		_, err := cli.Get(ts.URL)
		if err == nil {
			a.Fail("the request should be failed")
		}
		close(c)
	}()

	select {
	case <-time.After(2 * time.Second):
		a.Fail("timeout not applied")
	case <-c:
	}
}
