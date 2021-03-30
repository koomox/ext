package ext

import (
	"time"
	"strings"
)

const (
	customTimeFormat = "2006-01-02 15:04:05"
	frontTimeFormat = "2006-01-02T15:04:05Z"
	browserTimeFormat   = "2006/01/02 15:04"
)

type DateTime struct {
	Location *time.Location
	datetime string
	Time time.Time
}

func NewDateTime(datetime string) *DateTime {
	dt := &DateTime{datetime: datetime,}
	if dt.datetime == "" {
		dt.Time = time.Now()
		dt.Location, _ = time.LoadLocation("UTC")
		dt.datetime = dt.Time.In(dt.Location).Format(customTimeFormat)
	}

	return dt
}

func (dt *DateTime)Now() *DateTime {
	dt.Time = time.Now()
	return dt
}

func (dt *DateTime)UTC() *DateTime {
	dt.Location, _ = time.LoadLocation("UTC")
	return dt
}

func (dt *DateTime)CST() *DateTime {
	dt.Location, _ = time.LoadLocation("Asia/Shanghai")
	return dt
}

func (dt *DateTime)ParseTime() (*DateTime, error) {
	var err error
	if strings.Contains(dt.datetime, "/") {
		dt.Time, err = time.ParseInLocation(browserTimeFormat, dt.datetime, dt.Location)
		return dt, err
	}
	if strings.Contains(dt.datetime, "Z") {
		dt.Time, err = time.ParseInLocation(frontTimeFormat, dt.datetime, dt.Location)
		return dt, err
	}
	if strings.Contains(dt.datetime, "-") {
		dt.Time, err = time.ParseInLocation(customTimeFormat, dt.datetime, dt.Location)
		return dt, err
	}
	err = errors.New("parse time failed!")
	return dt, err
}

func (dt *DateTime)Parse() *DateTime {
	dt.Time, _ = time.ParseInLocation(customTimeFormat, dt.datetime, dt.Location)
	return dt
}

func (dt *DateTime)String() string {
	return dt.Time.In(dt.Location).Format(customTimeFormat)
}

func (dt *DateTime)Before(u *DateTime) bool {
	return dt.Time.Before(u.Time)
}

func NowNanosecond() int64 {
	return time.Now().UnixNano()
}

func DelayMillisecond(nanosecond int64) int {
	return int((NowNanosecond() - nanosecond)/1e6)
}