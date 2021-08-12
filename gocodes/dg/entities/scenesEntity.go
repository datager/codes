package entities

const ScenesTableName = "scenes"

type ScenesEntity struct {
	Timestamp    int64  `xorm:"ts"`
	SceneID      string `xorm:"scene_id not null pk VARCHAR(1024)"`
	SceneName    string `xorm:"scene_name"`
	SceneType    int    `xorm:"scene_type"`
	CreateType   int    `xorm:"create_type"`
	SceneComment string `xorm:"scene_comment"`
}
type ScenesEntityExtend struct {
	*ScenesEntity `xorm:"extends"`
	Tactics       []*PeopleTacticsEntity `xorm:"extends"`
}

func (ScenesEntity) TableName() string {
	return ScenesTableName
}

//func (entity *ScenesEntityExtend) ToPersonSceneModel() *models.PersonScene {
//	tmodels := make([]*models.PeopleTactics, 0)
//	for _, tactic := range entity.Tactics {
//		tmodels = append(tmodels, tactic.Tomodel())
//	}
//	return &models.PersonScene{
//		Timestamp:    entity.Timestamp,
//		SceneName:    entity.SceneName,
//		SceneComment: entity.SceneComment,
//		SceneID:      entity.SceneID,
//		SceneType:    entity.SceneType,
//		CreateType:   entity.CreateType,
//		Tactics:      tmodels,
//	}
//}
