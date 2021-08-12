package entities

type FaceRepoExtend struct {
	FaceRepos                             *FaceRepos                   `xorm:"extends"`
	MonitorRuleFaceRepoRelationRepository *MonitorRuleFaceRepoRelation `xorm:"extends"`
}
