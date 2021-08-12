package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type FileCountStatistics struct {
	ID    int64     `xorm:"id serial pk not null"`
	Ts    int64     `xorm:"not null default 0 BIGINT"`
	Uts   time.Time `xorm:"not null default 'now()' DATETIME"`
	Type  int       `xorm:"not null default 0 SMALLINT"`
	Count int64     `xorm:"not null default 0 BIGINT"`
}

func (f FileCountStatistics) TableName() string {
	return TableNameFileCountStatistics
}

func (f *FileCountStatistics) ToModel() *models.FileCountStatistics {
	if f == nil {
		return nil
	}

	return &models.FileCountStatistics{
		ID:    f.ID,
		Ts:    f.Ts,
		Uts:   f.Uts,
		Type:  f.Type,
		Count: f.Count,
	}
}

func FileCountStatisticsModelToEntity(model *models.FileCountStatistics) *FileCountStatistics {
	if model == nil {
		return nil
	}

	return &FileCountStatistics{
		ID:    model.ID,
		Ts:    model.Ts,
		Uts:   model.Uts,
		Type:  model.Type,
		Count: model.Count,
	}
}
