package wolf

import (
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils"
	"fmt"
	"github.com/golang/glog"
	"math/rand"
	"time"
)

var (
	statCapHourTables  []string
	statCapDayTables   []string
	statWarnHourTables []string
	statWarnDayTables  []string
)

func init() {
	statCapHourTables = []string{
		entities.TableNameFaceCaptureHourCount,
		entities.TableNameVehicleCaptureHourCount,
		entities.TableNameNonmotorCaptureHourCount,
		entities.TableNamePedestrianCaptureHourCount,
	}
	statCapDayTables = []string{
		entities.TableNameFaceCaptureDayCount,
		entities.TableNameVehicleCaptureDayCount,
		entities.TableNameNonmotorCaptureDayCount,
		entities.TableNamePedestrianCaptureDayCount,
	}

	statWarnHourTables = []string{
		entities.TableNameFaceWarnHourCount,
		entities.TableNameVehicleWarnHourCount,
	}
	statWarnDayTables = []string{
		entities.TableNameFaceWarnHourCount,
		entities.TableNameVehicleWarnHourCount,
	}
}

func InsertStatMockData() {
	// 建一个365天的数组, 365项, 每项是当天0点
	// 目标ts: 2019-06-05 ~ 2020-06-20
	for ts := 1559664000000; ts < 1592582400000; ts += utils.OneHourMillisecond {
		InsertStatHourTable(int64(ts))
	}
	for ts := 1559664000000; ts < 1592582400000; ts += utils.OneDayMillisecond {
		InsertStatDayTable(int64(ts))
	}

}

func InsertStatHourTable(ts int64) {
	for _, tableName := range statCapHourTables {
		glog.Infof("ts:%v, %v", utils.FormatTsByTsLOCAL(ts), tableName)
		sql := fmt.Sprintf("insert into %v (\"ts\", \"org_code\", \"sensor_int_id\", \"sensor_id\", \"count\") values (%v, 0, 10052, '9041c50f-7131-449e-a79d-68534ac758df', %v)", tableName, ts, pickMockHourCount())
		_, err := dbengine.GetStatisticsInstance().Exec(sql)
		if err != nil {
			glog.Warning(err)
		}
	}

	for _, tableName := range statWarnHourTables {
		glog.Infof("ts:%v, %v", utils.FormatTsByTsLOCAL(ts), tableName)
		sql := fmt.Sprintf("insert into %v (\"ts\", \"org_code\", \"sensor_int_id\", \"sensor_id\", \"repo_id\", \"count\") values (%v, 0, 10052, '9041c50f-7131-449e-a79d-68534ac758df', '0001', %v)", tableName, ts, pickMockHourCount())
		_, err := dbengine.GetStatisticsInstance().Exec(sql)
		if err != nil {
			glog.Warning(err)
		}
	}
}

func InsertStatDayTable(ts int64) {
	for _, tableName := range statCapDayTables {
		glog.Infof("ts:%v, %v", utils.FormatTsByTsLOCAL(ts), tableName)
		sql := fmt.Sprintf("insert into %v (\"ts\", \"org_code\", \"sensor_int_id\", \"sensor_id\", \"count\") values (%v, 0, 10052, '9041c50f-7131-449e-a79d-68534ac758df', %v)", tableName, ts, pickMockDayCount())
		_, err := dbengine.GetStatisticsInstance().Exec(sql)
		if err != nil {
			glog.Warning(err)
		}
	}

	for _, tableName := range statWarnDayTables {
		glog.Infof("ts:%v, %v", utils.FormatTsByTsLOCAL(ts), tableName)
		sql := fmt.Sprintf("insert into %v (\"ts\", \"org_code\", \"sensor_int_id\", \"sensor_id\", \"repo_id\", \"count\") values (%v, 0, 10052, '9041c50f-7131-449e-a79d-68534ac758df', '0001', %v)", tableName, ts, pickMockDayCount())
		_, err := dbengine.GetStatisticsInstance().Exec(sql)
		if err != nil {
			glog.Warning(err)
		}
	}
}

func pickMockHourCount() int64 {
	cnts := []int64{243, 633, 432, 245, 766, 435, 345, 457, 538, 396}

	rand.Seed(time.Now().Unix())
	return 5 * cnts[rand.Intn(10)]
}

func pickMockDayCount() int64 {
	return 12 * pickMockHourCount()
}
