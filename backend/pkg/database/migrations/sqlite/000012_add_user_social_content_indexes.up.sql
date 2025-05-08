CREATE INDEX idx_follows_following ON follows(follower_id);
CREATE INDEX idx_follows_followed ON follows(followed_id);
CREATE INDEX idx_post_user ON posts(user_id);
CREATE INDEX idx_comment_post ON comments(post_id);
CREATE INDEX idx_comment_user ON comments(user_id);
CREATE INDEX idx_message_sender ON messages(sender_id);
CREATE INDEX idx_sessions_expiry ON sessions(expiry);

-- Performance optimization indexes
CREATE INDEX idx_post_created_at ON posts(created_at DESC);
CREATE INDEX idx_message_created_at ON messages(created_at DESC);
CREATE INDEX idx_group_members_group ON group_members(group_id);
CREATE INDEX idx_group_members_user ON group_members(member_id);
CREATE INDEX idx_notification_receiver ON notification(sender_id);
CREATE INDEX idx_engagement_user ON engagement(user_id);
