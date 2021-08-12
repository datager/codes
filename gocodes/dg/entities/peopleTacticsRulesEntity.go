package entities

import "codes/gocodes/dg/models"

const TacticsRulesTableName = "people_tactics_rules"

type PeopleTacticsRulesEntity struct {
	Timestamp      int64              `xorm:"ts"`
	TacticRuleID   string             `xorm:"tactic_rule_id not null pk VARCHAR(1024)"`
	TacticID       string             `xorm:"tactic_id"`
	TacticRuleName string             `xorm:"tactic_rule_name"`
	IsOpen         bool               `xorm:"is_open"`
	WarnLevel      int                `xorm:" warn_level"`
	RepoRange      StringArray        `xorm:"repo_range"`
	AnalysisConfig *models.Parameters `xorm:"analysis_config json"`
	UserID         string             `xorm:"user_id"`
}

func (e *PeopleTacticsRulesEntity) EntityToModel() *models.PeopleTacticRules {
	m := &models.PeopleTacticRules{
		Timestamp:      e.Timestamp,
		UserID:         e.UserID,
		TacticID:       e.TacticID,
		TacticRuleID:   e.TacticRuleID,
		TacticRuleName: e.TacticRuleName,
		IsOpen:         e.IsOpen,
		WarningLevel:   e.WarnLevel,
		RepoIDs:        e.RepoRange,
		Parameters:     e.AnalysisConfig,
	}
	return m
}

func (PeopleTacticsRulesEntity) TableName() string {
	return TacticsRulesTableName
}
