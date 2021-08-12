package entities

const CoverFaceTableName = "cover_face_captures"

type CoverFaceEntity struct {
	Timestamp     int64  `xorm:"ts"`
	CaptureId     string `xorm:"capture_id"`
	SensorId      string `xorm:"sensor_id"`
	PlateText     string `xorm:"plate_text"`
	DriverType    int    `xorm:"driver_type"`
	PassengerType int    `xorm:"passenger_type"`
}

func (CoverFaceEntity) TableName() string {
	return CoverFaceTableName
}
