package entity

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID             int            `json:"id"`
	Avatar         sql.NullString `json:"avatar"`
	Nickname       string         `json:"nickname"`
	First          string         `json:"first_name"`
	Last           string         `json:"last_name"`
	UserName       sql.NullString `json:"username"`
	Content        string         `json:"content"`
	Image          string         `json:"image"`
	Comments       uint           `json:"comments"`
	CreatedAt      time.Time      `json:"created_at"`
	GroupName      string         `json:"groupe_name"`
	GroupID        uint
	Engagement     int       `json:"engagement"`
	LikesCount     int       `json:"likes_count"`
	Status         int       `json:"status"`
	AllowedViewers []int     `json:"allowed_viewers"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Target struct {
	Target uint `json:"target"`
	Limit  uint `json:"liimit"`
	Offset uint `json:"offset"`
}

// -- 0: custom, 1: friends, 2: global
const (
	PostStatusCustom = iota
	PostStatusFriends
	PostStatusGlobal
)

func (p *Post) Validate(users []string) error {
	if len(strings.TrimSpace(p.Content)) >= 1 && len(p.Content) <= 1000 {
		switch p.Status {
		case 0:
			if p.GroupID == 0 {
				return errors.New("must Have groupID in Private post")
			}
		case 1:
			if p.GroupID != 0 {
				return errors.New("must Have groupID in Private post")
			}
			if len(users) == 0 {
			}
		case 2:
			if p.GroupID == 0 {
				return errors.New("must Have'nt groupID in Public post")
			}
		}
	} else {
		return errors.New("content Lenght")
	}
	return nil
}
