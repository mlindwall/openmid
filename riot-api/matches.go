package riot_api

import (
	"net/http"
	"fmt"
)

func GetRecentMatches(w http.ResponseWriter, region string, accountId int) {
	getApiResult(region, fmt.Sprintf("/lol/match/v3/matchlists/by-account/%s/recent", accountId))
}