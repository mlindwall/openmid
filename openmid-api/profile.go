package openmid_api

import (
	"net/http"
	"strings"
	"github.com/mlindwall/openmid-api/riot-api"
	"encoding/json"
)

type Profile struct {
	Summoner riot_api.Summoner
	Matches riot_api.MatchList
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/summoner/")
	summoner := riot_api.GetSummoner("euw1", id)
	if summoner.AccountId == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	matches := riot_api.GetRecentMatches("euw1", summoner.AccountId)
	profile := Profile{summoner, matches}
	summonerJson, _ := json.Marshal(profile)
	w.Write([]byte(summonerJson))
}