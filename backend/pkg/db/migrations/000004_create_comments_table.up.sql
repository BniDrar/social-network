-- create_comments_table UP

CREATE TABLE comment (
    id SERIAL PRIMARY KEY,
    post_id INT REFERENCES post(id) ON DELETE CASCADE,
    user_id INT REFERENCES user(id) ON DELETE CASCADE,
    content TEXT
);