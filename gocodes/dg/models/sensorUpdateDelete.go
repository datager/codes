package models

// 删除/更改 接口暴露的返回值
type SensorUpdateDeleteResult struct {
	CheckIsSensorCanDeleteOrUpdate bool // 校验: 设备是否被占用: 起任务/被布控/被国标同步且国标正忙/处于非稳定状态(删除中/更新中)
	IsSensorStatusMarkSucceed      bool // 是否已置位 sensorStatus 为删除中/更新中
	BindStatusOfSensor             []*BindStatusOfSensor
}

func NewSensorUpdateDeleteResult() *SensorUpdateDeleteResult {
	return &SensorUpdateDeleteResult{
		CheckIsSensorCanDeleteOrUpdate: false,
		IsSensorStatusMarkSucceed:      false,
		BindStatusOfSensor:             make([]*BindStatusOfSensor, 0),
	}
}

// 设备是否被 bind 到其他, 导致无法删除/更新
type BindStatusOfSensor struct {
	SensorID                            string
	IsBindToTask                        bool
	IsBindToMonitorRule                 bool
	IsGBTypeAndPlatformSyncNotAvailable bool
	IsNonStableStatus                   bool
}
