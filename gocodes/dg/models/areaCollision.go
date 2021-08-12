package models

import (
	"strings"
	"time"
	"unicode/utf8"

	"codes/gocodes/dg/utils/config"
	"codes/gocodes/dg/utils/json"
)

const (
	MaxAreaCollisionTaskNumInAthena                                 = 10  // 所有 athena 里的区域碰撞 最多10个 任务(人+车)
	MaxAreaCollisionAnalyzeTaskNumOfOneUser                         = 10  // 单人最多区域碰撞 10个 任务(人+车)
	VehicleCollisionCreateTaskMaxSensorCount                        = 500 // 最多 500 sensor
	VehicleCollisionCreateTaskMaxDurationDaysBetweenStartTsAndEndTs = 31  // 最多 30 天
	CollisionCreateTaskMaxTaskNameLen                               = 20  // 任务名必须<=20
)

const (
	AreaCollisionTaskErr   = iota
	AreaCollisionTaskStop  //停止碰撞任务
	AreaCollisionTaskStart //启动碰撞任务
)

// 新建车辆碰撞+人员任务
type AreaCollisionTaskModel struct {
	TaskID      string
	TaskName    string
	TimeSpaces  []*TimeSpace
	CurrentUser *User
	UpdateType  int
}

type BatchDeleteAreaCollisionReqModel struct {
	TaskIDs []string
}
type TimeSpaces []*TimeSpace
type TimeSpace struct {
	SensorIds      []string
	StartTimestamp int64
	EndTimestamp   int64
	IDNo           int64
}

func TimeSpaceTimeToQueryArea(timespace TimeSpaces) [][]int64 {
	StartAndEndTimes := make([][]int64, 0)

	for _, value := range timespace {
		StartAndEndTime := make([]int64, 0)
		StartAndEndTime = append(StartAndEndTime, value.StartTimestamp)
		StartAndEndTime = append(StartAndEndTime, value.EndTimestamp)
		StartAndEndTimes = append(StartAndEndTimes, StartAndEndTime)
	}
	return StartAndEndTimes
}

type MinSensors []*MinSensor

func (s *MinSensors) ToDB() ([]byte, error) {
	return json.Marshal(s)
}

func (s *MinSensors) FromDB(b []byte) error {
	return json.Unmarshal(b, s)
}

type MinSensor struct {
	SensorID   string
	SensorName string
	Longitude  float64
	Latitude   float64
}

func (ac *AreaCollisionTaskModel) Valid() error {
	if err := ac.validTimeSpaces(); err != nil {
		return err
	}

	if err := ac.validOther(); err != nil {
		return err
	}
	return nil
}

