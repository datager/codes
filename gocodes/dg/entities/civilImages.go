package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	cacheKeyCivilImage = "civil-image-pk:%v:v1"
)

type CivilImages struct {
	Uts         time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts          int64     `xorm:"not null default 0 BIGINT"`
	CivilAttrId string    `xorm:"not null index VARCHAR(1024)"`
	ImageId     string    `xorm:"not null pk index VARCHAR(1024)"`
	ImageUri    string    `xorm:"default '''::character varying' VARCHAR"`
	Feature     string    `xorm:"default '''::text' TEXT"`
	ImageType   int       `xorm:"default 0 SMALLINT"`
	IsRanked    int       `xorm:"default 0 SMALLINT"`
	Status      int       `xorm:"not null default 1 SMALLINT"`
}

type CivilImagesExtend struct {
	CivilImage CivilImages `xorm:"extends"`
	CivilAttrs `xorm:"extends"`
}

func (CivilImagesExtend) TableName() string {
	return TableNameCivilImages
}

func (entity *CivilImages) ToModel() *models.ImageResult {
	if entity == nil {
		return nil
	}
	model := &models.ImageResult{}
	model.ImageUri = entity.ImageUri
	model.Feature = entity.Feature
	model.IsRanked = entity.IsRanked
	model.Status = models.TaskStatus(entity.Status)
	model.ImageId = entity.ImageId
	model.ImageType = entity.ImageType
	model.Feature = entity.Feature
	return model
}

func (entity *CivilImages) TableName() string {
	return TableNameCivilImages
}

func (entity *CivilImages) CacheKey() string {
	return fmt.Sprintf(cacheKeyCivilImage, entity.ImageId)
}

func (entity *CivilImages) CacheExpireTime() int64 {
	return redis.OneWeek * 2
}

func CivilImageHelper(ids []string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, id := range ids {
		slice[i] = &CivilImages{ImageId: id}
	}

	return slice
}
