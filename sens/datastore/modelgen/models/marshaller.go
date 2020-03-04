package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullString struct{ sql.NullString }
type NullTime struct{ sql.NullTime }

func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		x.Valid = true
		x.String = ""
	}
	return json.Marshal(x.String)
}

func (x *NullTime) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		x.Valid = true
		x.Time = time.Unix(0, 0)
	}
	return json.Marshal(x.Time)
}
