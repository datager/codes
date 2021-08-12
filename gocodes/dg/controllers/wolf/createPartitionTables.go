package wolf

import (
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils"
	"fmt"
	"github.com/golang/glog"
	"time"
)

var captureTableNames []string

func init() {
	captureTableNames = []string{
		entities.TableNameVehicleCaptureIndexEntity,
		entities.TableNameVehicleCaptureEntity,
		entities.TableNamePedestrianIndexCapture,
		entities.TableNamePedestrianCapture,
		entities.TableNameNonmotorIndexCapture,
		entities.TableNameNonmotorIndexCapture,
		entities.TableNameFacesIndex,
		entities.TableNameFaces,
	}
}

func getDay(array []int64, index int) int64 {
	i := index % len(array)
	if i >= len(array) {
		i = len(array) - 1
	}
	return array[i]
}

//select to_timestamp(min(ts/1000)), to_timestamp(max(ts/1000)) from vehicle_capture_index;
//select to_timestamp(min(ts/1000)), to_timestamp(max(ts/1000)) from faces_index;
// 2020-06-05 12:19:52+08 1591330792000
// 2020-06-05 22:41:40+08 1591368100000
// 功能把1天均匀分布到365天, 每30s更新到某一天
func ChangeCaptureTs() {
	// 建一个365天的数组, 365项, 每项是当天0点
	// 目标ts: 2019-06-05 ~ 2020-06-05
	targetDays := make([]int64, 0)
	for ts := 1559664000000; ts < 1591368100000; ts += utils.OneDayMillisecond {
		targetDays = append(targetDays, int64(ts))
	}

	//for ts := 1574006400000 - 7*utils.OneDayMillisecond; ts < 1591368100000+14*utils.OneDayMillisecond; ts += 7 * utils.OneDayMillisecond {
	//	CreatePartitionTable(int64(ts))
	//}

	// 抓拍库里的ts, 分布2020-06-05 12:19:52+08 ~ 2020-06-05 22:41:40+08
	var start int64 = 1591330792000
	var end int64 = 1591368100000
	var step int64 = 30 * 1000 // 30s
	i := 0
	for srcTs := start; srcTs < end; srcTs += step {
		i++

		targetDay := getDay(targetDays, i)
		updateCaptureTableTs(srcTs, targetDay)
	}
}

func updateCaptureTableTs(srcTs, targetDayTs int64) {

	for _, tableName := range captureTableNames {
		glog.Infof("ts:%v to %v, %v", utils.FormatTsByTsLOCAL(srcTs), utils.FormatTsByTsLOCAL(targetDayTs), tableName)
		var step int64 = 30 * 1000 // 30s
		sql := fmt.Sprintf("update %v set ts = %v where (ts >= %v and ts < %v)", tableName, targetDayTs, srcTs, srcTs+step)

		_, err := dbengine.GetV5Instance().Exec(sql)
		if err != nil {
			glog.Warning(err)
		}
	}
}

func CreatePartitionTable(ts int64) {
	faceSQL := "create table faces_%v partition of faces for values from (%v) to (%v);"
	indexSQL := "create table findex_%v partition of faces_index for values from (%v) to (%v)"
	whiteSQL := "create table white_warn_%v partition of white_warn for values from (%v) to (%v);"
	whiteIndexSQL := "create table white_warn_index_%v partition of white_warn_index for values from (%v) to (%v);"

	pedestrianSQL := "create table pedestrian_%v partition of pedestrian_capture for values from (%v) to (%v);"
	pedestrianIndexSQL := "create table pindex_%v partition of pedestrian_capture_index for values from (%v) to (%v);"
	nonmotorSQL := "create table nonmotor_%v partition of nonmotor_capture for values from (%v) to (%v);"
	nonmotorIndexSQL := "create table nindex_%v partition of nonmotor_capture_index for values from (%v) to (%v);"
	vehicleSQL := "create table vehicle_%v partition of vehicle_capture for values from (%v) to (%v);"
	vehicleIndexSQL := "create table vindex_%v partition of vehicle_capture_index for values from (%v) to (%v);"

	//now := time.Now()
	now := utils.TsToTime(ts)
	year, month, day := now.Date()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(year, month, day, 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	nextWeekStart := weekStart.AddDate(0, 0, 7)
	nextWeekEnd := nextWeekStart.AddDate(0, 0, 7)

	// end := nextMonth.AddDate(0, 1, 0)

	startTs := utils.TimeToTs(nextWeekStart)
	endTs := utils.TimeToTs(nextWeekEnd)

	date := nextWeekStart.Format("20060102")

	faceSQL = fmt.Sprintf(faceSQL, date, startTs, endTs)
	indexSQL = fmt.Sprintf(indexSQL, date, startTs, endTs)
	whiteSQL = fmt.Sprintf(whiteSQL, date, startTs, endTs)
	whiteIndexSQL = fmt.Sprintf(whiteIndexSQL, date, startTs, endTs)

	pedestrianSQL = fmt.Sprintf(pedestrianSQL, date, startTs, endTs)
	pedestrianIndexSQL = fmt.Sprintf(pedestrianIndexSQL, date, startTs, endTs)
	nonmotorSQL = fmt.Sprintf(nonmotorSQL, date, startTs, endTs)
	nonmotorIndexSQL = fmt.Sprintf(nonmotorIndexSQL, date, startTs, endTs)
	vehicleIndexSQL = fmt.Sprintf(vehicleIndexSQL, date, startTs, endTs)
	vehicleSQL = fmt.Sprintf(vehicleSQL, date, startTs, endTs)

	_, err := dbengine.GetV5Instance().Exec(faceSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(indexSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(pedestrianSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(pedestrianIndexSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(nonmotorSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(nonmotorIndexSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(vehicleIndexSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(vehicleSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(whiteSQL)
	if err != nil {
		glog.Warning(err)
	}

	_, err = dbengine.GetV5Instance().Exec(whiteIndexSQL)
	if err != nil {
		glog.Warning(err)
	}
}
