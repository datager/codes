package models

import "time"

// 用于 数据库表entity的 model
type PersonCollisionTaskModel struct {
	AthenaTask
	Uts            time.Time
	Ts             int64
	TaskID         string
	Sensors        []*MinSensors
	TimeSpace      []*TimeSpace
	StartTimestamp int64
	EndTimestamp   int64
}

type PersonCollisionTaskDatilModel struct {
	TaskName    string
	AreaSet     []*SensorInfor
	Total       int
	Plate       string
	ReIDCount   int64
	VID         string
	FaceReID    string
	FaceID      string
	SensorID    string
	SensorName  string
	Timestamp   int64
	TaskID      string `uri:"taskid"`
	ImageResult *ImageResult
	AreaValue   int64
	Pagination
}

type PersonCollisionTaskDatilModelResp struct {
	SensorIDs  []string
	Total      int64
	ReIDResult []*PersonCollisionTaskDatilModel
}
