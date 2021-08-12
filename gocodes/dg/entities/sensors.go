package entities

import (
	"codes/gocodes/dg/utils/redis"
	"fmt"
	"time"

	"codes/gocodes/dg/models"
)

const (
	sensorsCacheKey = "sensors-cache-key-sensorid:%v:v2"
)

type Sensors struct {
	Uts                 time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts                  int64     `xorm:"not null default 0 BIGINT"`
	SensorID            string    `json:"sensor_id" xorm:"sensor_id not null pk VARCHAR(1024)"`
	SensorIntID         int32     `xorm:"sensor_int_id autoincr"`
	OrgId               string    `json:"org_id" xorm:"not null default ''0000'::character varying' VARCHAR(1024)"`
	OrgCode             int64     `json:"org_code" xorm:"not null default 0 BIGINT"`
	SnId                string    `xorm:"not null unique VARCHAR(1024)"`
	SensorName          string    `xorm:"not null default '''::character varying' index VARCHAR(1024)"`
	Longitude           float64   `xorm:"default '-200'::real REAL"`
	Latitude            float64   `xorm:"default '-200'::real REAL"`
	Type                int       `xorm:"not null default 0 SMALLINT"`
	Ip                  string    `xorm:"default '''::character varying' VARCHAR"`
	Port                string    `xorm:"default '''::character varying' VARCHAR"`
	Url                 string    `xorm:"default '''::character varying' VARCHAR"`
	RenderedUrl         string    `xorm:"renderedurl default '''::character varying' VARCHAR"`
	RtmpUrl             string    `xorm:"rtmpurl default '''::character varying' VARCHAR"`
	FtpAddr             string    `xorm:"default '''::character varying' VARCHAR"`
	FtpDir              string    `xorm:"default '''::character varying' VARCHAR"`
	GateIP              string    `xorm:"gate_ip default '''::character varying' VARCHAR"`
	GateThreshold       float32   `xorm:"default 0 REAL"`
	Comment             string    `xorm:"default '''::character varying' VARCHAR"`
	AlgType             string    `xorm:"not null default ''EUR'::character varying' VARCHAR(1024)"`
	EurThreshold        float32   `xorm:"not null default 5.5 REAL"`
	CosThreshold        float32   `xorm:"not null default 0.8 REAL"`
	ProcessCapturedReid int       `xorm:"not null default 1 SMALLINT"`
	ProcessWarnedReid   int       `xorm:"not null default 1 SMALLINT"`
	WindowSize          int       `xorm:"not null default 20 SMALLINT"`
	OlympusId           string    `xorm:"not null VARCHAR(1024)"`
	Status              int       `xorm:"not null default 1 SMALLINT"`
	ConfigJson          string    `xorm:"VARCHAR"`
	MServerAddr         string    `xorm:"mserver_addr VARCHAR(1024)"`
	Speed               int       `xorm:"default 1 SMALLINT"`
	NvrAddr             string    `xorm:"default '''::character varying' VARCHAR"`
	FlatformSensor      int       `xorm:"not null default 1 SMALLINT"`
	OuterPlatformStatus int       `xorm:"not null default 0 SMALLINT"`
	AdditionalInfo      string    `xorm:"additional_info"`
}

type SensorsExtend struct {
	Sensors `xorm:"extends"`
	Org     OrgStructure `xorm:"extends"`
}

type FloraMinSensors struct {
	SensorID string `json:"sensor_id" xorm:"sensor_id not null pk VARCHAR(1024)"`
	OrgId    string `json:"org_id" xorm:"not null default ''0000'::character varying' VARCHAR(1024)"`
	OrgCode  int64  `json:"org_code" xorm:"not null default 0 BIGINT"`
}

