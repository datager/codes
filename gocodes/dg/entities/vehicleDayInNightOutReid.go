package entities

type VehicleDayInNightOutDetail struct {
	TaskID    string `xorm:"task_id  not null VARCHAR(1024) default ''"`
	Ts        int64  `xorm:"ts bigint not null default 0"`
	Reid      string `xorm:"reid" varchar ( 36 ) not null default ''`
	Vid       string `xorm:"vid" varchar ( 36 ) not null default ''`
	PlateText string `xorm:"plate_text" varchar ( 10 ) not null default ''`
	SharpDate int64  `xorm:"sharp_date" bigint not null default 0`
}

func (t *VehicleDayInNightOutDetail) TableName() string {
	return TableNameVehicleDayInNightOutDetail
}
