-- create_sessions_table UP

CREATE TABLE session (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES user(id) ON DELETE CASCADE,
    token TEXT,
    data TEXT,
    expiry TIMESTAMP
);