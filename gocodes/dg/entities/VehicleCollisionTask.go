package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type VehicleCollisionTask struct {
	Uts       time.Time            `xorm:"uts not null default 'now()' DATETIME"`
	Ts        int64                `xorm:"ts not null default 0 BIGINT"`
	TaskID    string               `xorm:"task_id pk not null VARCHAR(1024) default ''"`
	Sensors   []*models.MinSensors `xorm:"json"`
	AreaTimes models.TimeSpaces    `xorm:"json"`
}

func (t *VehicleCollisionTask) TableName() string {
	return TableNameVehicleCollisionTask
}

func VehicleCollisionTaskModelToEntity(model *models.VehicleCollisionTaskModel) *VehicleCollisionTask {
	return &VehicleCollisionTask{
		Uts:       model.Uts,
		Ts:        model.Ts,
		TaskID:    model.TaskID,
		Sensors:   model.Sensors,
		AreaTimes: model.TimeSpace,
	}
}
