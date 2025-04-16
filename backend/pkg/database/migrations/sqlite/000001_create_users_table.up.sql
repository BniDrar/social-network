CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    birthday DATETIME,
    avatar TEXT,
    nickname TEXT UNIQUE,
    about_me TEXT,
    status INTEGER DEFAULT 0  -- 0: private, 1: global
);