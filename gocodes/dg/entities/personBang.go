package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type PersonBang struct {
	Uts         time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts          int64     `xorm:"not null default 0 BIGINT"`
	Id          string    `xorm:"not null pk VARCHAR(1024)"`
	Vid         string    `xorm:"not null VARCHAR(1024)"`
	RepoId      string    `xorm:"not null index VARCHAR(1024)"`
	RepoName    string    `xorm:"not null VARCHAR(1024)"`
	CivilAttrId string    `xorm:"not null VARCHAR(1024)"`
	IdNo        string    `xorm:"VARCHAR(1024)"`
	MarkImg     string    `xorm:"not null VARCHAR(1024)"`
	Confidence  float32   `xorm:"not null default 0::real REAL"`
}

func (entity *PersonBang) ToModel() *models.PersonBang {
	if entity == nil {
		return nil
	}
	return &models.PersonBang{
		Ts:          entity.Ts,
		Vid:         entity.Vid,
		Id:          entity.Id,
		RepoId:      entity.RepoId,
		RepoName:    entity.RepoName,
		CivilAttrId: entity.CivilAttrId,
		IdNo:        entity.IdNo,
		MarkImg:     entity.MarkImg,
		Confidence:  entity.Confidence,
	}
}
