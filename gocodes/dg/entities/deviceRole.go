package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type DeviceRole struct {
	Uts            time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts             int64     `xorm:"not null default 0 BIGINT"`
	DeviceRoleId   string    `xorm:"not null pk index VARCHAR(1024)"`
	DeviceRoleName string    `xorm:"default '''::character varying' index VARCHAR(1024)"`
	OrgId          string    `xorm:"not null VARCHAR(1024)"`
	DeviceRoleType int       `xorm:"not null default 0 SMALLINT"`
	Comment        string    `xorm:"default '''::character varying' VARCHAR"`
	Status         int       `xorm:"not null default 1 SMALLINT"`
	UserCount      int       `xorm:"<- INT"` // readonly
}

type DeviceRoleExtend struct {
	DeviceRole      `xorm:"extends"`
	DeviceRoleList  DeviceRoleList          `xorm:"extends"`
	Sensor          Sensors                 `xorm:"extends"` // DeviceRoleList.Sensor
	DeviceRoleLists []*DeviceRoleListExtend // since xorm dose not support one-to-many, we will do it manually in DeviceRoleRepository
}

func (DeviceRoleExtend) TableName() string {
	return TableNameDeviceRole
}

func (entity *DeviceRole) ToModel() *models.DeviceRole {
	if entity == nil {
		return nil
	}
	return &models.DeviceRole{
		Ts:             entity.Ts,
		DeviceRoleId:   entity.DeviceRoleId,
		DeviceRoleName: entity.DeviceRoleName,
		OrgId:          entity.OrgId,
		DeviceRoleType: entity.DeviceRoleType,
		UserCount:      entity.UserCount,
		Comment:        entity.Comment,
	}
}

func DeviceRoleModelToEntity(model *models.DeviceRole) *DeviceRole {
	return &DeviceRole{
		DeviceRoleId:   model.DeviceRoleId,
		DeviceRoleName: model.DeviceRoleName,
		OrgId:          model.OrgId,
		DeviceRoleType: model.DeviceRoleType,
		Comment:        model.Comment,
		Status:         model.Status,
	}
}
