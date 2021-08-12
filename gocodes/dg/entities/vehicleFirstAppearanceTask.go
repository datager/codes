package entities

import (
	"codes/gocodes/dg/utils"
	"time"

	"codes/gocodes/dg/models"
)

type VehicleFirstAppearanceTask struct {
	Uts              time.Time                                   `xorm:"uts"`
	Ts               int64                                       `xorm:"ts"`
	TaskID           string                                      `xorm:"task_id"`
	TaskName         string                                      `xorm:"task_name"`
	Sensors          models.MiniSensorSlice                      `xorm:"sensors json"`
	OtherCondition   models.VehicleFirstAppearanceOtherCondition `xorm:"other json"`
	StartTimestamp   int64                                       `xorm:"start_timestamp"`
	EndTimestamp     int64                                       `xorm:"end_timestamp"`
	BacktrackingDays int64                                       `xorm:"backtracking_days"`
}

func (*VehicleFirstAppearanceTask) TableName() string {
	return TableNameVehicleFirstAppearanceTask
}

func NewVehicleFirstAppearanceTask(request *models.VehicleFirstAppearanceCreateRequest, minSensors models.MiniSensorSlice) *VehicleFirstAppearanceTask {
	return &VehicleFirstAppearanceTask{
		Uts:              time.Now(),
		Ts:               utils.GetNowTs(),
		TaskID:           request.TaskID,
		TaskName:         request.TaskName,
		Sensors:          minSensors,
		OtherCondition:   request.VehicleFirstAppearanceOtherCondition,
		StartTimestamp:   request.StartTimestamp,
		EndTimestamp:     request.EndTimestamp,
		BacktrackingDays: request.BacktrackingDays,
	}
}

func (s *VehicleFirstAppearanceTask) ToModel() *models.VehicleFirstAppearanceCreateRequest {
	return &models.VehicleFirstAppearanceCreateRequest{
		FirstAppearanceCreateRequest: models.FirstAppearanceCreateRequest{
			TaskID:   s.TaskID,
			TaskName: s.TaskName,
			Sensors:  s.Sensors,
			StartAndEndTimestamp: models.StartAndEndTimestamp{
				StartTimestamp: s.StartTimestamp,
				EndTimestamp:   s.EndTimestamp,
			},
			BacktrackingDays: s.BacktrackingDays,
		},
		VehicleFirstAppearanceOtherCondition: s.OtherCondition,
	}
}
