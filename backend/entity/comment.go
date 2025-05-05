package entity

type Comment struct {
	ID          int        `json:"id"`
	CreaterName string     `json:"creater_name"`
	PostID      int        `json:"post_id"`
	UserID      int        `json:"user_id"`
	Content     string     `json:"content"`
	Image       NullString `json:"image"`
	Avatar      NullString `json:"avatar"`
}
