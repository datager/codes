package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	facesCacheKey = "faces-cache-key-faceid:%v:v1"
)

type Faces struct {
	Uts               time.Time   `xorm:"not null default 'now()' DATETIME"`
	Ts                int64       `xorm:"not null default 0 index(face_time_spacial_idx) BIGINT"`
	SensorID          string      `xorm:"sensor_id not null index(face_time_spacial_idx) VARCHAR(1024)"`
	FaceID            string      `xorm:"face_id not null pk VARCHAR(1024)"`
	FaceReid          string      `xorm:"not null default '''::character varying' index VARCHAR(1024)"`
	Feature           string      `xorm:"default '''::text' TEXT"`
	Confidence        float32     `xorm:"not null default 0 REAL"`
	GenderID          int         `xorm:"gender_id default 0 SMALLINT"`
	GenderConfidence  float32     `xorm:"default 0 REAL"`
	AgeID             int         `xorm:"age_id default 0 SMALLINT"`
	AgeConfidence     float32     `xorm:"default 0 REAL"`
	NationID          int         `xorm:"nation_id default 0 SMALLINT"`
	NationConfidence  float32     `xorm:"default 0 REAL"`
	GlassID           int         `xorm:"glass_id default 0 SMALLINT"`
	GlassConfidence   float32     `xorm:"default 0 REAL"`
	MaskID            int         `xorm:"mask_id default 0 SMALLINT"`
	MaskConfidence    float32     `xorm:"default 0 REAL"`
	HatID             int         `xorm:"hat_id default 0 SMALLINT"`
	HatConfidence     float32     `xorm:"default 0 REAL"`
	HalmetID          int         `xorm:"halmet_id default 0 SMALLINT"`
	HalmetConfidence  float32     `xorm:"default 0 REAL"`
	ImageType         int         `xorm:"not null default 0 SMALLINT"`
	ImageURI          string      `xorm:"image_uri default '''::character varying' VARCHAR(256)"`
	ThumbnailImageURI string      `xorm:"thumbnail_image_uri default '''::character varying' VARCHAR(256)"`
	CutboardImageURI  string      `xorm:"cutboard_image_uri default '''::character varying' VARCHAR(256)"`
	CutboardX         int         `xorm:"default 0 INTEGER"`
	CutboardY         int         `xorm:"default 0 INTEGER"`
	CutboardWidth     int         `xorm:"default 0 INTEGER"`
	CutboardHeight    int         `xorm:"default 0 INTEGER"`
	CutboardResWidth  int         `xorm:"default 0 INTEGER"`
	CutboardResHeight int         `xorm:"default 0 INTEGER"`
	IsWarned          int         `xorm:"default 0 SMALLINT"`
	Status            int         `xorm:"not null default 1 SMALLINT"`
	RelationTypes     IntArray    `xorm:"text"`
	RelationIDs       StringArray `xorm:"relation_ids text"`
}

// func (entity *Faces) TableName() string {
// 	return TableNameFacesIndex
// }

type FacesExtend struct {
	Faces  `xorm:"extends"`
	Sensor Sensors `xorm:"extends"`
}

func (FacesExtend) TableName() string {
	return "Please specify table name explicitly"
}

func (entity *Faces) ToImageResultModel() *models.ImageResult {
	img := &models.ImageResult{}
	img.ImageUri = entity.ImageURI
	img.ThumbnailImageUri = entity.ThumbnailImageURI
	img.CutboardImageUri = entity.CutboardImageURI
	img.CutboardX = entity.CutboardX
	img.CutboardY = entity.CutboardY
	img.CutboardWidth = entity.CutboardWidth
	img.CutboardHeight = entity.CutboardHeight
	img.CutboardResWidth = entity.CutboardResWidth
	img.CutboardResHeight = entity.CutboardResHeight
	img.Feature = entity.Feature
	img.ImageType = entity.ImageType
	return img
}

func (entity *FacesExtend) ToModel() *models.CapturedFaceDetial {
	if entity == nil {
		return nil
	}
	model := &models.CapturedFaceDetial{}
	model.FaceID = entity.FaceID
	model.FaceReID = entity.FaceReid
	model.Timestamp = entity.Ts
	model.SensorID = entity.SensorID
	model.SensorName = entity.Sensor.SensorName
	model.SensorType = entity.Sensor.Type
	model.SensorURL = entity.Sensor.Url
	model.Longitude = entity.Sensor.Longitude
	model.Latitude = entity.Sensor.Latitude

	civilAttr := &models.CivilAttr{}
	civilAttr.GenderId = entity.GenderID
	civilAttr.GenderConfidence = entity.GenderConfidence
	civilAttr.AgeId = entity.AgeID
	civilAttr.AgeConfidence = entity.AgeConfidence
	civilAttr.NationId = entity.NationID
	civilAttr.NationConfidence = entity.NationConfidence
	civilAttr.GlassId = entity.GlassID
	civilAttr.GlassConfidence = entity.GlassConfidence
	civilAttr.HatId = entity.HatID
	civilAttr.HatConfidence = entity.HatConfidence
	civilAttr.HalmetId = entity.HalmetID
	civilAttr.HalmetConfidence = entity.HalmetConfidence
	civilAttr.MaskId = entity.MaskID
	civilAttr.MaskConfidence = entity.MaskConfidence
	model.CivilAttr = civilAttr

	model.ImageResult = entity.ToImageResultModel()
	model.AgeID = entity.AgeID
	model.GenderID = entity.GenderID
	model.GlassID = entity.GlassID
	model.RelationTypes = entity.Faces.RelationTypes
	model.RelationIDs = entity.Faces.RelationIDs
	return model
}

func (entity *FacesExtend) ToFacesExtendModel() *models.FacesExtend {
	model := &models.FacesExtend{}
	faceModel := entity.ToModel()
	model.CapturedFaceDetial = *faceModel

	sensorModel := entity.Sensor.ToModel()
	model.Sensor = *sensorModel

	return model
}

func (entity *Faces) TableName() string {
	return TableNameFaces
}

func (entity *Faces) CacheKey() string {
	return fmt.Sprintf(facesCacheKey, entity.FaceID)
}

func (entity *Faces) CacheExpireTime() int64 {
	return redis.OneWeek * 2
}

func FacesHelper(IDs []string) []redis.Entity {
	slice := make([]redis.Entity, len(IDs))
	for i, ID := range IDs {
		slice[i] = &Faces{FaceID: ID}
	}
	return slice
}
