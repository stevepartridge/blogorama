package blog

import "time"

// User holds the basic user information
type User struct {
	ID          int       `json:"id"`
	APIKey        string    `json:"api_key"`
	Name        string    `json:"name"`
	Active      bool      `json:"active"`
	CreatedByID int       `json:"created_by_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedByID int       `json:"updated_by_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
