package entity

type Engagement struct {
	ID        int    `json:"id"`
	CommentID int    `json:"comment_id"`
	PostID    int    `json:"post_id"`
	EventID   int    `json:"event_id"`
	UserID    int    `json:"user_id"`
	Status    int    `json:"status"`
}