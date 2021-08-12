package models

import "strings"

type Nonmotor struct {
	Timestamp       int64
	Speed           int32
	Direction       int32
	NonmotorColorID int32
	UpperColor      int32
	UpperStyle      int32
	ImageResult     ImageResult
	RepoID          string
	SensorID        string
	SensorName      string
	SensorType      int
	SensorURL       string
	Longitude       float64
	Latitude        float64
	NonmotorID      string
	NonmotorReID    string
	Favorite        *Favorite
}

type NonmotorIndex struct {
	Timestamp         int64
	Speed             int32
	Direction         int32
	NonmotorColorID   int32
	UpperColor        int32
	UpperStyle        int32
	ImageResult       *CutboardImage
	SensorID          string
	SensorIntID       int32
	OrgCode           int64
	SensorName        string
	Longitude         float64
	Latitude          float64
	NonmotorID        string
	NonmotorReID      string
	Favorite          *Favorite
	RelationTypeSlice []int
}

func (nonmotorCapture *Nonmotor) AddFavorite(favorite *Favorite) *Nonmotor {
	if favorite != nil {
		nonmotorCapture.Favorite = favorite
	} else {
		nonmotorCapture.Favorite = &Favorite{FavoriteTargetID: nonmotorCapture.NonmotorID}
	}
	return nonmotorCapture
}

func (nc *NonmotorIndex) AddFavorite(favorite *Favorite) *NonmotorIndex {
	if favorite != nil {
		nc.Favorite = favorite
	} else {
		nc.Favorite = &Favorite{FavoriteTargetID: nc.NonmotorID}
	}
	return nc
}

type NonmotorCaptureDetail struct {
	Nonmotor
	PlateColorID    int32
	PlateTypeID     int32
	PlateText       string
	FaceID          string
	NonmotorType    int32
	NonmotorGesture int32
	NonmotorColorID int32
	GenderID        int32
	AgeID           int32
	NationID        int32
	DescEye         int32
	DescHead        int32
	DescMouth       int32
	WithBackpack    int32
	WithShoulderBag int32
	UpperColor      int32
	UpperStyle      int32
	RelationTypes   []int
	RelationIDs     []string
	RelationURLs    []string
}

type NonmotorCaptureConfidence struct {
	Nonmotor
	Confidence       float32
	OriginImageIndex int
}

func (nonmotorCapture *Nonmotor) AddConfidence(featureResult *FeatureResult) *NonmotorCaptureConfidence {
	return &NonmotorCaptureConfidence{
		Nonmotor:         *nonmotorCapture,
		Confidence:       featureResult.Confidence,
		OriginImageIndex: featureResult.OriginImageIndex,
	}
}

type NonmotorConditionRequest struct {
	PlateColorID    []int32
	NonmotorColorID []int32
	NonmotorType    []int32
	NonmotorGesture []int32
	DescEye         int32
	DescHead        []int32
	DescMouth       int32
	WithShoulderBag int32
	WithBackpack    int32
	HasFace         []int
	Total           int
	Time            StartAndEndTimestamp
	Plate           string
	SensorID        []string
	OrgCodes        []int64
	GenderID        []int32
	NationID        []int32
	UpColor         []int32
	UpperStyle      []int32
	Speed           []int32
	Direction       []int32
	OrderAsc        bool   // default false
	OrderBy         string // default "ts"
	FromTailPage    bool
	Pagination
	TryLimit        int  // loki inner: limit + 1, for use of nextPage for fe
	AscWithFromPage bool // loki inner: OrderAsc + FromTailPage
	SensorIntIDs    []int32
	OrgCodeRanges   []*OrgCodeRange `json:"-"` // inner
}

func (req *NonmotorConditionRequest) GetLimit() int {
	return req.Limit
}

func (req *NonmotorConditionRequest) GetOffset() int {
	return req.Offset
}

func (req *NonmotorConditionRequest) SensorIDsVisibleOptimizerV1(userVisibleSensorIDs, allSensorIDsInPg []string) {
	req.SensorID = SensorIDsVisibleOptimizerV1(req.SensorID, userVisibleSensorIDs, allSensorIDsInPg)
}

func (req *NonmotorConditionRequest) SensorIDsVisibleOptimizerV2(userVisibleSensorIDs []string) {
	req.SensorID = SensorIDsVisibleOptimizerV2(req.SensorID, userVisibleSensorIDs)
}

func (req *NonmotorConditionRequest) PreProcess() {
	req.fix()
	req.TryLimit = req.Limit + 1
	req.AscWithFromPage = !req.OrderAsc && req.FromTailPage
}

func (req *NonmotorConditionRequest) fix() {
	if strings.EqualFold(req.OrderBy, "") {
		req.OrderBy = "ts"
		req.OrderAsc = false
	}
}
