package entities

import (
	"codes/gocodes/dg/models"
)

type VehicleDayInNightOutSummary struct {
	TaskID    string `xorm:"task_id  not null VARCHAR(1024) default ''"`
	Vid       string `xorm:"vid" not null VARCHAR(1024) default ''`
	PlateText string `xorm:"plate_text" not null VARCHAR(1024) default ''`
	ReidCount int    `xorm:"reid_count" bigint not null default 0`
}

func (t *VehicleDayInNightOutSummary) TableName() string {
	return TableNameVehicleDayInNightOutSummary
}

func (t *VehicleDayInNightOutSummary) SummerEntityToModel() *models.VehicleDayInNightOutTaskModel {

	return &models.VehicleDayInNightOutTaskModel{
		Plate:  t.PlateText,
		Count:  t.ReidCount,
		VID:    t.Vid,
		TaskID: t.TaskID,
	}
}
