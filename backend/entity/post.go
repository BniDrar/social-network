package entity

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Avatar    string    `json:"avatar"`
	UserName  string    `json:"username"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Comments  uint      `json:"comments"`
	Likes     uint      `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	GroupID   uint
	Status    int
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Post) Validate() error {
	if len(strings.TrimSpace(p.Content)) >= 1 && len(p.Content) <= 1000 {
		switch p.Status {
		case 0:
			if p.GroupID == 0 {
				return errors.New("must Have groupID in Private post")
			}
		case 1:
			if p.GroupID == 0 {
				return errors.New("must Have groupID in Private post")
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
