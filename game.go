package baseballSim

import (
	"time"
)

// Game holds single game data
type Game struct {
	ID                 string    `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	AtNeutralSite      bool      `json:"at_neutral_site"`
	Attendance         int       `json:"attendance"`
	AwayTeamOutcome    string    `json:"away_team_outcome"`
	AwayTeamScore      int       `json:"away_team_score"`
	Broadcast          string    `json:"broadcast"`
	Clock              string    `json:"clock"`
	ClockSecs          int       `json:"clock_secs"`
	Daytime            bool      `json:"daytime"`
	Duration           int       `json:"duration"`
	EndedAt            time.Time `json:"ended_at"`
	HomeTeamOutcome    string    `json:"home_team_outcome"`
	HomeTeamScore      int       `json:"home_team_score"`
	Humidity           int       `json:"humidity"`
	Interval           string    `json:"interval"`
	IntervalNumber     int       `json:"interval_number"`
	IntervalType       string    `json:"interval_type"`
	Label              string    `json:"label"`
	Name               string    `json:"name"`
	On                 string    `json:"on"`
	Period             int       `json:"period"`
	PeriodLabel        string    `json:"period_label"`
	Score              string    `json:"score"`
	ScoreDifferential  string    `json:"score_differential"`
	Scoreline          string    `json:"scoreline"`
	Slug               string    `json:"slug"`
	StartedAt          string    `json:"started_at"`
	Status             string    `json:"status"`
	InternetCoverage   string    `json:"internet_coverage"`
	SatelliteCoverage  string    `json:"satellite_coverage"`
	TelevisionCoverage string    `json:"television_coverage"`
	Temperature        float32   `json:"temperature"`
	TemperatureUnit    string    `json:"temperature_unit"`
	Timestamp          int       `json:"timestamp"`
	Title              string    `json:"title"`
	WeatherConditions  string    `json:"weather_conditions"`
	WindDirection      string    `json:"wind_direction"`
	WindSpeed          float32   `json:"wind_speed"`
	WindSpeedUnit      string    `json:"wind_speed_unit"`
	HomeTeamID         string    `json:"home_team_id"`
	AwayTeamID         string    `json:"away_team_id"`
	WinningTeamID      string    `json:"winning_team_id"`
	SeasonID           string    `json:"season_id"`
	VenueID            string    `json:"venue_id"`
	OfficialIDs        []string  `json:"official_ids"`
}
