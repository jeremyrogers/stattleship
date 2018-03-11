package baseballSim

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// StattleshipAPI interfaces with the API
type StattleshipAPI struct {
	APIToken string
	Sport    string
	League   string
}

// StattleshipJSONData stores Stattleship API calls
type StattleshipJSONData struct {
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
func APIInit(token string, sport string, league string) *StattleshipAPI {
	rv := new(StattleshipAPI)
	rv.APIToken = token
	rv.Sport = sport
	rv.League = league
	return rv
}

// APICall makes an API call and stores the info in a struct
func (s *StattleshipAPI) APICall() (*StattleshipJSONData, error) {
	req, err := http.NewRequest("GET", "https://api.stattleship.com/baseball/mlb/seasons", nil)
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
	data := new(StattleshipJSONData)
	json.Unmarshal(body, &data)

	return data, nil
}
