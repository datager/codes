package entities

import (
	"codes/gocodes/dg/models"
)

type VehiclesIndex struct {
	VehicleID        string   `xorm:"vehicle_id pk VARCHAR(36)"`
	VehicleReID      string   `xorm:"vehicle_reid"`
	VehicleVid       string   `xorm:"vehicle_vid"`
	Ts               int64    `xorm:"ts"`
	SensorID         string   `xorm:"sensor_id"`
	SensorIntID      int32    `xorm:"sensor_int_id"`
	OrgCode          int64    `xorm:"org_code not null default 0"`
	Speed            int      `xorm:"speed"`
	Direction        int      `xorm:"direction"`
	CutboardImageURI string   `xorm:"cutboard_image_uri"`
	BrandID          int      `xorm:"brand_id"`
	SubBrandID       int      `xorm:"sub_brand_id"`
	ModelYearID      int      `xorm:"model_year_id"`
	TypeID           int      `xorm:"type_id"`
	Side             int      `xorm:"side"`
	ColorID          int      `xorm:"color_id"`
	PlateText        string   `xorm:"plate_text"`
	PlateTypeID      int      `xorm:"plate_type_id"`
	PlateColorID     int      `xorm:"plate_color_id"`
	Symbols          IntArray `xorm:"text" json:"symbols" form:"symbols[]"`
	Specials         IntArray `xorm:"text" json:"specials" form:"specials[]"`
	RelationTypes    int      `xorm:"not null default 0"`
}

type VehiclesIndexExtend struct {
	VehiclesIndex `xorm:"extends"`
	Sensor        Sensors `xorm:"extends"`
}

func (VehiclesIndex) TableName() string {
	return TableNameVehicleCaptureIndexEntity
}

type VehicleAnalyzeData struct {
	VehicleVid  string `xorm:"vehicle_vid"`
	VehicleReID string `xorm:"vehicle_reid"`
	SensorID    string `xorm:"sensor_id"`
	Ts          int64
	Longitude   float64
	Latitude    float64
}

func (vd *VehicleAnalyzeData) ToModel() *models.VehicleAnalyzeData {
	if vd == nil {
		return nil
	}

	return &models.VehicleAnalyzeData{
		VehicleVid:  vd.VehicleVid,
		VehicleReID: vd.VehicleReID,
		SensorID:    vd.SensorID,
		Timestamp:   vd.Ts,
		Longitude:   vd.Longitude,
		Latitude:    vd.Latitude,
	}
}

func (vci *VehiclesIndex) ToModel() *models.VehicleCapture {
	return &models.VehicleCapture{
		Timestamp:         vci.Ts,
		SensorID:          vci.SensorID,
		SensorIntID:       vci.SensorIntID,
		OrgCode:           vci.OrgCode,
		VehicleID:         vci.VehicleID,
		VehicleVID:        vci.VehicleVid,
		VehicleReID:       vci.VehicleReID,
		ImageResult:       &models.CutboardImage{CutboardImageURI: vci.CutboardImageURI},
		PlateText:         vci.PlateText,
		PlateColorID:      int32(vci.PlateColorID),
		PlateTypeID:       int32(vci.PlateTypeID),
		ColorID:           int32(vci.ColorID),
		TypeID:            int32(vci.TypeID),
		BrandID:           int32(vci.BrandID),
		SubBrandID:        int32(vci.SubBrandID),
		ModelYearID:       int32(vci.ModelYearID),
		Symbols:           vci.Symbols,
		Specials:          vci.Specials,
		RelationTypeSlice: vci.ToRelationSlice(),
	}
}

func (vci *VehiclesIndex) ToRelationSlice() []int {
	ret := []int{}
	if vci.RelationTypes&models.DetTypeFace == models.DetTypeFace {
		ret = append(ret, models.DetTypeFace)
	}

	if vci.RelationTypes&models.DetTypeVehicle == models.DetTypeVehicle {
		ret = append(ret, models.DetTypeVehicle)
	}

	if vci.RelationTypes&models.DetTypeNonmotor == models.DetTypeNonmotor {
		ret = append(ret, models.DetTypeNonmotor)
	}

	if vci.RelationTypes&models.DetTypePedestrian == models.DetTypePedestrian {
		ret = append(ret, models.DetTypePedestrian)
	}

	return ret
}

func (ve *VehiclesIndexExtend) ToModel() *models.VehicleCapture {
	if ve == nil {
		return nil
	}

	return &models.VehicleCapture{
		Timestamp:         ve.Ts,
		SensorID:          ve.SensorID,
		SensorIntID:       ve.SensorIntID,
		OrgCode:           ve.OrgCode,
		SensorName:        ve.Sensor.SensorName,
		Longitude:         ve.Sensor.Longitude,
		Latitude:          ve.Sensor.Latitude,
		VehicleID:         ve.VehicleID,
		VehicleVID:        ve.VehicleVid,
		VehicleReID:       ve.VehicleReID,
		ImageResult:       &models.CutboardImage{CutboardImageURI: ve.CutboardImageURI},
		PlateText:         ve.PlateText,
		PlateColorID:      int32(ve.PlateColorID),
		PlateTypeID:       int32(ve.PlateTypeID),
		ColorID:           int32(ve.ColorID),
		TypeID:            int32(ve.TypeID),
		BrandID:           int32(ve.BrandID),
		SubBrandID:        int32(ve.SubBrandID),
		ModelYearID:       int32(ve.ModelYearID),
		Symbols:           ve.Symbols,
		Specials:          ve.Specials,
		RelationTypeSlice: ve.VehiclesIndex.ToRelationSlice(),
	}
}

func (ve *VehiclesIndexExtend) ToVehicleCaptureMode() *models.VehicleCapture {
	if ve == nil {
		return nil
	}

	return &models.VehicleCapture{
		Timestamp:         ve.Ts,
		SensorID:          ve.SensorID,
		SensorIntID:       ve.SensorIntID,
		OrgCode:           ve.OrgCode,
		SensorName:        ve.Sensor.SensorName,
		Longitude:         ve.Sensor.Longitude,
		Latitude:          ve.Sensor.Latitude,
		VehicleID:         ve.VehicleID,
		VehicleReID:       ve.VehicleReID,
		VehicleVID:        ve.VehicleVid,
		ImageResult:       &models.CutboardImage{CutboardImageURI: ve.CutboardImageURI},
		PlateText:         ve.PlateText,
		PlateColorID:      int32(ve.PlateColorID),
		PlateTypeID:       int32(ve.PlateTypeID),
		ColorID:           int32(ve.ColorID),
		TypeID:            int32(ve.TypeID),
		BrandID:           int32(ve.BrandID),
		SubBrandID:        int32(ve.SubBrandID),
		ModelYearID:       int32(ve.ModelYearID),
		RelationTypeSlice: ve.VehiclesIndex.ToRelationSlice(),
	}
}

func (vci *VehiclesIndex) ToMiniVehicle() *models.MiniVehicle {
	return &models.MiniVehicle{
		Timestamp:   vci.Ts,
		VehicleID:   vci.VehicleID,
		VehicleVID:  vci.VehicleVid,
		VehicleReID: vci.VehicleReID,
		ImageResult: models.CutboardImage{CutboardImageURI: vci.CutboardImageURI},
		PlateText:   vci.PlateText,
	}
}

func (ve *VehiclesIndexExtend) ToExtend() *models.VehiclesIndexExtend {
	if ve == nil {
		return nil
	}

	return &models.VehiclesIndexExtend{
		VehicleCapture: *(ve.VehiclesIndex.ToModel()),
		Sensor:         *(ve.Sensor.ToModel()),
	}
}
