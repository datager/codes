package models

import (
	"time"
)

const (
	AlarmLevelOne   = 1 // 1级最高
	AlarmLevelTwo   = 2 // 2级
	AlarmLevelThree = 3 // 3级最低
)

type FaceRuleExtend struct {
	ID         string
	SensorID   string
	FaceRepoID string
	UserID     string
	AlarmLevel int
	AlarmVoice string
	Ts         int64
	OrgID      string
	RuleID     string
	UpdatedAt  time.Time
}

func (fre FaceRuleExtend) ToRepoInRuleModel() *RepoInRule {
	r := &RepoInRule{
		RuleId: fre.RuleID,
		RepoId: fre.FaceRepoID,
		//RepoName: // todo
		//Confidence: // todo
		AlarmLevel: fre.AlarmLevel,
		AlarmVoice: fre.AlarmVoice,
		//Confidence: // todo
		//Status: // todo
	}
	return r
}

func RepoInRuleModelToFaceRuleExtendModel(rule *RepoInRule) *FaceRuleExtend {
	e := &FaceRuleExtend{
		RuleID:     rule.RuleId,
		FaceRepoID: rule.RepoId,
		AlarmLevel: rule.AlarmLevel,
		AlarmVoice: rule.AlarmVoice,
	}
	return e
}
