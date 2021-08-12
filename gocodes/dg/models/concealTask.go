package models

import (
	"strings"
	"time"
	"unicode/utf8"

	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/config"
)

const (
	MaxHiddenTaskNumOfOneUser                       = 10  // 单人最多 10个 任务
	HiddenCreateTaskMaxSensorCount                  = 500 // 最多 500 sensor
	HiddenTaskMaxDurationDaysBetweenStartTsAndEndTs = 31  // 最多 31 天
	HiddenCreateTaskMaxTaskNameLen                  = 20  // 任务名必须<=20
)

const (
	HiddenTaskErr   = iota
	HiddenTaskStop  //停止任务
	HiddenTaskStart //启动任务
)

type HiddenTaskUpdateStatusRequest struct {
	AthenaTask
	Uts                  time.Time
	Ts                   int64
	UTCOffset            int
	TaskID               string   `uri:"taskid"` //任务ID
	TaskName             string   //任务名字
	SensorIDs            []string //设备数组
	StartAndEndTimestamp []int64  //起止时间
	UpdateType           int      //更新的类型 1停止 2开始
	Sensors              MiniSensorSlice
	Plate                string
	Pagination
}

func (m *HiddenTaskUpdateStatusRequest) Check() error {
	if strings.EqualFold(m.TaskName, "") {
		return NewErrorVerificationFailed(ErrorCodeAthenaHiddenModelCreateTaskEmptyTaskNameError)
	}
	lenTaskName := utf8.RuneCountInString(m.TaskName)
	if lenTaskName > HiddenCreateTaskMaxTaskNameLen {
		return NewErrorVerificationFailed(ErrorCodeAthenaHiddenModelCreateTaskTaskNameTooLongError)
	}

	if len(m.SensorIDs) == 0 {
		return NewErrorVerificationFailed(ErrorCodeHiddenModelCreateTaskSensorIDsEmptyError)
	}

	maxSensorCount := int64(config.GetConfig().GetIntOrDefault("api.hidden.sensorCountLimit", HiddenCreateTaskMaxSensorCount))
	if len(m.SensorIDs) > int(maxSensorCount) {
		return NewErrorVerificationFailed(ErrorCodeHiddenModelCreateTaskSensorIDsTooManyError)
	}
	maxSelectedDays := int64(config.GetConfig().GetIntOrDefault("api.hidden.selectedDaysLimit", HiddenTaskMaxDurationDaysBetweenStartTsAndEndTs))
	if m.StartAndEndTimestamp[1] <= m.StartAndEndTimestamp[0] ||
		(m.StartAndEndTimestamp[1]-m.StartAndEndTimestamp[0]) > utils.DayToMillisecond(maxSelectedDays) {
		return NewErrorVerificationFailed(ErrorCodeHiddenModelCreateTaskSelectedTimeError)
	}
	return nil
}

type AthenaUnmarshalHiddenTaskStatusResp struct {
	ID         string // 任务ID
	Type       int    // 任务类别
	ConfigInfo *HiddenTaskConfigInfo
	StatusInfo *HiddenTaskStatusInfo
}

type QueryHiddenTaskListReq struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int
	CurrentUser      *User // loki内部使用
}

func (r *QueryHiddenTaskListReq) Check() error {
	if len(r.AthenaTaskStatus) == 0 {
		r.AthenaTaskStatus = GetAthenaStatusEnumByTaskType(AthenaTaskTypeVehicleHidden)
	}
	if len(r.AthenaTaskType) == 0 {
		r.AthenaTaskType = []int{AthenaTaskTypeVehicleHidden}
	}
	return nil
}

type HiddenTaskConfigInfo struct {
	SensorIDs      []string
	StartTimestamp int64
	EndTimestamp   int64
}

type HiddenTaskStatusInfo struct {
	//任务状态类型
	StateType int32

	// 任务排队信息
	PosInQueue  int
	QueueLength int

	//任务进度信息
	Done       int
	TotalItems int

	ResultCount int
	Err         string
}

type QueryHiddenTaskListResp struct {
	TaskID               string //任务ID
	TaskName             string // 任务名称
	ErrMsg               string
	ProcessPercent       float64 // 处理进度
	CreateTs             int64   // 任务创建时间
	AthenaTaskStatus     int     //任务状态
	ResultCount          int     //结果数量
	TaskType             int     //任务类型
	StartAndEndTimestamp []int64
}

type BatchDeleteReqHiddenModel struct {
	TaskIDs []string //删除的任务IDs
}

type BatchDelHiddenTasksResp struct {
	IsAllTasksDeleteSucceed bool           // 是否全部task删除成功
	DetailInfo              ResultMapError // 详情, key 为 taskID, 值为 err 原因
}

func NewBatchDelHiddenTasksResp() *BatchDelHiddenTasksResp {
	return &BatchDelHiddenTasksResp{
		IsAllTasksDeleteSucceed: true,
		DetailInfo:              make(map[string]error),
	}
}

type VehicleHiddenSummary struct {
	Total int64                       // 总条数
	Rets  []*VehicleHiddenSummaryItem // 数据
}

type VehicleHiddenSummaryItem struct {
	TaskID string // 时间戳
	Plate  string // 车牌
	Count  int64  // 次数
	VID    string
}

type VehicleHiddenDetailReq struct {
	TaskID string // 时间戳
	VID    string
	Pagination
}

type VehicleHiddenTaskDetailResp struct {
	TaskID      string
	TaskName    string
	SensorID    string
	SensorName  string
	Plate       string
	VehicleReID string
	VehicleID   string
	VID         string
	Timestamp   int64
	ImageResult *ImageResult
}

type VehicleHiddenDetails struct {
	Total int64                          // 总条数
	Rets  []*VehicleHiddenTaskDetailResp // 数据
}
