package entities

import "time"

type VehicleCollisionSummary struct {
	Uts              time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts               int64     `xorm:"ts not null BIGINT"`
	TaskID           string    `xorm:"task_id not null default '''::character varying' VARCHAR(1024)"`
	SensorName       string    `xorm:"sensor_name not null default '''::character varying' VARCHAR(1024)"`
	VehicleID        string    `xorm:"vehicle_id not null default '''::character varying' VARCHAR(36)"`
	VehicleReid      string    `xorm:"vehicle_reid not null default '''::character varying' VARCHAR(36)"`
	VehicleVid       string    `xorm:"vehicle_vid not null default '''::character varying' VARCHAR(36)"`
	PlateText        string    `xorm:"plate_text not null default '''::character varying' VARCHAR(10)"`
	SensorID         string    `xorm:"sensor_id not null default '''::character varying' VARCHAR(36)"`
	ImageID          string    `xorm:"image_id not null default '''::character varying' VARCHAR(36)"`
	ImageURL         string    `xorm:"image_uri default '''::text' TEXT"`
	CutboardImageURI string    `xorm:"cutboard_image_uri default '''::text' TEXT"`
	MinReidNum       int64     `xorm:"min_reid_num not null default 0"`
}

func (VehicleCollisionSummary) TableName() string {
	return TableNameVehicleCollisionSummary
}
