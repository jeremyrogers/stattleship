package stattleship

import (
	"time"
)

// Season contains information about the data's season
type Season struct {
	ID                  string    `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Name                string    `json:"name"`
	CurrentIntervalType string    `json:"current_interval_type"`
	StartsOn            string    `json:"starts_on"`
	EndsOn              string    `json:"ends_on"`
	Slug                string    `json:"slug"`
	LeagueID            string    `json:"league_id"`
}
