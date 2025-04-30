package entity

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) SetValid(value bool) {
	ns.Valid = value
	ns.NullString.Valid = value
}
// Make it work with database/sql
func (ns *NullString) Scan(value interface{}) error {
	return ns.NullString.Scan(value)
}

func (ns NullString) Value() (driver.Value, error) {
	return ns.NullString.Value()
}

// Already shown: JSON marshaling
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}
