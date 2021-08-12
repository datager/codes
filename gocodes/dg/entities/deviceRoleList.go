package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type DeviceRoleList struct {
	Uts                time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts                 int64     `xorm:"not null BIGINT"`
	DeviceSensorPairId string    `xorm:"not null pk VARCHAR(1024)"`
	DeviceRoleId       string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	SensorId           string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	AuthFlag           string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Status             int       `xorm:"not null default 1 SMALLINT"`
}

type DeviceRoleListExtend struct {
	DeviceRoleList `xorm:"extends"`
	Sensor         Sensors `xorm:"extends"`
}

func (DeviceRoleListExtend) TableName() string {
	return TableNameDeviceRoleList
}

func (entity *DeviceRoleListExtend) ToDeviceRoleSensorModel() *models.DeviceRoleSensor {
	return &models.DeviceRoleSensor{
		SensorId:   entity.Sensor.SensorID,
		SensorName: entity.Sensor.SensorName,
	}
}

func DeviceRoleSensorModelToDeviceRoleListEntity(model *models.DeviceRoleSensor) *DeviceRoleList {
	return &DeviceRoleList{
		DeviceSensorPairId: model.DeviceSensorPairId,
		SensorId:           model.SensorId,
		AuthFlag:           model.AuthFlag,
		Status:             model.Status,
	}
}

func DeviceRoleListExtendsToSensorIDs(listExtends []*DeviceRoleListExtend) []string {
	sensorIDs := make([]string, 0)
	for _, le := range listExtends {
		sensorIDs = append(sensorIDs, le.ToDeviceRoleSensorModel().ToSensorID())
	}
	return sensorIDs
}
