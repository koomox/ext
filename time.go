package ext

import "time"

const (
	customTimeFormat = "2006-01-02 15:04:05"
)

func TimeNowUTC() (ts string, err error) {
	var loc *time.Location
	tn := time.Now()
	if loc, err = time.LoadLocation("UTC"); err != nil {
		return
	}
	return tn.In(loc).Format(customTimeFormat), nil
}

func TimeNowCST() (ts string, err error) {
	var loc *time.Location
	tn := time.Now()
	if loc, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		return
	}
	return tn.In(loc).Format(customTimeFormat), nil
}

func ParseTimeUTC(ts string) (tc time.Time, err error) {
	var loc *time.Location
	if loc, err = time.LoadLocation("UTC"); err != nil {
		return
	}
	if tc, err = time.ParseInLocation(customTimeFormat, ts, loc); err != nil {
		return
	}
	return
}

func ParseTimeCST(ts string) (tc time.Time, err error) {
	var loc *time.Location
	if loc, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		return
	}
	if tc, err = time.ParseInLocation(customTimeFormat, ts, loc); err != nil {
		return
	}
	return
}

func NowNanosecond() int64 {
	return time.Now().UnixNano()
}

func DelayMillisecond(nanosecond int64) int {
	return int((NowNanosecond() - nanosecond)/1e6)
}