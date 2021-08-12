package entities

import (
	"codes/gocodes/dg/models"
)

type RepoDailyPicCount struct {
	Id       string `xorm:"not null pk index VARCHAR(1024)"`
	RepoId   string `xorm:"not null default  '''::character varying' VARCHAR(1024)"`
	PicCount int64  `xorm:"not null default 0 BIGINT"`
	Ts       int64  `xorm:"not null default 0 BIGINT"`
}

func (this *RepoDailyPicCount) ToModel() *models.RepoDailyPicCount {
	return &models.RepoDailyPicCount{
		Id:       this.Id,
		RepoId:   this.RepoId,
		PicCount: this.PicCount,
		Ts:       this.Ts,
	}
}

func RepoDailyPicCountModelToEntity(model *models.RepoDailyPicCount) *RepoDailyPicCount {
	return &RepoDailyPicCount{
		RepoId:   model.RepoId,
		PicCount: model.PicCount,
		Ts:       model.Ts,
	}
}
