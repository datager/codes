package utils

import (
	"fmt"
	"time"
)

const (
	DEFAULT_TIME_FORMAT = "2006-01-02 15:04:05.000"
)

func FormatPercent(val float32) string {
	return fmt.Sprintf("%.2f%%", val*100)
}

func FormatTimeDefault(t time.Time) string {
	return t.Format(DEFAULT_TIME_FORMAT)
}

func FormatTsDefault(ts int64, utcOffset int) string {
	t := TsToTime(ts)
	return FormatTimeDefault(UTCOffset(t, utcOffset))
}
