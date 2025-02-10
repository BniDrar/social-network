-- add_foreign_keys UP

ALTER TABLE post ADD CONSTRAINT fk_post_group FOREIGN KEY (group_id) REFERENCES group(id) ON DELETE CASCADE;
ALTER TABLE event ADD CONSTRAINT fk_event_group FOREIGN KEY (group_id) REFERENCES group(id) ON DELETE CASCADE;
ALTER TABLE notification ADD CONSTRAINT fk_notification_group FOREIGN KEY (group_id) REFERENCES group(id) ON DELETE CASCADE;
