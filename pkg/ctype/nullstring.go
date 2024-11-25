package ctype

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func NewNullString(value string) NullString {
	return NullString{sql.NullString{String: value, Valid: true}}
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.String)
}

func (n *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		n.Valid = false
		n.String = ""

		return nil
	}

	var str string
	err := json.Unmarshal(b, &str)
	n.String = str
	n.Valid = true

	return err
}

func (n NullString) IsNullOrEmpty() bool {
	return !n.Valid || n.String == ""
}

func (n NullString) IsEqual(v string) bool {
	return n.Valid && n.String == v
}
