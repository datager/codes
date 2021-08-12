package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	cacheKeyMonitorRule = "monitor-rule-pk:%v:v1"
)

type MonitorRule struct {
	Uts              time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts               int64     `xorm:"not null"`
	MonitorRuleID    string    `xorm:"'monitor_rule_id' not null pk VARCHAR(1024)"`
	MonitorRuleName  string    `xorm:"'monitor_rule_name' not null"`
	TimeIsLong       bool      `xorm:"'time_is_long' not null"`
	OrgID            string    `xorm:"'org_id' not null"`
	UserID           string    `xorm:"'user_id' not null"`
	LastUpdateOrgID  string    `xorm:"'last_update_org_id' not null"`
	LastUpdateUserID string    `xorm:"'last_update_user_id' not null"`
	Comment          string    `xorm:"'comment' not null"`
	Type             int       `xorm:"'type' not null"`
	Status           int       `xorm:"'status' not null"`
}

func (mr *MonitorRule) TableName() string {
	return TableNameMonitorRule
}

func (mr *MonitorRule) CacheKey() string {
	return fmt.Sprintf(cacheKeyMonitorRule, mr.MonitorRuleID)
}

func (mr *MonitorRule) CacheExpireTime() int64 {
	return redis.OneMonth
}

func MonitorRuleHelper(ids []string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, id := range ids {
		slice[i] = &MonitorRule{MonitorRuleID: id}
	}

	return slice
}

func (mr *MonitorRule) ToInfo() *models.MonitorRuleInfo {
	if mr == nil {
		return nil
	}

	return &models.MonitorRuleInfo{
		MonitorRuleID:   mr.MonitorRuleID,
		MonitorRuleName: mr.MonitorRuleName,
	}
}

func (mr *MonitorRule) ToModel() *models.MonitorRule {
	m := &models.MonitorRule{
		MonitorRuleId:   mr.MonitorRuleID,
		MonitorRuleName: mr.MonitorRuleName,
		Timestamp:       mr.Ts,
		TimeIsLong:      mr.TimeIsLong,
		Comment:         mr.Comment,
		Type:            models.MonitorRuleType(mr.Type),
		Status:          models.MonitorRuleStatus(mr.Status),
		OrgID:           mr.OrgID,
		UserID:          mr.UserID,
	}

	//m.Sensors = xx	// todo: need a way to change from monitorrule_sensor_relation table: from monitorruleid->to sensorIds, then from sensorId ->sensor
	//m.FaceRepos = yy  // todo: need a way to change from monitorrule_facerepo_relation table: from monitorruleid->to faceRepoIds, then from facerepoId -> faceRepo
	return m
}

func MonitorRuleModelToEntity(mr *models.MonitorRule) *MonitorRule {
	m := &MonitorRule{
		Uts:              time.Now(),
		Ts:               mr.Timestamp,
		MonitorRuleID:    mr.MonitorRuleId,
		MonitorRuleName:  mr.MonitorRuleName,
		TimeIsLong:       mr.TimeIsLong,
		OrgID:            mr.OrgID,
		UserID:           mr.UserID,
		LastUpdateOrgID:  mr.LastUpdateOrgID,
		LastUpdateUserID: mr.LastUpdateUserID,
		Comment:          mr.Comment,
		Type:             int(mr.Type),
		Status:           int(mr.Status),
	}

	return m
}

func MonitorRuleEntitiesToModels(ens []*MonitorRule) []*models.MonitorRule {
	models := make([]*models.MonitorRule, 0)
	for _, en := range ens {
		models = append(models, en.ToModel())
	}
	return models
}
