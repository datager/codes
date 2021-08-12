package entities

import (
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils"
	"time"
)

type VehicleDayInNightOutTask struct {
	Uts        time.Time         `xorm:"uts not null default 'now()' DATETIME"`
	Ts         int64             `xorm:"ts not null default 0 BIGINT"`
	TaskID     string            `xorm:"task_id pk not null VARCHAR(1024) default ''"`
	TaskName   string            `xorm:"task_name"`
	Sensors    models.MinSensors `xorm:"sensors json"`
	PlateText  string            `xorm:"plate_text"`
	DayTime    []int64           `xorm:"day_time"`
	BrightTime []int64           `xorm:"bright_time"`
	NightTime  []int64           `xorm:"night_time"`
}

func (t *VehicleDayInNightOutTask) TableName() string {
	return TableNameVehicleDayInNightOutTask
}

func VehicleDayInNightOutTaskModelToEntity(model *models.VehicleDayInNightOutTaskModel, minSensors models.MinSensors) *VehicleDayInNightOutTask {
	return &VehicleDayInNightOutTask{
		Uts:        time.Now(),
		Ts:         utils.GetNowTs(),
		TaskID:     model.TaskID,
		TaskName:   model.TaskName,
		PlateText:  model.Plate,
		DayTime:    model.StartAndEndTimestamp,
		NightTime:  model.NightStartAndEndTime,
		BrightTime: model.DayStartAndEndTime,
		Sensors:    minSensors,
	}
}

func (t *VehicleDayInNightOutTask) EntiytToModel() *models.VehicleDayInNightOutTaskModel {
	return &models.VehicleDayInNightOutTaskModel{
		Uts:                  t.Uts,
		Ts:                   t.Ts,
		TaskID:               t.TaskID,
		TaskName:             t.TaskName,
		Plate:                t.PlateText,
		StartAndEndTimestamp: t.DayTime,
		NightStartAndEndTime: t.NightTime,
		DayStartAndEndTime:   t.BrightTime,
		Sensors:              &t.Sensors,
	}
}
