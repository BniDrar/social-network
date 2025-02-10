-- create_engagement_table UP

CREATE TABLE engagement (
    id SERIAL PRIMARY KEY,
    comment_id INT REFERENCES comment(id) ON DELETE CASCADE,
    post_id INT REFERENCES post(id) ON DELETE CASCADE,
    event_id INT REFERENCES event(id) ON DELETE CASCADE,
    user_id INT REFERENCES user(id) ON DELETE CASCADE,
    status INT
);