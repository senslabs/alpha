package datastore

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type NullString struct{ sql.NullString }
type NullTime struct{ sql.NullTime }
type RawMessage struct{ json.RawMessage }

func (this *NullString) MarshalJSON() ([]byte, error) {
	if !this.Valid {
		this.Valid = true
		this.String = ""
	}
	return json.Marshal(this.String)
}

func (this *NullString) UnmarshalJSON(data []byte) error {
	if !this.Valid {
		this.Valid = true
		this.String = ""
	}
	return json.Unmarshal(data, &this.String)
}

func (this *NullTime) MarshalJSON() ([]byte, error) {
	if !this.Valid {
		this.Valid = true
		this.Time = time.Unix(0, 0)
	}
	return json.Marshal(this.Time)
}

func (this *NullTime) UnmarshalJSON(data []byte) error {
	if !this.Valid {
		this.Valid = true
		this.Time = time.Unix(0, 0)
	}
	return json.Unmarshal(data, &this.Time)
}

func (this RawMessage) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *RawMessage) Scan(value interface{}) error {
	if value == nil {
		value = []byte("{}")
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &this)
}
