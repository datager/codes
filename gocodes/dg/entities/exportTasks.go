package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type ExportTasks struct {
	Uts            time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts             int64     `xorm:"not null default 0 BIGINT"`
	TaskId         string    `xorm:"not null pk VARCHAR(1024)"`
	Type           int       `xorm:"not null SMALLINT"`
	FileNamePrefix string    `xorm:"not null"`
	Query          string    `xorm:"not null TEXT"`
	Total          int       `xorm:"not null default 0 INT"`
	Processed      int       `xorm:"not null default 0 INT"`
	TaskStatus     int       `xorm:"not null default 0 SMALLINT"`
	StartAt        int64     `xorm:"not null default 0 BIGINT"`
	UpdateAt       int64     `xorm:"not null default 0 BIGINT"`
	Status         int       `xorm:"not null default 1 SMALLINT"`
}

func (entity *ExportTasks) ToModel() *models.ExportTask {
	return &models.ExportTask{
		Id: entity.TaskId,
		Request: &models.ExportRequest{
			Type:           entity.Type,
			FileNamePrefix: entity.FileNamePrefix,
			Query:          entity.Query,
		},
		Response: &models.ExportResponse{
			Total:     entity.Total,
			Processed: entity.Processed,
			Status:    models.ExportTaskStatus(entity.TaskStatus),
			StartAt:   entity.StartAt,
			UpdateAt:  entity.UpdateAt,
		},
		Timestamp: entity.Ts,
	}
}
