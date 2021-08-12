package sensor

import (
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
)

func GetAllSensorIDs() ([]string, error) {
	var sIDs []string
	err := dbengine.GetV5Instance().Select("sensor_id").Table("sensors").Find(&sIDs)
	return sIDs, err
}

func GetSensorIDsByLimit(off, limit int) ([]string, error) {
	var sIDs []string
	err := dbengine.GetV5Instance().Select("sensor_id").Table("sensors").Limit(limit, off).Find(&sIDs)
	return sIDs, err
}

func GetTotalSensorCount() (int64, error) {
	sensor := new(entities.Sensors)
	return dbengine.GetV5Instance().Count(sensor)
}

func GetBySensorIntID(sIntID int64) (*entities.Sensors, error) {
	s := new(entities.Sensors)
	_, err := dbengine.GetV5Instance().Table(entities.TableNameSensors).Where("sensor_int_id = ?", sIntID).Get(s)
	return s, err
}
