package entity

/*
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
*/
type User struct {
	ID             uint   `json:"id,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	Avatar         []byte `json:"avatar,omitempty"`
	First          string `json:"first,omitempty"`
	Last           string `json:"last,omitempty"`
	DateOfBirth    string `json:"date_of_birth,omitempty"`
	AboutMe        string `json:"about_me,omitempty"`
	Status         uint   `json:"status,omitempty"`
	FollowersCount uint   `json:"followers_count,omitempty"`
	FollowingCount uint   `json:"following_count,omitempty"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
