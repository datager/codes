package entities

import "time"

type VehicleCollisionDetail struct {
	Uts              time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts               int64     `xorm:"ts not null BIGINT"`
	TaskID           string    `xorm:"task_id not null default '''::character varying' VARCHAR(1024)"`
	SensorName       string    `xorm:"sensor_name not null default '''::character varying' VARCHAR(1024)"`
	VehicleID        string    `xorm:"vehicle_id not null default '''::character varying' VARCHAR(36)"`
	VehicleReid      string    `xorm:"vehicle_reid not null default '''::character varying' VARCHAR(36)"`
	VehicleVid       string    `xorm:"vehicle_vid not null default '''::character varying' VARCHAR(36)"`
	SensorID         string    `xorm:"sensor_id not null default '''::character varying' VARCHAR(36)"`
	CutboardImageUri string    `xorm:"cutboard_image_uri default '''::text' TEXT"`
	TimespaceIDno    int64     `xorm:"timespace_idno smallint not null default 0"`
}

func (VehicleCollisionDetail) TableName() string {
	return TableNameVehicleCollisionDetail
}
