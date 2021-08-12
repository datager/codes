package models

import "time"

// 用于 jormougand
// 频次分析任务 详情 响应体(人)
type PersonFrequencyTaskDetailResp struct {
	QueryFrequencyTaskListResp
	IDNo    string    // 身份证
	Sensors []*Sensor // 新建任务时所选的设备(包含 id, name, geo), 列表接口不返回, 详情接口返回
}

// 新建人员频次分析任务请求体
type CreatePersonFrequencyTaskReq struct {
	CreateFrequencyTaskReq
	IDNo string // 身份证号
}

// 用于 数据库表entity的 model
type PersonFrequencyTask struct {
	Uts            time.Time
	Ts             int64
	TaskID         string
	SensorIDs      []string // not in use after #FACE-4459
	Sensors        Sensors  // in use after #FACE-4459
	TimesThreshold int64    // 频次
	IDNo           string   // 身份证号
}

// 用于 数据库表entity的 model
type PersonFrequencyTaskExtend struct {
	AthenaTask
	SensorIDs      []string // not in use after #FACE-4459
	Sensors        Sensors  // in use after #FACE-4459
	TimesThreshold int64
	IDNo           string
}
