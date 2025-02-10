-- create_follows_table UP

CREATE TABLE follows (
    following_user_id INT REFERENCES user(id) ON DELETE CASCADE,
    followed_user_id INT REFERENCES user(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (following_user_id, followed_user_id)
);