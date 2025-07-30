package commons

import (
	"database/sql"
)

type ICommonMapper interface {
	ToNullString(str string) sql.NullString
	FromNullString(ns sql.NullString) string
}

// commonMapper handles common data transformations.
type commonMapper struct{}

// NewMapper creates a new instance of CommonMapper.
func NewMapper() ICommonMapper {
	return &commonMapper{}
}

// ToNullString converts a string to sql.NullString.
func (m *commonMapper) ToNullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: str != ""}
}

// FromNullString converts sql.NullString to string, returns empty string if null.
func (m *commonMapper) FromNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
