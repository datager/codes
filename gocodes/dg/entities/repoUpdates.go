package entities

import "codes/gocodes/dg/models"

const (
	KindAdd = iota + 1
	KindDelete
)

type RepoUpdates struct {
	Id         string `xorm:"not null pk VARCHAR(1024)"`
	RepoId     string `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	UpdateKind int    `xorm:"not null default 0 SMALLINT"`
	Value      int    `xorm:"not null default 0 "`
	Ts         int64  `xorm:"not null default 0 BIGINT"`
}

func RepoUpdatesModelToEntity(model *models.RepoUpdates) *RepoUpdates {
	return &RepoUpdates{
		Id:         model.Id,
		RepoId:     model.RepoId,
		UpdateKind: model.Kind,
		Value:      model.Value,
		Ts:         model.Ts,
	}
}

func (this *RepoUpdates) ToModel() *models.RepoUpdates {
	return &models.RepoUpdates{
		Id:     this.Id,
		RepoId: this.RepoId,
		Kind:   this.UpdateKind,
		Value:  this.Value,
		Ts:     this.Ts,
	}
}
