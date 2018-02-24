package main

import (
	"net/http"

	"github.com/ice-judge/ICE-api/api"
)

func main() {
	API := api.NewAPI()
	http.ListenAndServe("127.0.0.1:8080", API.GetRouter())
}
