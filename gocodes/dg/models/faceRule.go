package models

import "time"

type FaceRule struct {
	RuleId    string
	Timestamp int64
	Sensors   []*SensorInRule
	FaceRepos []*RepoInRule
	Times     *TimeInRule
	Comment   string
	Status    TaskStatus
	UserID    string
	OrgId     string
}

type SensorInRule struct {
	SensorId    string
	SensorName  string
	Rois        []*Roi
	OrgId       string
	VideoWidth  float32
	VideoHeight float32
	// Switcher   SwitcherStatus
	Status     TaskStatus
	TaskStatus string
}

type Roi struct {
	Points []*Point
}

type Point struct {
	X float32
	Y float32
}

// 布控库报警声音等级 model
// composed of FaceRuleExtend model; MonitorRuleFaceRepoRelation model
type RepoInRule struct {
	RuleId     string // 布控
	RepoId     string // 库
	OrgId      string // 组织
	RepoName   string
	Confidence float32 // 阈值
	AlarmLevel int     // 报警等级
	AlarmVoice string  // 报警声音
	CountLimit int     // not in use
	Type       int64
	Status     TaskStatus
}

type TimeInRule struct {
	StartDate int64
	EndDate   int64
	StartTime int64
	EndTime   int64
	IsLong    bool
}

func (this *TimeInRule) GetStartTs() int64 {
	return this.StartDate*1000 + this.StartTime
}

func (this *TimeInRule) GetEndTs() int64 {
	return this.EndDate*1000 + this.EndTime
}

func (rir *RepoInRule) ToFaceRuleExtendModel() *FaceRuleExtend {
	fr := &FaceRuleExtend{
		FaceRepoID: rir.RepoId,
		AlarmLevel: rir.AlarmLevel,
		AlarmVoice: rir.AlarmVoice,
		RuleID:     rir.RuleId,
		OrgID:      rir.OrgId,
	}
	return fr
}

func (rir *RepoInRule) ToFaceRepo() *FaceRepo {
	fr := &FaceRepo{
		RepoID:   rir.RepoId,
		RepoName: rir.RepoName,
		//IssuedRepo    *IssuedRepo
		//IssuedTo      []*IssuedRepo
		OrgId: rir.OrgId,
		//OrgName       string
		//Timestamp     int64
		//Comment       string
		//PicCount      int64
		//ErrorPicCount int64
		//NameListAttr  int
		Type:   rir.Type,
		Status: rir.Status,
		//SystemTag     int
	}
	return fr
}

func (rir *RepoInRule) ToMonitorRuleFaceRepoRelationEntity() *MonitorRuleFaceRepoRelation {
	r := &MonitorRuleFaceRepoRelation{
		Uts:           time.Now(),
		Ts:            time.Now().Unix(),
		MonitorRuleID: rir.RuleId,
		FaceRepoID:    rir.RepoId,
		Threshold:     rir.Confidence,
	}
	return r
}

func RepoInRuleToRepoIDs(repoInRules []*RepoInRule) []string {
	rIDs := make([]string, 0)
	for _, r := range repoInRules {
		rIDs = append(rIDs, r.RepoId)
	}
	return rIDs
}
