package models

import (
	"codes/gocodes/dg/utils/json"
	"time"
)

// 用于 jormougand
// 频次分析任务 详情 响应体(车)
type VehicleFrequencyTaskDetailResp struct {
	QueryFrequencyTaskListResp
	Plate   string    // 车牌号
	Sensors []*Sensor // 新建任务时所选的设备(包含 id, name, geo), 列表接口不返回, 详情接口返回
}

// 新建车辆频次分析任务请求体
type CreateVehicleFrequencyTaskReq struct {
	CreateFrequencyTaskReq
	Plate string // 车牌号
}

// 用于 数据库表entity的 model
type VehicleFrequencyTask struct {
	Uts            time.Time
	Ts             int64
	TaskID         string
	SensorIDs      []string // not in use after #FACE-4459
	Sensors        Sensors  // in use after #FACE-4459
	TimesThreshold int64    // 频次
	Plate          string   // 车牌号
}

// 用于 数据库表entity的 model
type VehicleFrequencyTaskExtend struct {
	AthenaTask
	SensorIDs      []string // not in use after #FACE-4459
	Sensors        Sensors  // in use after #FACE-4459
	TimesThreshold int64
	Plate          string
}

// VehicleFrequencyTask的 Sensors 以 json类型存
type Sensors []*Sensor

func (sensors *Sensors) ToDB() ([]byte, error) {
	return json.Marshal(sensors)
}

func (sensors *Sensors) FromDB(b []byte) error {
	return json.Unmarshal(b, sensors)
}
