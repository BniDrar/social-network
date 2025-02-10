-- create_posts_table UP

CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    title TEXT,
    image BYTEA,
    content TEXT,
    user_id INT REFERENCES user(id) ON DELETE CASCADE,
    status INT,
    group_id INT, -- Foreign key added later
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);