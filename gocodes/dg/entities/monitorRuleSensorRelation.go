package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type MonitorRuleSensorRelation struct {
	Uts           time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts            int64     `xorm:"not null BIGINT"`
	MonitorRuleID string    `xorm:"monitor_rule_id not null default '''::character varying' VARCHAR(1024)"`
	SensorID      string    `xorm:"sensor_id not null default '''::character varying' VARCHAR(1024)"`
}

func (mfr *MonitorRuleSensorRelation) TableName() string {
	return TableNameMonitorRuleSensorRelation
}

func (mfr *MonitorRuleSensorRelation) ToModel() *models.MonitorRuleSensorRelation {
	if mfr == nil {
		return nil
	}
	r := &models.MonitorRuleSensorRelation{
		Uts:           mfr.Uts,
		Ts:            mfr.Ts,
		MonitorRuleID: mfr.MonitorRuleID,
		SensorID:      mfr.SensorID,
	}
	return r
}

func MonitorRuleSensorRelationModelToEntity(mfr *models.MonitorRuleSensorRelation) *MonitorRuleSensorRelation {
	if mfr == nil {
		return nil
	}
	r := &MonitorRuleSensorRelation{
		Uts:           time.Now(),
		Ts:            mfr.Ts,
		MonitorRuleID: mfr.MonitorRuleID,
		SensorID:      mfr.SensorID,
	}
	return r
}

func MonitorRuleSensorRelationsToSensorIDsMap(monitorRuleSensorRelations []*MonitorRuleSensorRelation) map[string]*MonitorRuleSensorRelation {
	monitorRuleSensorRelationMap := make(map[string]*MonitorRuleSensorRelation)
	for _, mrsr := range monitorRuleSensorRelations {
		monitorRuleSensorRelationMap[mrsr.SensorID] = mrsr
	}
	return monitorRuleSensorRelationMap
}
