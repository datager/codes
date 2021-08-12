package entities

import (
	"fmt"
	"strconv"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	cacheKeyCivilAttrs = "civil-attrs-pk:%v:v1"
)

type CivilAttrs struct {
	Uts         time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts          int64     `xorm:"not null default 0 BIGINT"`
	CivilAttrId string    `xorm:"not null pk VARCHAR(1024)"`
	RepoId      string    `xorm:"not null index VARCHAR(1024)"`
	Name        string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Birthday    string    `xorm:"VARCHAR(1024)"`
	IdType      int       `xorm:"default 0 index SMALLINT"`
	IdNo        string    `xorm:"default '''::character varying' index VARCHAR(1024)"`
	Address     string    `xorm:"default '''::character varying' index VARCHAR(1024)"`
	GenderId    int       `xorm:"default 0 SMALLINT"`
	NationId    int       `xorm:"default 0 SMALLINT"`
	Comment     string    `xorm:"default '''::text' TEXT"`
	ImageUris   string    `xorm:"not null default '''::text' TEXT"`
	Status      int       `xorm:"not null default 1 SMALLINT"`
	SearchID    string    `xorm:"search_id varchar(1024)"`
}

type CivilAttrsExtend struct {
	CivilAttrs   `xorm:"extends"`
	CivilImage   CivilImages    `xorm:"extends"`
	FaceRepo     FaceRepos      `xorm:"extends"`
	ImageResults []*CivilImages // since xorm dose not support one-to-many, we will do it manually in CivilAttrRepository
}

type CivilAttrsSimpleExtend struct {
	CivilAttrs `xorm:"extends"`
	FaceRepo   FaceRepos `xorm:"extends"`
}

func (CivilAttrsExtend) TableName() string {
	return TableNameCivilAttrs
}

func (entity *CivilAttrs) ToModel() *models.CivilAttr {
	var birthday *int64
	if entity.Birthday == "" {
		birthday = nil
	} else {
		n, _ := strconv.Atoi(entity.Birthday)
		c := int64(n)
		birthday = &c
	}
	model := &models.CivilAttr{}
	model.Name = entity.Name
	model.IdType = models.IdType(entity.IdType)
	model.IdNo = entity.IdNo
	model.GenderId = entity.GenderId
	model.NationId = entity.NationId
	model.Addr = entity.Address
	model.Birthday = birthday
	model.Comment = entity.Comment
	model.Status = models.TaskStatus(entity.Status)
	model.AttrId = entity.CivilAttrId
	model.SaveStatus = models.CivilSaveStatusSuccess
	model.SearchID = entity.SearchID
	return model
}

func (entity *CivilAttrsExtend) ToModel() *models.Civil {
	model := &models.Civil{}
	model.CivilAttr = entity.CivilAttrs.ToModel()
	model.FaceRepoId = entity.FaceRepo.RepoId
	model.FaceRepoName = entity.FaceRepo.RepoName
	model.NameListAttr = entity.FaceRepo.NameListAttr
	model.Status = models.TaskStatus(entity.Status)
	model.Timestamp = entity.Ts
	// model.Confidence // todo ?
	imgs := make([]*models.ImageResult, len(entity.ImageResults))
	for i, img := range entity.ImageResults {
		imgs[i] = img.ToModel()
	}
	model.ImageResults = imgs
	return model
}

func (entity *CivilAttrsSimpleExtend) ToModel() *models.Civil {
	model := &models.Civil{}
	model.CivilAttr = entity.CivilAttrs.ToModel()
	model.FaceRepoId = entity.FaceRepo.RepoId
	model.FaceRepoName = entity.FaceRepo.RepoName
	model.NameListAttr = entity.FaceRepo.NameListAttr
	model.Status = models.TaskStatus(entity.Status)
	model.Timestamp = entity.Ts

	image := models.ImageResult{
		ImageUri: entity.ImageUris,
	}

	model.ImageResults = []*models.ImageResult{&image}
	return model
}

func (entity *CivilAttrsSimpleExtend) ToCivilAttrsExtend() *CivilAttrsExtend {
	if entity == nil {
		return nil
	}

	return &CivilAttrsExtend{
		CivilAttrs: entity.CivilAttrs,
		FaceRepo:   entity.FaceRepo,
		ImageResults: []*CivilImages{
			&CivilImages{
				ImageUri: entity.CivilAttrs.ImageUris,
			},
		},
	}
}

func (entity *CivilAttrs) TableName() string {
	return TableNameCivilAttrs
}

func (entity *CivilAttrs) CacheKey() string {
	return fmt.Sprintf(cacheKeyCivilAttrs, entity.CivilAttrId)
}

func (entity *CivilAttrs) CacheExpireTime() int64 {
	return redis.OneWeek * 2
}

func CivilAttrHelper(ids []string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, id := range ids {
		slice[i] = &CivilAttrs{CivilAttrId: id}
	}

	return slice
}
