package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type CaptureCountStatistics struct {
	ID       int64     `xorm:"id serial"`
	Uts      time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts       int64     `xorm:"not null default 0 BIGINT"`
	SensorID string    `xorm:"sensor_id not null VARCHAR(1024)"`
	Type     int       `xorm:"not null default 0 SMALLINT"`
	Count    int64     `xorm:"not null default 0 BIGINT"`
}

func (CaptureCountStatistics) TableName() string {
	return TableNameCaptureCountStatistics
}

func (ccs *CaptureCountStatistics) Build() *models.CaptureStatisticsItem {
	result := &models.CaptureStatisticsItem{Ts: ccs.Ts}
	switch ccs.Type {
	case models.DetTypeFace:
		result.Face += ccs.Count
	case models.DetTypePedestrian:
		result.Pedestrian += ccs.Count
	case models.DetTypeVehicle:
		result.Vehicle += ccs.Count
	case models.DetTypeNonmotor:
		result.Nonmotor += ccs.Count
	}
	return result
}

func (ccs *CaptureCountStatistics) ToModel() *models.CaptureCountStatistics {
	if ccs == nil {
		return nil
	}

	return &models.CaptureCountStatistics{
		ID:       ccs.ID,
		Uts:      ccs.Uts,
		Ts:       ccs.Ts,
		SensorID: ccs.SensorID,
		Type:     ccs.Type,
		Count:    ccs.Count,
	}
}
