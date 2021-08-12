package entities

import (
	"codes/gocodes/dg/utils/slice"
	"time"

	"codes/gocodes/dg/models"
)

type VehicleFrequencyTask struct {
	Uts            time.Time      `xorm:"uts not null default 'now()' DATETIME"`
	Ts             int64          `xorm:"ts not null default 0 BIGINT"`
	TaskID         string         `xorm:"task_id pk not null VARCHAR(1024) default ''"`
	SensorIDs      string         `xorm:"sensor_ids not null text default ''"`
	Sensors        models.Sensors `xorm:"json"`
	TimesThreshold int64          `xorm:"times_threshold not null default 0 BIGINT"`
	Plate          string         `xorm:"plate not null VARCHAR(1024) default ''"`
}

func (t *VehicleFrequencyTask) TableName() string {
	return TableNameVehicleFrequencyTask
}

func (t *VehicleFrequencyTask) ToModel() *models.VehicleFrequencyTask {
	return &models.VehicleFrequencyTask{
		Uts:            t.Uts,
		Ts:             t.Ts,
		TaskID:         t.TaskID,
		SensorIDs:      slice.ParseIDsByRemoveComma(t.SensorIDs),
		Sensors:        t.Sensors,
		TimesThreshold: t.TimesThreshold,
		Plate:          t.Plate,
	}
}

func VehicleFrequencyTaskModelToEntity(model *models.VehicleFrequencyTask) *VehicleFrequencyTask {
	return &VehicleFrequencyTask{
		Uts:            model.Uts,
		Ts:             model.Ts,
		TaskID:         model.TaskID,
		SensorIDs:      slice.FormatIDsByAddComma(model.SensorIDs),
		Sensors:        model.Sensors,
		TimesThreshold: model.TimesThreshold,
		Plate:          model.Plate,
	}
}

func VehicleFrequencyTaskEntitiesToModels(entity []*VehicleFrequencyTask) []*models.VehicleFrequencyTask {
	taskModels := make([]*models.VehicleFrequencyTask, 0)
	for _, e := range entity {
		taskModels = append(taskModels, e.ToModel())
	}
	return taskModels
}
