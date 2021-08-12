package wolf

import (
	"fmt"
	"strings"
	"time"

	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/utils"

	"github.com/golang/glog"
)

// 用于一次性创建 一段时间 内的分区表(按战狼loki的命名规则, 每周1个)
// startTs: 2020-03-19 09:00:01+08
// endTs: 2020-09-19 09:00:01+08
// 左闭右开
func AutoCreatePartitionTable(startTs, endTs int64, targetDBStr string) {
	targetDB := dbengine.CreateEngine(targetDBStr)

	// sql
	faceSQL := "create table faces_%v partition of faces for values from (%v) to (%v);"
	faceIndexSQL := "create table findex_%v partition of faces_index for values from (%v) to (%v)"
	//facesStrangerSQL := "create table faces_stranger_%v partition of faces_stranger for values from (%v) to (%v);"
	//facesStrangerIndexSQL := "create table faces_stranger_index_%v partition of faces_stranger_index for values from (%v) to (%v);"
	pedestrianSQL := "create table pedestrian_%v partition of pedestrian_capture for values from (%v) to (%v);"
	pedestrianIndexSQL := "create table pindex_%v partition of pedestrian_capture_index for values from (%v) to (%v);"
	nonmotorSQL := "create table nonmotor_%v partition of nonmotor_capture for values from (%v) to (%v);"
	nonmotorIndexSQL := "create table nindex_%v partition of nonmotor_capture_index for values from (%v) to (%v);"
	vehicleSQL := "create table vehicle_%v partition of vehicle_capture for values from (%v) to (%v);"
	vehicleIndexSQL := "create table vindex_%v partition of vehicle_capture_index for values from (%v) to (%v);"
	sqls := make([]string, 0)
	sqls = append(sqls, faceSQL, faceIndexSQL, pedestrianSQL, pedestrianIndexSQL, nonmotorSQL, nonmotorIndexSQL, vehicleSQL, vehicleIndexSQL)

	// 找到每周一
	startTime, endTime := utils.TsToTime(startTs), utils.TsToTime(endTs)
	startMonday, endMonday := utils.GetMondayOfWeek(startTime), utils.GetMondayOfWeek(endTime) // 获取start周1, end周1
	for tm := startMonday; tm.Before(endMonday); tm = tm.Add(time.Hour * 24 * 7) {
		partitionNameStr := tm.Format("20060102")

		thisWeekMondayTs := utils.TimeToTs(tm)
		nextWeekMondayTs := utils.TimeToTs(tm.Add(time.Hour * 24 * 7))

		glog.Infof("start建立%v的分区: 从%v到%v", partitionNameStr, thisWeekMondayTs, nextWeekMondayTs)
		for _, sql := range sqls {
			sql = fmt.Sprintf(sql, partitionNameStr, thisWeekMondayTs, nextWeekMondayTs)
			_, err := targetDB.Exec(sql)
			if err != nil {
				errMsg := err.Error()
				if !strings.Contains(errMsg, "already exists") {
					glog.Fatal(err)
				}
			}
		}
		glog.Infof("done建立分区")
	}

	if err := targetDB.Close(); err != nil {
		glog.Warning(err)
	}
}
