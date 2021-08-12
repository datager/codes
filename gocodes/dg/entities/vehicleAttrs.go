package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type VehicleAttrs struct {
	ID           string    `xorm:"id not null pk index varchar(1024)"`
	RepoID       string    `xorm:"repo_id varchar(1024) not null default ''"`
	OrgID        string    `xorm:"org_id varchar(1024) not null default ''"`
	PlateText    string    `xorm:"plate_text not null index varchar(100)"`
	PlateColorID string    `xorm:"plate_color_id"`
	ColorID      string    `xorm:"color_id"`
	BrandID      string    `xorm:"brand_id"`
	SubBrandID   string    `xorm:"sub_brand_id"`
	ModelYearID  string    `xorm:"model_year_id"`
	OwnerName    string    `xorm:"owner_name varchar(100)"`
	PhoneNumber  string    `xorm:"phone_number varchar(100)"`
	IDType       int       `xorm:"id_type default 0 index smallint"`
	IDNo         string    `xorm:"id_no default '''::character varying' index varchar(1024)"`
	Comment      string    `xorm:"default '''::text' TEXT"`
	Ts           int64     `xorm:"not null default 0 BIGINT"`
	Uts          time.Time `xorm:"not null default 'now()' DATETIME"`
}

type VehicleAttrsExtend struct {
	VehicleAttrs `xorm:"extends"`
	Repo         FaceRepos `xorm:"extends"`
}

func (v VehicleAttrs) TableName() string {
	return TableNameVehicleAttrs
}

func (e *VehicleAttrsExtend) ToExtend() *models.VehicleAttrsExtend {
	if e == nil {
		return nil
	}

	return &models.VehicleAttrsExtend{
		Attr: e.VehicleAttrs.ToModel(),
		Repo: e.Repo.ToModel(),
	}
}

func (v *VehicleAttrs) ToModel() *models.VehicleAttrs {
	if v == nil {
		return nil
	}

	return &models.VehicleAttrs{
		ID:           v.ID,
		RepoID:       v.RepoID,
		OrgID:        v.OrgID,
		PlateText:    v.PlateText,
		PlateColorID: v.PlateColorID,
		ColorID:      v.ColorID,
		BrandID:      v.BrandID,
		SubBrandID:   v.SubBrandID,
		ModelYearID:  v.ModelYearID,
		OwnerName:    v.OwnerName,
		PhoneNumber:  v.PhoneNumber,
		IDType:       v.IDType,
		IDNo:         v.IDNo,
		Comment:      v.Comment,
		Ts:           v.Ts,
		Uts:          v.Uts,
	}
}

func VehicleAttrsModelToEntity(model *models.VehicleAttrs) *VehicleAttrs {
	if model == nil {
		return nil
	}

	return &VehicleAttrs{
		ID:           model.ID,
		RepoID:       model.RepoID,
		OrgID:        model.OrgID,
		PlateText:    model.PlateText,
		PlateColorID: model.PlateColorID,
		ColorID:      model.ColorID,
		BrandID:      model.BrandID,
		SubBrandID:   model.SubBrandID,
		ModelYearID:  model.ModelYearID,
		OwnerName:    model.OwnerName,
		PhoneNumber:  model.PhoneNumber,
		IDType:       model.IDType,
		IDNo:         model.IDNo,
		Comment:      model.Comment,
		Ts:           model.Ts,
		Uts:          model.Uts,
	}
}
