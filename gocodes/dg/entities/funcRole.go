package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type FuncRole struct {
	Uts          time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts           int64     `xorm:"not null default 0 BIGINT"`
	FuncRoleId   string    `xorm:"not null pk index VARCHAR(1024)"`
	OrgId        string    `xorm:"not null default ''0000'::character varying' VARCHAR(1024)"`
	FuncRoleName string    `xorm:"default '''::character varying' index VARCHAR(1024)"`
	FuncAuthFlag string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	FuncRoleType int       `xorm:"not null default 0 SMALLINT"`
	Comment      string    `xorm:"default '''::character varying' VARCHAR"`
	Status       int       `xorm:"not null default 1 SMALLINT"`
	UserCount    int       `xorm:"<- INT"`     // readonly
	OrgName      string    `xorm:"<- VARCHAR"` // readonly
}

func (entity *FuncRole) ToModel() *models.FuncRole {
	if entity == nil {
		return nil
	}
	return &models.FuncRole{
		Ts:           entity.Ts,
		FuncRoleId:   entity.FuncRoleId,
		FuncRoleName: entity.FuncRoleName,
		OrgId:        entity.OrgId,
		FuncAuthFlag: entity.FuncAuthFlag,
		FuncRoleType: entity.FuncRoleType,
		Comment:      entity.Comment,
		UserCount:    entity.UserCount,
		OrgName:      entity.OrgName,
	}
}

func FuncRoleModelToEntity(model *models.FuncRole) *FuncRole {
	return &FuncRole{
		Ts:           model.Ts,
		FuncRoleId:   model.FuncRoleId,
		FuncRoleName: model.FuncRoleName,
		OrgId:        model.OrgId,
		FuncAuthFlag: model.FuncAuthFlag,
		FuncRoleType: model.FuncRoleType,
		Comment:      model.Comment,
	}
}
