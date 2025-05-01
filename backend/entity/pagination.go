package entity

import "time"

type Cursor struct {
	Time *time.Time `json:"creation_time"`
	Id *int `json:"id"`
}

const LIMIT = 10