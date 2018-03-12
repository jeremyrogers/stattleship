package stattleship

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// API interfaces with the API
type API struct {
	APIToken string
	Sport    string
	League   string
}

// JSONData stores Stattleship API calls
type JSONData struct {
	Games        []Game     `json:"games"`
	HomeTeams    []Team     `json:"home_teams"`
	AwayTeams    []Team     `json:"away_teams"`
	Leagues      []League   `json:"leagues"`
	WinningTeams []Team     `json:"winning_teams"`
	Seasons      []Season   `json:"seasons"`
	Venues       []Venue    `json:"venues"`
	Players      []Player   `json:"players"`
	Teams        []Team     `json:"teams"`
	Officials    []Official `json:"officials"`
}

// APIInit initialize the API with your token
func APIInit(token string, sport string, league string) *API {
	rv := new(API)
	rv.APIToken = token
	rv.Sport = sport
	rv.League = league
	return rv
}

// APICall makes an API call and stores the info in a struct
func (s *API) APICall(value string) (*JSONData, error) {
	apiEndpoint := fmt.Sprintf("https://api.stattleship.com/%v/%v/%v/", s.Sport, s.League, value)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatal("Error in NewRequest!")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token token="+s.APIToken)
	req.Header.Set("Accept", "application/vnd.stattleship.com; version=1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error in request!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var data *JSONData
	e := json.Unmarshal(body, &data)
	if e != nil {
		log.Fatal(e)
	}

	return data, nil
}
