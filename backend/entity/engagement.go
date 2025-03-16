package entity

/*
CREATE TABLE IF NOT EXISTS engagement (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    comment_id INTEGER,
    post_id INTEGER,
    event_id INTEGER,
    user_id INTEGER,
    status INTEGER, -- 0: none, 1: like, 2: dislike
    FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
*/

type Engagement struct {
	ID        int    `json:"id"`
	CommentID int    `json:"comment_id"`
	PostID    int    `json:"post_id"`
	EventID   int    `json:"event_id"`
	UserID    int    `json:"user_id"`
	Status    int    `json:"status"`
}