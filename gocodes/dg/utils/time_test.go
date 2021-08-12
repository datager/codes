package utils

import (
	"testing"
)

func TestTruncateDay(t *testing.T) {
	floor := TruncateToFloorDayByTs(GetNowTs())
	ceil := TruncateToCeilDayByTs(GetNowTs())
	xx := TsToTime(floor).Format("2006-01-02 15:04:05")
	yy := TsToTime(ceil).Format("2006-01-02 15:04:05")
	_ = xx
	_ = yy
}
