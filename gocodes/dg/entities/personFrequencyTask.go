package entities

import (
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/slice"
)

type PersonFrequencyTask struct {
	Uts            time.Time      `xorm:"uts not null default 'now()' DATETIME"`
	Ts             int64          `xorm:"ts not null default 0 BIGINT"`
	TaskID         string         `xorm:"task_id pk not null VARCHAR(1024) default ''"`
	SensorIDs      string         `xorm:"sensor_ids not null text default ''"`
	Sensors        models.Sensors `xorm:"json"`
	TimesThreshold int64          `xorm:"times_threshold not null default 0 BIGINT"`
	IDNo           string         `xorm:"id_no not null VARCHAR(1024) default ''"`
}

func (t *PersonFrequencyTask) TableName() string {
	return TableNameVehicleFrequencyTask
}

func (t *PersonFrequencyTask) ToModel() *models.PersonFrequencyTask {
	return &models.PersonFrequencyTask{
		Uts:            t.Uts,
		Ts:             t.Ts,
		TaskID:         t.TaskID,
		SensorIDs:      slice.ParseIDsByRemoveComma(t.SensorIDs),
		Sensors:        t.Sensors,
		TimesThreshold: t.TimesThreshold,
		IDNo:           t.IDNo,
	}
}

func PersonFrequencyTaskModelToEntity(model *models.PersonFrequencyTask) *PersonFrequencyTask {
	return &PersonFrequencyTask{
		Uts:            model.Uts,
		Ts:             model.Ts,
		TaskID:         model.TaskID,
		SensorIDs:      slice.FormatIDsByAddComma(model.SensorIDs),
		Sensors:        model.Sensors,
		TimesThreshold: model.TimesThreshold,
		IDNo:           model.IDNo,
	}
}

func PersonFrequencyTaskEntitiesToModels(entity []*PersonFrequencyTask) []*models.PersonFrequencyTask {
	taskModels := make([]*models.PersonFrequencyTask, 0)
	for _, e := range entity {
		taskModels = append(taskModels, e.ToModel())
	}
	return taskModels
}
