package entity

type Notification struct {
	Id int `json:"id"`
	Type int `json:"type"`
	GroupId int `json:"group_id"`
	SenderId int `json:"sender_id"`
	Message string `json:"message"`
	Accepted bool `json:"accepted"`
}

const (
	//notification answer
	NotificationWithoutAnswer = 0
	NotificationAccepted = 1
	NotificationNotAccepted = 2
	//type of notification
	FollowingNotification = 0
	EventNotification = 1
	GroupInvitationNotification = 2
	GroupParticipationNotification = 3
)