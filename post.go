package blog

import "time"

// Post holds the basic user information
type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Private     bool      `json:"active"`
	CreatedByID int       `json:"created_by_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedByID int       `json:"updated_by_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
