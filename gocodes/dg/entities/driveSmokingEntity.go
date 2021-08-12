package entities

const DriveSmokingTableName = "driving_smoking_captures"

type DriveSmokingEntity struct {
	Timestamp int64  `xorm:"ts"`
	CaptureId string `xorm:"capture_id"`
	SensorId  string `xorm:"sensor_id"`
	PlateText string `xorm:"plate_text"`
}

func (DriveSmokingEntity) TableName() string {
	return DriveSmokingTableName
}
