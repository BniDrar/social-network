CREATE INDEX idx_follows_following ON follows(follower_id);
CREATE INDEX idx_follows_followed ON follows(followed_id);
CREATE INDEX idx_post_user ON posts(user_id);
CREATE INDEX idx_comment_post ON comments(post_id);
CREATE INDEX idx_comment_user ON comments(user_id);
-- CREATE INDEX idx_message_chat ON message(chat_id);
CREATE INDEX idx_message_sender ON messages(sender_id);
-- Create an index on the expiry column for faster queries
CREATE INDEX sessions_expiry_idx ON sessions (expiry);

-- Consider adding these performance-improving indexes if needed:
-- CREATE INDEX idx_post_created_at ON post(created_at DESC);
-- CREATE INDEX idx_message_created_at ON message(created_at DESC);