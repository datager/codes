package entities

import (
	"time"

	"codes/gocodes/dg/utils/json"

	"codes/gocodes/dg/models"
)

type OperRecord struct {
	Uts           time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts            int64     `xorm:"not null default 0 BIGINT"`
	RecordId      string    `xorm:"not null pk VARCHAR(1024)"`
	UserId        string    `xorm:"not null VARCHAR(1024)"`
	OperField     string    `xorm:"not null index VARCHAR(1024)"`
	OperType      string    `xorm:"not null index VARCHAR(1024)"`
	OperContent   string    `xorm:"default '''::text' index TEXT"`
	Url           string    `xorm:"default '''::text' TEXT"`
	IpAddr        string    `xorm:"default '''::character varying' VARCHAR"`
	Result        string    `xorm:"default '''::character varying' index VARCHAR"`
	ResultContent string    `xorm:"default '''::text' TEXT"`
	Status        int       `xorm:"not null default 1 SMALLINT"`
}
type OperRecordExtend struct {
	OperRecord `xorm:"extends"`
	Account    Account `xorm:"extends"`
}

func (OperRecordExtend) TableName() string {
	return TableNameOperRecord
}

func (entity *OperRecordExtend) ToModel() *models.Record {
	urls := make([]string, 0)
	if entity.Url != "" {
		json.Unmarshal([]byte(entity.Url), &urls) // ignore error
	}
	return &models.Record{
		Timestamp:   entity.Ts,
		RecordId:    entity.RecordId,
		UserName:    entity.Account.UserName,
		UserId:      entity.UserId,
		OperField:   entity.OperField,
		OperType:    entity.OperType,
		OperContent: entity.OperContent,
		// BinDate:     entity.BinDate, // todo
		Url:           urls,
		IpAddr:        entity.IpAddr,
		Result:        entity.Result,
		ResultContent: entity.ResultContent,
		Status:        entity.Status,
	}
}
