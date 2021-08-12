package entities

import (
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils"
	"time"
)

type VehicleHiddenTask struct {
	Uts      time.Time              `xorm:"uts not null default 'now()' DATETIME"`
	Ts       int64                  `xorm:"ts not null default 0 BIGINT"`
	TaskID   string                 `xorm:"task_id pk not null VARCHAR(1024) default ''"`
	TaskName string                 `xorm:"task_name"`
	Sensors  models.MiniSensorSlice `xorm:"sensors json"`
	DayTime  []int64                `xorm:"day_time"`
}

func (t *VehicleHiddenTask) TableName() string {
	return TableNameVehicleHiddenTask
}

func VehicleHiddenTaskModelToEntity(model *models.HiddenTaskUpdateStatusRequest, minSensors models.MiniSensorSlice) *VehicleHiddenTask {
	return &VehicleHiddenTask{
		Uts:      time.Now(),
		Ts:       utils.GetNowTs(),
		TaskID:   model.TaskID,
		TaskName: model.TaskName,
		DayTime:  model.StartAndEndTimestamp,
		Sensors:  minSensors,
	}
}

func (t *VehicleHiddenTask) ToModel() *models.HiddenTaskUpdateStatusRequest {
	return &models.HiddenTaskUpdateStatusRequest{
		TaskID:               t.TaskID,
		TaskName:             t.TaskName,
		Sensors:              t.Sensors,
		StartAndEndTimestamp: t.DayTime,
	}
}
