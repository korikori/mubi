package mubi

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

// The only way I know how to get a list of the current films of the day is to
// scrape some JSON out of https://mubi.com/film-of-the-day

type FilmOfTheDay struct {
	Id          int    `json:"id"`
	AvailableAt string `json:"available_at"`
	ExpiresAt   string `json:"expires_at"`
	Film        Film   `json:"film"`
}

type FilmsOfTheDayData struct {
	Properties struct {
		InitialState struct {
			FilmOfTheDayContainer struct {
				FilmsOfTheDay []FilmOfTheDay `json:"filmProgrammings"`
			} `json:"filmProgramming"`
		} `json:"initialState"`
	} `json:"props"`
}

func (api *MubiAPI) GetFilmsOfTheDay() []FilmOfTheDay {
	body := api.GetResponseBody("https://mubi.com/film-of-the-day")

	// Use the goquery library to find the
	// <script id="__NEXT_DATA__" ...> element containing the JSON data
	doc, loadErr := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if loadErr != nil {
		log.Fatal(loadErr)
	}
	element := doc.Find("script[id='__NEXT_DATA__']")
	jsonContent := element.Text()

	var filmsData FilmsOfTheDayData
	jsonErr := json.Unmarshal([]byte(jsonContent), &filmsData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return filmsData.Properties.InitialState.FilmOfTheDayContainer.FilmsOfTheDay
}
