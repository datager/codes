package models

//-----------------------------------------------------------------------
// 车辆频次 + 人员频次通用
//-----------------------------------------------------------------------
type AthenaCommonFrequencyAnalysisStatusInfo struct {
	StateType  int    // 任务状态
	Done       int    // 同步完成设备
	TotalItems int    // 需要同步的总设备
	Err        string // 错误信息, 在 athena 的 error 状态时需要处理
}

//-----------------------------------------------------------------------
// 车辆频次
//-----------------------------------------------------------------------
// 新建车辆频次分析 for athena
type AthenaVehicleFrequencyAnalysisConfig struct {
	StartTs        int64    // 起始时间
	EndTs          int64    // 起始时间
	TimesThreshold int64    // 频次阈值
	LicencePlateNo string   // 车牌号
	SensorIds      []string // 设备id
}

// 由clients.StatusResponse 解析后的 车辆频次的Athena 结构体
type AthenaUnmarshalVehicleFrequencyTaskStatusResp struct {
	ID         string                                   // 任务ID
	Type       int                                      // 任务类别
	ConfigInfo *AthenaVehicleFrequencyAnalysisConfig    // 任务配置信息
	StatusInfo *AthenaCommonFrequencyAnalysisStatusInfo // 任务状态信息
}

//-----------------------------------------------------------------------
// 人员频次
//-----------------------------------------------------------------------
// 新建人员频次分析 for athena
type AthenaPersonFrequencyAnalysisConfig struct {
	StartTs        int64    // 起始时间
	EndTs          int64    // 起始时间
	TimesThreshold int64    // 频次阈值
	IDNo           string   // 身份证号
	SensorIds      []string // 设备id
}

// 由clients.StatusResponse 解析后的 人员频次的Athena 结构体
type AthenaUnmarshalPersonFrequencyTaskStatusResp struct {
	ID         string                                   // 任务ID
	Type       int                                      // 任务类别
	ConfigInfo *AthenaPersonFrequencyAnalysisConfig     // 任务配置信息
	StatusInfo *AthenaCommonFrequencyAnalysisStatusInfo // 任务状态信息
}
