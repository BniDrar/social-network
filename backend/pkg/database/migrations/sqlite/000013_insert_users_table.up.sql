-- Step 1: Create 10 users with random status (0 or 1)
INSERT INTO users (email, password, avatar, first_name, last_name, birthday, nickname, about_me, status)
VALUES
  ('user1@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First1', 'Last1', '1990-01-01', 'user1', 'About user 1', ROUND(RANDOM()) % 2),
  ('user2@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First2', 'Last2', '1992-02-02', 'user2', 'About user 2', ROUND(RANDOM()) % 2),
  ('user3@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First3', 'Last3', '1989-03-03', 'user3', 'About user 3', ROUND(RANDOM()) % 2),
  ('user4@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First4', 'Last4', '1995-04-04', 'user4', 'About user 4', ROUND(RANDOM()) % 2),
  ('user5@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First5', 'Last5', '1997-05-05', 'user5', 'About user 5', ROUND(RANDOM()) % 2),
  ('user6@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First6', 'Last6', '2000-06-06', 'user6', 'About user 6', ROUND(RANDOM()) % 2),
  ('user7@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First7', 'Last7', '1987-07-07', 'user7', 'About user 7', ROUND(RANDOM()) % 2),
  ('user8@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First8', 'Last8', '1993-08-08', 'user8', 'About user 8', ROUND(RANDOM()) % 2),
  ('user9@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First9', 'Last9', '1991-09-09', 'user9', 'About user 9', ROUND(RANDOM()) % 2),
  ('user10@example.com', '$2a$10$aJZQ6rpegN36ZVmJufmEquVQ2xaxJ.Z5na5tP4QdWBp43idLD8AUK', 'media/profile/profile.jpeg', 'First10', 'Last10', '1985-10-10', 'user10', 'About user 10', ROUND(RANDOM()) % 2);
