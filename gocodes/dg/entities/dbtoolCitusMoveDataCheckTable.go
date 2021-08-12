package entities

import "time"

const DBToolCitusMoveDataCheckTableName = "dbcheck_citus_move_data_check_table"

type DBToolCitusMoveDataCheckTable struct {
	TableToStatistics string
	DayStartTsHuman   time.Time
	DayStartTs        int64

	FirstStatisticsTime        time.Time
	FirstCountOfDay            int64
	FirstPkOfMinTsOfDayStartTs string

	SecondCountOfDay            int64
	SecondPkOfMinTsOfDayStartTs string

	IsEqualBetweenPreviousAndCurrent int64
}

func (d *DBToolCitusMoveDataCheckTable) TableName() string {
	return DBToolCitusMoveDataCheckTableName
}
