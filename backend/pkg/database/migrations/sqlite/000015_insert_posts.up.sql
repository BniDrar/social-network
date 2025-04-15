INSERT INTO posts (title, image, content, user_id, status, group_id, created_at, updated_at)
VALUES
('Post 1', X'89504E470D0A1A0A', 'This is the content of post 1.', 1, 0, 1, datetime('now'), datetime('now')),
('Post 2', X'89504E470D0A1A0A', 'This is the content of post 2.', 2, 1, 1, datetime('now'), datetime('now')),
('Post 3', X'89504E470D0A1A0A', 'This is the content of post 3.', 3, 2, 2, datetime('now'), datetime('now')),
('Post 4', X'89504E470D0A1A0A', 'This is the content of post 4.', 1, 0, 2, datetime('now'), datetime('now')),
('Post 5', X'89504E470D0A1A0A', 'This is the content of post 5.', 2, 1, 3, datetime('now'), datetime('now')),
('Post 6', X'89504E470D0A1A0A', 'This is the content of post 6.', 3, 2, 3, datetime('now'), datetime('now')),
('Post 7', X'89504E470D0A1A0A', 'This is the content of post 7.', 1, 1, 1, datetime('now'), datetime('now')),
('Post 8', X'89504E470D0A1A0A', 'This is the content of post 8.', 2, 2, 2, datetime('now'), datetime('now')),
('Post 9', X'89504E470D0A1A0A', 'This is the content of post 9.', 3, 0, 3, datetime('now'), datetime('now')),
('Post 10', X'89504E470D0A1A0A', 'This is the content of post 10.', 1, 1, 1, datetime('now'), datetime('now'));


INSERT INTO groups (name, type, description, admin) VALUES 
('Adventure Seekers', 0, 'A group for outdoor and travel enthusiasts.', 1),
('Bot Brigade', 1, 'Totally real people doing totally real stuff.', 2),
('Message Overflow', 2, 'Archive of daily message dumps.', 3),
('Tech Talk', 0, 'Discuss the latest in technology and gadgets.', 4),
('Phantom Squad', 1, 'Ghost group with no real members.', NULL),
('Inbox Infinity', 2, 'All messages, all the time.', 1),
('Foodie Club', 0, 'Sharing recipes, reviews, and food photos.', 2),
('AI Enthusiasts', 0, 'For fans of artificial intelligence and machine learning.', 3),
('Echo Chamber', 1, 'Where fake accounts echo fake news.', 4),
('Chat Archives', 2, 'Auto-generated logs of past discussions.', NULL);

INSERT INTO group_members (member_id, group_id) VALUES 
(1, 1),
(2, 1),
(3, 1),
(4, 2),
(1, 2),
(5, 2),
(6, 3),
(7, 3),
(2, 3),
(3, 4),
(5, 4),
(6, 4),
(1, 5),
(7, 5),
(8, 6),
(9, 6),
(10, 6),
(2, 7),
(4, 7),
(6, 8),
(7, 8),
(3, 9),
(8, 9),
(10, 10),
(1, 10);

