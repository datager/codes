package entities

import (
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils"
	"time"
)

type PlatformSyncPVG struct {
	Uts          time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts           int64     `xorm:"ts not null default 0 BIGINT"`
	DeviceID     string    `xorm:"device_id not null pk index VARCHAR(1024)"`
	SensorID     string    `xorm:"sensor_id not null VARCHAR(1024)"`
	URL          string    `xorm:"url not null VARCHAR(1024)"`
	Name         string    `xorm:"name not null default '' VARCHAR(1024)"`
	AVType       int64     `xorm:"av_type not null default 0 INT"`
	Host         string    `xorm:"host not null default '' VARCHAR(1024)"`
	HostTitle    string    `xorm:"host_title not null default '' VARCHAR(1024)"`
	HostType     int64     `xorm:"host_type not null default 0 INT"`
	IsCurNode    int64     `xorm:"is_cur_node not null default 0 INT"`
	Level        int64     `xorm:"level not null default 0 INT"`
	Org          string    `xorm:"org not null default '' VARCHAR(1024)"`
	OrgPath      string    `xorm:"org_path not null default '' VARCHAR(1024)"`
	Path         string    `xorm:"path not null default '' VARCHAR(1024)"`
	Title        string    `xorm:"title not null default '' VARCHAR(1024)"`
	Longitude    float64   `xorm:"longitude default '-200'::real REAL"`
	Latitude     float64   `xorm:"latitude default '-200'::real REAL"`
	IsDistribute int       `xorm:"is_distribute not null default 0 SMALLINT"`
}

func (*PlatformSyncPVG) TableName() string {
	return TableNamePlatformSyncPVGEntity
}

func (pst *PlatformSyncPVG) ToSensorEntity(orgID string) *Sensors {
	return &Sensors{
		Uts:                 pst.Uts,
		Ts:                  utils.GetNowTs(),
		OrgId:               orgID,
		SnId:                pst.DeviceID,
		SensorID:            pst.SensorID,
		SensorName:          pst.Title,
		Longitude:           pst.Longitude,
		Latitude:            pst.Latitude,
		Url:                 pst.URL,
		Status:              models.SensorStatusCreated,
		Type:                models.SensorType_Sensor_Type_Netposa_PVG,
		OuterPlatformStatus: models.OuterPlatformStatusNormal,
	}
}

func (pst *PlatformSyncPVG) ToModel() *models.PlatformSyncPVG {
	return &models.PlatformSyncPVG{
		Uts:          pst.Uts,
		Ts:           pst.Ts,
		DeviceID:     pst.DeviceID,
		DeviceName:   pst.Title,
		SensorID:     pst.SensorID,
		URL:          pst.URL,
		Name:         pst.Name,
		AVType:       pst.AVType,
		Host:         pst.Host,
		HostTitle:    pst.HostTitle,
		HostType:     pst.HostType,
		IsCurNode:    pst.IsCurNode,
		Level:        pst.Level,
		Org:          pst.Org,
		OrgPath:      pst.OrgPath,
		Path:         pst.Path,
		Title:        pst.Title,
		Longitude:    pst.Longitude,
		Latitude:     pst.Latitude,
		IsDistribute: pst.IsDistribute,
	}
}
