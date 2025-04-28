-- Step 1: Create 10 users
INSERT INTO users (email, password, first_name, last_name, birthday, avatar, nickname, about_me, status)
VALUES
  ('user1@example.com', 'media/profile/default.png', 'First1', 'Last1', '1990-01-01', NULL, 'user1', 'About user 1'),
  ('user2@example.com', 'media/profile/default.png', 'First2', 'Last2', '1992-02-02', NULL, 'user2', 'About user 2'),
  ('user3@example.com', 'media/profile/default.png', 'First3', 'Last3', '1989-03-03', NULL, 'user3', 'About user 3'),
  ('user4@example.com', 'media/profile/default.png', 'First4', 'Last4', '1995-04-04', NULL, 'user4', 'About user 4'),
  ('user5@example.com', 'media/profile/default.png', 'First5', 'Last5', '1997-05-05', NULL, 'user5', 'About user 5'),
  ('user6@example.com', 'media/profile/default.png', 'First6', 'Last6', '2000-06-06', NULL, 'user6', 'About user 6'),
  ('user7@example.com', 'media/profile/default.png', 'First7', 'Last7', '1987-07-07', NULL, 'user7', 'About user 7'),
  ('user8@example.com', 'media/profile/default.png', 'First8', 'Last8', '1993-08-08', NULL, 'user8', 'About user 8'),
  ('user9@example.com', 'media/profile/default.png', 'First9', 'Last9', '1991-09-09', NULL, 'user9', 'About user 9'),
  ('user10@example.com', 'media/profile/default.png', 'First10', 'Last10', '1985-10-10', NULL, 'user10', 'About user 10');