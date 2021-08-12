package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
)

type FaceEvents struct {
	Uts                time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts                 int64     `xorm:"not null index BIGINT"`
	EventId            string    `xorm:"not null pk VARCHAR(1024)"`
	EventReid          string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	SensorId           string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	RuleId             string    `xorm:"not null VARCHAR(1024)"`
	FaceId             string    `xorm:"not null VARCHAR(1024)"`
	RepoId             string    `xorm:"not null VARCHAR(1024)"`
	ImageId            string    `xorm:"not null index VARCHAR(1024)"`
	Confidence         float32   `xorm:"default 0 index REAL"`
	UserIds            string    `xorm:"default '''::text' TEXT"`
	EventType          int       `xorm:"not null default 0 SMALLINT"`
	IsChecked          int       `xorm:"default 0 index SMALLINT"`
	EventIsChecked     int       `xorm:"default 0 SMALLINT"`
	MaxConfidenceEvent bool      `xorm:"default false BOOL"`
	Comment            string    `xorm:"default '''::text' TEXT"`
	Status             int       `xorm:"not null default 1 SMALLINT"`

	SensorName string
	RepoName   string

	CapturedImageUri          string `xorm:"default '''::character varying' VARCHAR(256)"`
	CapturedThumbnailImageUri string `xorm:"default '''::character varying' VARCHAR(256)"`
	CapturedCutboardImageUri  string `xorm:"default '''::character varying' VARCHAR(256)"`
	CapturedCutboardX         int    `xorm:"default 0 INTEGER"`
	CapturedCutboardY         int    `xorm:"default 0 INTEGER"`
	CapturedCutboardWidth     int    `xorm:"default 0 INTEGER"`
	CapturedCutboardHeight    int    `xorm:"default 0 INTEGER"`
	CapturedCutboardResWidth  int    `xorm:"default 0 INTEGER"`
	CapturedCutboardResHeight int    `xorm:"default 0 INTEGER"`

	CivilAttrId    string `xorm:"not null  VARCHAR(1024)"`
	CivilName      string `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	CivilBirthday  int64  `xorm:"default 0 bigint"`
	CivilIdType    int    `xorm:"default 0  SMALLINT"`
	CivilIdNo      string `xorm:"default '''::character varying' index VARCHAR(1024)"`
	CivilAddress   string `xorm:"default '''::character varying' index VARCHAR(1024)"`
	CivilGenderId  int    `xorm:"default 0 SMALLINT"`
	CivilNationId  int    `xorm:"default 0 SMALLINT"`
	CivilImageUris string `xorm:"not null default '''::text' TEXT"`

	CivilImageUri string `xorm:"default '''::character varying'  VARCHAR(1024)"`
	CivilComment  string `xorm:"default '''::character varying' varchar(1024)"`

	RuleName    string `xorm:"default '''::character varying' varchar(1024)"`
	SensorIntID int32  `xorm:"sensor_int_id"`
	OrgCode     int64
}

type FaceEventsExtend struct {
	FaceEvents      `xorm:"extends"`
	SensorName      string
	FaceRepo        FaceRepos    `xorm:"extends"`
	RegisteredImage *CivilImages `xorm:"extends"`
	CivilAttr       *CivilAttrs  `xorm:"extends"`
	CapturedImage   *Faces       `xorm:"extends"`
}

type FaceEventsCivilAttr struct {
	CivilAttrID string `xorm:"civil_attr_id"`
}

func (FaceEventsExtend) TableName() string {
	return "Please specify table name explicitly"
}

func (entity *FaceEventsExtend) ToModel() *models.FaceEvent {
	model := &models.FaceEvent{}
	model.EventId = entity.EventId
	model.EventReId = entity.EventReid
	model.Timestamp = entity.Ts
	model.SensorId = entity.SensorId
	model.SensorName = entity.SensorName
	model.FaceRepoId = entity.RepoId
	model.FaceRepoName = entity.FaceRepo.RepoName
	model.NameListAttr = entity.FaceRepo.NameListAttr
	model.FaceId = entity.FaceId
	model.FaceReId = entity.CapturedImage.FaceReid
	model.CapturedImage = entity.CapturedImage.ToImageResultModel()
	model.ImageId = entity.ImageId
	model.CivilAttr = entity.CivilAttr.ToModel()
	model.RegisteredImage = entity.RegisteredImage.ToModel()
	model.RuleId = entity.RuleId
	// model.Rule // useless ?
	model.Confidence = entity.Confidence
	model.IsChecked = entity.IsChecked
	model.UserIds = entity.UserIds
	model.EventIsChecked = entity.EventIsChecked
	model.Comment = entity.Comment
	model.Status = models.TaskStatus(entity.Status)
	model.RuleName = entity.RuleName
	model.OrgCode = entity.OrgCode
	model.SensorIntID = entity.SensorIntID
	return model
}

func FaceEventsListToExtends(data []*FaceEvents) []*FaceEventsExtend {
	ret := make([]*FaceEventsExtend, len(data))
	for i := 0; i < len(data); i++ {
		ret[i] = data[i].ToExtend()
	}

	return ret
}

func (entity *FaceEvents) ToExtend() *FaceEventsExtend {
	captured := &Faces{
		ImageURI:          entity.CapturedImageUri,
		ThumbnailImageURI: entity.CapturedThumbnailImageUri,
		CutboardImageURI:  entity.CapturedCutboardImageUri,
		CutboardX:         entity.CapturedCutboardX,
		CutboardY:         entity.CapturedCutboardY,
		CutboardWidth:     entity.CapturedCutboardWidth,
		CutboardHeight:    entity.CapturedCutboardHeight,
		CutboardResWidth:  entity.CapturedCutboardResWidth,
		CutboardResHeight: entity.CapturedCutboardResHeight,
	}

	civilAttr := &CivilAttrs{
		CivilAttrId: entity.CivilAttrId,
		IdNo:        entity.CivilIdNo,
		IdType:      entity.CivilIdType,
		Name:        entity.CivilName,
		Address:     entity.CivilAddress,
		Birthday:    fmt.Sprintf("%v", entity.CivilBirthday),
		GenderId:    entity.CivilGenderId,
		NationId:    entity.CivilNationId,
		ImageUris:   entity.CivilImageUris,
		Comment:     entity.CivilComment,
	}

	register := &CivilImages{
		ImageId:  entity.ImageId,
		ImageUri: entity.CivilImageUri,
	}

	repo := FaceRepos{
		RepoId:   entity.RepoId,
		RepoName: entity.RepoName,
	}

	a := &FaceEventsExtend{
		FaceEvents:      *entity,
		SensorName:      entity.SensorName,
		FaceRepo:        repo,
		RegisteredImage: register,
		CivilAttr:       civilAttr,
		CapturedImage:   captured,
	}

	return a
}
