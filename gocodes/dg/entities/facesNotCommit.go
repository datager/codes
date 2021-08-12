package entities

import (
	"time"
)

type FacesNotCommit struct {
	Uts           time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts            int64     `xorm:"not null default 0 BIGINT"`
	PairId        string    `xorm:"not null pk VARCHAR(1024)"`
	GateFaceId    string    `xorm:"default '''::character varying' VARCHAR(1024)"`
	CaptureFaceId string    `xorm:"default '''::character varying' VARCHAR(1024)"`
	Confidence    float32   `xorm:"not null default 0 REAL"`
	CheckValid    int       `xorm:"default 0 SMALLINT"`
	Name          string    `xorm:"default '''::character varying' VARCHAR(256)"`
	IdType        int       `xorm:"default 0 SMALLINT"`
	IdNo          string    `xorm:"default '''::character varying' VARCHAR(1024)"`
	Address       string    `xorm:"default '''::character varying' VARCHAR(1024)"`
	Comment       string    `xorm:"default '''::text' TEXT"`
	Status        int       `xorm:"not null default 1 SMALLINT"`
}

func (FacesNotCommit) TableName() string {
	return TableNameUnCommitFace
}
