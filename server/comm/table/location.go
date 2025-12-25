package table

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// UnmarshalJSON handles JSON parsing for LocalTime.
// It checks for empty strings and null values before attempting to parse the time.
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	// Convert the byte slice to a string and trim quotes
	timeStr := strings.Trim(string(data), `"`)

	// Check for empty values (empty string or JSON null)
	if timeStr == "" || timeStr == "null" {
		// Set to zero value if empty
		*t = LocalTime(time.Time{})
		return nil
	}

	// Attempt to parse the time string using the expected format
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return err
	}
	*t = LocalTime(parsedTime)
	return nil
}
