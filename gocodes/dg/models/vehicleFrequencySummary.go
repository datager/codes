package models

import "time"

// 分析结果的summary 响应体
type VehicleFrequencySummary struct {
	Uts        time.Time
	Ts         int64
	TaskID     string
	SensorID   string
	SensorName string
	Count      int64 // 每个 sensor 分析出了 count 个结果
}
