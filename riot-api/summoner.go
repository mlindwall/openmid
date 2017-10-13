package riot_api

import "encoding/json"

type Summoner struct {
	Id int				`json: "id"`
	AccountId int		`json: "accountId"`
	Name string 		`json: "name"`
	ProfileIconId int	`json: "profileIconId"`
	RevisionDate int	`json: "revisionDate"`
	SummonerLevel int	`json: "summonerLevel"`
}

func GetSummoner(region, name string) Summoner {
	res := getApiResult(region, "summoner/v3/summoners/by-name/" + name)
	summoner := Summoner{}
	json.Unmarshal([]byte(res), &summoner)
	return summoner
}