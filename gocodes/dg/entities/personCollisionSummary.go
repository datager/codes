package entities

import "time"

type PersonCollisionSummary struct {
	Uts              time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts               int64     `xorm:"ts not null BIGINT"`
	TaskID           string    `xorm:"task_id not null default '''::character varying' VARCHAR(1024)"`
	SensorName       string    `xorm:"sensor_name not null default '''::character varying' VARCHAR(1024)"`
	FaceID           string    `xorm:"face_id not null default '''::character varying' VARCHAR(36)"`
	FaceReid         string    `xorm:"face_reid not null default '''::character varying' VARCHAR(36)"`
	FaceVid          string    `xorm:"face_vid not null default '''::character varying' VARCHAR(36)"`
	SensorID         string    `xorm:"sensor_id not null default '''::character varying' VARCHAR(36)"`
	ImageID          string    `xorm:"image_id not null default '''::character varying' VARCHAR(36)"`
	ImageURL         string    `xorm:"image_uri default '''::text' TEXT"`
	CutboardImageURI string    `xorm:"cutboard_image_uri default '''::text' TEXT"`
	MinReidNum       int64     `xorm:"min_reid_num not null default 0"`
}

func (PersonCollisionSummary) TableName() string {
	return TableNamePersonCollisionSummary
}
