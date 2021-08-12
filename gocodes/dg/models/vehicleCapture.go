package models

import (
	"encoding/json"
	"strings"

	"codes/gocodes/dg/dg.model"
)

const (
	PlateTypeNoPlate = 2000 //无牌车
)

type VehicleCapture struct {
	Timestamp         int64
	SensorID          string
	SensorIntID       int32
	OrgCode           int64
	SensorName        string
	Longitude         float64
	Latitude          float64
	VehicleID         string
	VehicleReID       string
	VehicleVID        string
	ImageResult       *CutboardImage
	PlateText         string
	PlateColorID      int32
	PlateTypeID       int32
	ColorID           int32
	TypeID            int32
	BrandID           int32
	SubBrandID        int32
	ModelYearID       int32
	Symbols           []int
	Specials          []int
	Favorite          *Favorite
	RelationTypeSlice []int
}

type VehicleCaptureDetail struct {
	VehicleCapture
	SensorType    int
	SensorURL     string
	ImageResult   *ImageResult
	Speed         int32
	Direction     int32
	FaceID        string
	BrandID       int32
	SubBrandID    int32
	ModelYearID   int32
	TypeID        int32
	Side          int32
	ColorID       int32
	Symbols       SymbolSlice
	SymbolsDesc   string
	Specials      []int64
	Lane          string
	RelationTypes []int
	RelationIDs   []string
	RelationURLs  []string
}

func (vehicleCapture *VehicleCapture) AddFavorite(favorite *Favorite) *VehicleCapture {
	if favorite != nil {
		vehicleCapture.Favorite = favorite
	} else {
		vehicleCapture.Favorite = &Favorite{FavoriteTargetID: vehicleCapture.VehicleID}
	}
	return vehicleCapture
}

func (vehicleCapture *VehicleCaptureDetail) AddFavorite(favorite *Favorite) *VehicleCaptureDetail {
	if favorite != nil {
		vehicleCapture.Favorite = favorite
	} else {
		vehicleCapture.Favorite = &Favorite{FavoriteTargetID: vehicleCapture.VehicleID}
	}
	return vehicleCapture
}

type VehicleCaptureConfidence struct {
	VehicleCaptureDetail
	Confidence       float32
	OriginImageIndex int
}

func (vehicleCapture *VehicleCaptureDetail) AddConfidence(featureResult *FeatureResult) *VehicleCaptureConfidence {
	return &VehicleCaptureConfidence{
		VehicleCaptureDetail: *vehicleCapture,
		Confidence:           featureResult.Confidence,
		OriginImageIndex:     featureResult.OriginImageIndex,
	}
}

type VehicleConditionRequest struct {
	LocatedSensor bool
	Vid           string
	Total         int
	Plate         string
	TypeID        []int
	ColorID       []int
	PlateTypeID   []int
	PlateColorID  []int
	Direction     []int
	Speed         []int
	Side          []int
	Specials      []int
	SpecialInt    int `json:"-"`
	Symbols       []int
	SymbolInt     int `json:"-"`
	HasFace       []int
	SensorID      []string
	OrgCodes      []int64
	Brands        []struct {
		Brand    string
		SubBrand string
		Year     string
	}
	Time         StartAndEndTimestamp
	OrderAsc     bool   // default false
	OrderBy      string // default "ts"
	FromTailPage bool
	Pagination
	TryLimit        int             // loki inner: limit + 1, for use of nextPage for fe
	AscWithFromPage bool            // loki inner: OrderAsc + FromTailPage
	SensorIntIDs    []int           `json:"-"`
	OrgCodeRanges   []*OrgCodeRange `json:"-"` // inner
	SensorArgs      string
}

func (query *VehicleConditionRequest) SensorIDsVisibleOptimizerV1(userVisibleSensorIDs, allSensorIDsInPg []string) {
	query.SensorID = SensorIDsVisibleOptimizerV1(query.SensorID, userVisibleSensorIDs, allSensorIDsInPg)
}

func (query *VehicleConditionRequest) SensorIDsVisibleOptimizerV2(userVisibleSensorIDs []string) {
	query.SensorID = SensorIDsVisibleOptimizerV2(query.SensorID, userVisibleSensorIDs)
}

func (query *VehicleConditionRequest) GetLimit() int {
	return query.Limit
}

func (query *VehicleConditionRequest) GetOffset() int {
	return query.Offset
}

func (query *VehicleConditionRequest) GetStartTimestamp() int64 {
	return query.Time.StartTimestamp
}

func (query *VehicleConditionRequest) GetEndTimestamp() int64 {
	return query.Time.EndTimestamp
}

func (query *VehicleConditionRequest) GetSortBy() string {
	return "ts"
}

func (query *VehicleConditionRequest) GetSortAsc() bool {
	return false
}

//把symbol转成symbol_int
func (query *VehicleConditionRequest) MixConditions() {
	query.SymbolInt = SymbolSliceToInt(query.Symbols)
}

func SymbolIntToSlice(symbolInt int64) []int64 {
	rs := make([]int64, 0, 64)
	for i := uint64(0); i < 64; i++ {
		k := int64(1 << i)
		if symbolInt&k == k {
			rs = append(rs, int64(i))
		}
	}
	return rs
}

func SymbolSliceToInt(symbols []int) int {
	rs := int(0)
	for _, k := range symbols {
		rs = rs | int(1<<uint(k))
	}
	return rs
}

type SymbolSlice []*dg_model.VehicleSymbol

func (symbolSlice *SymbolSlice) ToDB() ([]byte, error) {
	return json.Marshal(symbolSlice)
}

func (symbolSlice *SymbolSlice) FromDB(b []byte) error {
	if string(b) == "{}" {
		return nil
	}
	return json.Unmarshal(b, symbolSlice)
}

func (query *VehicleConditionRequest) PreProcess() {
	query.fix()
	query.TryLimit = query.Limit + 1
	query.AscWithFromPage = !query.OrderAsc && query.FromTailPage
}

func (query *VehicleConditionRequest) fix() {
	if strings.EqualFold(query.OrderBy, "") {
		query.OrderBy = "ts"
		query.OrderAsc = false
	}
}
