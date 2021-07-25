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
	)
	timeArray := []string{"0000", "00", "00", "00", "00", "00"}
	timeOffset := []int{4, 2, 2, 2, 2, 2}
	length := len(dt.datetime)
	if dt.datetime == "" {
		err = errors.New("date time is null")
		return dt, err
	}
	offset := 0
	for i := 0; i < len(timeArray); i++ {
		if length <= 0 {
			break
		}
		if !IsNumber(dt.datetime[offset]) {
			offset++
			length--
		}
		if length < timeOffset[i] {
			break
		}
		timeArray[i] = dt.datetime[offset: offset+timeOffset[i]]
		for _, v := range timeArray[i] {
			if !IsNumber(v) {
				err = errors.New("date time is bad format string")
				return dt, err
			}
		}
		offset += timeOffset[i]
		length -= timeOffset[i]
	}
	
	if timeArray[0] == "0000" || timeArray[1] == "00" || timeArray[2] == "00" {
		err = errors.New("date time is bad format string")
		return dt, err
	}
	if timeArray[0] == "0001" && timeArray[1] == "01" && timeArray[2] == "01" {
		err = errors.New("date time is bad format string")
		return dt, err
	}
	
	dt.datetime = fmt.Sprintf("%v-%v-%v %v:%v:%v", timeArray[0], timeArray[1], timeArray[2], timeArray[3], timeArray[4], timeArray[5])

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