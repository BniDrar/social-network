package entity

type User struct {
	ID             uint       `json:"id,omitempty"`
	Nickname       NullString `json:"nickname,omitempty"`
	Email          string     `json:"email,omitempty"`
	Password       string     `json:"password,omitempty"`
	Avatar         NullString `json:"avatar,omitempty"`
	First          string     `json:"first,omitempty"`
	Last           string     `json:"last,omitempty"`
	DateOfBirth    string     `json:"birthday,omitempty"`
	AboutMe        string     `json:"about_me,omitempty"`
	Status         uint       `json:"status,omitempty"`
	ProfileOwner   bool       `json:"profile_owner,omitempty"`
	IsFollowing     bool       `json:"is_following"` // current user follows the target user
	IsFollowed     bool       `json:"is_followed"`  // target user follows the current user
	FollowersCount uint       `json:"followers_count,omitempty"`
	FollowingCount uint       `json:"following_count,omitempty"`
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
