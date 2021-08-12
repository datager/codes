package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

// ------------------------------------------
// VehicleFrequencyDetail
// ------------------------------------------
// 车辆频率分析详情
type VehicleFrequencyDetail struct {
	Uts              time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts               int64     `xorm:"ts not null default 0 BIGINT"`
	TaskID           string    `xorm:"task_id not null VARCHAR(1024)"`
	SensorName       string    `xorm:"sensor_name not null VARCHAR(1024)"`
	MaxTsTag         int64     `xorm:"max_ts_tag"`
	ReIDCount        int64     `xorm:"reid_count bigint not null default 0"`
	SensorID         string    `xorm:"sensor_id"`
	VehicleID        string    `xorm:"vehicle_id"`
	VehicleReID      string    `xorm:"vehicle_reid"`
	CutboardImageURI string    `xorm:"cutboard_image_uri"`
	PlateText        string    `xorm:"plate_text"`
	VehicleVID       string    `xorm:"vehicle_vid"`
}

func (vfd *VehicleFrequencyDetail) TableName() string {
	return TableNameVehicleFrequencyDetail
}

func (vfd *VehicleFrequencyDetail) ToModel() *models.VehicleFrequencyDetail {
	return &models.VehicleFrequencyDetail{
		Uts:    vfd.Uts,
		Ts:     vfd.Ts,
		TaskID: vfd.TaskID,
		//MaxTsTag:  vfd.MaxTsTag,
		//ReIDCount: vfd.ReIDCount,

		// 以下来自vehicle_capture_index表
		VehicleCapture: models.VehicleCapture{
			Timestamp:   vfd.Ts,
			SensorID:    vfd.SensorID,
			SensorName:  vfd.SensorName,
			VehicleID:   vfd.VehicleID,
			VehicleReID: vfd.VehicleReID,
			VehicleVID:  vfd.VehicleVID,
			ImageResult: &models.CutboardImage{
				CutboardImageURI: vfd.CutboardImageURI,
			},
			PlateText: vfd.PlateText,
		},
	}
}

func VehicleFrequencyDetailEntitiesToModels(list []*VehicleFrequencyDetail) []*models.VehicleFrequencyDetail {
	model := make([]*models.VehicleFrequencyDetail, 0)
	for _, l := range list {
		model = append(model, l.ToModel())
	}
	return model
}

// ------------------------------------------
// VehicleFrequencyDetailWithReIDCountExtend
// ------------------------------------------
type VehicleFrequencyDetailWithReIDCountExtend struct {
	VehicleFrequencyDetail `xorm:"extends"`
	ReIDCount              int64 `xorm:"reid_count"` // group by 的自定义命名
}

func (e *VehicleFrequencyDetailWithReIDCountExtend) TableName() string {
	return e.VehicleFrequencyDetail.TableName()
}

func (e *VehicleFrequencyDetailWithReIDCountExtend) ToModel() *models.VehicleFrequencyDetailWithReIDCountExtend {
	return &models.VehicleFrequencyDetailWithReIDCountExtend{
		LatestTsOfCaptureVID: e.VehicleFrequencyDetail.ToModel(),
		ReIDCount:            e.ReIDCount,
	}
}

func VehicleFrequencyDetailWithReIDCountExtendEntitiesToModels(entities []*VehicleFrequencyDetailWithReIDCountExtend, sensorMap map[string]*Sensors) []*models.VehicleFrequencyDetailWithReIDCountExtend {
	rets := make([]*models.VehicleFrequencyDetailWithReIDCountExtend, 0)
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
