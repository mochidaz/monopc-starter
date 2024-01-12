package utils

import (
	"database/sql"
	"time"
)

func StringToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func ParseDate(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}
