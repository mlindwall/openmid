package main

import (
	"net/http"
	"github.com/mlindwall/openmid-api/riot-api"
)

func main() {
	http.HandleFunc("/summoner/", riot_api.GetSummoner)
	http.HandleFunc("/favicon.ico", disregard)
	http.ListenAndServe(":8080", nil)
}

func disregard(w http.ResponseWriter, r *http.Request) {}