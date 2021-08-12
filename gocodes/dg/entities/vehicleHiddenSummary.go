package entities

import "codes/gocodes/dg/models"

type VehicleHiddenSummary struct {
	TaskID     string `xorm:"task_id  not null VARCHAR(1024) default ''"`
	VehicleVid string `xorm:"vehicle_vid" not null VARCHAR(36) default ''`
	PlateText  string `xorm:"plate_text" not null VARCHAR(1024) default ''`
	ReidCount  int    `xorm:"reid_count" bigint not null default 0`
}

func (t *VehicleHiddenSummary) TableName() string {
	return TableNameVehicleHiddenSummary
}

func (t *VehicleHiddenSummary) ToModel() *models.VehicleHiddenSummaryItem {

	return &models.VehicleHiddenSummaryItem{
		Plate:  t.PlateText,
		Count:  int64(t.ReidCount),
		VID:    t.VehicleVid,
		TaskID: t.TaskID,
	}
}
