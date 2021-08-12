package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type PersonCollisionTask struct {
	Uts       time.Time            `xorm:"uts not null default 'now()' DATETIME"`
	Ts        int64                `xorm:"ts not null default 0 BIGINT"`
	TaskID    string               `xorm:"task_id pk not null VARCHAR(1024) default ''"`
	Sensors   []*models.MinSensors `xorm:"json"`
	AreaTimes models.TimeSpaces    `xorm:"json"`
}

func (t *PersonCollisionTask) TableName() string {
	return TableNamePersonCollisionTask
}

func PersonCollisionTaskModelToEntity(model *models.PersonCollisionTaskModel) *PersonCollisionTask {
	return &PersonCollisionTask{
		Uts:       model.Uts,
		Ts:        model.Ts,
		TaskID:    model.TaskID,
		Sensors:   model.Sensors,
		AreaTimes: model.TimeSpace,
	}
}
