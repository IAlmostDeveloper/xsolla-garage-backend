package dto

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

const DateFormat = "2006-01-02 15:04:05"

type TimeJson time.Time

func (t TimeJson) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(DateFormat))
	return []byte(stamp), nil
}

func (t *TimeJson) UnmarshalJSON(src []byte) error {
	timeString := string(src)
	timeString = strings.Trim(timeString, "\"")
	parsedTime, err := time.Parse(DateFormat, timeString)
	if err != nil {
		return err
	}
	*t = TimeJson(parsedTime)
	return nil
}

func (t TimeJson) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Scan assigns a value from a database driver.
//
// The src value will be of one of the following types:
//
//    int64
//    float64
//    bool
//    []byte
//    string
//    time.Time
//    nil - for NULL values
//
// An error should be returned if the value cannot be stored
// without loss of information.
//
// Reference types such as []byte are only valid until the next call to Scan
// and should not be retained. Their underlying memory is owned by the driver.
// If retention is necessary, copy their values before the next call to Scan.
func (t *TimeJson) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	log.Print(string(bytes))
	if !ok {
		return errors.New("cannot convert value to string")
	}
	parsedTime, err := time.Parse(DateFormat, string(bytes))
	if err != nil {
		return err
	}
	*t = TimeJson(parsedTime)
	return nil
}
