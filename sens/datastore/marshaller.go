package datastore

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/senslabs/alpha/sens/logger"
)

type NullString struct{ sql.NullString }
type NullInt64 struct{ sql.NullInt64 }
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

func (this *NullInt64) MarshalJSON() ([]byte, error) {
	logger.Debugf("%#v", *this)
	if !this.Valid {
		return []byte("0"), nil
	}
	return json.Marshal(this.Int64)
}

func (this *NullInt64) UnmarshalJSON(b []byte) error {
	if !this.Valid {
		this.Valid = true
		this.Int64 = 0
	}
	return json.Unmarshal(b, &this.Int64)
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
