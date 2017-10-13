package riot_api

import (
	"fmt"
	"encoding/json"
	"strings"
)

type MatchList struct {
	Matches []Match `json: "matches"`
}

type Match struct {
	PlatformId string `json: "platformId"`
	GameId     int    `json: "gameId"`
	Champion   int    `json: "champion"`
	Queue      int    `json: "queue"`
	Season     int    `json: "season"`
	Timestamp  int    `json: "timestamp"`
	Role       string `json: "role"`
	Lane       string `json: "lane"`
}

func GetRecentMatches(region string, accountId int) MatchList {
	res := getApiResult(region, fmt.Sprintf("match/v3/matchlists/by-account/%v/recent", accountId))
	r := new(MatchList)
	json.NewDecoder(strings.NewReader(res)).Decode(r)
	return *r
}