package entity

import (
	"time"
)

type Chat struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FriendID  int       `json:"friend_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Chats []Chat

type groupoup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Type        int       `json:"type"` // 0 - public chat group, 1 - private chat group
	Admin       int       `json:"admin"`
	MemberCount int       `json:"member_count"`
	Members     []int     `json:"members"`
	Messages    []Message `json:"messages"`
}

type Message struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"from"`
	ReceiverID *int      `json:"to,omitempty"`       // nil if group message
	GroupID    *int      `json:"group_id,omitempty"` // nil if private message
	Content    string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
	IsGroup    bool      `json:"is_group"`
}
