package entities

import (
	"fmt"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	nonmotorCaptureCacheKey = "nonmotor-capture-cache-key-nonmotorid:%v:v1"
)

type Nonmotor struct {
	Ts                int64       `xorm:"ts"`
	Confidence        float32     `xorm:"confidence"`
	Speed             int32       `xorm:"speed"`
	Direction         int32       `xorm:"direction"`
	CutboardX         int32       `xorm:"cutboard_x"`
	CutboardY         int32       `xorm:"cutboard_y"`
	CutboardWidth     int32       `xorm:"cutboard_width"`
	CutboardHeight    int32       `xorm:"cutboard_height"`
	CutboardResWidth  int32       `xorm:"cutboard_res_width"`
	CutboardResHeight int32       `xorm:"cutboard_res_height"`
	NonmotorType      int32       `xorm:"nonmotor_type"`
	NonmotorGesture   int32       `xorm:"nonmotor_gesture"`
	NonmotorColorID   int32       `xorm:"nonmotor_color_id"`
	PlateColorID      int32       `xorm:"plate_color_id"`
	PlateTypeID       int32       `xorm:"plate_type_id"`
	GenderID          int32       `xorm:"gender_id"`
	AgeID             int32       `xorm:"age_id"`
	NationID          int32       `xorm:"nation_id"`
	DescEye           int32       `xorm:"desc_eye"`
	DescHead          int32       `xorm:"desc_head"`
	DescMouth         int32       `xorm:"desc_mouth"`
	WithBackpack      int32       `xorm:"with_backpack"`
	WithShoulderBag   int32       `xorm:"with_shoulder_bag"`
	UpperColor        int32       `xorm:"upper_color"`
	UpperStyle        int32       `xorm:"upper_style"`
	RepoID            string      `xorm:"repo_id"`
	SensorID          string      `xorm:"sensor_id"`
	NonmotorID        string      `xorm:"nonmotor_id"`
	NonmotorReID      string      `xorm:"nonmotor_reid"`
	FaceID            string      `xorm:"face_id"`
	ImageURI          string      `xorm:"image_uri"`
	CutboardImageURI  string      `xorm:"cutboard_image_uri"`
	PlateText         string      `xorm:"plate_text"`
	Lane              string      `xorm:"lane"`
	AliasTableName    string      `xorm:"-"`
	RelationTypes     IntArray    `xorm:"text"`
	RelationIDs       StringArray `xorm:"relation_ids text"`
}

type NonmotorCaptureExtend struct {
	Nonmotor `xorm:"extends"`
	Sensor   Sensors `xorm:"extends"`
}

func (entity *Nonmotor) TableName() string {
	return TableNameNonmotorCapture
}

func (entity *Nonmotor) ToImageResult() models.ImageResult {
	return models.ImageResult{
		ImageUri:          entity.ImageURI,
		CutboardImageUri:  entity.CutboardImageURI,
		CutboardX:         int(entity.CutboardX),
		CutboardY:         int(entity.CutboardY),
		CutboardWidth:     int(entity.CutboardWidth),
		CutboardHeight:    int(entity.CutboardHeight),
		CutboardResWidth:  int(entity.CutboardResWidth),
		CutboardResHeight: int(entity.CutboardResHeight),
		Confidence:        entity.Confidence,
	}
}

func (entity *Nonmotor) ToModel() *models.Nonmotor {
	nonMotor := &models.Nonmotor{
		Timestamp:       entity.Ts,
		Speed:           entity.Speed,
		Direction:       entity.Direction,
		ImageResult:     entity.ToImageResult(),
		RepoID:          entity.RepoID,
		SensorID:        entity.SensorID,
		NonmotorID:      entity.NonmotorID,
		NonmotorReID:    entity.NonmotorReID,
		NonmotorColorID: entity.NonmotorColorID,
		UpperColor:      entity.UpperColor,
		UpperStyle:      entity.UpperStyle,
	}
	return nonMotor
}

func (entity *NonmotorCaptureExtend) ToModel() *models.Nonmotor {
	nonMotor := &models.Nonmotor{
		Timestamp:    entity.Ts,
		Speed:        entity.Speed,
		Direction:    entity.Direction,
		ImageResult:  entity.Nonmotor.ToImageResult(),
		RepoID:       entity.RepoID,
		SensorID:     entity.SensorID,
		SensorName:   entity.Sensor.SensorName,
		Latitude:     entity.Sensor.Latitude,
		Longitude:    entity.Sensor.Longitude,
		NonmotorID:   entity.NonmotorID,
		NonmotorReID: entity.NonmotorReID,
	}
	return nonMotor
}

func (entity *Nonmotor) ToDetailModel() *models.NonmotorCaptureDetail {
	nonMotorDetail := &models.NonmotorCaptureDetail{
		Nonmotor:        *entity.ToModel(),
		PlateText:       entity.PlateText,
		PlateTypeID:     entity.PlateTypeID,
		PlateColorID:    entity.PlateColorID,
		FaceID:          entity.FaceID,
		NonmotorType:    entity.NonmotorType,
		NonmotorGesture: entity.NonmotorGesture,
		NonmotorColorID: entity.NonmotorColorID,
		GenderID:        entity.GenderID,
		AgeID:           entity.AgeID,
		NationID:        entity.NationID,
		DescEye:         entity.DescEye,
		DescHead:        entity.DescHead,
		DescMouth:       entity.DescMouth,
		WithBackpack:    entity.WithBackpack,
		WithShoulderBag: entity.WithShoulderBag,
		UpperColor:      entity.UpperColor,
		UpperStyle:      entity.UpperStyle,
		RelationIDs:     entity.RelationIDs,
		RelationTypes:   entity.RelationTypes,
	}
	return nonMotorDetail
}

func (entity *Nonmotor) CacheKey() string {
	return fmt.Sprintf(nonmotorCaptureCacheKey, entity.NonmotorID)
}

func (entity *Nonmotor) CacheExpireTime() int64 {
	return redis.OneWeek * 2
}

func NonmotorCaptureHelper(IDs []string) []redis.Entity {
	slice := make([]redis.Entity, len(IDs))
	for i, ID := range IDs {
		slice[i] = &Nonmotor{NonmotorID: ID}
	}
	return slice
}

func NonmotorCaptureHelperV2(ids []string, extra ...string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))

	for i, id := range ids {
		slice[i] = &Nonmotor{NonmotorID: id, AliasTableName: extra[0]}
	}
	return slice
}
