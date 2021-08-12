package wolf

import (
	"fmt"
	"strings"
	"sync"

	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils"

	"github.com/golang/glog"
	"xorm.io/xorm"
)

// 从db1灌数据到db2(可以跨机器), 一般用于从142灌到47
func QueryDataFromDB1AndInsertIntoDB2(sourceDBStr, targetDBStr string, tableNames []string, startTs, endTs int64) {
	sourceDB := dbengine.CreateEngine(sourceDBStr)
	targetDB := dbengine.CreateEngine(targetDBStr)

	tsStep := int64(100 * 1000) // 单位millSeconds
	threadChan := make(chan struct{}, 1000)
	wg := sync.WaitGroup{}

	for ts := startTs; ts < endTs; ts += tsStep {
		for _, tableName := range tableNames {

			// 捕获一下, 其他的变量都不会变
			ts := ts
			tableName := tableName
			threadChan <- struct{}{}
			wg.Add(1)
			go func() {
				defer func() {
					<-threadChan
					wg.Done()
				}()
				switch tableName {
				case entities.TableNameFaces:
					queryInsertFaces(sourceDB, targetDB, ts, ts+tsStep)
				case entities.TableNameFacesIndex:
					queryInsertFacesIndex(sourceDB, targetDB, ts, ts+tsStep)
				case entities.TableNameVehicleCaptureEntity:
					queryInsertVehicle(sourceDB, targetDB, ts, ts+tsStep)
				case entities.TableNameVehicleCaptureIndexEntity:
					queryInsertVehicleIndex(sourceDB, targetDB, ts, ts+tsStep)
				case entities.TableNameNonmotorCapture:
					queryInsertNonmotor(sourceDB, targetDB, ts, ts+tsStep)
				case entities.TableNameNonmotorIndexCapture:
					queryInsertNonmotorIndex(sourceDB, targetDB, ts, ts+tsStep)
				case entities.TableNamePedestrianCapture:
					queryInsertPedestrian(sourceDB, targetDB, ts, ts+tsStep)
				case entities.TableNamePedestrianIndexCapture:
					queryInsertPedestrianIndex(sourceDB, targetDB, ts, ts+tsStep)
				default:
					panic("default table name?")
				}
			}()
		}
	}
	wg.Wait()
	glog.Infof("QueryDataFromDB1AndInsertIntoDB2 success!!!")
}

func queryInsertFaces(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNameFaces

	data := make([]*entities.Faces, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}

func queryInsertFacesIndex(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNameFacesIndex

	data := make([]*entities.FacesIndex, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}

func queryInsertVehicle(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNameVehicleCaptureEntity

	data := make([]*entities.Vehicles, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}

func queryInsertVehicleIndex(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNameVehicleCaptureIndexEntity
	data := make([]*entities.VehicleCaptureIndex, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}

func queryInsertNonmotor(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNameNonmotorCapture

	data := make([]*entities.Nonmotor, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}

func queryInsertNonmotorIndex(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNameNonmotorIndexCapture

	data := make([]*entities.NonmotorIndex, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}

func queryInsertPedestrian(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNamePedestrianCapture

	data := make([]*entities.Pedestrian, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}

func queryInsertPedestrianIndex(sourceDB, targetDB *xorm.Engine, st, et int64) {
	tableName := entities.TableNamePedestrianIndexCapture

	data := make([]*entities.PedestrianIndex, 0)
	sourceSession := sourceDB.NewSession()
	defer sourceSession.Clone()
	err := sourceSession.SQL(fmt.Sprintf("select * from %v where ts >= %v and ts < %v", tableName, st, et)).Find(&data)
	if err != nil {
		glog.Error(err)
	}

	if len(data) == 0 {
		glog.V(11).Infof("skip   插入跳过%v条, %v表, 从%v到%v", len(data), tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
		return
	}

	targetSession := targetDB.NewSession()
	defer targetSession.Clone()
	ret, err := targetDB.Table(tableName).InsertMulti(&data)
	glog.Infof("success插入成功%v条, %v表, 从%v到%v", ret, tableName, utils.FormatTsByTsLOCAL(st), utils.FormatTsByTsLOCAL(et))
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			glog.Error(err)
		}
	}
}
