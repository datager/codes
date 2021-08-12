package entities

import "codes/gocodes/dg/models"

type VehicleFakePlateDetail struct {
	Ts        int64                   `xorm:"ts"`
	TaskID    string                  `xorm:"task_id"`
	PlateText string                  `xorm:"plate_text"`
	Vehicles  models.MiniVehicleSlice `xorm:"vehicles json"`
}

func (*VehicleFakePlateDetail) TableName() string {
	return TableNameVehicleFakePlateDetail
}
