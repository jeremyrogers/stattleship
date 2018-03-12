package stattleship

import "time"

// League holds information about the league (e.g. NFL, NBA, MLB)
type League struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Abbreviation     string    `json:"abbreviation"`
	Color            string    `json:"color"`
	MinutesPerPeriod int       `json:"minutes_per_period"`
	Name             string    `json:"name"`
	Periods          int       `json:"periods"`
	Slug             string    `json:"slug"`
	Sport            string    `json:"sport"`
}
