SELECT user.nickname, user.avatar From followers
INNER JOIN user ON user.id = follower_id