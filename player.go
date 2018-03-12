package stattleship

import "time"

// Player holds information about a specific player
type Player struct {
	ID                   string    `json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Active               bool      `json:"active"`
	BirthDate            string    `json:"birth_date"`
	Captain              string    `json:"captain"`
	City                 string    `json:"city"`
	Country              string    `json:"country"`
	DraftSeason          string    `json:"draft_season"`
	DraftRound           string    `json:"draft_round"`
	DraftOverallPick     string    `json:"draft_overall_pick"`
	DraftTeamName        string    `json:"draft_team_name"`
	FirstName            string    `json:"first_name"`
	Handedness           string    `json:"handedness"`
	Bats                 string    `json:"bats"`
	Height               float32   `json:"height"`
	HighSchool           string    `json:"high_school"`
	UnitOfHeight         string    `json:"unit_of_height"`
	LastName             string    `json:"last_name"`
	Name                 string    `json:"name"`
	Nickname             string    `json:"nickname"`
	PositionAbbreviation string    `json:"position_abbreviation"`
	PositionName         string    `json:"reliever"`
	Salary               int       `json:"salary"`
	HumanizedSalary      string    `json:"humanized_salary"`
	SalaryCurrency       string    `json:"salary_currency"`
	School               string    `json:"school"`
	Slug                 string    `json:"slug"`
	Sport                string    `json:"sport"`
	State                string    `json:"state"`
	Weight               float32   `json:"weight"`
	UniformNumber        int       `json:"uniform_number"`
	UnitOfWeight         string    `json:"unit_of_weight"`
	YearsOfExperience    int       `json:"years_of_experience"`
	ProDebut             string    `json:"pro_debut"`
	LeagueID             string    `json:"league_id"`
	PlayingPositionID    string    `json:"playing_position_id"`
	TeamID               string    `json:"team_id"`
}
