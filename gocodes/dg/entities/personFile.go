package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/json"
	"codes/gocodes/dg/utils/redis"
	"codes/gocodes/dg/utils/systemtags"
)

const (
	personFileCacheKey = "person-file-cache-key-vid:%v:v1"
)

type PersonFile struct {
	Uts              time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts               int64     `xorm:"not null default 0 BIGINT"`
	Vid              string    `xorm:"not null pk VARCHAR(1024)"`
	Name             string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Birthday         int64     `xorm:"default 0 BIGINT"`
	IdType           int       `xorm:"default 0 index SMALLINT"`
	IdNo             string    `xorm:"default '''::character varying' index VARCHAR(1024)"`
	Address          string    `xorm:"default '''::character varying' index VARCHAR(1024)"`
	GenderId         int       `xorm:"default 0 SMALLINT"`
	NationId         int       `xorm:"default 0 SMALLINT"`
	Comment          string    `xorm:"default '''::text' TEXT"`
	ImageUrl         string    `xorm:"not null TEXT"`
	MarkState        int       `xorm:"default 0 SMALLINT"`
	IsNeedAutoUpdate bool      `xorm:"default false BOOL"`
	Status           int       `xorm:"not null default 1 SMALLINT"`
	Confidence       float32   `xorm:"-"` // ignore
	IsImpacted       bool      `xorm:"default false BOOL"`
	Tags             int       `xorm:"default 0 SMALLINT"`
}

type PersonFileExtend struct {
	PersonFile   `xorm:"extends"`
	CapturedFace FacesIndexExtend `xorm:"extends"`
}

func (entity *PersonFile) ToModel() *models.PersonFile {
	if entity == nil {
		return nil
	}

	repos := make([]*models.ReposData, 0)
	systemTags := make([]string, 0)
	customTags := make([]string, 0)
	if entity.Comment != "" {
		commentJson := new(models.CommentJSON)
		err := json.Unmarshal([]byte(entity.Comment), commentJson)
		if err == nil {
			for id, name := range commentJson.Repos {
				d := &models.ReposData{
					RepoID:   id,
					RepoName: name,
				}

				repos = append(repos, d)
			}
			Tags := commentJson.Tags
			systemTags, customTags = systemtags.TagsSplit(Tags)
		}
	}

	return &models.PersonFile{
		Ts:               entity.Ts,
		Vid:              entity.Vid,
		Name:             entity.Name,
		Birthday:         entity.Birthday,
		IdType:           entity.IdType,
		IdNo:             entity.IdNo,
		Address:          entity.Address,
		GenderId:         entity.GenderId,
		NationId:         entity.NationId,
		Comment:          entity.Comment,
		ImageUrl:         entity.ImageUrl,
		MarkState:        entity.MarkState,
		IsNeedAutoUpdate: entity.IsNeedAutoUpdate,
		Status:           entity.Status,
		Confidence:       entity.Confidence,
		IsImpacted:       entity.IsImpacted,
		ExistInRepos:     repos,
		SystemTags:       systemTags,
		CustomTags:       customTags,
	}
}

func (entity *PersonFileExtend) ToModel() *models.PersonFile {
	if entity == nil {
		return nil
	}

	model := entity.PersonFile.ToModel()
	capture := entity.CapturedFace.ToCapturedFace()
	model.CapturedFace = capture
	return model
}

func (entity *PersonFile) Set(model *models.PersonFile) {
	if entity == nil {
		return
	}
	entity.Ts = model.Ts
	entity.Vid = model.Vid
	entity.Name = model.Name
	entity.Birthday = model.Birthday
	entity.IdType = model.IdType
	entity.IdNo = model.IdNo
	entity.Address = model.Address
	entity.GenderId = model.GenderId
	entity.NationId = model.NationId
	entity.Comment = model.Comment
	entity.ImageUrl = model.ImageUrl
	entity.MarkState = model.MarkState
	entity.IsNeedAutoUpdate = model.IsNeedAutoUpdate
	entity.Status = model.Status
	entity.Confidence = model.Confidence
	entity.IsImpacted = model.IsImpacted
}

func (entity *PersonFile) TableName() string {
	return TableNamePersonFile
}

func (entity *PersonFile) CacheKey() string {
	return fmt.Sprintf(personFileCacheKey, entity.Vid)
}

func (entity *PersonFile) CacheExpireTime() int64 {
	return redis.OneWeek * 2
}

func PersonFileHelper(ids []string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, id := range ids {
		slice[i] = &PersonFile{Vid: id}
	}
	return slice
}
