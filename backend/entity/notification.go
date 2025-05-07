package entity

import "time"

// NotificationType represents the type of notification
type NotificationType string

const (
	// Notification types
	FollowRequest    NotificationType = "follow_request"
	GroupInvite      NotificationType = "group_invite"
	GroupJoinRequest NotificationType = "group_join_request"
	NewEvent         NotificationType = "new_event"
)

// NotificationMessage represents a WebSocket notification message
type NotificationMessage struct {
	Type      NotificationType `json:"type"`
	ID        int              `json:"id"`
	SenderID  int              `json:"sender_id"`
	Content   string           `json:"content"`
	CreatedAt time.Time        `json:"created_at"`
	Data      interface{}      `json:"data,omitempty"` // Additional data specific to notification type
}

type Notification struct {
	Id         int    `json:"id"`
	Type       int    `json:"type"`
	GroupId    int    `json:"group_id"`
	SenderId   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	EventID    int    `json:"event_id"`
	Message    string `json:"message"`
	Accepted   bool   `json:"accepted"`
}

const (
	//notification answer
	NotificationWithoutAnswer = iota
	NotificationAccepted
	NotificationNotAccepted
)

const (
	//type of notification
	FollowingNotification = iota
	EventNotification
	GroupInvitationNotification
	GroupParticipationNotification
)
