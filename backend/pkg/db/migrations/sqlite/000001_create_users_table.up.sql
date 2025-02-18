CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Email TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL,
    First TEXT,
    Last TEXT,
    Date_Of_Birth DATETIME,
    Avatar BLOB,
    Nickname TEXT UNIQUE,
    About_Me TEXT,
    status INTEGER DEFAULT 0  -- 0: private, 1: global
);