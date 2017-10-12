package riot_api

import (
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"fmt"
)

type Summoner struct {
	Id int				`json: "id"`
	AccountId int		`json: "accountId"`
	Name string 		`json: "name"`
	ProfileIconId int	`json: "profileIconId"`
	RevisionDate int	`json: "revisionDate"`
	SummonerLevel int	`json: "summonerLevel"`
}

var key = "" // Riot Games API Key

func GetSummoner(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/summoner/")
	apiRequest(w, "euw1", "summoner/v3/summoners/by-name/" + id)
}

func apiRequest(w http.ResponseWriter, region string, endpoint string) {
	url := "https://" + region + ".api.riotgames.com/lol/" + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	req.Header.Set("X-Riot-Token", key)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	result, _ := ioutil.ReadAll(res.Body)
	summoner := Summoner{}
	json.Unmarshal([]byte(result), &summoner)
	fmt.Println("Summoner ID:", summoner.Name)

	w.Write(result)
}