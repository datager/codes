package models

import (
	"strings"
	"time"
	"unicode/utf8"

	"codes/gocodes/dg/utils/config"
)

const (
	MaxFrequencyAnalyzeTaskNumInAthena                              = 10  // 所有 athena 里的频次分析 最多10个 任务(人+车)
	MaxFrequencyAnalyzeTaskNumOfOneUser                             = 10  // 单人最多频次分析 10个 任务(人+车)
	VehicleFrequencyCreateTaskMaxSensorCount                        = 500 // 最多 500 sensor
	VehicleFrequencyCreateTaskMaxDurationDaysBetweenStartTsAndEndTs = 31  // 最多 30 天
	FrequencyAnalyzeCreateTaskMaxTaskNameLen                        = 20  // 任务名必须<=20
	FrequencyAnalyzeCreateTaskMinTimesThreshold                     = 2   // 频次必须>=2
)

// 人车通用

// -----------------------------------------------------------------
// 查询 任务 列表
// -----------------------------------------------------------------
// 用于 jormougand
// 频次分析任务列表请求体
type QueryFrequencyTaskListReq struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int // 任务类型(人员频次分析、车辆频次分析)
	CurrentUser      *User // loki内部使用
}

// t: frequencyTypeToBeFilter(只保留 参数 t 的 type)
func (query *QueryFrequencyTaskListReq) FilterByFrequencyType(t int) *QueryFrequencyTaskListReq {
	filteredTypes := make([]int, 0)
	for _, tp := range query.AthenaTaskType {
		if tp == t {
			filteredTypes = append(filteredTypes, t)
		}
	}

	filtered := &QueryFrequencyTaskListReq{
		AthenaTaskStatus: query.AthenaTaskStatus,
		AthenaTaskType:   filteredTypes,
		CurrentUser:      query.CurrentUser,
	}
	return filtered
}

// 用于 jormougand
// 频次分析任务 列表/详情 响应体
type QueryFrequencyTaskListResp struct {
	TaskID                   string
	TaskName                 string  // 任务名称
	ErrMsg                   string  // 任务错误原因, 是额外透传的信息
	StartAndEndTimestamps    []int64 // 起始时间+结束时间
	TimesThreshold           int64   // 频次阈值
	AnalyzeProcessPercent    float64 // 处理进度 = ProcessedSensorCnt/TotalSensorCnt
	ProcessedSensorCnt       int     // 处理过的 sensor 数量, 是额外透传的信息
	TotalSensorCnt           int     // 需要频次分析的总 sensor数量, 是额外透传的信息
	CreateTs                 int64   // 任务创建时间
	AthenaTaskStatus         int     // 任务状态(排队中、进行中、已完成、暂停中)
	FrequencyAnalyzeTaskType int     // 任务类型(人员频次分析、车辆频次分析)
}

type QueryFrequencyTaskListResps []*QueryFrequencyTaskListResp

func NewFrequencyTaskListQueryResps() QueryFrequencyTaskListResps {
	return make(QueryFrequencyTaskListResps, 0)
}

func (resps QueryFrequencyTaskListResps) Len() int {
	return len(resps)
}

// CreateTs 降序
func (resps QueryFrequencyTaskListResps) Less(i, j int) bool {
	return resps[i].CreateTs > resps[j].CreateTs
}

func (resps QueryFrequencyTaskListResps) Swap(i, j int) {
	resps[i], resps[j] = resps[j], resps[i]
}

// -----------------------------------------------------------------
// 创建 任务
// -----------------------------------------------------------------
// 新建车辆频次分析任务请求体
type CreateFrequencyTaskReq struct {
	TaskName    string // 任务名称
	CurrentUser *User  // loki内部使用

	// 以下字段来自 athena models.AthenaVehicleFrequencyAnalysisConfig, 用于发送给 athena
	StartTimestamp int64    // 起始时间, ms
	EndTimestamp   int64    // 起始时间, ms
	TimesThreshold int64    // 频次阈值
	SensorIDs      []string `json:"SensorIDs"` // 设备id
}

func (createReq *CreateFrequencyTaskReq) Valid() error {
	if err := createReq.validSensor(); err != nil {
		return err
	}
	if err := createReq.validTs(); err != nil {
		return err
	}
	if err := createReq.validOther(); err != nil {
		return err
	}
	return nil
}

func (createReq *CreateFrequencyTaskReq) validSensor() error {
	if len(createReq.SensorIDs) == 0 {
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskCouldNotHaveEmptySensorsError, "len(sensorIDs) == 0")
	}

	maxTaskSensorCountLimit := config.GetConfig().GetIntOrDefault("api.frequencyAnalyze.common.createTaskMaxSensorCount", VehicleFrequencyCreateTaskMaxSensorCount)
	if len(createReq.SensorIDs) > maxTaskSensorCountLimit {
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskSensorCountReachMaxLimitError, "len(sensorIDs): %d > maxSensorCountLimit: %d", len(createReq.SensorIDs), maxTaskSensorCountLimit)
	}
	return nil
}

func (createReq *CreateFrequencyTaskReq) validTs() error {
	if createReq.StartTimestamp <= 0 || createReq.EndTimestamp <= 0 {
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskIllegalTsError, "startts <= 0 || endts <= 0")
	}

	if createReq.StartTimestamp >= createReq.EndTimestamp {
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskStartTsGreaterEqualThanEndTsError, "startts >= endts")
	}

	maxTaskTsDuration := config.GetConfig().GetIntOrDefault("api.frequencyAnalyze.common.createTaskMaxTsDurationDays", VehicleFrequencyCreateTaskMaxDurationDaysBetweenStartTsAndEndTs)
	if (createReq.EndTimestamp - createReq.StartTimestamp) > (int64(maxTaskTsDuration) * 24 * int64(time.Hour) / 1e6) {
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError, "(endts - startts) > %d days, startts: %d, endts: %d", maxTaskTsDuration, createReq.StartTimestamp, createReq.EndTimestamp)
	}
	return nil
}

func (createReq *CreateFrequencyTaskReq) validOther() error {
	if strings.EqualFold(createReq.TaskName, "") {
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskEmptyTaskNameError, "illegal: task name should not be empty")
	}

	lenTaskName := utf8.RuneCountInString(createReq.TaskName)
	if lenTaskName > FrequencyAnalyzeCreateTaskMaxTaskNameLen { // 任务最大名称
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskTaskNameTooLongError, "illegal: len(task name): %d should not > %d", lenTaskName, FrequencyAnalyzeCreateTaskMaxTaskNameLen)
	}
	if createReq.TimesThreshold < FrequencyAnalyzeCreateTaskMinTimesThreshold {
		return NewErrorf(ErrorCodeAthenaVehicleFrequencyModelCreateTaskIllegalTimesThresholdError, "illegal: times threshold should not < %d", FrequencyAnalyzeCreateTaskMinTimesThreshold)
	}
	return nil
}

// -----------------------------------------------------------------
// 删除 任务
// -----------------------------------------------------------------
type ResultMapError map[string]error

// 批量删频次任务的响应
type BatchDelFrequencyTasksResp struct {
	IsAllTasksDeleteSucceed bool           // 是否全部task删除成功
	DetailInfo              ResultMapError // 详情, key 为 taskID, 值为 err 原因
}

func NewBatchDelFrequencyTasksResp() *BatchDelFrequencyTasksResp {
	return &BatchDelFrequencyTasksResp{
		IsAllTasksDeleteSucceed: true,
		DetailInfo:              make(map[string]error),
	}
}
