package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type RepoCollision struct {
	Uts        time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts         int64     `xorm:"not null default 0 BIGINT"`
	TaskId     string    `xorm:"not null pk index VARCHAR(1024)"`
	TaskName   string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	UserId     string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	TargetRepo string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	SrcRepo    string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Threshold  float32   `xorm:"not null default 0 REAL"`
	Status     int       `xorm:"not null default 0 SMALLINT"`
}

func (entity *RepoCollision) ToModel() *models.TaskCollisionModel {
	if entity == nil {
		return nil
	}
	return &models.TaskCollisionModel{
		TaskCollision: models.TaskCollision{
			TaskId:     entity.TaskId,
			SrcRepo:    entity.SrcRepo,
			TargetRepo: entity.TargetRepo,
			TaskName:   entity.TaskName,
			Threshold:  entity.Threshold * 100,
		},
		Timestamp: entity.Ts,
		Status:    entity.Status,
	}
}
func RepoCollisionsModelToEntity(model *models.TaskCollisionModel) *RepoCollision {
	return &RepoCollision{
		Uts:        time.Now(),
		TaskId:     model.TaskId,
		TaskName:   model.TaskName,
		Threshold:  model.Threshold / 100,
		TargetRepo: model.TargetRepo,
		SrcRepo:    model.SrcRepo,
		Status:     model.Status,
	}
}
