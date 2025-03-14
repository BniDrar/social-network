package entity

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Image     []byte    `json:"image"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	Status    int       `json:"status"` // 0: private, 1: friends, 2: global
	GroupID   int       `json:"group"`  // Group ID, can be null
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Reaction struct {
	ID     int
	Status int
}
