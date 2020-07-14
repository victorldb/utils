package utils

import (
	"errors"
	"time"
)

// AfterNextDayByHour 下一刻时间必须是大于今天最后一刻
func AfterNextDayByHour(hour int) (subTime time.Duration) {
	now := time.Now()
	return time.Duration(24+hour)*time.Hour -
		time.Duration((now.Hour()*3600+now.Minute()*60+now.Second())*1e9+now.Nanosecond())
}

// NextStartDay --
func NextStartDay(t time.Time) (nt time.Time) {
	year, month, day := t.Date()
	nt = time.Date(year, month, day, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1)
	return nt
}

// EqualDate --
func EqualDate(s, d time.Time) bool {
	syear, smonth, sday := s.Date()
	dyear, dmonth, dday := d.Date()
	return syear == dyear && smonth == dmonth && sday == dday
}

//ConventToLocalTime 转换成本地时间
func ConventToLocalTime(t time.Time) (localTime time.Time) {
	LocalTimeZone, LocalTimeOffset := time.Now().Zone()
	year, month, day := t.Date()
	if t.IsZero() || (year == 1 && month == 1 && day == 1) {
		return time.Time{}.Local()
	}

	//如果是本地时间就直接返回
	zoneName, offset := t.Zone()
	if (zoneName == LocalTimeZone && offset == LocalTimeOffset) || t.Location().String() == "Local" {
		return t
	}

	return time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
}

// NSToTime --
func NSToTime(ns int64) (time.Time, error) {
	if ns < 0 {
		return time.Time{}, errors.New("ns is err")
	}
	if ns == 0 {
		return time.Time{}, nil
	}
	sec := ns / 1e9
	nsec := ns - sec*1e9
	return time.Unix(sec, nsec), nil
}

// MinDayTime 获取这一天中最小的时间
func MinDayTime(t time.Time) (nt time.Time) {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
