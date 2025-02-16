CREATE TABLE IF NOT EXISTS post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    image BLOB,
    content TEXT,
    user_id INTEGER,
    status INTEGER DEFAULT 0, -- 0: private, 1: friends, 2: global
    "group" INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY ("group") REFERENCES "group"(id) ON DELETE CASCADE
);