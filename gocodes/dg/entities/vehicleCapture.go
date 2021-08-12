package entities

import (
	"fmt"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	vehicleCaptureCacheKey = "vehicle-capture-cache-key-vehicleid:%v:v2"
)

type Vehicles struct {
	Ts                  int64              `xorm:"ts"`
	RepoID              string             `xorm:"repo_id"`
	SensorID            string             `xorm:"sensor_id"`
	VehicleID           string             `xorm:"vehicle_id"`
	VehicleReID         string             `xorm:"vehicle_reid"`
	Confidence          float32            `xorm:"confidence"`
	Speed               int32              `xorm:"speed"`
	Direction           int32              `xorm:"direction"`
	FaceID              string             `xorm:"face_id"`
	ImageURI            string             `xorm:"image_uri"`
	CutboardImageURI    string             `xorm:"cutboard_image_uri"`
	CutboardX           int32              `xorm:"cutboard_x"`
	CutboardY           int32              `xorm:"cutboard_y"`
	CutboardWidth       int32              `xorm:"cutboard_width"`
	CutboardHeight      int32              `xorm:"cutboard_height"`
	CutboardResWidth    int32              `xorm:"cutboard_res_width"`
	CutboardResHeight   int32              `xorm:"cutboard_res_height"`
	BrandID             int32              `xorm:"brand_id"`
	SubBrandID          int32              `xorm:"sub_brand_id"`
	ModelYearID         int32              `xorm:"model_year_id"`
	TypeID              int32              `xorm:"type_id"`
	Side                int32              `xorm:"side"`
	ColorID             int32              `xorm:"color_id"`
	PlateText           string             `xorm:"plate_text"`
	PlateTypeID         int32              `xorm:"plate_type_id"`
	PlateColorID        int32              `xorm:"plate_color_id"`
	SymbolInt           int64              `xorm:"symbol_int"`
	SymbolStr           string             `xorm:"symbol_str"`
	IllegalInt          int64              `xorm:"illegal_int"`
	IllegalStr          string             `xorm:"illegal_str"`
	CoillegalInt        int64              `xorm:"coillegal_int"`
	CoillegalStr        string             `xorm:"coillegal_str"`
	SpecialInt          int64              `xorm:"special_int"`
	SpecialStr          string             `xorm:"special_str"`
	DriverOnThePhone    bool               `xorm:"driver_on_the_phone"`
	DriverWithoutBelt   bool               `xorm:"driver_without_belt"`
	CodriverWithoutBelt bool               `xorm:"codriver_without_belt"`
	Lane                string             `xorm:"lane"` // TODO 比较复杂-稍后再看
	Symbols             models.SymbolSlice `xorm:"json"`
	SymbolsDesc         string             `xorm:"symbols_desc"`
	AliasTableName      string             `xorm:"-"`
	RelationTypes       IntArray           `xorm:"text"`
	RelationIDs         StringArray        `xorm:"relation_ids text"`
}

type VehicleCaptureExtend struct {
	Vehicles `xorm:"extends"`
	Sensor   Sensors `xorm:"extends"`
}

func (entity *Vehicles) TableName() string {
	if entity.AliasTableName != "" {
		return entity.AliasTableName
	}

	return TableNameVehicleCaptureEntity
}

