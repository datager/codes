package entities

import (
	"codes/gocodes/dg/models"
)

const ()

type NonmotorIndex struct {
	NonmotorID       string   `xorm:"nonmotor_id"`
	NonmotorReID     string   `xorm:"nonmotor_reid"`
	SensorID         string   `xorm:"sensor_id"`
	SensorIntID      int32    `xorm:"sensor_int_id"`
	OrgCode          int64    `xorm:"org_code not null default 0"`
	Ts               int64    `xorm:"ts"`
	Speed            int32    `xorm:"speed"`
	Direction        int32    `xorm:"direction"`
	CutboardImageURI string   `xorm:"cutboard_image_uri"`
	NonmotorType     int32    `xorm:"nonmotor_type"`
	NonmotorGesture  int32    `xorm:"nonmotor_gesture"`
	NonmotorColorID  int32    `xorm:"nonmotor_color_id"`
	PlateColorID     int32    `xorm:"plate_color_id"`
	PlateTypeID      int32    `xorm:"plate_type_id"`
	GenderID         int32    `xorm:"gender_id"`
	AgeID            int32    `xorm:"age_id"`
	NationID         int32    `xorm:"nation_id"`
	HeadItems        IntArray `xorm:"text"`
	Accessories      IntArray `xorm:"text"`
	Bags             IntArray `xorm:"text"`
	UpperColor       int32    `xorm:"upper_color"`
	UpperStyle       int32    `xorm:"upper_style"`
	PlateText        string   `xorm:"plate_text"`
	AliasTableName   string   `xorm:"-"`
	RelationTypes    int      `xorm:"not null default 0"`
}

type NonmotorCaptureIndexExtend struct {
	NonmotorIndex `xorm:"extends"`
	Sensor        Sensors `xorm:"extends"`
}

func (entity *NonmotorIndex) TableName() string {
	return TableNameNonmotorIndexCapture
}

func (entity *NonmotorIndex) ToModel() *models.NonmotorIndex {
	nonMotor := &models.NonmotorIndex{
		Timestamp:         entity.Ts,
		Speed:             entity.Speed,
		Direction:         entity.Direction,
		ImageResult:       &models.CutboardImage{CutboardImageURI: entity.CutboardImageURI},
		SensorID:          entity.SensorID,
		SensorIntID:       entity.SensorIntID,
		OrgCode:           entity.OrgCode,
		NonmotorID:        entity.NonmotorID,
		NonmotorReID:      entity.NonmotorReID,
		NonmotorColorID:   entity.NonmotorColorID,
		UpperColor:        entity.UpperColor,
		UpperStyle:        entity.UpperStyle,
		RelationTypeSlice: entity.ToRelationSlice(),
	}
	return nonMotor
}

func (entity *NonmotorIndex) ToRelationSlice() []int {
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

func (entity *NonmotorCaptureIndexExtend) ToModel() *models.NonmotorIndex {
	nonMotor := &models.NonmotorIndex{
		Timestamp:         entity.Ts,
		Speed:             entity.Speed,
		Direction:         entity.Direction,
		NonmotorColorID:   entity.NonmotorColorID,
		UpperColor:        entity.UpperColor,
		UpperStyle:        entity.UpperStyle,
		ImageResult:       &models.CutboardImage{CutboardImageURI: entity.CutboardImageURI},
		SensorID:          entity.SensorID,
		SensorIntID:       entity.SensorIntID,
		OrgCode:           entity.OrgCode,
		SensorName:        entity.Sensor.SensorName,
		Latitude:          entity.Sensor.Latitude,
		Longitude:         entity.Sensor.Longitude,
		NonmotorID:        entity.NonmotorID,
		NonmotorReID:      entity.NonmotorReID,
		RelationTypeSlice: entity.NonmotorIndex.ToRelationSlice(),
	}
	return nonMotor
}
