package entities

const NonmotorWithPeopleTableName = "nonmotor_with_people_captures"

type NonmotorWithPeopleEntity struct {
	Timestamp              int64  `xorm:"ts"`
	CaptureId              string `xorm:"capture_id"`
	SensorId               string `xorm:"sensor_id"`
	NonmotorWithPeopleType int    `xorm:"nonmotor_with_people_type"`
}

func (NonmotorWithPeopleEntity) TableName() string {
	return NonmotorWithPeopleTableName
}
