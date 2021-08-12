package entities

import "codes/gocodes/dg/models"

type VehiclePathAnalyzeSummary struct {
	TaskID     string `xorm:"task_id  not null VARCHAR(1024) default ''"`
	VehicleVid string `xorm:"vehicle_vid" not null VARCHAR(36) default ''`
	PlateText  string `xorm:"plate_text" not null VARCHAR(1024) default ''`
	ReidCount  int    `xorm:"reid_count" bigint not null default 0`
}

func (t *VehiclePathAnalyzeSummary) TableName() string {
	return TableNameVehiclePathAnalyzeSummary
}

func (t *VehiclePathAnalyzeSummary) ToModel() *models.PathAnalyzeSummaryItem {

	return &models.PathAnalyzeSummaryItem{
		Plate:  t.PlateText,
		Count:  int64(t.ReidCount),
		VID:    t.VehicleVid,
		TaskID: t.TaskID,
	}
}
