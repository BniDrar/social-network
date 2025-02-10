-- create_messages_table UP

CREATE TABLE message (
    id SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chat(id) ON DELETE CASCADE,
    receiver_id INT REFERENCES user(id) ON DELETE CASCADE,
    sender_id INT REFERENCES user(id) ON DELETE CASCADE,
    group_id INT REFERENCES group(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    content TEXT
);