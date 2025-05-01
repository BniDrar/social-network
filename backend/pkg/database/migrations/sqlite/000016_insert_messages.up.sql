-- 000016_insert_messages.up.sql

INSERT INTO messages (sender_id, receiver_id, group_id, content)
VALUES 
    (1, 2, NULL, 'Hey, how are you?'),
    (2, 1, NULL, 'I am good, thanks! How about you?'),
    (1, 3, NULL, 'Are you coming to the event tomorrow?'),
    (3, 1, NULL, 'Yes, I will be there.'),
    (2, NULL, 1, 'Hello group! This is my first message.'),
    (3, NULL, 1, 'Welcome to the group chat.'),
    (1, NULL, 1, 'Thanks! Glad to be here.');

-- Assumptions:
-- sender_id and receiver_id refer to valid users (IDs 1, 2, 3)
-- group_id refers to an existing group with ID = 1 (type = messages group or fake group)
