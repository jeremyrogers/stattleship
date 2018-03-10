package baseballSim

import "time"

// Team holds data for a single team
type Team struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Color      string    `json:"color"`
	Colors     []string  `json:"colors"`
	Hashtag    string    `json:"hashtag"`
	Hashtags   []string  `json:"hashtags"`
	Location   string    `json:"location"`
	Name       string    `json:"name"`
	Nickname   string    `json:"nickname"`
	Latitude   float32   `json:"latitude"`
	Longitude  float32   `json:"longitude"`
	Slug       string    `json:"slug"`
	DivisionID string    `json:"division_id"`
	LeagueID   string    `json:"league_id"`
}
