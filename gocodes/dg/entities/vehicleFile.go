package entities

import (
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/json"
)

type VehicleFile struct {
	Vid          string    `xorm:"not null pk varchar(36)"`
	Ts           int64     `xorm:"not null default 0 bigint"`
	Uts          time.Time `xorm:"not null default 'now()' datetime"`
	Comment      string    `xorm:"default '''::text' text"`
	ImageURL     string    `xorm:"image_url not null text"`
	PlateText    string    `xorm:"default '''::text' text"`
	PlateColorID int32     `xorm:"plate_color_id"`
	ColorID      int32     `xorm:"color_id"`
	BrandID      int32     `xorm:"brand_id"`
	SubBrandID   int32     `xorm:"sub_brand_id"`
	ModelYearID  int32     `xorm:"model_year_id"`
	TypeID       int32     `xorm:"type_id"`
	Tag          int       `xorm:"default 0 smallint"`
	Status       int       `xorm:"not null default 1 smallint"`
}

type VehicleFileExtend struct {
	VehicleFile     `xorm:"extends"`
	CapturedVehicle VehiclesIndexExtend `xorm:"extends"`
}

func (vf VehicleFile) TableName() string {
	return TableNameVehicleFile
}

func (vf *VehicleFile) ToModel() *models.VehicleFile {
	if vf == nil {
		return nil
	}

	var systemTagNames []string
	var systemTagValues []int
	var customTags []string
	if vf.Comment != "" {
		commentJSON := new(models.VehicleFileTag)
		err := json.Unmarshal([]byte(vf.Comment), commentJSON)
		if err == nil {
			systemTagNames = commentJSON.SystemName
			systemTagValues = commentJSON.SystemValue
			customTags = commentJSON.Custom
		}
	}

	if systemTagNames == nil {
		systemTagNames = []string{}
	}

	if systemTagValues == nil {
		systemTagValues = []int{}
	}

	if customTags == nil {
		customTags = []string{}
	}

	return &models.VehicleFile{
		Vid:             vf.Vid,
		Ts:              vf.Ts,
		Uts:             vf.Uts,
		Comment:         vf.Comment,
		ImageURL:        vf.ImageURL,
		PlateText:       vf.PlateText,
		PlateColorID:    vf.PlateColorID,
		ColorID:         vf.ColorID,
		BrandID:         vf.BrandID,
		SubBrandID:      vf.SubBrandID,
		ModelYearID:     vf.ModelYearID,
		TypeID:          vf.TypeID,
		Tag:             vf.Tag,
		Status:          vf.Status,
		SystemTagNames:  systemTagNames,
		SystemTagValues: systemTagValues,
		CustomTags:      customTags,
	}
}

func VehicleFileModelToEntity(model *models.VehicleFile, systemTagMap map[int]string) *VehicleFile {
	if model == nil {
		return nil
	}

	var comment string
	var tag int

	commentJSON := new(models.VehicleFileTag)
	// if model.Comment != "" {
	// 	err := json.Unmarshal([]byte(model.Comment), commentJSON)
	// if err == nil {
	commentJSON.Custom = model.CustomTags
	commentJSON.SystemValue = model.SystemTagValues

	nameSlice := make([]string, len(model.SystemTagValues))
	for index, v := range model.SystemTagValues {
		tag += v
		nameSlice[index] = systemTagMap[v]
	}

	commentJSON.SystemName = nameSlice
	// 	}
	// }

	byteComment, err := json.Marshal(commentJSON)
	if err == nil {
		comment = string(byteComment)
	}

	return &VehicleFile{
		Vid:          model.Vid,
		Ts:           model.Ts,
		Uts:          time.Now(),
		Comment:      comment,
		ImageURL:     model.ImageURL,
		PlateText:    model.PlateText,
		PlateColorID: model.PlateColorID,
		ColorID:      model.ColorID,
		BrandID:      model.BrandID,
		SubBrandID:   model.SubBrandID,
		ModelYearID:  model.ModelYearID,
		TypeID:       model.TypeID,
		Tag:          tag,
		Status:       model.Status,
	}
}
