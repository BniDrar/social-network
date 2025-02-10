-- create_notifications_table UP

CREATE TABLE notification (
    id SERIAL PRIMARY KEY,
    type INT,
    group_id INT, -- Foreign key added later
    sender_id INT REFERENCES user(id) ON DELETE CASCADE,
    accepted BOOLEAN DEFAULT FALSE
);