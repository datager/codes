package models

import "encoding/json"

type MiniVehicle struct {
	Timestamp   int64         // 时间戳
	VehicleID   string        // 车辆ID
	VehicleReID string        // 车辆ReID
	VehicleVID  string        // 车辆VID
	PlateText   string        // 车牌
	ImageResult CutboardImage // 图片信息
}

func (s *MiniVehicle) ToDB() ([]byte, error) {
	return json.Marshal(s)
}

func (s *MiniVehicle) FromDB(b []byte) error {
	return json.Unmarshal(b, s)
}

type MiniVehicleSlice []*MiniVehicle

func (s *MiniVehicleSlice) ToDB() ([]byte, error) {
	return json.Marshal(s)
}

func (s *MiniVehicleSlice) FromDB(b []byte) error {
	return json.Unmarshal(b, s)
}

type MiniVehicleWithMiniSensor struct {
	MiniVehicle
	MiniSensor
}
