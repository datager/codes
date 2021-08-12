package models

import "time"

const (
	AreaValueAll = iota
	AreaValueOne
	AreaValueTwo
	AreaValueThree
	AreaValueFour
	AreaValueFive
)

// 用于 数据库表entity的 model
type VehicleCollisionTaskModel struct {
	AthenaTask
	Uts            time.Time
	Ts             int64
	TaskID         string
	Sensors        []*MinSensors
	TimeSpace      []*TimeSpace
	StartTimestamp int64
	EndTimestamp   int64
}

type VehicleCollisionTaskDatilModel struct {
	TaskName    string
	AreaSet     []*SensorInfor
	Total       int
	Plate       string
	ReIDCount   int64
	VID         string
	VehicleReID string
	VehicleID   string
	SensorID    string
	SensorName  string
	Timestamp   int64
	TaskID      string `uri:"taskid"`
	ImageResult *ImageResult
	AreaValue   int64
	Pagination
}

type ListResult struct {
	Total  int64
	Rets   interface{}
	TaskID string
}

type VehicleCollisionTaskDatilModelResp struct {
	SensorIDs  []string
	Total      int64
	ReIDResult []*VehicleCollisionTaskDatilModel
}

type SensorInfor struct {
	StartTimestamp int64
	EndTimestamp   int64
	AreaSensors    *MinSensors
}
