package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

// 车辆频率分析概述
type VehicleFrequencySummary struct {
	Uts        time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts         int64     `xorm:"ts not null BIGINT"`
	TaskID     string    `xorm:"task_id not null default '''::character varying' VARCHAR(1024)"`
	SensorID   string    `xorm:"sensor_id not null default '''::character varying' VARCHAR(1024)"`
	SensorName string    `xorm:"sensor_name not null default '''::character varying' VARCHAR(1024)"`
	Count      int64     `xorm:"count not null default 0 BIGINT"` // 每个 sensor 分析出了 count 个结果
}

func (VehicleFrequencySummary) TableName() string {
	return TableNameVehicleFrequencySummary
}

func (vfa *VehicleFrequencySummary) ToModel() *models.VehicleFrequencySummary {
	return &models.VehicleFrequencySummary{
		Uts:        vfa.Uts,
		Ts:         vfa.Ts,
		TaskID:     vfa.TaskID,
		SensorID:   vfa.SensorID,
		SensorName: vfa.SensorName,
		Count:      vfa.Count,
	}
}

func VehicleFrequencySummaryEntitiesToModels(list []*VehicleFrequencySummary) []*models.VehicleFrequencySummary {
	model := make([]*models.VehicleFrequencySummary, 0)
	for _, l := range list {
		model = append(model, l.ToModel())
	}
	return model
}
