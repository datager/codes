package models

import (
	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/config"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	MaxPathAnalyzeTaskNumOfOneUser                              = 10  // 单人最多 10个 任务
	VehiclePathAnalyzeCreateTaskMaxSensorCount                  = 500 // 最多 500 sensor
	VehiclePathAnalyzeTaskMaxDurationDaysBetweenStartTsAndEndTs = 31  // 最多 31 天
	PathAnalyzeCreateTaskMaxTaskNameLen                         = 20  // 任务名必须<=20
)

const (
	PathAnalyzeTaskErr   = iota
	PathAnalyzeTaskStop  //停止任务
	PathAnalyzeTaskStart //启动任务
)

type PathAnalyzeTaskUpdateStatusRequest struct {
	AthenaTask
	Uts                  time.Time
	Ts                   int64
	UTCOffset            int
	TaskID               string              `uri:"taskid"` //任务ID
	TaskName             string              //任务名字
	PositionSensorIDsMap map[string][]string //轨迹点设备字典
	Trail                []string            //轨迹点数组
	PathConfidence       int64               //轨迹阈值
	StartAndEndTimestamp []int64             //起止时间
	UpdateType           int                 //更新的类型 1停止 2开始
	TaskType             int                 //用来区分车辆任务和人任务
	Plate                string
	Sensors              []MiniSensorSlice
	Pagination
}

func (m *PathAnalyzeTaskUpdateStatusRequest) Check() error {
	if strings.EqualFold(m.TaskName, "") {
		return NewErrorVerificationFailed(ErrorCodeAthenaPathAnalyzeModelCreateTaskEmptyTaskNameError)
	}
	lenTaskName := utf8.RuneCountInString(m.TaskName)
	if lenTaskName > PathAnalyzeCreateTaskMaxTaskNameLen {
		return NewErrorVerificationFailed(ErrorCodeAthenaPathAnalyzeModelCreateTaskTaskNameTooLongError)
	}

	if len(m.PositionSensorIDsMap) == 0 {
		return NewErrorVerificationFailed(ErrorCodePathAnalyzeModelCreateTaskSensorIDsEmptyError)
	}

	if len(m.StartAndEndTimestamp) != 2 {
		return NewErrorVerificationFailed(ErrorCodePathAnalyzeModelCreateTaskSelectedTimeError)
	}

	if m.TaskType != AthenaTaskTypeVehiclePathAnalyze && m.TaskType != AthenaTaskTypePersonPathAnalyze {
		return NewErrorVerificationFailed(ErrorCodePathAnalyzeModelCreateTaskTypeError)
	}
	count := 0
	for _, value := range m.PositionSensorIDsMap {
		count = count + len(value)
	}

	maxSensorCount := int64(config.GetConfig().GetIntOrDefault("api.pathanalyze.sensorCountLimit", VehiclePathAnalyzeCreateTaskMaxSensorCount))
	if count > int(maxSensorCount) {
		return NewErrorVerificationFailed(ErrorCodePathAnalyzeModelCreateTaskSensorIDsTooManyError)
	}

	maxSelectedDays := int64(config.GetConfig().GetIntOrDefault("api.pathanalyze.selectedDaysLimit", VehiclePathAnalyzeTaskMaxDurationDaysBetweenStartTsAndEndTs))
	if m.StartAndEndTimestamp[1] <= m.StartAndEndTimestamp[0] ||
		(m.StartAndEndTimestamp[1]-m.StartAndEndTimestamp[0]) > utils.DayToMillisecond(maxSelectedDays) {
		return NewErrorVerificationFailed(ErrorCodePathAnalyzeModelCreateTaskSelectedTimeError)
	}
	return nil
}

type AthenaUnmarshalPathAnalyzeTaskStatusResp struct {
	ID         string // 任务ID
	Type       int    // 任务类别
	ConfigInfo *PathAnalyzeTaskConfigInfo
	StatusInfo *PathAnalyzeTaskStatusInfo
}

type QueryPathAnalyzeTaskListReq struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int
	CurrentUser      *User // loki内部使用
}

func (r *QueryPathAnalyzeTaskListReq) Check() error {
	if len(r.AthenaTaskStatus) == 0 {
		r.AthenaTaskStatus = GetAthenaStatusEnumByTaskType(AthenaTaskTypeVehiclePathAnalyze)
	}
	if len(r.AthenaTaskType) == 0 {
		r.AthenaTaskType = []int{AthenaTaskTypeVehiclePathAnalyze, AthenaTaskTypePersonPathAnalyze}
	}
	return nil
}

type PathAnalyzeTaskConfigInfo struct {
	PositionSensorIDsMap map[string][]string
	Trail                []string
	StartTimestamp       int64
	EndTimestamp         int64
	PathConfidence       int64
}

type PathAnalyzeTaskStatusInfo struct {
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

type QueryPathAnalyzeTaskListResp struct {
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

type BatchDeleteReqPathAnalyzeModel struct {
	TaskIDs []string //删除的任务IDs
}

type BatchDelPathAnalyzeTasksResp struct {
	IsAllTasksDeleteSucceed bool           // 是否全部task删除成功
	DetailInfo              ResultMapError // 详情, key 为 taskID, 值为 err 原因
}

func NewBatchDelPathAnalyzeTasksResp() *BatchDelPathAnalyzeTasksResp {
	return &BatchDelPathAnalyzeTasksResp{
		IsAllTasksDeleteSucceed: true,
		DetailInfo:              make(map[string]error),
	}
}

type PathAnalyzeSummary struct {
	Total int64                     // 总条数
	Rets  []*PathAnalyzeSummaryItem // 数据
}

type PathAnalyzeSummaryItem struct {
	TaskID string // 时间戳
	Plate  string // 车牌
	Count  int64  // 次数
	VID    string
}

type PathAnalyzeDetailReq struct {
	TaskID string // 时间戳
	VID    string
	Pagination
}

type PathAnalyzeTaskDetailResp struct {
	TaskID      string
	TaskName    string
	SensorID    string
	SensorName  string
	Plate       string
	VehicleReID string
	VehicleID   string
	FaceReID    string
	FaceID      string
	VID         string
	Timestamp   int64
	ImageResult *ImageResult
}

type PathAnalyzeDetails struct {
	Total int64                        // 总条数
	Rets  []*PathAnalyzeTaskDetailResp // 数据
}

type SensorIDs []string
