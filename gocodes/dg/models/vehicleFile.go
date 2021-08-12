package models

import "time"

type VehicleFile struct {
	Vid                  string
	Ts                   int64
	Uts                  time.Time
	Comment              string
	ImageURL             string `json:"ImageUrl"`
	PlateText            string
	PlateColorID         int32
	ColorID              int32
	BrandID              int32
	SubBrandID           int32
	ModelYearID          int32
	TypeID               int32
	Tag                  int
	Status               int
	SystemTagNames       []string
	SystemTagValues      []int
	CustomTags           []string
	CapturedVehicleCount int
}

type VehicleFileExtend struct {
	VehicleFile
	CapturedVehicle *VehicleCapture
}

type VehicleFileQuery struct {
	Limit     int
	Offset    int
	Total     int64
	PlateText string
	Vids      []interface{} `json:"-"`
	OrderAsc  bool
	OrderBy   string
}

type VehicleFileTag struct {
	SystemName  []string
	SystemValue []int
	Custom      []string
}

func (v *VehicleFileQuery) ShallowClone() *VehicleFileQuery {
	return &VehicleFileQuery{
		Limit:     v.Limit,
		Offset:    v.Offset,
		Total:     v.Total,
		PlateText: v.PlateText,
		Vids:      v.Vids,
		OrderAsc:  v.OrderAsc,
		OrderBy:   v.OrderBy,
	}
}

func (v VehicleFileQuery) GetSortBy() string {
	return v.OrderBy
}

func (v VehicleFileQuery) GetSortAsc() bool {
	return v.OrderAsc
}
