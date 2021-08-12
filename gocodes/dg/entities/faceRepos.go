package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	cacheKeyFaceRepos = "face-repos-pk:%v:v1"
)

type FaceRepos struct {
	Uts           time.Time `xorm:"not null default 'now()' DATETIME"`                          //最后一次更新时间戳
	Ts            int64     `xorm:"not null default 0 BIGINT"`                                  //创建时间
	RepoId        string    `xorm:"not null pk index VARCHAR(1024)"`                            //库ID
	OrgId         string    `xorm:"not null default ''0000'::character varying' VARCHAR(1024)"` //组织ID
	RepoName      string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`     //库名称
	PicCount      int64     `xorm:"default 0 BIGINT"`                                           //图片数量
	ErrorPicCount int64     `xorm:"default 0 BIGINT"`                                           //Error图片数量
	NameListAttr  int       `xorm:"not null default 0 SMALLINT"`                                //黑白名单???
	Comment       string    `xorm:"default '''::character varying' VARCHAR"`                    //备注
	Type          int64     `xorm:"'type' not null"`                                            //比对库性质???
	Status        int       `xorm:"not null default 1 SMALLINT"`                                //库状态
	SystemTag     int       `xorm:"not null default 2 SMALLINT"`                                //系统默认库或者用户建立库
}

func (entity *FaceRepos) ToModel() *models.FaceRepo {
	return &models.FaceRepo{
		Timestamp:     entity.Ts,
		RepoID:        entity.RepoId,
		OrgId:         entity.OrgId,
		RepoName:      entity.RepoName,
		PicCount:      entity.PicCount,
		ErrorPicCount: entity.ErrorPicCount,
		NameListAttr:  entity.NameListAttr,
		Comment:       entity.Comment,
		Type:          entity.Type,
		Status:        models.TaskStatus(entity.Status),
		SystemTag:     entity.SystemTag,
	}
}

func (entity *FaceRepos) ToInfo() *models.RepoInfo {
	if entity == nil {
		return nil
	}

	return &models.RepoInfo{
		RepoID:       entity.RepoId,
		RepoName:     entity.RepoName,
		NameListAttr: entity.NameListAttr,
		Type:         int(entity.Type),
	}
}

func FaceRepoModelToEntity(repo *models.FaceRepo) *FaceRepos {
	if repo == nil {
		return nil
	}

	return &FaceRepos{
		Uts:           time.Now(),
		Ts:            repo.Timestamp,
		RepoId:        repo.RepoID,
		RepoName:      repo.RepoName,
		PicCount:      repo.PicCount,
		ErrorPicCount: repo.ErrorPicCount,
		NameListAttr:  repo.NameListAttr,
		Comment:       repo.Comment,
		Type:          repo.Type,
		Status:        int(repo.Status),
		SystemTag:     repo.SystemTag,
	}
}

func FaceRepoEntitiesToModels(entities []*FaceRepos) []*models.FaceRepo {
	models := make([]*models.FaceRepo, 0)
	for _, r := range entities {
		models = append(models, r.ToModel())
	}
	return models
}

func (entity *FaceRepos) TableName() string {
	return TableNameFaceRepos
}

func (entity *FaceRepos) CacheKey() string {
	return fmt.Sprintf(cacheKeyFaceRepos, entity.RepoId)
}

func (entity *FaceRepos) CacheExpireTime() int64 {
	return redis.OneMonth
}

func FaceReposHelper(ids []string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, id := range ids {
		slice[i] = &FaceRepos{RepoId: id}
	}

	return slice
}
