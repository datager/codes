package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	orgStructureCacheKey = "org-structure-cache-key-org-id:%v:v2"
)

type OrgStructure struct {
	Uts           time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts            int64     `xorm:"not null default 0 BIGINT"`
	OrgId         string    `xorm:"not null pk VARCHAR(1024)"`
	OrgCode       int64     `xorm:"not null default -1 BIGINT"`
	OrgName       string    `xorm:"default '''::character varying' VARCHAR(1024)"`
	OrgLevel      int64     `xorm:"not null default 0 BIGINT"`
	SuperiorOrgId string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Comment       string    `xorm:"default '''::character varying' VARCHAR"`
	Status        int       `xorm:"not null default 1 SMALLINT"`
}

func (entity *OrgStructure) Build() *models.OrgDict {
	if entity == nil {
		return nil
	}

	return &models.OrgDict{
		OrgID:         entity.OrgId,
		OrgCode:       entity.OrgCode,
		OrgName:       entity.OrgName,
		OrgLevel:      entity.OrgLevel,
		SuperiorOrgID: entity.SuperiorOrgId,
		Comment:       entity.Comment,
		Status:        entity.Status,
		Ts:            entity.Ts,
	}
}

func (entity *OrgStructure) ToModel() *models.Org {
	if entity == nil {
		return nil
	}
	return &models.Org{
		Ts:            entity.Ts,
		OrgID:         entity.OrgId,
		OrgCode:       entity.OrgCode,
		OrgName:       entity.OrgName,
		OrgLevel:      entity.OrgLevel,
		SuperiorOrgId: entity.SuperiorOrgId,
		Comment:       entity.Comment,
		Status:        entity.Status,
	}
}

func OrgModelToEntity(model *models.Org) *OrgStructure {
	return &OrgStructure{
		Ts:            model.Ts,
		OrgId:         model.OrgID,
		OrgCode:       model.OrgCode,
		OrgName:       model.OrgName,
		OrgLevel:      model.OrgLevel,
		SuperiorOrgId: model.SuperiorOrgId,
		Comment:       model.Comment,
		Status:        model.Status,
	}
}

func OrgToOrgIDs(orgs []*OrgStructure) []string {
	orgIDs := make([]string, 0)
	for _, o := range orgs {
		orgIDs = append(orgIDs, o.OrgId)
	}
	return orgIDs
}

func OrgEntitiesToModels(entities []*OrgStructure) []*models.Org {
	models := make([]*models.Org, 0)
	for _, r := range entities {
		models = append(models, r.ToModel())
	}
	return models
}

func (entity *OrgStructure) TableName() string {
	return TableNameOrgStructure
}

func (entity *OrgStructure) CacheKey() string {
	return fmt.Sprintf(orgStructureCacheKey, entity.OrgId)
}

func (entity *OrgStructure) CacheExpireTime() int64 {
	return redis.OneMonth
}

func OrgStructureHelpler(ids []string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, id := range ids {
		slice[i] = &OrgStructure{
			OrgId: id,
		}
	}

	return slice
}
