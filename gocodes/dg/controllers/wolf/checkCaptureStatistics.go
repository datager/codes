package wolf

import (
	"fmt"

	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils"

	"github.com/golang/glog"
	"xorm.io/xorm"
)

// 把统计结果写入dbcheck_citus_move_data_check_table表里
type StatisticsType string

const (
	FirstStatisticsType  StatisticsType = "FirstStatisticsType"
	SecondStatisticsType StatisticsType = "SecondStatisticsType"
)

// 用途:
// 扩容前, 调一下这个函数, 记录一下数据库情况
// 扩容后, 调一下这个函数, 记录一下数据库情况, 和扩容前记录的比较一下看是否相等
// 查每天的count数
// 查每天ts最小的face_id
func CheckIsDBStatisticsEqual(kind StatisticsType, dbStr string, tableNames []string, startTs, endTs int64) {
	dbInstance := dbengine.CreateEngine(dbStr)

	if kind == FirstStatisticsType {
		sql := "drop table if exists dbcheck_citus_move_data_check_table;\ncreate table dbcheck_citus_move_data_check_table\n(\n    table_to_statistics                   text      not null default '', -- 哪个表(8大抓拍表)\n    day_start_ts_human                    timestamp not null,            -- 哪一天的ts(当天00:00)\n    day_start_ts                          bigint    not null default -1, -- 那一天的开始00:00\n\n    first_statistics_time                 timestamp not null,            -- 上一次: 统计的时间\n    first_count_of_day                    bigint    not null default -1, -- 上一次: 那一天有几个抓拍\n    first_pk_of_min_ts_of_day_start_ts    text      not null default '', -- 上一次: 那一天最小的ts对应的主键是什么(比如faces表就是07:30的那条face_id, faces_index表就是07:35的face_reid)\n\n    second_count_of_day                   bigint    not null default -1, -- 这一次: 那一天有几个抓拍\n    second_pk_of_min_ts_of_day_start_ts   text      not null default '', -- 这一次: 那一天最小的ts对应的主键是什么(比如faces表就是07:30的那条face_id, faces_index表就是07:35的face_reid)\n\n    is_equal_between_previous_and_current integer   not null default -1  -- 上一次统计和这一次统计的指标是不是都相等(1相等, 0不相等, 默认-1)\n);\ncreate unique index uidx_table_dayts on dbcheck_citus_move_data_check_table (table_to_statistics, day_start_ts);"
		if _, err := dbInstance.Exec(sql); err != nil {
			panic(err)
		}
	}

	oneDayStep := int64(24 * 60 * 60 * 1000) // 单位millSeconds

	for ts := startTs; ts < endTs; ts += oneDayStep {
		for _, tableName := range tableNames {
			switch tableName {
			case entities.TableNameFaces:
				statisticsForDBTool(kind, entities.TableNameFaces, "face_id", ts, ts+oneDayStep, dbInstance)
			case entities.TableNameFacesIndex:
				statisticsForDBTool(kind, entities.TableNameFacesIndex, "face_reid", ts, ts+oneDayStep, dbInstance)
			case entities.TableNameVehicleCaptureEntity:
				statisticsForDBTool(kind, entities.TableNameVehicleCaptureEntity, "vehicle_id", ts, ts+oneDayStep, dbInstance)
			case entities.TableNameVehicleCaptureIndexEntity:
				statisticsForDBTool(kind, entities.TableNameVehicleCaptureIndexEntity, "vehicle_reid", ts, ts+oneDayStep, dbInstance)
			case entities.TableNameNonmotorCapture:
				statisticsForDBTool(kind, entities.TableNameNonmotorCapture, "nonmotor_id", ts, ts+oneDayStep, dbInstance)
			case entities.TableNameNonmotorIndexCapture:
				statisticsForDBTool(kind, entities.TableNameNonmotorIndexCapture, "nonmotor_reid", ts, ts+oneDayStep, dbInstance)
			case entities.TableNamePedestrianCapture:
				statisticsForDBTool(kind, entities.TableNamePedestrianCapture, "pedestrian_id", ts, ts+oneDayStep, dbInstance)
			case entities.TableNamePedestrianIndexCapture:
				statisticsForDBTool(kind, entities.TableNamePedestrianIndexCapture, "pedestrian_reid", ts, ts+oneDayStep, dbInstance)
			case entities.TableNamePetrolVehicles:
				statisticsForDBTool(kind, entities.TableNamePetrolVehicles, "vehicle_id", ts, ts+oneDayStep, dbInstance)
			default:
				panic("default table name?")
			}
		}
	}
	glog.Info("查完了 success")
}

