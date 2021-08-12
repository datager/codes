package models

import (
	"errors"
	"strings"
	"unicode/utf8"

	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/config"
)

// 套牌
const (
	FakePlateUserTaskLimit = 10 // 单人最多 10 个 任务(人+车)
	FakePlateTaskNameLimit = 20 // 任务名必须 <=20

	VehicleFakePlateSensorCountLimit  = 500 // 最多 500 sensor
	VehicleFakePlateSelectedDaysLimit = 90  // 最多 90 天
)

const (
	FakePlateTaskErr = iota
	FakePlateStop
	FakePlateStart
)

type FakePlateListRequest struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int // 任务类型(人员频次分析、车辆频次分析)
}

type FakePlateCreateRequest struct {
	TaskID    string          // 任务ID(前端不需要传递)
	TaskName  string          // 任务名称
	SensorIDs []string        // 选中设备ID
	Sensors   MiniSensorSlice // 选中设备(前端不需要传递)
	StartAndEndTimestamp
}

type FakePlateCreateResponse struct {
	TaskID string // 任务ID
}

type FakePlateTaskListRequest struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int // 任务类型(车辆、人员)
}

func (r *FakePlateTaskListRequest) Check() error {
	if len(r.AthenaTaskStatus) == 0 {
		r.AthenaTaskStatus = GetAthenaStatusEnumByTaskType(AthenaTaskTypeVehicleFakePlate)
	}
	if len(r.AthenaTaskType) == 0 {
		r.AthenaTaskType = []int{AthenaTaskTypeVehicleFakePlate}
	}
	return nil
}

type FakePlateTaskListResponse struct {
	TaskID                string  // 任务ID
	TaskName              string  // 任务名称
	ErrMsg                string  // 任务错误原因
	ProcessPercent        float64 // 处理进度
	CreateTs              int64   // 任务创建时间
	AthenaTaskType        int     // 任务类型
	AthenaTaskStatus      int     // 任务状态
	ResultCount           int     // 结果数量
	StartAndEndTimestamps []int64
}

type FakePlateTaskUpdateRequest struct {
	TaskID     string
	UpdateType int // 停止:1,开启:2
}

type FakePlateTaskBatchDeleteRequest []string

type FakePlateTaskBatchDeleteResponse struct {
	IsAllTasksDeleteSucceed bool             // 是否全部删除成功
	DetailInfo              map[string]error // 详情
}

type VehicleFakePlateAthenaConfigInfo struct {
	PlateText string
	SensorIDs []string
	StartAndEndTimestamp
}

type VehicleFakePlateAthenaStatusInfo struct {
	StateType   int    // 任务状态类型
	PosInQueue  int    // 任务排队信息
	QueueLength int    // 任务排队信息
	Done        int    // 任务完成数
	TotalItems  int    // 任务总数
	ResultCount int    // 结果数
	Err         string // 错误信息
}

type VehicleFakePlateAthenaTaskStatus struct {
	ID         string                            // 任务ID
	Type       int                               // 任务类别
	ConfigInfo *VehicleFakePlateAthenaConfigInfo // 任务配置信息
	StatusInfo *VehicleFakePlateAthenaStatusInfo // 任务状态信息
}

type VehicleFakePlateCreateRequest struct {
	FakePlateCreateRequest
	PlateText string // 车牌
}

func (m *FakePlateCreateRequest) Check() error {
	if strings.EqualFold(m.TaskName, "") {
		return NewErrorVerificationFailed(ErrorCodeAthenaCollisionModelCreateTaskEmptyTaskNameError)
	}
	lenTaskName := utf8.RuneCountInString(m.TaskName)
	if lenTaskName > FakePlateTaskNameLimit {
		return NewErrorVerificationFailed(ErrorCodeAthenaCollisionModelCreateTaskTaskNameTooLongError)
	}
	if len(m.SensorIDs) == 0 {
		return NewErrorVerificationFailed(ErrorCodeFakePlateModelCreateTaskSensorIDsEmptyError)
	}
	maxSensorCount := int64(config.GetConfig().GetIntOrDefault("api.fakePlate.sensorCountLimit", VehicleFakePlateSensorCountLimit))
	if len(m.SensorIDs) > int(maxSensorCount) {
		return NewErrorVerificationFailed(ErrorCodeFakePlateModelCreateTaskSensorIDsTooManyError)
	}
	maxSelectedDays := int64(config.GetConfig().GetIntOrDefault("api.fakePlate.selectedDaysLimit", VehicleFakePlateSelectedDaysLimit))
	if m.EndTimestamp <= m.StartTimestamp ||
		(m.EndTimestamp-m.StartTimestamp) > utils.DayToMillisecond(maxSelectedDays) {
		return NewErrorVerificationFailed(ErrorCodeFakePlateModelCreateTaskSelectedTimeError)
	}
	return nil
}

type VehicleFakePlateResultRequest struct {
	TaskID    string // 任务ID
	PlateText string // 车牌
	Pagination
}

func (m *VehicleFakePlateResultRequest) DetailCheck() error {
	if strings.EqualFold(m.PlateText, "") {
		return errors.New("query detail need plate text")
	}
	return nil
}

type VehicleFakePlateSummaryItem struct {
	Timestamp int64  // 时间戳
	PlateText string // 车牌
	Num       int64  // 次数
}

type VehicleFakePlateSummaryResponse struct {
	Total int64                          // 总条数
	Data  []*VehicleFakePlateSummaryItem // 数据
}

type VehicleFakePlateDetailItem struct {
	Timestamp int64                        // 时间戳
	Vehicles  []*MiniVehicleWithMiniSensor // 车辆列表
}

type VehicleFakePlateDetailResponse struct {
	Total int64                         // 总条数
	Data  []*VehicleFakePlateDetailItem // 数据
}
