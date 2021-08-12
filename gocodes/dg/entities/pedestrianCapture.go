package entities

import (
	"fmt"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	ExplainCountLine          = 50000
	pedestrianCaptureCacheKey = "pedestrian-capture-cache-key-pedestrianid:%v:v1"
)

type Pedestrian struct {
	Ts                int64       `xorm:"ts"`
	RepoID            string      `xorm:"repo_id"`
	SensorID          string      `xorm:"sensor_id"`
	PedestrianID      string      `xorm:"pedestrian_id"`
	PedestrianReID    string      `xorm:"pedestrian_reid"`
	Confidence        float32     `xorm:"confidence"`
	Speed             int32       `xorm:"speed"`
	Direction         int32       `xorm:"direction"`
	FaceID            string      `xorm:"face_id"`
	ImageID           string      `xorm:"image_id"`
	ImageURI          string      `xorm:"image_uri"`
	CutboardImageURI  string      `xorm:"cutboard_image_uri"`
	CutboardX         int64       `xorm:"cutboard_x"`
	CutboardY         int64       `xorm:"cutboard_y"`
	CutboardWidth     int64       `xorm:"cutboard_width"`
	CutboardHeight    int64       `xorm:"cutboard_height"`
	CutboardResWidth  int64       `xorm:"cutboard_res_width"`
	CutboardResHeight int64       `xorm:"cutboard_res_height"`
	GenderID          int32       `xorm:"gender_id"`
	AgeID             int32       `xorm:"age_id"`
	NationID          int32       `xorm:"nation_id"`
	DescEye           int32       `xorm:"desc_eye"`
	DescHead          int32       `xorm:"desc_head"`
	DescMouth         int32       `xorm:"desc_mouth"`
	WithBackpack      int32       `xorm:"with_backpack"`
	WithShoulderBag   int32       `xorm:"with_shoulder_bag"`
	WithHandbag       int32       `xorm:"with_handbag"`
	WithHandCarry     int32       `xorm:"with_hand_carry"`
	WithPram          int32       `xorm:"with_pram"`
	WithLuggage       int32       `xorm:"with_luggage"`
	WithTrolley       int32       `xorm:"with_trolley"`
	WithUmbrella      int32       `xorm:"with_umbrella"`
	WithHoldBaby      int32       `xorm:"with_hold_baby"`
	WithScarf         int32       `xorm:"with_scarf"`
	UpperColor        int32       `xorm:"upper_color"`
	UpperTexture      int32       `xorm:"upper_texture"`
	UpperStyle        int32       `xorm:"upper_style"`
	LowerColor        int32       `xorm:"lower_color"`
	LowerStyle        int32       `xorm:"lower_style"`
	ShoesColor        int32       `xorm:"shoes_color"`
	ShoesStyle        int32       `xorm:"shoes_style"`
	HairStyle         int32       `xorm:"hair_style"`
	RelationTypes     IntArray    `xorm:"text"`
	RelationIDs       StringArray `xorm:"relation_ids text"`
}

type PedestrianCaptureExtend struct {
	Pedestrian `xorm:"extends"`
	Sensor     Sensors `xorm:"extends"`
}

func (*Pedestrian) TableName() string {
	return TableNamePedestrianCapture
}

func (entity *Pedestrian) ToImageResult() models.ImageResult {
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

func (entity *Pedestrian) ToModel() *models.Pedestrian {
	return &models.Pedestrian{
		PedestrianID:   entity.PedestrianID,
		PedestrianReID: entity.PedestrianReID,
		Timestamp:      entity.Ts,
		SensorID:       entity.SensorID,
		AgeID:          entity.AgeID,
		GenderID:       entity.GenderID,
		NationID:       entity.NationID,
		ImageResult:    entity.ToImageResult(),
		UpperColor:     entity.UpperColor,
		LowerColor:     entity.LowerColor,
		HairStyle:      entity.HairStyle,
	}
}

func (entity *PedestrianCaptureExtend) ToModel() *models.Pedestrian {
	return &models.Pedestrian{
		PedestrianID:   entity.PedestrianID,
		PedestrianReID: entity.PedestrianReID,
		Timestamp:      entity.Ts,
		SensorID:       entity.SensorID,
		SensorName:     entity.Sensor.SensorName,
		Latitude:       entity.Sensor.Latitude,
		Longitude:      entity.Sensor.Longitude,
		AgeID:          entity.AgeID,
		GenderID:       entity.GenderID,
		NationID:       entity.NationID,
		ImageResult:    entity.Pedestrian.ToImageResult(),
	}
}

func (entity *Pedestrian) ToDetailModel() *models.PedestrianCaptureDetail {
	return &models.PedestrianCaptureDetail{
		Pedestrian:      *entity.ToModel(),
		FaceID:          entity.FaceID,
		Speed:           entity.Speed,
		Direction:       entity.Direction,
		DescEye:         entity.DescEye,
		DescHead:        entity.DescHead,
		DescMouth:       entity.DescMouth,
		WithBackpack:    entity.WithBackpack,
		WithShoulderBag: entity.WithShoulderBag,
		WithHandbag:     entity.WithHandbag,
		WithHandCarry:   entity.WithHandCarry,
		WithPram:        entity.WithPram,
		WithLuggage:     entity.WithLuggage,
		WithTrolley:     entity.WithTrolley,
		WithUmbrella:    entity.WithUmbrella,
		WithHoldBaby:    entity.WithHoldBaby,
		WithScarf:       entity.WithScarf,
		UpperColor:      entity.UpperColor,
		UpperTexture:    entity.UpperTexture,
		UpperStyle:      entity.UpperStyle,
		LowerColor:      entity.LowerColor,
		LowerStyle:      entity.LowerStyle,
		ShoesColor:      entity.ShoesColor,
		ShoesStyle:      entity.ShoesStyle,
		HairStyle:       entity.HairStyle,
		RelationIDs:     entity.RelationIDs,
		RelationTypes:   entity.RelationTypes,
	}
}

func (entity *Pedestrian) CacheKey() string {
	return fmt.Sprintf(pedestrianCaptureCacheKey, entity.PedestrianID)
}

func (entity *Pedestrian) CacheExpireTime() int64 {
	return redis.OneWeek * 2
}

func PedestrianCaptureHelper(IDs []string) []redis.Entity {
	slice := make([]redis.Entity, len(IDs))
	for i, ID := range IDs {
		slice[i] = &Pedestrian{PedestrianID: ID}
	}
	return slice
}
