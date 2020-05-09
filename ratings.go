package mubi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Rating struct {
	Overall   int   `json:"overall"`
	Timestamp int64 `json:"updated_at"`
	Film      Film
}

type RatingsAPI struct {
	Client *http.Client
}

func (api *RatingsAPI) GetRatings(userId int64) []Rating {
	url := fmt.Sprintf("https://mubi.com/services/api/ratings?user_id=%d", userId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "jdennes/mubi")
	res, getErr := api.Client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	ratings := make([]Rating, 0)
	jsonErr := json.Unmarshal(body, &ratings)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return ratings
}
