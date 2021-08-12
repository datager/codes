package models

import (
	"time"
)

const (
	// scheduler timer
	SensorPreviewSchedulerTime = 1 * time.Minute // 1min

	// keys by set
	CacheKeySensorPreviewKeysOfLatestTs = "sensor-preview-keys-of-latest-ts:v2" // k: regular string, v: sensor_ids(of CacheKeySensorPreviewLatestTs), format: redis set

	// key by kv
	CacheKeySensorPreviewLatestTs          = "sensor-preview-latest-ts:sensor:%v:v2" // k: sensor_id, v: latest_ts(of preview api heartbeat)
	CacheExpireTimeOfSensorPreviewLatestTs = 2 * 60                                  // 2 min= 2*60=2*SensorPreviewMaxKeepAliveTime
	SensorPreviewMaxKeepAliveTime          = 1 * time.Minute                         // 1min
)

type SensorPreview struct {
	Uts      time.Time
	Ts       int64
	SensorID string
	URL      string
}

type SensorPreviewReq struct {
	SensorID string
}

type SensorPreviewResp struct {
	URL string `json:"url"`
}
