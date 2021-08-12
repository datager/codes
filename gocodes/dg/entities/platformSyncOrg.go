package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type PlatformSyncOrg struct {
	Uts           time.Time `xorm:"uts not null default 'now()' DATETIME"`
	Ts            int64     `xorm:"ts not null default 0 BIGINT"`
	OrgID         string    `xorm:"org_id not null VARCHAR(1024)"`
	SuperiorOrgID string    `xorm:"superior_org_id not null VARCHAR(1024)"`
	OrgName       string    `xorm:"org_name not null VARCHAR(1024)"`
	Type          int       `xorm:"type not null default 0 SMALLINT"`
}

func (*PlatformSyncOrg) TableName() string {
	return TableNamePlatformSyncOrgEntity
}

func (pst *PlatformSyncOrg) ToModel() *models.PlatformSyncOrgResponse {
	return &models.PlatformSyncOrgResponse{
		OrgID:         pst.OrgID,
		OrgName:       pst.OrgName,
		SuperiorOrgID: pst.SuperiorOrgID,
		Orgs:          make([]*models.PlatformSyncOrgResponse, 0),
		Devices:       make([]*models.PlatformSyncDevice, 0),
	}
}
