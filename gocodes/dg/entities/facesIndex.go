package entities

import (
	"codes/gocodes/dg/models"
)

type FacesIndex struct {
	FaceID           string `xorm:"face_id not null pk VARCHAR(36)"`
	SensorID         string `xorm:"sensor_id not null VARCHAR(36)"`
	SensorIntID      int32  `xorm:"sensor_int_id"`
	OrgCode          int64  `xorm:"org_code not null default 0"`
	FaceReID         string `xorm:"face_reid not null VARCHAR(36)"`
	FaceVid          string `xorm:"face_vid not null VARCHAR(36)"`
	GenderID         int    `xorm:"gender_id default 0 "`
	AgeID            int    `xorm:"age_id default 0 "`
	NationID         int    `xorm:"nation_id default 0"`
	GlassID          int    `xorm:"glass_id default 0 "`
	MaskID           int    `xorm:"mask_id default 0 "`
	HatID            int    `xorm:"hat_id default 0 "`
	HalmetID         int    `xorm:"halmet_id default 0 "`
	CutboardImageURI string `xorm:"cutboard_image_uri VARCHAR(256)"`
	Ts               int64  `xorm:"not null default 0 BIGINT"`
	RelationTypes    int    `xorm:"not null default 0"`
}

type FacesIndexExtend struct {
	FacesIndex `xorm:"extends"`
	Sensor     Sensors `xorm:"extends"`
}

func (f *FacesIndex) TableName() string {
	return TableNameFacesIndex
}

func (fe *FacesIndexExtend) ToCapturedFace() *models.CapturedFace {
	if fe == nil {
		return nil
	}

	return &models.CapturedFace{
		FaceID:            fe.FaceID,
		FaceReID:          fe.FaceReID,
		Timestamp:         fe.Ts,
		SensorID:          fe.SensorID,
		SensorIntID:       fe.SensorIntID,
		OrgCode:           fe.OrgCode,
		SensorName:        fe.Sensor.SensorName,
		Longitude:         fe.Sensor.Longitude,
		Latitude:          fe.Sensor.Latitude,
		AgeID:             fe.AgeID,
		GenderID:          fe.GenderID,
		GlassID:           fe.GlassID,
		FaceVid:           fe.FaceVid,
		ImageResult:       &models.CutboardImage{CutboardImageURI: fe.CutboardImageURI},
		RelationTypeSlice: fe.FacesIndex.ToRelationSlice(),
	}
}

func (f *FacesIndex) ToRelationSlice() []int {
	ret := []int{}
	if f.RelationTypes&models.DetTypeFace == models.DetTypeFace {
		ret = append(ret, models.DetTypeFace)
	}

	if f.RelationTypes&models.DetTypeVehicle == models.DetTypeVehicle {
		ret = append(ret, models.DetTypeVehicle)
	}

	if f.RelationTypes&models.DetTypeNonmotor == models.DetTypeNonmotor {
		ret = append(ret, models.DetTypeNonmotor)
	}

	if f.RelationTypes&models.DetTypePedestrian == models.DetTypePedestrian {
		ret = append(ret, models.DetTypePedestrian)
	}

	return ret
}

func (fe *FacesIndexExtend) ToFacesIndexExtendModel() *models.FacesIndexExtend {
	model := &models.FacesIndexExtend{}
	faceModel := fe.ToCapturedFace()
	model.CapturedFace = *faceModel

	sensorModel := fe.Sensor.ToModel()
	model.Sensor = *sensorModel

	return model
}

func (fe *FacesIndexExtend) ToModel() *models.CapturedFaceDetial {
	if fe == nil {
		return nil
	}
	entity := fe
	model := &models.CapturedFaceDetial{}
	model.FaceID = entity.FaceID
	model.FaceReID = entity.FaceReID
	model.Timestamp = entity.Ts
	model.SensorID = entity.SensorID
	model.OrgCode = entity.OrgCode
	model.SensorName = entity.Sensor.SensorName
	model.Longitude = entity.Sensor.Longitude
	model.Latitude = entity.Sensor.Latitude
	model.FaceVid = entity.FaceVid
	model.ImageResult = &models.ImageResult{CutboardImageUri: entity.CutboardImageURI}
	model.AgeID = entity.AgeID
	model.GenderID = entity.GenderID
	model.GlassID = entity.GlassID
	return model
}
