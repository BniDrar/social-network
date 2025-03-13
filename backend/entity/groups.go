package entity

type Groups []Group

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Admin       int    `json:"admin"`
	UserID      int    `json:"user_id"`
	MemberCount int    `json:"member_count"`
	PostCount   int    `json:"post_count"`
	Members     []int  `json:"members"`
	Posts       []Post `json:"posts"`
}
