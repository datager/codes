package models

import "time"

type MonitorRuleSensorRelation struct {
	Uts           time.Time
	Ts            int64
	MonitorRuleID string
	SensorID      string
}

func NewMonitorRuleSensorRelation(monitorRuleID, sensorID string) *MonitorRuleSensorRelation {
	now := time.Now()
	return &MonitorRuleSensorRelation{
		Uts:           now,
		Ts:            now.Unix() * 1000,
		MonitorRuleID: monitorRuleID,
		SensorID:      sensorID,
	}
}
