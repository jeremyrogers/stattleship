package stattleship

import "time"

// Official holds information about a referee or official
type Official struct {
	ID            string    `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Name          string    `json:"name"`
	Role          string    `json:"role"`
	RoleLabel     string    `json:"role_label"`
	UniformNumber string    `json:"uniform_number"`
}
