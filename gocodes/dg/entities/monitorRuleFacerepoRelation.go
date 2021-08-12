package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type MonitorRuleFaceRepoRelation struct {
	Uts           time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts            int64     `xorm:"not null BIGINT"`
	MonitorRuleID string    `xorm:"monitor_rule_id not null default '''::character varying' VARCHAR(1024)"`
	FaceRepoID    string    `xorm:"face_repo_id not null default '''::character varying' VARCHAR(1024)"`
	Threshold     float32   `xorm:"not null default 0"` // 布控库的阈值, 同一个库, 在多个布控下, 可能有多个报警阈值
}

func (mfr *MonitorRuleFaceRepoRelation) TableName() string {
	return TableNameMonitorRuleFaceRepoRelation
}

func (mfr *MonitorRuleFaceRepoRelation) ToModel() *models.MonitorRuleFaceRepoRelation {

	r := &models.MonitorRuleFaceRepoRelation{
		Uts:           mfr.Uts,
		Ts:            mfr.Ts,
		MonitorRuleID: mfr.MonitorRuleID,
		FaceRepoID:    mfr.FaceRepoID,
		Threshold:     mfr.Threshold,
	}
	return r
}

func MonitorRuleFaceRepoRelationModelToEntity(mfr *models.MonitorRuleFaceRepoRelation) *MonitorRuleFaceRepoRelation {
	r := &MonitorRuleFaceRepoRelation{
		Uts:           time.Now(),
		Ts:            mfr.Ts,
		MonitorRuleID: mfr.MonitorRuleID,
		FaceRepoID:    mfr.FaceRepoID,
		Threshold:     mfr.Threshold,
	}
	return r
}
