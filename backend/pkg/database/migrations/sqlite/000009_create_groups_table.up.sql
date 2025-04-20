CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    type INTEGER NOT NULL, -- 0: real group, 1: fake, 2: messages group
    description TEXT,
    admin INTEGER,
    FOREIGN KEY (admin) REFERENCES users(id) ON DELETE SET NULL
);