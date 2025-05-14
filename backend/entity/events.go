package entity

type Event struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	CreatedAt   string `json:"created_at"`
	Going       int    `json:"going"`
}

const (
	EventStatusNotVoted = iota
	EventStatusNotGoing
	EventStatusGoing
)
