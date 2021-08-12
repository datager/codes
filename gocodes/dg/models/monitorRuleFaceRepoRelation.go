package models

import "time"

type MonitorRuleFaceRepoRelation struct {
	Uts           time.Time
	Ts            int64
	MonitorRuleID string
	FaceRepoID    string
	Threshold     float32 // 布控库的阈值, 同一个库, 在多个布控下, 可能有多个报警阈值
}
