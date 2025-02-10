-- create_events_table UP

CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    group_id INT, -- Foreign key added later
    user_id INT REFERENCES user(id) ON DELETE CASCADE,
    title TEXT,
    description TEXT,
    event_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);