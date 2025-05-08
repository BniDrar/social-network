package entity

type Event struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Location    string `json:"location"`
	GroupID     int    `json:"group_id"`
	CreatedAt   string `json:"created_at"`
	Going       int   `json:"going"`
}
