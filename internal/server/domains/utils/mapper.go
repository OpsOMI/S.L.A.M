package utils

import (
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
)

// IUtilMapper defines utility methods for mapping primitive values to pointer and nullable types.
type IUtilMapper interface {
	// String utils
	FromStrToPtrStr(s string) *string
	FromPtrStrToNullStr(s *string) sql.NullString
	FromStrToNullStr(s string) sql.NullString
	FromNullStrToPtrStr(s sql.NullString) *string
	TrimIfNotNil(s *string) *string

	// Int32 utils
	FromInt32ToPtrInt32(i int32) *int32
	FromPtrInt32ToNullInt32(i *int32) sql.NullInt32
	FromInt32ToNullInt32(v int32) sql.NullInt32
	FromNullInt32ToPtrInt32(v sql.NullInt32) *int32

	// Bool utils
	FromPtrBoolToNullBool(b *bool) sql.NullBool
	FromStringToNullBool(s string) sql.NullBool
	FromNullBoolToPtrBool(b sql.NullBool) *bool

	// Time utils
	FromTimeToPtrTime(t time.Time) *time.Time
	FromPtrTimeToNullTime(t *time.Time) sql.NullTime
	FromNullTimeToPtrTime(input sql.NullTime) *time.Time

	// UUID utils
	FromUUIDToNullUUID(uid uuid.UUID) uuid.NullUUID
	FromNullUUIDToUUIDPtr(uid uuid.NullUUID) *uuid.UUID
	FromUUIDPtrToUUIDNull(uidPtr *uuid.UUID) uuid.NullUUID
}

// mapper handles transformation between primitive values and their pointer/nullable equivalents.
type mapper struct{}

// NewMapper returns a new instance of utilMapper.
func NewMapper() IUtilMapper {
	return &mapper{}
}

//
// ======================
// String helpers
// ======================

func (m *mapper) FromStrToPtrStr(s string) *string {
	if s == "" {
		return nil
	}
	trimmed := strings.TrimSpace(s)
	return &trimmed
}

func (m *mapper) FromPtrStrToNullStr(s *string) sql.NullString {
	if s == nil || *s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: strings.TrimSpace(*s), Valid: true}
}

func (m *mapper) FromStrToNullStr(s string) sql.NullString {
	if strings.TrimSpace(s) == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: strings.TrimSpace(s), Valid: true}
}

func (m *mapper) FromNullStrToPtrStr(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}

	return &s.String
}

func (m *mapper) TrimIfNotNil(s *string) *string {
	if s == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*s)
	return &trimmed
}

//
// ======================
// Int32 helpers
// ======================

func (m *mapper) FromInt32ToPtrInt32(i int32) *int32 {
	if i == 0 {
		return nil
	}
	return &i
}

func (m *mapper) FromPtrInt32ToNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func (m *mapper) FromInt32ToNullInt32(v int32) sql.NullInt32 {
	if v == 0 {
		return sql.NullInt32{Int32: v, Valid: false}
	}
	return sql.NullInt32{Int32: v, Valid: true}
}

func (m *mapper) FromNullInt32ToPtrInt32(v sql.NullInt32) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}

//
// ======================
// Bool helpers
// ======================

func (u *mapper) FromPtrBoolToNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}

func (u *mapper) FromStringToNullBool(s string) sql.NullBool {
	s = strings.TrimSpace(strings.ToLower(s))
	if s == "" {
		return sql.NullBool{Valid: false}
	}
	if s == "true" {
		return sql.NullBool{Bool: true, Valid: true}
	}
	if s == "false" {
		return sql.NullBool{Bool: false, Valid: true}
	}
	return sql.NullBool{Valid: false}
}

func (u *mapper) FromNullBoolToPtrBool(b sql.NullBool) *bool {
	if !b.Valid {
		return nil
	}
	return &b.Bool
}

//
// ======================
// Time helpers
// ======================

func (m *mapper) FromTimeToPtrTime(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func (m *mapper) FromPtrTimeToNullTime(t *time.Time) sql.NullTime {
	if t == nil || t.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func (u *mapper) FromNullTimeToPtrTime(input sql.NullTime) *time.Time {
	if !input.Valid {
		return nil
	}
	return &input.Time
}

//
// ======================
// UUID helpers
// ======================

func (u *mapper) FromUUIDToNullUUID(uid uuid.UUID) uuid.NullUUID {
	if uid == uuid.Nil {
		return uuid.NullUUID{UUID: uid, Valid: false}
	}
	return uuid.NullUUID{UUID: uid, Valid: true}
}

func (u *mapper) FromNullUUIDToUUIDPtr(uid uuid.NullUUID) *uuid.UUID {
	if !uid.Valid || uid.UUID == uuid.Nil {
		return nil
	}

	return &uid.UUID
}

func (u *mapper) FromUUIDPtrToUUIDNull(uidPtr *uuid.UUID) uuid.NullUUID {
	if uidPtr == nil {
		return uuid.NullUUID{Valid: false}
	}

	return uuid.NullUUID{UUID: *uidPtr, Valid: true}
}
