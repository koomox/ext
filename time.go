package ext

import (
	"time"
	"strings"
	"errors"
)

const (
	dateTimeFormat = "2006-01-02"
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
	dt := &DateTime{datetime: strings.Trim(datetime, " "),}
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

func (dt *DateTime)parser() (*DateTime, error) {
	var (
		err error
		year = ""
		month = ""
		day = ""
		hour = "00"
		minute = "00"
		second = "00"
	)
	length := len(dt.datetime)
	if dt.datetime == "" || length < 8 {
		err = errors.New("parse date time failed")
		return dt, err
	}
	offset := 4
	year = dt.datetime[:offset]
	if !IsNumber(dt.datetime[offset]) {
		offset += 1
	}
	month = dt.datetime[offset:offset+2]
	offset += 2
	if !IsNumber(dt.datetime[offset]) {
		offset += 1
	}
	day = dt.datetime[offset:offset+2]
	offset += 2
	if length >= 16 {
		offset += 1
		hour = dt.datetime[offset:offset+2]
		offset += 2
		if !IsNumber(dt.datetime[offset]) {
			offset += 1
		}
		minute = dt.datetime[offset:offset+2]
		offset += 2
	}
	if length >= 18 {
		if !IsNumber(dt.datetime[offset]) {
			offset += 1
		}
		second = dt.datetime[offset:offset+2]
	}


	ch := strings.Join([]string{year, month, day, hour, minute, second}, "")
	length = len(ch)
	for _, v := range ch {
		if !IsNumber(v) {
			err = errors.New("parse date time failed")
			return dt, err
		}
		if v == '0' {
			length--
		}
	}
	if length == 0 {
		err = errors.New("parse date time failed")
		return dt, err
	}
	dt.datetime = fmt.Sprintf("%v-%v-%v %v:%v:%v", year, month, day, hour, minute, second)

	return dt, err
}

func (dt *DateTime)ParseTime() (*DateTime, error) {
	return dt.parser()
}

func (dt *DateTime)Parse() *DateTime {
	dt.Time, _ = time.ParseInLocation(customTimeFormat, dt.datetime, dt.Location)
	return dt
}

func (dt *DateTime)Format(layout string) string {
	return dt.Time.In(dt.Location).Format(layout)
}

func (dt *DateTime)String() string {
	return dt.Time.In(dt.Location).Format(customTimeFormat)
}

func (dt *DateTime)DateString() string {
	return dt.Time.In(dt.Location).Format(dateTimeFormat)
}

func (dt *DateTime)Before(u *DateTime) bool {
	return dt.Time.Before(u.Time)
}

func (dt *DateTime)Sub(u *DateTime) time.Duration {
	return dt.Time.Sub(u.Time)
}

func NowNanosecond() int64 {
	return time.Now().UnixNano()
}

func DelayMillisecond(nanosecond int64) int {
	return int((NowNanosecond() - nanosecond)/1e6)
}