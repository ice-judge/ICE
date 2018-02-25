package main

import (
	"net/http"

	"github.com/ice-judge/ICE-api/api"
	"github.com/ice-judge/ICE-api/ice"
)

func main() {
	ICE := &ice.ICE{}
	API := api.NewAPI(ICE)
	http.ListenAndServe("127.0.0.1:8080", API.GetRouter())
}
