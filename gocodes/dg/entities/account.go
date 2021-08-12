package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type Account struct {
	Uts           time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts            int64     `xorm:"not null default 0 BIGINT"`
	UserId        string    `xorm:"not null pk index VARCHAR(1024)"`
	UserName      string    `xorm:"not null default '''::character varying' index unique VARCHAR(1024)"`
	UserPasswd    string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	OrgId         string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	FuncRoleId    string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	DeviceRoleId  string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	IsValid       bool      `xorm:"not null default true BOOL"`
	RealName      string    `xorm:"default '''::character varying' VARCHAR(1024)"`
	Comment       string    `xorm:"default '''::character varying' VARCHAR"`
	Status        int       `xorm:"not null default 1 SMALLINT"`
	SecurityToken string    `xorm:"not null index unique VARCHAR(1024)"`
}

type AccountExtend struct {
	Account    `xorm:"extends"`
	Org        OrgStructure `xorm:"extends"`
	FuncRole   FuncRole     `xorm:"extends"`
	DeviceRole DeviceRole   `xorm:"extends"`
}

func (AccountExtend) TableName() string {
	return TableNameAccount
}

func (entity *Account) ToModel() *models.User {
	return &models.User{
		Ts:       entity.Ts,
		UserId:   entity.UserId,
		UserName: entity.UserName,
		// ignore UserPasswd
		OrgId:         entity.OrgId,
		IsValid:       entity.IsValid,
		RealName:      entity.RealName,
		Comment:       entity.Comment,
		Status:        entity.Status,
		SecurityToken: entity.SecurityToken,
	}
}

func (entity *AccountExtend) ToModel() *models.User {
	if entity == nil {
		return nil
	}
	user := entity.Account.ToModel()
	user.OrgName = entity.Org.OrgName
	user.UserFuncRole = entity.FuncRole.ToModel()
	user.UserFuncRole.OrgName = entity.Org.OrgName
	user.UserDeviceRole = entity.DeviceRole.ToModel()
	user.UserDeviceRole.OrgName = entity.Org.OrgName
	return user
}

func UserModelToEntity(model *models.User) *Account {
	return &Account{
		Ts:       model.Ts,
		UserId:   model.UserId,
		UserName: model.UserName,
		// ignore UserPassword
		OrgId:         model.OrgId,
		IsValid:       model.IsValid,
		RealName:      model.RealName,
		Comment:       model.Comment,
		Status:        model.Status,
		SecurityToken: model.SecurityToken,
		FuncRoleId:    model.UserFuncRole.FuncRoleId,
		DeviceRoleId:  model.UserDeviceRole.DeviceRoleId,
	}
}

type UserCount struct {
	Count        int    `xorm:"count"`
	FuncRoleID   string `xorm:"func_role_id"`
	DeviceRoleID string `xorm:"device_role_id"`
}