func (ac *AreaCollisionTaskModel) validTimeSpaces() error {

	if len(ac.TimeSpaces) == 0 || len(ac.TimeSpaces) > 5 {
		return NewErrorf(ErrorCodeAthenaCollisionModelCreateTaskCouldNotHaveEmptyTimeSpacesError, "len(TimeSpaces) == 0 or len(TimeSpaces) > 5")
	}
	validMap1 := make(map[string]*TimeSpace, len(ac.TimeSpaces))
	validMap2 := make(map[string]*TimeSpace, len(ac.TimeSpaces))
	validMap3 := make(map[string]*TimeSpace, len(ac.TimeSpaces))
	validMap4 := make(map[string]*TimeSpace, len(ac.TimeSpaces))
	validMap5 := make(map[string]*TimeSpace, len(ac.TimeSpaces))

	for k, TimeSpace := range ac.TimeSpaces {
		maxTaskSensorCountLimit := config.GetConfig().GetIntOrDefault("api.areacollision.sensorCountLimit", VehicleCollisionCreateTaskMaxSensorCount)
		if len(TimeSpace.SensorIds) > maxTaskSensorCountLimit {
			return NewErrorf(ErrorCodeAthenaCollisionModelCreateTaskSensorCountReachMaxLimitError, "len(sensorIDs): %d > maxSensorCountLimit: %d", len(TimeSpace.SensorIds), maxTaskSensorCountLimit)
		}

		if TimeSpace.StartTimestamp <= 0 || TimeSpace.EndTimestamp <= 0 {
			return NewErrorf(ErrorCodeAthenaCollisionModelCreateTaskIllegalTsError, "startts <= 0 || endts <= 0")
		}

		if TimeSpace.StartTimestamp >= TimeSpace.EndTimestamp {
			return NewErrorf(ErrorCodeAthenaCollisionModelCreateTaskStartTsGreaterEqualThanEndTsError, "startts >= endts")
		}

		maxTaskTsDuration := config.GetConfig().GetIntOrDefault("api.areacollision.selectedDaysLimit", VehicleCollisionCreateTaskMaxDurationDaysBetweenStartTsAndEndTs)
		if (TimeSpace.EndTimestamp - TimeSpace.StartTimestamp) > (int64(maxTaskTsDuration) * 24 * int64(time.Hour) / 1e6) {
			return NewErrorf(ErrorCodeAthenaCollisionModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError, "(endts - startts) > %d days, startts: %d, endts: %d", maxTaskTsDuration, TimeSpace.StartTimestamp, TimeSpace.EndTimestamp)
		}
		switch k {
		case 0:
			for _, sensorid := range TimeSpace.SensorIds {
				validMap1[sensorid] = TimeSpace
			}
		case 1:
			for _, sensorid := range TimeSpace.SensorIds {
				validMap2[sensorid] = TimeSpace
			}
		case 2:
			for _, sensorid := range TimeSpace.SensorIds {
				validMap3[sensorid] = TimeSpace
			}
		case 3:
			for _, sensorid := range TimeSpace.SensorIds {
				validMap4[sensorid] = TimeSpace
			}
		case 4:
			for _, sensorid := range TimeSpace.SensorIds {
				validMap5[sensorid] = TimeSpace
			}
		}
		TimeSpace.IDNo = int64(k + 1)
	}

	for sensorid, timesp := range validMap1 {
		if value, ok := validMap2[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap3[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap4[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap5[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
	}

	for sensorid, timesp := range validMap2 {
		if value, ok := validMap1[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap3[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap4[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap5[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
	}
	for sensorid, timesp := range validMap3 {
		if value, ok := validMap2[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap1[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap4[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap5[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
	}
	for sensorid, timesp := range validMap4 {
		if value, ok := validMap2[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap3[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap1[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap5[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
	}
	for sensorid, timesp := range validMap5 {
		if value, ok := validMap2[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap3[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap4[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
		if value, ok := validMap1[sensorid]; ok {
			if timesp.StartTimestamp < value.EndTimestamp && timesp.EndTimestamp > value.StartTimestamp {
				return NewErrorf(ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition, "time and spaces superposition")
			}
		}
	}

	return nil
}

func (ac *AreaCollisionTaskModel) validOther() error {
	if strings.EqualFold(ac.TaskName, "") {
		return NewErrorf(ErrorCodeAthenaCollisionModelCreateTaskEmptyTaskNameError, "illegal: task name should not be empty")
	}

	lenTaskName := utf8.RuneCountInString(ac.TaskName)
	if lenTaskName > CollisionCreateTaskMaxTaskNameLen { // 任务最大名称
		return NewErrorf(ErrorCodeAthenaCollisionModelCreateTaskTaskNameTooLongError, "illegal: len(task name): %d should not > %d", lenTaskName, CollisionCreateTaskMaxTaskNameLen)
	}

	return nil
}

//接收前端的查询请求
type QueryAreaCollisionTaskListReq struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int // 任务类型(人员频次分析、车辆频次分析)
	CurrentUser      *User // loki内部使用
}

//返回
type QueryAreaCollisionTaskListResp struct {
	TaskID                  string
	TaskName                string  // 任务名称
	ErrMsg                  string  // 任务错误原因, 是额外透传的信息
	CollisionProcessPercent float64 // 处理进度 = ProcessedSensorCnt/TotalSensorCnt
	CreateTs                int64   // 任务创建时间
	AthenaTaskStatus        int     //任务状态
	ResultCount             int     //结果数量
	AreaCollisionTaskType   int     //任务类型
	StartAndEndTimestamps   [][]int64
}

type QueryAreaCollisionTaskListResps []*QueryAreaCollisionTaskListResp

func NewAreaCollisionTaskListQueryResps() QueryAreaCollisionTaskListResps {
	return make(QueryAreaCollisionTaskListResps, 0)
}
func (resps QueryAreaCollisionTaskListResps) Len() int {
	return len(resps)
}

// CreateTs 降序
func (resps QueryAreaCollisionTaskListResps) Less(i, j int) bool {
	return resps[i].CreateTs > resps[j].CreateTs
}

func (resps QueryAreaCollisionTaskListResps) Swap(i, j int) {
	resps[i], resps[j] = resps[j], resps[i]
}

func (query *QueryAreaCollisionTaskListReq) FilterByquencyType(t int) *QueryAreaCollisionTaskListReq {
	filteredTypes := make([]int, 0)
	for _, tp := range query.AthenaTaskType {
		if tp == t {
			filteredTypes = append(filteredTypes, t)
		}
	}

	filtered := &QueryAreaCollisionTaskListReq{
		AthenaTaskStatus: query.AthenaTaskStatus,
		AthenaTaskType:   filteredTypes,
		CurrentUser:      query.CurrentUser,
	}
	return filtered
}

// 批量删频次任务的响应
type BatchDelCollisionTasksResp struct {
	IsAllTasksDeleteSucceed bool           // 是否全部task删除成功
	DetailInfo              ResultMapError // 详情, key 为 taskID, 值为 err 原因
}

func NewBatchDelCollisionTasksResp() *BatchDelCollisionTasksResp {
	return &BatchDelCollisionTasksResp{
		IsAllTasksDeleteSucceed: true,
		DetailInfo:              make(map[string]error),
	}
}
