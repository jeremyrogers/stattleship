package baseballSim

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
	StartsOn            time.Time `json:"starts_on"`
	EndsOn              time.Time `json:"ends_on"`
	Slug                string    `json:"slug"`
	LeagueID            string    `json:"league_id"`
}
