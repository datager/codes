package entities

import "codes/gocodes/dg/models"

type MonitorRuleExtend struct {
	MonitorRule                 *MonitorRule                 `xorm:"extends"`
	MonitorRuleSensorRelation   *MonitorRuleSensorRelation   `xorm:"extends"`
	MonitorRuleFaceRepoRelation *MonitorRuleFaceRepoRelation `xorm:"extends"`
}

// not in using
type RelationExtendCount struct {
	MonitorRule   *MonitorRule `xorm:"extends"`
	SensorCount   int64
	FaceRepoCount int64
}

func (e *MonitorRuleExtend) TableName() string {
	return TableNameMonitorRule
}

func (e *MonitorRuleExtend) ToModel() *models.MonitorRuleExtend {
	respList := &models.MonitorRuleExtend{
		MonitorRule:                 e.MonitorRule.ToModel(),
		MonitorRuleSensorRelation:   e.MonitorRuleSensorRelation.ToModel(),
		MonitorRuleFaceRepoRelation: e.MonitorRuleFaceRepoRelation.ToModel(),
	}
	return respList
}