func (entity *Vehicles) ToImageResult() *models.ImageResult {
	return &models.ImageResult{
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

func (entity *Vehicles) ToModel() *models.VehicleCaptureDetail {
	return &models.VehicleCaptureDetail{
		VehicleCapture: models.VehicleCapture{
			Timestamp:    entity.Ts,
			SensorID:     entity.SensorID,
			VehicleID:    entity.VehicleID,
			VehicleReID:  entity.VehicleReID,
			PlateText:    entity.PlateText,
			PlateColorID: entity.PlateColorID,
			PlateTypeID:  entity.PlateTypeID,
		},
		ImageResult:   entity.ToImageResult(),
		Symbols:       entity.Symbols,
		SymbolsDesc:   entity.SymbolsDesc,
		Specials:      models.SymbolIntToSlice(entity.SpecialInt),
		Speed:         entity.Speed,
		Direction:     entity.Direction,
		FaceID:        entity.FaceID,
		BrandID:       entity.BrandID,
		SubBrandID:    entity.SubBrandID,
		ModelYearID:   entity.ModelYearID,
		TypeID:        entity.TypeID,
		Side:          entity.Side,
		ColorID:       entity.ColorID,
		Lane:          entity.Lane,
		RelationIDs:   entity.RelationIDs,
		RelationTypes: entity.RelationTypes,
	}
}

func (entity VehicleCaptureExtend) ToModel() *models.VehicleCaptureDetail {
	return &models.VehicleCaptureDetail{
		VehicleCapture: models.VehicleCapture{
			Timestamp:    entity.Ts,
			SensorID:     entity.SensorID,
			SensorName:   entity.Sensor.SensorName,
			Longitude:    entity.Sensor.Longitude,
			Latitude:     entity.Sensor.Latitude,
			VehicleID:    entity.VehicleID,
			VehicleReID:  entity.VehicleReID,
			PlateText:    entity.PlateText,
			PlateColorID: entity.PlateColorID,
			PlateTypeID:  entity.PlateTypeID,
		},
		ImageResult:   entity.ToImageResult(),
		SensorType:    entity.Sensor.Type,
		SensorURL:     entity.Sensor.Url,
		Symbols:       entity.Symbols,
		SymbolsDesc:   entity.SymbolsDesc,
		Specials:      models.SymbolIntToSlice(entity.SpecialInt),
		Speed:         entity.Speed,
		Direction:     entity.Direction,
		FaceID:        entity.FaceID,
		BrandID:       entity.BrandID,
		SubBrandID:    entity.SubBrandID,
		ModelYearID:   entity.ModelYearID,
		TypeID:        entity.TypeID,
		Side:          entity.Side,
		ColorID:       entity.ColorID,
		Lane:          entity.Lane,
		RelationIDs:   entity.RelationIDs,
		RelationTypes: entity.RelationTypes,
	}
}

func (entity *VehicleCaptureExtend) ToVehicleIndexModel() *models.VehicleCapture {
	if entity == nil {
		return nil
	}

	return &models.VehicleCapture{
		Timestamp:    entity.Ts,
		SensorID:     entity.SensorID,
		SensorName:   entity.Sensor.SensorName,
		Longitude:    entity.Sensor.Longitude,
		Latitude:     entity.Sensor.Latitude,
		VehicleID:    entity.VehicleID,
		VehicleReID:  entity.VehicleReID,
		VehicleVID:   "",
		ImageResult:  &models.CutboardImage{CutboardImageURI: entity.CutboardImageURI},
		PlateText:    entity.PlateText,
		PlateColorID: entity.PlateColorID,
		PlateTypeID:  entity.PlateTypeID,
	}
}

func (entity *Vehicles) CacheKey() string {
	return fmt.Sprintf(vehicleCaptureCacheKey, entity.VehicleID)
}

func (entity *Vehicles) CacheExpireTime() int64 {
	return redis.OneWeek * 2
}

func VehicleCaptureHelper(IDs []string) []redis.Entity {
	slice := make([]redis.Entity, len(IDs))
	for i, ID := range IDs {
		slice[i] = &Vehicles{VehicleID: ID}
	}
	return slice
}

func VehicleCaptureHelperV2(ids []string, extra ...string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, id := range ids {
		slice[i] = &Vehicles{VehicleID: id, AliasTableName: extra[0]}
	}
	return slice
}

type VehicleCaptureIndex struct {
	VehicleID        string   `xorm:"vehicle_id"`
	VehicleReID      string   `xorm:"vehicle_reid"`
	VehicleVID       string   `xorm:"vehicle_vid"`
	RelationTypes    IntArray `xorm:"text"`
	Ts               int64    `xorm:"ts"`
	SensorID         string   `xorm:"sensor_id"`
	SensorIntID      int64    `xorm:"sensor_int_id"`
	OrgCode          int64    `xorm:"org_code"`
	Speed            int32    `xorm:"speed"`
	Direction        int32    `xorm:"direction"`
	CutboardImageURI string   `xorm:"cutboard_image_uri"`
	BrandID          int32    `xorm:"brand_id"`
	SubBrandID       int32    `xorm:"sub_brand_id"`
	ModelYearID      int32    `xorm:"model_year_id"`
	TypeID           int32    `xorm:"type_id"`
	Side             int32    `xorm:"side"`
	ColorID          int32    `xorm:"color_id"`
	PlateText        string   `xorm:"plate_text"`
	PlateTypeID      int32    `xorm:"plate_type_id"`
	PlateColorID     int32    `xorm:"plate_color_id"`
	//Symbols              `xorm:"json"`
	//Specials             . `xorm:"json"`
}
