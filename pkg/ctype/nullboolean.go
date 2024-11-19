package ctype

import (
	"database/sql"
	"encoding/json"
)

type NullBoolean struct {
	sql.NullBool
}

func NewNullBoolean(value bool) NullBoolean {
	return NullBoolean{sql.NullBool{Bool: value, Valid: true}}
}

func (n NullBoolean) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Bool)
}

func (n *NullBoolean) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		n.Valid = false
		n.Bool = false

		return nil
	}

	var value bool
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	*n = NewNullBoolean(value)

	return nil
}
