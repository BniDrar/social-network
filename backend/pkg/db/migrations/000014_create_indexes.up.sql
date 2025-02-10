-- create_indexes UP

CREATE INDEX idx_user_email ON user(Email);
CREATE INDEX idx_follows_following_user_id ON follows(following_user_id);
CREATE INDEX idx_follows_followed_user_id ON follows(followed_user_id);
CREATE INDEX idx_post_user_id ON post(user_id);
CREATE INDEX idx_comment_post_id ON comment(post_id);
CREATE INDEX idx_message_chat_id ON message(chat_id);