package entities

import "codes/gocodes/dg/models"

type VehicleFirstAppearanceDetail struct {
	Ts          int64  `xorm:"ts"`
	TaskID      string `xorm:"task_id"`
	VehicleReID string `xorm:"vehicle_reid"`
}

func (*VehicleFirstAppearanceDetail) TableName() string {
	return TableNameVehicleFirstAppearanceDetail
}

func (s *VehicleFirstAppearanceDetail) ToModel(miniSensor *models.MiniSensor, vehicleCaptureIndexEntity *VehiclesIndex) *models.MiniVehicleWithMiniSensor {
	return &models.MiniVehicleWithMiniSensor{
		MiniVehicle: *vehicleCaptureIndexEntity.ToMiniVehicle(),
		MiniSensor:  *miniSensor,
	}
}
