package entities

import "codes/gocodes/dg/models"

// userID 把 repoID 下发给了 orgID
type IssuedRepo struct {
	Ts     int64  `xorm:"not null default 0 BIGINT"`
	Id     string `xorm:"not null pk index VARCHAR(1024)"`
	RepoId string `xorm:"not null default ''0000'::character varying' VARCHAR(1024)"`
	OrgId  string `xorm:"not null default ''0000'::character varying' VARCHAR(1024)"`
	UserId string `xorm:"not null default ''0000'::character varying' VARCHAR(1024)"`
}

func (entity *IssuedRepo) ToModel() *models.IssuedRepo {
	if entity == nil {
		return nil
	}
	return &models.IssuedRepo{
		Ts:     entity.Ts,
		Id:     entity.Id,
		RepoId: entity.RepoId,
		OrgId:  entity.OrgId,
		UserId: entity.UserId,
	}
}

func IssuedRepoToIssuedRepoIDs(issuedRepos []*IssuedRepo) []string {
	issuedRepoIDs := make([]string, 0)
	for _, r := range issuedRepos {
		issuedRepoIDs = append(issuedRepoIDs, r.RepoId)
	}
	return issuedRepoIDs
}

func IssuedRepoEntitiesToModels(entities []*IssuedRepo) []*models.IssuedRepo {
	models := make([]*models.IssuedRepo, 0)
	for _, r := range entities {
		models = append(models, r.ToModel())
	}
	return models
}
