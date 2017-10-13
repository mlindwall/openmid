package main

import (
	"net/http"
	"github.com/mlindwall/openmid-api/openmid-api"
)

func main() {
	http.HandleFunc("/summoner/", openmid_api.GetProfile)
	http.HandleFunc("/favicon.ico", disregard)
	http.ListenAndServe(":8080", nil)
}

func disregard(w http.ResponseWriter, r *http.Request) {}