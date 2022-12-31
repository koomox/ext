package ext

import (
	"errors"
	"fmt"
	"time"
)

const (
	dateTimeFormat    = "2006-01-02"
	customTimeFormat  = "2006-01-02 15:04:05"
	frontTimeFormat   = "2006-01-02T15:04:05Z"
	browserTimeFormat = "2006/01/02 15:04"
)

type DateTime struct {
	Location *time.Location
	datetime string
	Time     time.Time
}

func NewDateTime() *DateTime {
	loc, _ := time.LoadLocation("UTC")
	return &DateTime{Location: loc, Time: time.Now()}
}

func FromDateTime(datetime string) (*DateTime, error) {
	if datetime == "" {
		return nil, errors.New("date time is null")
	}
	timeArray := []string{"0000", "00", "00", "00", "00", "00"}
	timeOffset := []int{4, 2, 2, 2, 2, 2}
	length := len(datetime)
	offset := 0
	for i := 0; i < len(timeArray); i++ {
		if length <= 0 {
			break
		}
		if !IsNumber(byte(datetime[offset])) {
			offset++
			length--
		}
		if length < timeOffset[i] {
			break
		}
		timeArray[i] = datetime[offset : offset+timeOffset[i]]
		for _, v := range timeArray[i] {
			if !IsNumber(byte(v)) {
				return nil, errors.New("date time is bad format string")
			}
		}
		offset += timeOffset[i]
		length -= timeOffset[i]
	}

	if timeArray[0] == "0000" || timeArray[1] == "00" || timeArray[2] == "00" {
		return nil, errors.New("date time is bad format string")
	}
	if timeArray[0] == "0001" && timeArray[1] == "01" && timeArray[2] == "01" {
		return nil, errors.New("date time is bad format string")
	}

	return &DateTime{datetime: fmt.Sprintf("%v-%v-%v %v:%v:%v", timeArray[0], timeArray[1], timeArray[2], timeArray[3], timeArray[4], timeArray[5])}, nil
}

func (dt *DateTime) UTC() *DateTime {
	dt.Location, _ = time.LoadLocation("UTC")
	return dt
}

func (dt *DateTime) CST() *DateTime {
	dt.Location, _ = time.LoadLocation("Asia/Shanghai")
	return dt
}

func (dt *DateTime) Parser() *DateTime {
	dt.Time, _ = time.ParseInLocation("2006-01-02 15:04:05", dt.datetime, dt.Location)
	return dt
}

func (dt *DateTime) Format(layout string) string {
	return dt.Time.In(dt.Location).Format(layout)
}

func (dt *DateTime) String() string {
	return dt.Time.In(dt.Location).Format("2006-01-02 15:04:05")
}

func (dt *DateTime) Date() string {
	return dt.Time.In(dt.Location).Format("2006-01-02")
}

func (dt *DateTime) Year() string {
	return dt.Time.In(dt.Location).Format("2006")
}

func (dt *DateTime) Before(u *DateTime) bool {
	return dt.Time.Before(u.Time)
}

func (dt *DateTime) Sub(u *DateTime) time.Duration {
	return dt.Time.Sub(u.Time)
}

func NowNanosecond() int64 {
	return time.Now().UnixNano()
}

func DelayMillisecond(nanosecond int64) int {
	return int((NowNanosecond() - nanosecond) / 1e6)
}

func Timestamp() int64 {
	return time.Now().Unix()
}
