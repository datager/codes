package entities

const UnbeltedTableName = "unbelted_captures"

type UnbeltedEntity struct {
	Timestamp     int64  `xorm:"ts"`
	CaptureId     string `xorm:"capture_id"`
	SensorId      string `xorm:"sensor_id"`
	PlateText     string `xorm:"plate_text"`
	DriverType    int    `xorm:"driver_type"`
	PassengerType int    `xorm:"passenger_type"`
}

func (UnbeltedEntity) TableName() string {
	return UnbeltedTableName
}
