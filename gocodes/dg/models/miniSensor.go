package models

import "encoding/json"

type MiniSensor struct {
	SensorID   string  // 设备ID
	SensorName string  // 设备名称
	Longitude  float64 // 经纬度
	Latitude   float64 // 经纬度
}

func (s *MiniSensor) ToDB() ([]byte, error) {
	return json.Marshal(s)
}

func (s *MiniSensor) FromDB(b []byte) error {
	return json.Unmarshal(b, s)
}

type MiniSensorSlice []*MiniSensor

func (s *MiniSensorSlice) ToDB() ([]byte, error) {
	return json.Marshal(s)
}

func (s *MiniSensorSlice) FromDB(b []byte) error {
	return json.Unmarshal(b, s)
}
