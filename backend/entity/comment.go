package entity

import "time"

type Comment struct {
	ID          int        `json:"id"`
	CreaterName NullString `json:"creater_name"`
	PostID      int        `json:"post_id"`
	UserID      int        `json:"user_id"`
	Content     string     `json:"content"`
	Image       NullString `json:"image"`
	Avatar      NullString `json:"avatar"`
	CreatedAt   time.Time  `json:"created_at"`
}
