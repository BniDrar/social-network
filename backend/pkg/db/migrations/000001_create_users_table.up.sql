-- create_users_table UP

CREATE IF NOT EXIST TABLE user (
    id SERIAL PRIMARY KEY,
    Email TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL,
    First TEXT,
    Last TEXT,
    Date_Of_Birth TIME,
    Avatar BYTEA,
    Nickname TEXT,
    About_Me TEXT,
    status INT
);