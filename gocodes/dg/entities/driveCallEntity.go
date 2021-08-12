package entities

const DriveCallTableName = "driving_calling_captures"

type DriveCallEntity struct {
	Timestamp int64  `xorm:"ts"`
	CaptureId string `xorm:"capture_id"`
	SensorId  string `xorm:"sensor_id"`
	PlateText string `xorm:"plate_text"`
}

func (DriveCallEntity) TableName() string {
	return DriveCallTableName
}
