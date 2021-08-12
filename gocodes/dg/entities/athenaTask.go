package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type AthenaTask struct {
	Uts       time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts        int64     `xorm:"ts not null default 0 BIGINT"`
	TaskID    string    `xorm:"task_id not null pk VARCHAR(1024)"`
	Name      string    `xorm:"name not null VARCHAR(1024)"`
	UserID    string    `xorm:"user_id not null VARCHAR(1024)"`
	OrgID     string    `xorm:"org_id not null VARCHAR(1024)"`
	Type      int       `xorm:"type not null default 0 SMALLINT"`
	InnerType int       `xorm:"inner_type not null default 0 SMALLINT"`
	StartTS   int64     `xorm:"start_ts not null default 0 BIGINT"`
	EndTS     int64     `xorm:"end_ts not null default 0 BIGINT"`
}

func (*AthenaTask) TableName() string {
	return TableNameAthenaTaskEntity
}

func (at *AthenaTask) NeedUpdateEndTS() bool {
	return at.EndTS == 0
}

func (at *AthenaTask) ToModel() *models.AthenaTask {
	return &models.AthenaTask{
		Uts:       at.Uts,
		Ts:        at.Ts,
		TaskID:    at.TaskID,
		Name:      at.Name,
		UserID:    at.UserID,
		OrgID:     at.OrgID,
		Type:      at.Type,
		InnerType: at.InnerType,
		StartTS:   at.StartTS,
		EndTS:     at.EndTS,
	}
}

func AthenaTaskModelToEntity(model *models.AthenaTask) *AthenaTask {
	return &AthenaTask{
		Uts:     model.Uts,
		Ts:      model.Ts,
		Name:    model.Name,
		UserID:  model.UserID,
		OrgID:   model.OrgID,
		TaskID:  model.TaskID,
		Type:    model.Type,
		StartTS: model.StartTS,
		EndTS:   model.EndTS,
	}
}

func AthenaTaskEntitiesToModels(entity []*AthenaTask) []*models.AthenaTask {
	taskModels := make([]*models.AthenaTask, 0)
	for _, e := range entity {
		taskModels = append(taskModels, e.ToModel())
	}
	return taskModels
}
