package utils

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

const (
	OneDayMillisecond  = 86400 * 1000
	OneHourMillisecond = 3600 * 1000
)

func TimeToTs(t time.Time) int64 {
	return t.UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
}

func TsToTime(ts int64) time.Time {
	nsec := ts % 1000 * int64(time.Microsecond)
	return time.Unix(ts/1000, nsec)
}

func GetNowTs() int64 {
	t := time.Now()
	return TimeToTs(t)
}

func GetHourStartTs() int64 {
	oneHour := int64(3600 * 1000)
	nowTs := GetNowTs()
	return nowTs - nowTs%oneHour
}

func GetDayStartTs(ts int64) int64 {
	oneHour := int64(3600 * 1000)
	return ts - ts%oneHour
}

func GetNextDayStartTs(ts int64) int64 {
	oneHour := int64(3600 * 1000)
	oneDay := int64(86400 * 1000)
	return (ts - ts%oneHour) + oneDay
}

func UTCOffset(t time.Time, offset int) time.Time {
	return t.UTC().Add(time.Duration(offset) * time.Minute)
}

func SecondsToTime(seconds int) string {
	hour := seconds / 3600
	minute := seconds % 3600

	var strHour, strMinute string
	if hour < 10 {
		strHour = fmt.Sprintf("0%v", hour)
	} else {
		strHour = fmt.Sprintf("%v", hour)
	}

	minute /= 60
	if minute < 10 {
		strMinute = fmt.Sprintf("0%v", minute)
	} else {
		strMinute = fmt.Sprintf("%v", minute)
	}

	return fmt.Sprintf("%v:%v", strHour, strMinute)
}

// SplitTimeZone -- 按照时间间隔切分24小时
func SplitTimeZone(span int) [][]int {
	begin := 0
	end := 24 * 60 * 60

	ret := make([][]int, 0)
	for {
		if begin >= end {
			break
		}

		d := []int{begin, begin + span}
		ret = append(ret, d)

		begin += span
	}

	return ret
}

func FindPositionInTimeZone(timeZone [][]int, timestamp int64) int {
	t := TsToTime(timestamp)
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	duration := (hour*60+minute)*60 + second

	for i := 0; i < len(timeZone); i++ {
		start := timeZone[i][0]
		end := timeZone[i][1]

		if duration >= start && duration <= end {
			return i
		}
	}

	return -1
}

func DayToMillisecond(day int64) int64 {
	return day * OneDayMillisecond
}

func HourToMillisecond(hour int64) int64 {
	return hour * OneHourMillisecond
}

func WeekStart(timestamp int64) string {
	t := TsToTime(timestamp)
	year, month, day := t.Date()
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(year, month, day, 0, 0, 0, 0, t.Location()).AddDate(0, 0, offset)

	dateString := weekStart.Format("20060102")

	return dateString
}

func NewCarbon(ts int64) *carbon.Carbon {
	return carbon.NewCarbon(TsToTime(ts))
}

// 2019-11-24 16:24:09 => 2019-11-24 16:00:00
func TruncateToFloorHour(tm time.Time) time.Time {
	return tm.Truncate(time.Hour * 1)
}

func TruncateToFloorHourByTs(tm int64) int64 {
	return TimeToTs(TruncateToFloorHour(TsToTime(tm)))
}

// 2019-11-24 16:24:09 => 2019-11-24 17:00:00
func TruncateToCeilHour(tm time.Time) time.Time {
	return tm.Add(time.Hour).Truncate(time.Hour * 1)
}

func TruncateToCeilHourByTs(tm int64) int64 {
	return TimeToTs(TruncateToCeilHour(TsToTime(tm)))
}

// 2019/11/01-10:30 => 2019/11/01-00:00
func TruncateToFloorDayByTs(tm int64) int64 {
	return TimeToTs(NewCarbon(tm).StartOfDay().Time)
}

// 2019/11/01-10:30 => 2019/11/02-00:00
func TruncateToCeilDayByTs(tm int64) int64 {
	return TimeToTs(NewCarbon(tm).StartOfDay().AddDays(1).Time)
}

// 2006-01-02 15:04:05 => "15:00"
func FormatTsToHourStr(tm int64) string {
	return SubstrLength(TsToTime(tm).Format("2006-01-02 15:04:05"), 11, 5)
}

//TimeToZero 时间戳转零点
func TimeToZero(changeTime int64, updown int) int64 {

	t1 := changeTime / 1000
	t2 := changeTime - t1*1000
	tm := time.Unix(t1, t2*10^6)
	zeroTime := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
	backTime := zeroTime.AddDate(0, 0, updown).UnixNano() / 1e6
	if updown == 1 && backTime == (changeTime+(3600*1000*24)) {
		return changeTime
	}

	return backTime
}

func FormatTsByTs(ts int64) string {
	//return TsToTime(ts).UTC().Format(time.RFC3339)
	return fmt.Sprintf("{UTC} %v, {LOCAL} %v", TsToTime(ts).UTC(), TsToTime(ts))
}

func FormatTsByTsLOCAL(ts int64) string {
	//return TsToTime(ts).UTC().Format(time.RFC3339)
	return fmt.Sprintf("%v", TsToTime(ts))
}

const BasisTime = "2006-01-02 15:04:05"

func ParseLocalTimeStr2Ts(tsStr string) int64 {
	_tsTime, err := time.ParseInLocation(BasisTime, tsStr, time.Local)
	if err != nil {
		panic(fmt.Errorf("ParseLocalTs, %v", err.Error()))
	}
	return TimeToTs(_tsTime)
}

func GetMondayOfWeek(tm time.Time) time.Time {
	// 和周1差几天
	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}

	monday := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return monday
}
