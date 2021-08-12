package entities

import (
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils"
	"fmt"
	"time"
)

type PlatformSyncGB struct {
	Uts          time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts           int64     `xorm:"ts not null default 0 BIGINT"`
	DeviceID     string    `xorm:"device_id not null pk index VARCHAR(1024)"`
	SensorID     string    `xorm:"sensor_id not null VARCHAR(1024)"`
	Online       int       `xorm:"online not null default 1 SMALLINT"`
	ParentID     string    `xorm:"parent_id not null default '' VARCHAR(1024)"`
	Name         string    `xorm:"name not null default '' VARCHAR(1024)"`
	Longitude    float64   `xorm:"longitude default '-200'::real REAL"`
	Latitude     float64   `xorm:"latitude default '-200'::real REAL"`
	Manufacturer string    `xorm:"manufacturer not null default '' VARCHAR(1024)"`
	Model        string    `xorm:"model not null default '' VARCHAR(1024)"`
	Owner        string    `xorm:"owner not null default '' VARCHAR(1024)"`
	CivilCode    string    `xorm:"civil_code not null default '' VARCHAR(1024)"`
	Address      string    `xorm:"address not null default '' VARCHAR(1024)"`
	IPAddress    string    `xorm:"ip_address not null default '' VARCHAR(1024)"`
	Port         int       `xorm:"port not null default 0 SMALLINT"`
	Parental     int       `xorm:"parental not null default 0 SMALLINT"`
	Secrecy      int       `xorm:"secrecy not null default 0 SMALLINT"`
	IsDistribute int       `xorm:"is_distribute not null default 0 SMALLINT"`
}

func (*PlatformSyncGB) TableName() string {
	return TableNamePlatformSyncGBEntity
}

func (pst *PlatformSyncGB) ToSensorEntity(orgID, ip string, port int) *Sensors {
	return &Sensors{
		Uts:                 pst.Uts,
		Ts:                  utils.GetNowTs(),
		OrgId:               orgID,
		SnId:                pst.DeviceID,
		SensorID:            pst.SensorID,
		SensorName:          pst.Name,
		Longitude:           pst.Longitude,
		Latitude:            pst.Latitude,
		Url:                 fmt.Sprintf("gb28181://username:password@%s:%d/%s", ip, port, pst.DeviceID),
		Status:              models.SensorStatusCreated,
		Type:                models.SensorType_Sensor_Type_GB28181,
		OuterPlatformStatus: models.OuterPlatformStatusNormal,
	}
}

func (pst *PlatformSyncGB) ToModel() *models.PlatformSyncGB {
	return &models.PlatformSyncGB{
		Uts:          pst.Uts,
		Ts:           pst.Ts,
		DeviceID:     pst.DeviceID,
		DeviceName:   pst.Name,
		SensorID:     pst.SensorID,
		Online:       pst.Online,
		ParentID:     pst.ParentID,
		Name:         pst.Name,
		Longitude:    pst.Longitude,
		Latitude:     pst.Latitude,
		Manufacturer: pst.Manufacturer,
		Model:        pst.Model,
		Owner:        pst.Owner,
		CivilCode:    pst.CivilCode,
		Address:      pst.Address,
		IPAddress:    pst.IPAddress,
		Port:         pst.Port,
		Parental:     pst.Parental,
		Secrecy:      pst.Secrecy,
		IsDistribute: pst.IsDistribute,
	}
}
