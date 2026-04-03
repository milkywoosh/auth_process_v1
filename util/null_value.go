package util

import (
	"database/sql"
	"time"
)

func NullableInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: i, Valid: true}
}

func NullableInt(n sql.NullInt64) *int {
	if !n.Valid {
		return nil
	}
	v := int(n.Int64)
	return &v
}

func NullableString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}

	v := string(s.String)
	return &v
}

func NullableTime(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}

	v := t.Time
	return &v
}
