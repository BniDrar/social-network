CREATE INDEX idx_follows_following ON follows(following_user_id);
CREATE INDEX idx_follows_followed ON follows(followed_user_id);
CREATE INDEX idx_post_user ON post(user_id);
CREATE INDEX idx_comment_post ON comment(post_id);
CREATE INDEX idx_comment_user ON comment(user_id);
CREATE INDEX idx_message_chat ON message(chat_id);
CREATE INDEX idx_message_sender ON message(sender_id);
CREATE INDEX idx_message_receiver ON message(receiver_id);
-- Create an index on the expiry column for faster queries
CREATE INDEX sessions_expiry_idx ON sessions (expiry);

-- Consider adding these performance-improving indexes if needed:
-- CREATE INDEX idx_post_created_at ON post(created_at DESC);
-- CREATE INDEX idx_message_created_at ON message(created_at DESC);