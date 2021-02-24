package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// DefaultTimeFmt 默认时间序列化
const DefaultTimeFmt = "2006-01-02 15:04:05.000"

// Time json time
type Time struct {
	time.Time
}

// MarshalJSON makes Time implements json.Marshaler interface
func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, t.Format(DefaultTimeFmt))
	return []byte(formatted), nil
}

// UnmarshalJSON makes Time implements json.Unmarshaler interface
func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == "NULL" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	value, err := time.Parse(`"`+DefaultTimeFmt+`"`, string(data))
	if err != nil {
		return err
	}
	*t = Time{value}
	return nil
}

// Value makes Time implements drive.Valuer interface
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan makes Time implements sql.Scanner interface
func (t *Time) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = Time{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to time.Time", v)
}

// Now returns the current local time
func Now() Time {
	return Time{Time: time.Now()}
}
