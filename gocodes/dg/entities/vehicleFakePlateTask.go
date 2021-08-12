package entities

import (
	"codes/gocodes/dg/utils"
	"time"

	"codes/gocodes/dg/models"
)

type VehicleFakePlateTask struct {
	Uts            time.Time              `xorm:"uts"`
	Ts             int64                  `xorm:"ts"`
	TaskID         string                 `xorm:"task_id"`
	TaskName       string                 `xorm:"task_name"`
	Sensors        models.MiniSensorSlice `xorm:"sensors json"`
	StartTimestamp int64                  `xorm:"start_timestamp"`
	EndTimestamp   int64                  `xorm:"end_timestamp"`
	PlateText      string                 `xorm:"plate_text"`
}

func (*VehicleFakePlateTask) TableName() string {
	return TableNameVehicleFakePlateTask
}

func NewVehicleFakePlateTask(request *models.VehicleFakePlateCreateRequest, minSensors models.MiniSensorSlice) *VehicleFakePlateTask {
	return &VehicleFakePlateTask{
		Uts:            time.Now(),
		Ts:             utils.GetNowTs(),
		TaskID:         request.TaskID,
		TaskName:       request.TaskName,
		Sensors:        minSensors,
		StartTimestamp: request.StartTimestamp,
		EndTimestamp:   request.EndTimestamp,
		PlateText:      request.PlateText,
	}
}

func (s *VehicleFakePlateTask) ToModel() *models.VehicleFakePlateCreateRequest {
	return &models.VehicleFakePlateCreateRequest{
		FakePlateCreateRequest: models.FakePlateCreateRequest{
			TaskID:   s.TaskID,
			TaskName: s.TaskName,
			Sensors:  s.Sensors,
			StartAndEndTimestamp: models.StartAndEndTimestamp{
				StartTimestamp: s.StartTimestamp,
				EndTimestamp:   s.EndTimestamp,
			},
		},
		PlateText: s.PlateText,
	}
}
