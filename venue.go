package stattleship

import "time"

// Venue holds information about a stadium or arena
type Venue struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Abbreviation string    `json:"abbreviation"`
	Capacity     int       `json:"capacity"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	FieldType    string    `json:"field_type"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	State        string    `json:"state"`
	StadiumType  string    `json:"stadium_type"`
	TimeZone     string    `json:"time_zone"`
	Latitude     float32   `json:"latitude"`
	Longitude    float32   `json:"longitude"`
}
