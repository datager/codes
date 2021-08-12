package entities

type VehicleHiddenDetail struct {
	TaskID      string `xorm:"task_id  not null VARCHAR(1024) default ''"`
	Ts          int64  `xorm:"ts bigint not null default 0"`
	VehicleReid string `xorm:"vehicle_reid" varchar ( 36 ) not null default ''`
	VehicleVid  string `xorm:"vehicle_vid" varchar ( 36 ) not null default ''`
	PlateText   string `xorm:"plate_text" varchar ( 1024 ) not null default ''`
}

func (t *VehicleHiddenDetail) TableName() string {
	return TableNameVehicleHiddenDetail
}
