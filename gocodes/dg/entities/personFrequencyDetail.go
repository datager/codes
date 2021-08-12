package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

// ------------------------------------------
// PersonFrequencyDetail
// ------------------------------------------
// 车辆频率分析详情
type PersonFrequencyDetail struct {
	Uts              time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts               int64     `xorm:"ts not null default 0 BIGINT"`
	TaskID           string    `xorm:"task_id not null VARCHAR(1024)"`
	SensorName       string    `xorm:"sensor_name not null VARCHAR(1024)"`
	MaxTsTag         int64     `xorm:"max_ts_tag int2 not null default 0"`
	ReIDCount        int64     `xorm:"reid_count bigint not null default 0"`
	SensorID         string    `xorm:"sensor_id"`
	FaceID           string    `xorm:"face_id"`
	FaceReID         string    `xorm:"face_reid"`
	FaceVID          string    `xorm:"face_vid"`
	CutboardImageURI string    `xorm:"cutboard_image_uri"`
}

func (vfd *PersonFrequencyDetail) TableName() string {
	return TableNamePersonFrequencyDetail
}

func (vfd *PersonFrequencyDetail) ToModel() *models.PersonFrequencyDetail {
	return &models.PersonFrequencyDetail{
		Uts:    vfd.Uts,
		Ts:     vfd.Ts,
		TaskID: vfd.TaskID,
		//MaxTsTag:  vfd.MaxTsTag,
		//ReIDCount: vfd.ReIDCount,

		// 以下来自faces_index表
		CapturedFace: models.CapturedFace{
			FaceID:     vfd.FaceID,
			FaceReID:   vfd.FaceReID,
			Timestamp:  vfd.Ts,
			SensorID:   vfd.SensorID,
			SensorName: vfd.SensorName,
			Longitude:  0.0,
			Latitude:   0.0,
			ImageResult: &models.CutboardImage{
				CutboardImageURI: vfd.CutboardImageURI,
			},
			FaceVid: vfd.FaceVID,
		},
	}
}

func PersonFrequencyDetailEntitiesToModels(list []*PersonFrequencyDetail) []*models.PersonFrequencyDetail {
	model := make([]*models.PersonFrequencyDetail, 0)
	for _, l := range list {
		model = append(model, l.ToModel())
	}
	return model
}

// ------------------------------------------
// PersonFrequencyDetailWithReIDCountExtend
// ------------------------------------------
type PersonFrequencyDetailWithReIDCountExtend struct {
	PersonFrequencyDetail `xorm:"extends"`
	ReIDCount             int64 `xorm:"reid_count"` // group by 的自定义命名
}

func (e *PersonFrequencyDetailWithReIDCountExtend) TableName() string {
	return e.PersonFrequencyDetail.TableName()
}

func (e *PersonFrequencyDetailWithReIDCountExtend) ToModel() *models.PersonFrequencyDetailWithReIDCountExtend {
	return &models.PersonFrequencyDetailWithReIDCountExtend{
		LatestTsOfCaptureVID: e.PersonFrequencyDetail.ToModel(),
		ReIDCount:            e.ReIDCount,
	}
}

func PersonFrequencyDetailWithReIDCountExtendEntitiesToModels(entities []*PersonFrequencyDetailWithReIDCountExtend, sensorMap map[string]*Sensors) []*models.PersonFrequencyDetailWithReIDCountExtend {
	rets := make([]*models.PersonFrequencyDetailWithReIDCountExtend, 0)
	for _, en := range entities {
		if sensor, ok := sensorMap[en.SensorID]; ok {
			rModel := en.ToModel()
			rModel.LatestTsOfCaptureVID.SensorName = sensor.SensorName
			rModel.LatestTsOfCaptureVID.Longitude = sensor.Longitude
			rModel.LatestTsOfCaptureVID.Latitude = sensor.Latitude
			rModel.ReIDCount = en.ReIDCount // todo: 整合VehicleFrequencyDetailWithReIDCountExtend 与 VehicleFrequencyDetail 的 model

			rets = append(rets, rModel)
		} else {
			rets = append(rets, nil)
		}
	}
	return rets
}

/*
func PersonFrequencyDetailWithReIDCountExtendEntitiesToModels(entities []*PersonFrequencyDetailWithReIDCountExtend) []*models.PersonFrequencyDetailWithReIDCountExtend {
	models := make([]*models.PersonFrequencyDetailWithReIDCountExtend, 0)
	for _, r := range entities {
		rmodel := r.ToModel()
		rmodel.ReIDCount = r.ReIDCount // todo: 整合VehicleFrequencyDetailWithReIDCountExtend 与 VehicleFrequencyDetail 的 model

		models = append(models, rmodel)
	}
	return models
}
*/
