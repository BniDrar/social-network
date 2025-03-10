CREATE TABLE IF NOT EXISTS followers (
    follower_user_id INTEGER,
    followed_user_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (followed_user_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (follower_user_id, followed_user_id)
);