CREATE TABLE IF NOT EXISTS follows (
    following_user_id INTEGER,
    followed_user_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (following_user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (followed_user_id) REFERENCES user(id) ON DELETE CASCADE,
    PRIMARY KEY (following_user_id, followed_user_id)
);