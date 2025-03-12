INSERT OR IGNORE INTO users (
    email, 
    password, 
    first_name, 
    last_name, 
    birthday, 
    avatar, 
    nickname, 
    about_me, 
    status
) 
VALUES (
    'existuser@example.com',
    '$2a$10$lWWMbAFxfFFVv5XiMIcRjuu/ovU8burJnM6qSTaFfBfDk4.T3QsRC',
    'exist', 
    'user', 
    '2000-01-01 00:00:00', 
    NULL,
    'existuser',
    'I love coding and tea.',
    1 
);
