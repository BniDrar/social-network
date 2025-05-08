-- Follower-user 1 follows user 2
INSERT INTO follows (follower_id, followed_id) VALUES (1, 2);
-- Follower-user 1 follows user 3
INSERT INTO follows (follower_id, followed_id) VALUES (1, 3);
-- Follower-user 2 follows user 3
INSERT INTO follows (follower_id, followed_id) VALUES (2, 3);
-- Follower-user 3 follows user 1
INSERT INTO follows (follower_id, followed_id) VALUES (3, 1);
-- Follower-user 4 follows user 2
INSERT INTO follows (follower_id, followed_id) VALUES (4, 2);
-- Follower-user 5 follows user 6
INSERT INTO follows (follower_id, followed_id) VALUES (5, 6);
-- Follower-user 6 follows user 5
INSERT INTO follows (follower_id, followed_id) VALUES (6, 5);
-- Follower-user 7 follows user 8
INSERT INTO follows (follower_id, followed_id) VALUES (7, 8);
-- Follower-user 8 follows user 9
INSERT INTO follows (follower_id, followed_id) VALUES (8, 9);
-- Follower-user 9 follows user 10
INSERT INTO follows (follower_id, followed_id) VALUES (9, 10);
-- Follower-user 10 follows user 7

-- Additional random follows (you can continue with more follows in a similar fashion)
-- Randomly add a few more follow relationships
INSERT INTO follows (follower_id, followed_id) VALUES (2, 7);
INSERT INTO follows (follower_id, followed_id) VALUES (4, 6);
INSERT INTO follows (follower_id, followed_id) VALUES (3, 9);
INSERT INTO follows (follower_id, followed_id) VALUES (1, 5);
INSERT INTO follows (follower_id, followed_id) VALUES (6, 4);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 1);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 2);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 3);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 4);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 5);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 6);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 7);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 8);
INSERT INTO follows (follower_id, followed_id) VALUES (10, 9);
INSERT INTO follows (follower_id, followed_id) VALUES (9, 1);
INSERT INTO follows (follower_id, followed_id) VALUES (1, 9);