func (entity *Sensors) ToModel() *models.Sensor {
	return &models.Sensor{
		Uts:                 entity.Uts,
		Timestamp:           entity.Ts,
		OrgId:               entity.OrgId,
		OrgCode:             entity.OrgCode,
		SerialID:            entity.SnId,
		SensorID:            entity.SensorID,
		SensorIntID:         entity.SensorIntID,
		SensorName:          entity.SensorName,
		Type:                entity.Type,
		Status:              entity.Status,
		Longitude:           entity.Longitude,
		Latitude:            entity.Latitude,
		Ip:                  entity.Ip,
		Port:                entity.Port,
		Url:                 entity.Url,
		RenderedUrl:         entity.RenderedUrl,
		RtmpUrl:             entity.RtmpUrl,
		FtpAddr:             entity.FtpAddr,
		FtpDir:              entity.FtpDir,
		GateIP:              entity.GateIP,
		Comment:             entity.Comment,
		ConfigJson:          entity.ConfigJson,
		OlympusId:           entity.OlympusId,
		GateThreshold:       entity.GateThreshold,
		MServerAddr:         entity.MServerAddr,
		AlgType:             entity.AlgType,
		EurThreshold:        entity.EurThreshold,
		CosThreshold:        entity.CosThreshold,
		ProcessCapturedReid: entity.ProcessCapturedReid,
		ProcessWarnedReid:   entity.ProcessWarnedReid,
		WindowSize:          entity.WindowSize,
		NvrAddr:             entity.NvrAddr,
		FlatformSensor:      entity.FlatformSensor,
		OuterPlatformStatus: entity.OuterPlatformStatus,
		AdditionalInfo:      entity.AdditionalInfo,
	}
}

func SensorModelToEntity(model *models.Sensor) *Sensors {
	return &Sensors{
		Uts:                 model.Uts,
		Ts:                  model.Timestamp,
		OrgId:               model.OrgId,
		OrgCode:             model.OrgCode,
		SnId:                model.SerialID,
		SensorID:            model.SensorID,
		SensorName:          model.SensorName,
		Type:                model.Type,
		Status:              int(model.Status),
		Longitude:           model.Longitude,
		Latitude:            model.Latitude,
		Ip:                  model.Ip,
		Port:                model.Port,
		Url:                 model.Url,
		RenderedUrl:         model.RenderedUrl,
		RtmpUrl:             model.RtmpUrl,
		FtpAddr:             model.FtpAddr,
		FtpDir:              model.FtpDir,
		GateIP:              model.GateIP,
		Comment:             model.Comment,
		ConfigJson:          model.ConfigJson,
		OlympusId:           model.OlympusId,
		GateThreshold:       model.GateThreshold,
		MServerAddr:         model.MServerAddr,
		AlgType:             model.AlgType,
		EurThreshold:        model.EurThreshold,
		CosThreshold:        model.CosThreshold,
		ProcessCapturedReid: model.ProcessCapturedReid,
		ProcessWarnedReid:   model.ProcessWarnedReid,
		WindowSize:          model.WindowSize,
		NvrAddr:             model.NvrAddr,
		FlatformSensor:      model.FlatformSensor,
		OuterPlatformStatus: model.OuterPlatformStatus,
		AdditionalInfo:      model.AdditionalInfo,
	}
}

func SensorEntitiesToModels(entities []*Sensors) []*models.Sensor {
	models := make([]*models.Sensor, 0)
	for _, r := range entities {
		models = append(models, r.ToModel())
	}
	return models
}

func SensorToSensorIDs(sensors []*Sensors) []string {
	sensorIDs := make([]string, 0)
	for _, s := range sensors {
		sensorIDs = append(sensorIDs, s.SensorID)
	}
	return sensorIDs
}

func (entity *Sensors) ToMinSensor() *models.MinSensor {
	return &models.MinSensor{
		SensorID:   entity.SensorID,
		SensorName: entity.SensorName,
		Latitude:   entity.Latitude,
		Longitude:  entity.Longitude,
	}
}

func (entity *Sensors) ToMiniSensor() *models.MiniSensor {
	return &models.MiniSensor{
		SensorID:   entity.SensorID,
		SensorName: entity.SensorName,
		Latitude:   entity.Latitude,
		Longitude:  entity.Longitude,
	}
}

func (entity *Sensors) ToInfo() *models.SensorInfo {
	if entity == nil {
		return nil
	}

	return &models.SensorInfo{
		SensorID:       entity.SensorID,
		SensorName:     entity.SensorName,
		SensorIntID:    entity.SensorIntID,
		AdditionalInfo: entity.AdditionalInfo,
		Longitude:      entity.Longitude,
		Latitude:       entity.Latitude,
	}
}

func (entity *Sensors) TableName() string {
	return TableNameSensors
}

func (entity *Sensors) CacheKey() string {
	return fmt.Sprintf(sensorsCacheKey, entity.SensorID)
}

func (entity *Sensors) CacheExpireTime() int64 {
	return redis.OneMonth
}

func (entity *Sensors) IsNil() bool {
	return entity.SensorID == ""
}

func SensorsHelper(IDs []string) []redis.Entity {
	slice := make([]redis.Entity, len(IDs))
	for i, ID := range IDs {
		slice[i] = &Sensors{SensorID: ID}
	}
	return slice
}