func statisticsForDBTool(kind StatisticsType, tableName, orderByColumnName string, dayStart, nextDayStart int64, dbInstance *xorm.Engine) {
	pkOfMinTsOfDayStartTs := ""
	sql := fmt.Sprintf("select %v from %v where ts >= %v and ts < %v order by ts, %v asc limit 1", orderByColumnName, tableName, dayStart, nextDayStart, orderByColumnName)
	gotPk, err := dbInstance.SQL(sql).Get(&pkOfMinTsOfDayStartTs)
	if err != nil {
		glog.Error(err, sql)
	}
	if !gotPk {
		glog.Warningf("!got: %v, 从%v到%v", sql, utils.FormatTsByTsLOCAL(dayStart), utils.FormatTsByTsLOCAL(nextDayStart))
		return
	}

	countOfDay := int64(0)
	if gotPk { // 如果没查到pk limit 1的话, 再去查count也肯定是0
		sql := fmt.Sprintf("select count(*) from %v where ts >= %v and ts < %v", tableName, dayStart, nextDayStart)
		gotCnt, err := dbInstance.SQL(sql).Get(&countOfDay)
		if err != nil {
			glog.Error(err, sql)
		}
		if !gotCnt {
			glog.Warningf("!got: %v, 从%v到%v", sql, utils.FormatTsByTsLOCAL(dayStart), utils.FormatTsByTsLOCAL(nextDayStart))
		}
	}
	glog.Infof("okok: %v, %v, 从%v到%v", pkOfMinTsOfDayStartTs, countOfDay, utils.FormatTsByTsLOCAL(dayStart), utils.FormatTsByTsLOCAL(nextDayStart))

	if kind == FirstStatisticsType {
		newStatisticsData := &entities.DBToolCitusMoveDataCheckTable{
			TableToStatistics:          tableName,
			DayStartTsHuman:            utils.TsToTime(dayStart),
			DayStartTs:                 dayStart,
			FirstStatisticsTime:        utils.TsToTime(utils.GetNowTs()),
			FirstCountOfDay:            countOfDay,
			FirstPkOfMinTsOfDayStartTs: pkOfMinTsOfDayStartTs,
			//SecondCountOfDay            int64
			//SecondPkOfMinTsOfDayStartTs string
			//
			//IsEqualBetweenPreviousAndCurrent int64
		}

		if _, err = dbInstance.Insert(newStatisticsData); err != nil {
			glog.Error(err)
		}
	} else if kind == SecondStatisticsType {
		oldStatisticsData := entities.DBToolCitusMoveDataCheckTable{}
		_, err := dbInstance.SQL(fmt.Sprintf("select * from %v where table_to_statistics = '%v' and day_start_ts = %v", entities.DBToolCitusMoveDataCheckTableName, tableName, dayStart)).Get(&oldStatisticsData)
		if err != nil {
			glog.Error(err)
		}

		isEqualBetweenPreviousAndCurrent := 0
		if (oldStatisticsData.FirstCountOfDay == countOfDay) && (oldStatisticsData.FirstPkOfMinTsOfDayStartTs == pkOfMinTsOfDayStartTs) {
			isEqualBetweenPreviousAndCurrent = 1
		} else {
			isEqualBetweenPreviousAndCurrent = -1
		}

		sql := fmt.Sprintf("update %v "+
			" set second_count_of_day = %v, "+
			" second_pk_of_min_ts_of_day_start_ts = '%v', "+
			" is_equal_between_previous_and_current = %v "+
			" where table_to_statistics = '%v' and day_start_ts = %v", entities.DBToolCitusMoveDataCheckTableName, countOfDay, pkOfMinTsOfDayStartTs, isEqualBetweenPreviousAndCurrent, tableName, dayStart)
		if _, err = dbInstance.Exec(sql); err != nil {
			glog.Error(err)
		}
	}
}
