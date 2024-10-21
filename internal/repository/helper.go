package repository

import (
	"encoding/base64"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00"
)

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	bytes, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}
	timeString := string(bytes)
	t, _ := time.Parse(timeFormat, timeString)
	return t, err
}

func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)
	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
