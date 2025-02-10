-- create_groups_table UP

CREATE TABLE group (
    id SERIAL PRIMARY KEY,
    name TEXT,
    type INT,
    admin INT REFERENCES user(id) ON DELETE CASCADE
);