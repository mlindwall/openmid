package riot_api

import (
	"net/http"
	"io/ioutil"
)

const key = "" // Riot Games API Key

func apiRequest(w http.ResponseWriter, region, endpoint string) {
	result := getApiResult(region, endpoint)
	if result == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte(result))
}

func getApiResult(region, endpoint string) string {
	url := "https://" + region + ".api.riotgames.com/lol/" + endpoint

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return ""
	}

	req.Header.Set("X-Riot-Token", key)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil || res.StatusCode != 200 {
		return ""
	}

	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return ""
	}

	return string(result)
}