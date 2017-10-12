package riot_api

import (
	"net/http"
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

func GetSummoner(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/summoner/")
	result := getApiResult("euw1", "summoner/v3/summoners/by-name/" + id)

	summoner := Summoner{}
	json.Unmarshal([]byte(result), &summoner)
	fmt.Println("Summoner ID:", summoner.Name)

	w.Write([]byte(result))
}
