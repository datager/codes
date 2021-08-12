package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/json"
)

type FaceRuleSensors struct {
	Uts       time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts        int64     `xorm:"not null default 0 BIGINT"`
	RuleId    string    `xorm:"not null pk VARCHAR(1024)"`
	IsLong    bool      `xorm:"default false BOOL"`
	StartDate int64     `xorm:"not null default 0 BIGINT"`
	EndDate   int64     `xorm:"not null default 0 BIGINT"`
	StartTime int64     `xorm:"not null default 0 BIGINT"`
	EndTime   int64     `xorm:"not null default 0 BIGINT"`
	SensorId  string    `xorm:"not null default '''::character varying' index VARCHAR(1024)"`
	Rois      string    `xorm:"default '''::text' TEXT"`
	RepoId    string    `xorm:"VARCHAR(1024)"`
	Rule      string    `xorm:"not null default '''::text' TEXT"`
	Comment   string    `xorm:"not null default '''::character varying' VARCHAR"`
	UserIds   string    `xorm:"default '''::text' TEXT"`
	Switcher  int       `xorm:"not null default 1 SMALLINT"`
	Status    int       `xorm:"not null default 1 SMALLINT"`
	OrgId     string    `xorm:"not null default '''::character varying"`
}

type FaceRule struct {
	FaceRuleSensors `xorm:"extends"`
	FaceRuleExtend  `xorm:"extends"`
	Sensors         `xorm:"extends"`
}

func (FaceRule) TableName() string {
	return TableNameFaceRuleSensors
}

type FaceRules []*FaceRule

func (entities FaceRules) ToModel() (*models.FaceRule, error) {
	if len(entities) == 0 {
		return nil, nil
	}
	e := entities[0]
	sensorId := e.FaceRuleSensors.SensorId
	m := &models.FaceRule{
		RuleId:    e.FaceRuleSensors.RuleId,
		Timestamp: e.FaceRuleSensors.Ts,
		Sensors: []*models.SensorInRule{
			&models.SensorInRule{
				SensorId:   sensorId,
				SensorName: e.SensorName,
				OrgId:      e.Sensors.OrgId,
				Status:     models.TaskStatus(e.FaceRuleSensors.Status),
			},
		},
		FaceRepos: make([]*models.RepoInRule, len(entities)), // fill the slice later
		Times: &models.TimeInRule{
			StartDate: e.StartDate,
			EndDate:   e.EndDate,
			StartTime: e.StartTime,
			EndTime:   e.EndTime,
			IsLong:    e.IsLong,
		},
		Comment: e.FaceRuleSensors.Comment,
		Status:  models.TaskStatus(e.FaceRuleSensors.Status),
		UserID:  e.UserID,
		OrgId:   e.FaceRuleExtend.OrgID,
	}
	// face repos
	for i, e := range entities {
		if e.FaceRuleSensors.SensorId != sensorId {
			return nil, fmt.Errorf("Sensor id must be distinct")
		}
		m.FaceRepos[i] = &models.RepoInRule{
			RuleId:     e.FaceRuleSensors.RuleId,
			RepoId:     e.RepoId,
			AlarmLevel: e.AlarmLevel,
			AlarmVoice: e.AlarmVoice,
			Status:     models.TaskStatus(e.FaceRuleSensors.Status),
		}
		if e.Rule != "" {
			gotConfidence := false
			strs := strings.Split(e.Rule, ";")
			for _, str := range strs {
				idx := strings.Index(str, "repos=")
				if idx >= 0 {
					res := gjson.Get(str[idx:], e.RepoId)
					if res.Exists() {
						m.FaceRepos[i].Confidence = float32(res.Float())
						gotConfidence = true
						break
					}
				}
			}
			if !gotConfidence {
				return nil, fmt.Errorf("Failed to unmarshal rule %v", e.Rule)
			}
		}
	}
	if e.Rois != "" {
		var rois []*models.Roi
		if err := json.Unmarshal([]byte(e.Rois), &rois); err != nil {
			return nil, errors.WithStack(err)
		}
		m.Sensors[0].Rois = rois
	}
	return m, nil
}
