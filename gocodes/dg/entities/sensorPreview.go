package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	CacheKeySensorPreview = "sensor-preview:sensor:%v:v1" // k: sensorID, v: url
)

type SensorPreview struct {
	Uts      time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts       int64     `xorm:"ts not null default 0 BIGINT"`
	SensorID string    `xorm:"sensor_id not null pk VARCHAR(1024)"`
	URL      string    `xorm:"url default '''::character varying' VARCHAR"`
}

func (spe *SensorPreview) TableName() string {
	return TableNameSensorPreview
}

func (spe *SensorPreview) CacheKey() string {
	return fmt.Sprintf(CacheKeySensorPreview, spe.SensorID)
}

func (spe *SensorPreview) CacheExpireTime() int64 {
	return redis.OneWeek
}

func (spe *SensorPreview) ToModel() *models.SensorPreview {
	return &models.SensorPreview{
		Uts:      spe.Uts,
		SensorID: spe.SensorID,
		Ts:       spe.Ts,
		URL:      spe.URL,
	}
}

func SensorPreViewModelToEntity(m *models.SensorPreview) *SensorPreview {
	return &SensorPreview{
		Uts:      m.Uts,
		SensorID: m.SensorID,
		Ts:       m.Ts,
		URL:      m.URL,
	}
}

func SensorPreviewsHelper(sensorIDs []string) []redis.Entity {
	slice := make([]redis.Entity, len(sensorIDs))
	for i, id := range sensorIDs {
		slice[i] = &SensorPreview{SensorID: id}
	}

	return slice
}
