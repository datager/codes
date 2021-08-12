package entities

import "codes/gocodes/dg/models"

type VehicleFakePlateSummary struct {
	Ts        int64  `xorm:"ts"`
	TaskID    string `xorm:"task_id"`
	PlateText string `xorm:"plate_text"`
	Num       int64  `xorm:"num"`
}

func (*VehicleFakePlateSummary) TableName() string {
	return TableNameVehicleFakePlateSummary
}

func (s *VehicleFakePlateSummary) ToModel() *models.VehicleFakePlateSummaryItem {
	return &models.VehicleFakePlateSummaryItem{
		Timestamp: s.Ts,
		PlateText: s.PlateText,
		Num:       s.Num,
	}
}
