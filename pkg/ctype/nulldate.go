package ctype

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullDate struct {
	sql.NullTime
}

func NewNullDate(value time.Time) NullDate {
	return NullDate{sql.NullTime{Time: value, Valid: true}}
}

func (n NullDate) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	t := n.Time.Format(TimeFormatYMD)

	return json.Marshal(t)
}

func (n *NullDate) UnmarshalJSON(b []byte) error {
	var strDate string

	err := json.Unmarshal(b, &strDate)
	if err != nil {
		return err
	}

	if strDate == "" {
		n.Valid = false
		return nil
	}

	t, err := time.Parse(TimeFormatYMD, strDate)
	if err != nil {
		return err
	}

	n.Valid = true
	n.Time = t

	return nil
}

func (n *NullDate) String() string {
	return n.Time.Format(TimeFormatYMD)
}
