package entity

import "time"

type Cursor struct {
	ID      int        `json:"id"`
	Time    *time.Time `json:"creation_time"`
	LastId  *int       `json:"last_id"`
	IsGroup bool       `json:"is_group"`
	Limit   int        `json:"limit"`
}
