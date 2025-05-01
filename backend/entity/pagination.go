package entity

import "time"

type Cursor struct {
	UserId int        `json:"user_id"`
	Time   *time.Time `json:"creation_time"`
	LastId     *int       `json:"last_id"`
}

const LIMIT = 10
