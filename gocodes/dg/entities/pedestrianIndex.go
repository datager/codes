package entities

import (
	"codes/gocodes/dg/models"
)

type PedestrianIndex struct {
	PedestrianID     string   `xorm:"pedestrian_id"`
	PedestrianReID   string   `xorm:"pedestrian_reid"`
	Ts               int64    `xorm:"ts"`
	SensorID         string   `xorm:"sensor_id"`
	SensorIntID      int32    `xorm:"sensor_int_id"`
	OrgCode          int64    `xorm:"org_code not null default 0"`
	Speed            int      `xorm:"speed"`
	Direction        int      `xorm:"direction"`
	CutboardImageURI string   `xorm:"cutboard_image_uri"`
	GenderID         int      `xorm:"gender_id"`
	AgeID            int      `xorm:"age_id"`
	NationID         int      `xorm:"nation_id"`
	HeadItems        IntArray `xorm:"text"`
	Accessories      IntArray `xorm:"text"`
	Bags             IntArray `xorm:"text"`
	UpperColor       int      `xorm:"upper_color"`
	UpperTexture     int      `xorm:"upper_texture"`
	UpperStyle       int      `xorm:"upper_style"`
	LowerColor       int      `xorm:"lower_color"`
	LowerStyle       int      `xorm:"lower_style"`
	ShoesColor       int      `xorm:"shoes_color"`
	ShoesStyle       int      `xorm:"shoes_style"`
	HairStyle        int      `xorm:"hair_style"`
	RelationTypes    int      `xorm:"not null default 0"`
}

type PedestrianIndexExtend struct {
	PedestrianIndex `xorm:"extends"`
	Sensor          Sensors `xorm:"extends"`
}

func (*PedestrianIndex) TableName() string {
	return TableNamePedestrianIndexCapture
}

func (entity *PedestrianIndex) ToModel() *models.PedestrianIndex {
	return &models.PedestrianIndex{
		PedestrianID:      entity.PedestrianID,
		PedestrianReID:    entity.PedestrianReID,
		Timestamp:         entity.Ts,
		SensorID:          entity.SensorID,
		SensorIntID:       entity.SensorIntID,
		OrgCode:           entity.OrgCode,
		AgeID:             entity.AgeID,
		GenderID:          entity.GenderID,
		NationID:          entity.NationID,
		ImageResult:       &models.CutboardImage{CutboardImageURI: entity.CutboardImageURI},
		UpperColor:        entity.UpperColor,
		LowerColor:        entity.LowerColor,
		HairStyle:         entity.HairStyle,
		RelationTypeSlice: entity.ToRelationSlice(),
	}
}

func (entity *PedestrianIndex) ToRelationSlice() []int {
	ret := []int{}
	if entity.RelationTypes&models.DetTypeFace == models.DetTypeFace {
		ret = append(ret, models.DetTypeFace)
	}

	if entity.RelationTypes&models.DetTypeVehicle == models.DetTypeVehicle {
		ret = append(ret, models.DetTypeVehicle)
	}

	if entity.RelationTypes&models.DetTypeNonmotor == models.DetTypeNonmotor {
		ret = append(ret, models.DetTypeNonmotor)
	}

	if entity.RelationTypes&models.DetTypePedestrian == models.DetTypePedestrian {
		ret = append(ret, models.DetTypePedestrian)
	}

	return ret
}

func (entity *PedestrianIndexExtend) ToModel() *models.PedestrianIndex {
	return &models.PedestrianIndex{
		PedestrianID:      entity.PedestrianID,
		PedestrianReID:    entity.PedestrianReID,
		Timestamp:         entity.Ts,
		SensorID:          entity.SensorID,
		SensorIntID:       entity.SensorIntID,
		OrgCode:           entity.OrgCode,
		SensorName:        entity.Sensor.SensorName,
		Latitude:          entity.Sensor.Latitude,
		Longitude:         entity.Sensor.Longitude,
		AgeID:             entity.AgeID,
		GenderID:          entity.GenderID,
		NationID:          entity.NationID,
		ImageResult:       &models.CutboardImage{CutboardImageURI: entity.CutboardImageURI},
		UpperColor:        entity.UpperColor,
		LowerColor:        entity.LowerColor,
		HairStyle:         entity.HairStyle,
		RelationTypeSlice: entity.PedestrianIndex.ToRelationSlice(),
	}
}
