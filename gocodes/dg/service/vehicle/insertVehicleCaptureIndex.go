package vehicle

import (
	"codes/gocodes/dg/configs"
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/loglevel"
	"codes/gocodes/dg/utils/random"
	"codes/gocodes/dg/utils/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/golang/glog"
	//"golang.org/x/sync/errgroup"
)

// captureCnt: 待写数量
func InsertVehicleCaptureIndexOfOneSensor(sID string, sIntID int32, orgCode int64) {
	cPool := make(chan struct{}, configs.GetInstance().CaptureConfig.Concurrency)
	var g errgroup.Group

	captureNumToInsertOfOneSensor := int(configs.GetInstance().CaptureConfig.CaptureNumToInsertOfOneSensor)
	numInsideOneInsertBatch := int(configs.GetInstance().CaptureConfig.NumInsideOneInsertBatch)

	batches := captureNumToInsertOfOneSensor / numInsideOneInsertBatch
	for bn := 0; bn < batches; bn++ { // bn: batchNum
		cPool <- struct{}{}
		g.Go(func() error {
			defer func() {
				<-cPool
			}()

			captures := make([]*entities.VehiclesIndex, 0)
			for i := 0; i < numInsideOneInsertBatch; i++ {
				random.RandSeed()
				c := &entities.VehiclesIndex{
					VehicleID:        uuid.NewV4().String(),
					VehicleReID:      uuid.NewV4().String(),
					VehicleVid:       uuid.NewV4().String(),
					Ts:               random.RandInt64(configs.GetInstance().CaptureConfig.Ts.TsStart, configs.GetInstance().CaptureConfig.Ts.TsEnd-1),
					SensorID:         sID,
					SensorIntID:      sIntID,
					OrgCode:          orgCode,
					Speed:            random.RandInt(0, 3), // 范围详见dg.model/common.proto
					Direction:        random.RandInt(0, 8), // 范围详见dg.model/common.proto
					CutboardImageURI: configs.GetInstance().CaptureConfig.URL.VehicleURL,
					BrandID:          random.RandInt(0, 350), // 范围详见INDEX_BRAND.json
					SubBrandID:       random.RandInt(0, 200), // 范围详见INDEX_BRAND.json, 没找到暂按200算
					ModelYearID:      random.RandInt(0, 20),  // 范围详见INDEX_BRAND.json, 没找到暂按200算
					TypeID:           random.RandInt(0, 20),  // 应该是在数据流结构里, 222上是[0,20]区间 (select max(type_id), min(type_id) from vehicle_capture_index;)
					Side:             random.RandInt(0, 2),   // 同type_id, [0,2], select max(side), min(side) from vehicle_capture_index;
					ColorID:          random.RandInt(0, 20),  // 同type_id, [0,12], select max(color_id), min(color_id) from vehicle_capture_index;
					PlateText:        "京A111111",
					PlateTypeID:      random.RandInt(-1, 1000),           // 同type_id, [-1,1000], select max(plate_type_id), min(plate_type_id) from vehicle_capture_index;
					PlateColorID:     random.RandInt(0, 20),              // 同type_id, [0,12], select max(color_id), min(color_id) from vehicle_capture_index;
					Symbols:          random.RandIntArray(1, 200, 1, 20), // 根据random.RandIntArray(1,200, 1, 20)查的, 肯定比实际多
					Specials:         random.RandIntArray(1, 200, 1, 20), // 同上, 实际基本都只有一个'{99}'
					RelationTypes:    0,
				}
				captures = append(captures, c)
				//glog.Infof("%+v", c)
			}
			err := BatchInsertVehicleCaptureIndex(captures)
			if err != nil {
				glog.Error(loglevel.VehicleCapture, err)
			}
			return nil
		})
		if err := g.Wait(); err != nil {
			glog.Error(loglevel.VehicleCapture, err)
		}
	}
	//glog.Infof("sIntID:%v write captureNumToInsertOfOneSensor:%v done", sIntID, captureNumToInsertOfOneSensor)
}

// 因为并发所以要开session??? 需要查一下
func BatchInsertVehicleCaptureIndex(caps []*entities.VehiclesIndex) error {
	session := dbengine.GetV5Instance().NewSession()
	defer session.Close()

	l, err := session.Table(entities.TableNameVehicleCaptureIndexEntity).InsertMulti(&caps)
	if err != nil {
		return err
	}
	if l != int64(len(caps)) {
		glog.Error(loglevel.VehicleCapture, "insert l != int64(len(caps)")
	}
	return nil
}
