package models

import (
	"encoding/json"
	"strings"
	"unicode/utf8"

	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/config"
)

// 首次入城
const (
	FirstAppearanceUserTaskLimit = 10 // 单人最多 10 个 任务(人+车)
	FirstAppearanceTaskNameLimit = 20 // 任务名必须 <=20

	VehicleFirstAppearanceSensorCountLimit      = 500 // 最多 500 sensor
	VehicleFirstAppearanceSelectedDaysLimit     = 90  // 最多 90 天
	VehicleFirstAppearanceBacktrackingDaysLimit = 90  // 最多 90 天
)

const (
	FirstAppearanceTaskErr = iota
	FirstAppearanceStop
	FirstAppearanceStart
)

type FirstAppearanceListRequest struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int // 任务类型(人员频次分析、车辆频次分析)
}

type FirstAppearanceCreateRequest struct {
	TaskID           string          // 任务ID(前端不需要传递)
	TaskName         string          // 任务名称
	SensorIDs        []string        // 选中设备ID
	Sensors          MiniSensorSlice // 选中设备(前端不需要传递)
	BacktrackingDays int64           // 回溯时长
	StartAndEndTimestamp
}

type FirstAppearanceCreateResponse struct {
	TaskID string // 任务ID
}

type FirstAppearanceTaskListRequest struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int // 任务类型(车辆、人员)
}

func (r *FirstAppearanceTaskListRequest) Check() error {
	if len(r.AthenaTaskStatus) == 0 {
		r.AthenaTaskStatus = GetAthenaStatusEnumByTaskType(AthenaTaskTypeVehicleFirstAppearance)
	}
	if len(r.AthenaTaskType) == 0 {
		r.AthenaTaskType = []int{AthenaTaskTypeVehicleFirstAppearance}
	}
	return nil
}

type FirstAppearanceTaskListResponse struct {
	TaskID                string  // 任务ID
	TaskName              string  // 任务名称
	ErrMsg                string  // 任务错误原因
	ProcessPercent        float64 // 处理进度
	CreateTs              int64   // 任务创建时间
	AthenaTaskType        int     // 任务类型
	AthenaTaskStatus      int     // 任务状态
	ResultCount           int     // 结果数量
	Backtracking          int64   // 回溯时长
	StartAndEndTimestamps []int64
}

type FirstAppearanceTaskUpdateRequest struct {
	TaskID     string
	UpdateType int // 停止:1,开启:2
}

type FirstAppearanceTaskBatchDeleteRequest []string

type FirstAppearanceTaskBatchDeleteResponse struct {
	IsAllTasksDeleteSucceed bool             // 是否全部删除成功
	DetailInfo              map[string]error // 详情
}

// 车辆首次入城
type VehicleFirstAppearanceAthenaConfigInfo struct {
	SensorIDs        []string
	BacktrackingDays int64
	VehicleCondition VehicleFirstAppearanceOtherCondition
	StartAndEndTimestamp
}

type VehicleFirstAppearanceAthenaStatusInfo struct {
	StateType   int    // 任务状态类型
	PosInQueue  int    // 任务排队信息
	QueueLength int    // 任务排队信息
	Done        int    // 任务完成数
	TotalItems  int    // 任务总数
	ResultCount int    // 结果数
	Err         string // 错误信息
}

type VehicleFirstAppearanceAthenaTaskStatus struct {
	ID         string                                  // 任务ID
	Type       int                                     // 任务类别
	ConfigInfo *VehicleFirstAppearanceAthenaConfigInfo // 任务配置信息
	StatusInfo *VehicleFirstAppearanceAthenaStatusInfo // 任务状态信息
}

type VehicleFirstAppearanceOtherCondition struct {
	PlateText    string // 车牌
	BrandID      *int32 // 主品牌
	SubBrandID   *int32 // 子品牌
	TypeID       *int32 // 车辆类型
	ColorID      *int32 // 车辆颜色
	PlateTypeID  *int32 // 车牌类型
	PlateColorID *int32 // 车牌颜色
	Symbol       *int   // 标志物
	Symbols      []int  // 标志物
	Special      *int   // 特殊车辆
	Specials     []int  // 特殊车辆
}

func (s *VehicleFirstAppearanceOtherCondition) ToDB() ([]byte, error) {
	return json.Marshal(s)
}

func (s *VehicleFirstAppearanceOtherCondition) FromDB(b []byte) error {
	return json.Unmarshal(b, s)
}

type VehicleFirstAppearanceCreateRequest struct {
	FirstAppearanceCreateRequest
	VehicleFirstAppearanceOtherCondition
}

func (m *FirstAppearanceCreateRequest) Check() error {
	if strings.EqualFold(m.TaskName, "") {
		return NewErrorVerificationFailed(ErrorCodeFirstAppearanceModelCreateTaskEmptyTaskNameError)
	}
	lenTaskName := utf8.RuneCountInString(m.TaskName)
	if lenTaskName > FirstAppearanceTaskNameLimit {
		return NewErrorVerificationFailed(ErrorCodeAthenaCollisionModelCreateTaskTaskNameTooLongError)
	}
	if len(m.SensorIDs) == 0 {
		return NewErrorVerificationFailed(ErrorCodeFirstAppearanceModelCreateTaskSensorIDsEmptyError)
	}
	maxSensorCount := int64(config.GetConfig().GetIntOrDefault("api.firstAppearance.sensorCountLimit", VehicleFirstAppearanceSensorCountLimit))
	if len(m.SensorIDs) > int(maxSensorCount) {
		return NewErrorVerificationFailed(ErrorCodeFirstAppearanceModelCreateTaskSensorIDsTooManyError)
	}
	maxSelectedDays := int64(config.GetConfig().GetIntOrDefault("api.firstAppearance.selectedDaysLimit", VehicleFirstAppearanceSelectedDaysLimit))
	if m.EndTimestamp <= m.StartTimestamp ||
		(m.EndTimestamp-m.StartTimestamp) > utils.DayToMillisecond(maxSelectedDays) {
		return NewErrorVerificationFailed(ErrorCodeFirstAppearanceModelCreateTaskSelectedTimeError)
	}
	maxBacktrackingDays := int64(config.GetConfig().GetIntOrDefault("api.firstAppearance.backtrackingDaysLimit", VehicleFirstAppearanceBacktrackingDaysLimit))
	if m.BacktrackingDays <= 0 ||
		m.BacktrackingDays > utils.DayToMillisecond(maxBacktrackingDays) {
		return NewErrorVerificationFailed(ErrorCodeFirstAppearanceModelCreateTaskBacktrackingTimeError)
	}
	return nil
}

func (m *VehicleFirstAppearanceCreateRequest) Check() error {
	if err := m.FirstAppearanceCreateRequest.Check(); err != nil {
		return err
	}
	if strings.EqualFold(m.VehicleFirstAppearanceOtherCondition.PlateText, "") {
		return NewErrorVerificationFailed(ErrorCodeFirstAppearanceModelCreateTaskPlateTextEmptyError)
	}
	if len(m.Specials) > 0 {
		m.Special = new(int)
		*m.Special = SymbolSliceToInt(m.Specials)
	}
	if len(m.Symbols) > 0 {
		m.Symbol = new(int)
		*m.Symbol = SymbolSliceToInt(m.Symbols)
	}
	return nil
}

type VehicleFirstAppearanceDetailRequest struct {
	TaskID string // 任务ID
	Pagination
}

type VehicleFirstAppearanceDetailResponse struct {
	Total int64                        // 总条数
	Data  []*MiniVehicleWithMiniSensor // 数据
}
