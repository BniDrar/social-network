package entity

type User struct {
	ID             uint       `json:"id"`
	Nickname       NullString `json:"nickname"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	Avatar         NullString `json:"avatar"`
	First          string     `json:"first"`
	Last           string     `json:"last"`
	DateOfBirth    string     `json:"birthday"`
	AboutMe        string     `json:"about_me"`
	Status         uint       `json:"status"`
	ProfileOwner   bool       `json:"profile_owner"`
	IsFollowing    bool       `json:"is_following"` // current user follows the target user
	IsFollowed     bool       `json:"is_followed"`  // target user follows the current user
	FollowersCount uint       `json:"followers_count"`
	FollowingCount uint       `json:"following_count"`
	IsAdmin        bool       `json:"is_admin"` // whether the user is an admin of a group
}

type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	GroupName string `json:"group_name"`
	Avatar    string `json:"avatar"`
	Online    bool
}

type Follows struct {
	Followers []User
	Following []User
}

type Credentials struct {
	Username string `json:"nickname"`
	Password string `json:"password"`
}

const (
	PublicUser  = 0
	PrivateUser = 1
)
