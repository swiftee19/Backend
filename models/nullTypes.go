package models

import (
	"database/sql"
	"encoding/json"
)

// NullString is a wrapper around sql.NullString for custom JSON marshaling.
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString to return null instead of {"String": "", "Valid": false}.
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// Scan implements the sql.Scanner interface to handle database values.
func (ns *NullString) Scan(value interface{}) error {
	return ns.NullString.Scan(value)
}

// NullTime is a wrapper around sql.NullTime for custom JSON marshaling.
type NullTime struct {
	sql.NullTime
}

// MarshalJSON for NullTime to return null instead of {"Time": "0001-01-01T00:00:00Z", "Valid": false}.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}

// Scan implements the sql.Scanner interface to handle database values.
func (nt *NullTime) Scan(value interface{}) error {
	return nt.NullTime.Scan(value)
}
