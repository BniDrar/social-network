package entity

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	Status    int       `json:"status"`
	GroupID   int       `json:"group"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
