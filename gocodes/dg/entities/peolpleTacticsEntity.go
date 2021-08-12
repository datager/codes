package entities

const PeopleTacticsTableName = "people_tactics"

type PeopleTacticsEntity struct {
	Timestamp     int64    `xorm:"ts"`
	SceneID       string   `xorm:"scene_id"`
	TacticID      string   `xorm:"tactic_id not null pk VARCHAR(1024)"`
	TacticName    string   `xorm:"tactic_name"`
	RepoTags      IntArray `xorm:"repo_tags`
	TacticComment string   `xorm:"tactic_comment"`
	PersonComment string   `xorm:"person_comment"`
	AnalysisType  int      `xorm:"analysis_type"`
	AnalysisItem  IntArray `xorm:"analysis_item"`
	PageStyle     string   `xorm:"page_style"`
}

//func (entity *PeopleTacticsEntity) Tomodel() *models.PeopleTactics {
//	return &models.PeopleTactics{
//		Timestamp:     entity.Timestamp,
//		SceneID:       entity.SceneID,
//		TacticID:      entity.TacticID,
//		TacticName:    entity.TacticName,
//		RepoTagIDs:    entity.RepoTags,
//		TacticComment: entity.TacticComment,
//		PersonComment: entity.PersonComment,
//		AnalysisType:  entity.AnalysisType,
//		AnalysisItem:  entity.AnalysisItem,
//		PageStyle:     entity.PageStyle,
//	}
//}

func (PeopleTacticsEntity) TableName() string {
	return PeopleTacticsTableName
}
