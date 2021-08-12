package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type DeleteTask struct {
	TaskID    string    `xorm:"task_id not null pk index varchar(1024)"`
	TaskType  string    `xorm:"not null varchar(20)"`
	Status    int       `xorm:"not null default 0"`
	DeleteID  string    `xorm:"delete_id not null varchar(1024)"`
	Ts        int64     `xorm:"not null default 0 bigint"`
	EndTime   int64     `xorm:"not null default 0 bigint"`
	Note      string    `xorm:"not null varchar(1024)"`
	UpdatedAt time.Time `xorm:"not null default 'now()' datetime"`
}

func (dt *DeleteTask) ToModel() *models.DeleteTask {
	return &models.DeleteTask{
		TaskID:    dt.TaskID,
		TaskType:  dt.TaskType,
		Status:    dt.Status,
		DeleteID:  dt.DeleteID,
		Ts:        dt.Ts,
		EndTime:   dt.EndTime,
		Note:      dt.Note,
		UpdatedAt: dt.UpdatedAt,
	}
}

func DeleteTaskModelToEntity(model *models.DeleteTask) *DeleteTask {
	if model == nil {
		return nil
	}

	return &DeleteTask{
		TaskID:    model.TaskID,
		TaskType:  model.TaskType,
		Status:    model.Status,
		DeleteID:  model.DeleteID,
		Ts:        model.Ts,
		EndTime:   model.EndTime,
		Note:      model.Note,
		UpdatedAt: model.UpdatedAt,
	}
}
