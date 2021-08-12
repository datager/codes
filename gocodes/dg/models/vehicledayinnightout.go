package models

import (
	"codes/gocodes/dg/utils/config"
	"strings"
	"time"
	"unicode/utf8"
)

//新建任务
type VehicleDayInNightOutTaskModel struct {
	AthenaTask
	Uts                  time.Time
	Ts                   int64
	UTCOffset            int         //时区
	TaskID               string      `uri:"taskid"` //任务ID
	TaskName             string      //任务名字
	SensorIDs            []string    //设备数组
	Plate                string      //车牌号
	StartAndEndTimestamp []int64     //天的起止时间
	DayStartAndEndTime   []int64     //白天的时间段
	NightStartAndEndTime []int64     //夜晚的时间段
	CurrentUser          *User       //loki自用
	UpdateType           int         //更新的类型 1停止 2开始
	Sensors              *MinSensors //设备信息
	Count                int         //个数
	VID                  string
	Pagination
}

func (createReq *VehicleDayInNightOutTaskModel) Valid() error {
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

func (createReq *VehicleDayInNightOutTaskModel) validSensor() error {
	if len(createReq.SensorIDs) == 0 {
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskCouldNotHaveEmptySensorsError, "len(sensorIDs) == 0")
	}

	maxTaskSensorCountLimit := config.GetConfig().GetIntOrDefault("api.dayinnightout.sensorCountLimit", VehicleDayInNightOutCreateTaskMaxSensorCount)
	if len(createReq.SensorIDs) > maxTaskSensorCountLimit {
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskSensorCountReachMaxLimitError, "len(sensorIDs): %d > maxSensorCountLimit: %d", len(createReq.SensorIDs), maxTaskSensorCountLimit)
	}
	return nil
}

func (createReq *VehicleDayInNightOutTaskModel) validTs() error {
	if len(createReq.StartAndEndTimestamp) != 2 && len(createReq.DayStartAndEndTime) != 2 && len(createReq.NightStartAndEndTime) != 2 {
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskIllegalTsError, "len(req)!=2")
	}

	if createReq.StartAndEndTimestamp[0] <= 0 || createReq.StartAndEndTimestamp[1] <= 0 {
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskIllegalTsError, "startts <= 0 || endts <= 0")
	}

	if createReq.StartAndEndTimestamp[0] > createReq.StartAndEndTimestamp[1] {
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskIllegalTsError, "startts > endts")
	}

	if (createReq.DayStartAndEndTime[0] <= createReq.NightStartAndEndTime[1] && createReq.NightStartAndEndTime[1] <= createReq.DayStartAndEndTime[1]) ||
		(createReq.DayStartAndEndTime[0] <= createReq.NightStartAndEndTime[0] && createReq.NightStartAndEndTime[0] <= createReq.DayStartAndEndTime[1]) {
		return NewErrorf(ErrorCodeAthenaDayInNightOutTaskCreateTaskTsSuperposition, "time superposition")
	}

	maxTaskTsDuration := config.GetConfig().GetIntOrDefault("api.dayinnightout.selectedDaysLimit", VehicleDayInNightOutTaskMaxDurationDaysBetweenStartTsAndEndTs)
	if (createReq.StartAndEndTimestamp[0] - createReq.StartAndEndTimestamp[1]) > (int64(maxTaskTsDuration) * 24 * int64(time.Hour) / 1e6) {
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError, "(endts - startts) > %d days, startts: %d, endts: %d", maxTaskTsDuration, createReq.StartAndEndTimestamp[0], createReq.StartAndEndTimestamp[1])
	}
	return nil
}

func (createReq *VehicleDayInNightOutTaskModel) validOther() error {
	if strings.EqualFold(createReq.TaskName, "") {
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskEmptyTaskNameError, "illegal: task name should not be empty")
	}

	lenTaskName := utf8.RuneCountInString(createReq.TaskName)
	if lenTaskName > DayInNightOutCreateTaskMaxTaskNameLen { // 任务最大名称
		return NewErrorf(ErrorCodeAthenaDayInNightOutModelCreateTaskTaskNameTooLongError, "illegal: len(task name): %d should not > %d", lenTaskName, DayInNightOutCreateTaskMaxTaskNameLen)
	}
	return nil
}

type VehicleDayInNightOutTaskDetail struct {
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
