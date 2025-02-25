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
	ID       uint `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar    []byte `json:"avatar"`
	First    string `json:"first"`
	Last     string `json:"last"`
	DateOfBirth string `json:"date_of_birth"`
	AboutMe  string `json:"about_me"`
	Status   uint `json:"status"`
	FollowersCount uint `json:"followers_count"`
	FollowingCount uint `json:"following_count"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}