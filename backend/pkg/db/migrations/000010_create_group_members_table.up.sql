-- create_group_members_table UP

CREATE TABLE group_members (
    id SERIAL PRIMARY KEY,
    member_id INT REFERENCES user(id) ON DELETE CASCADE,
    group_id INT REFERENCES group(id) ON DELETE CASCADE,
    UNIQUE (member_id, group_id)
);