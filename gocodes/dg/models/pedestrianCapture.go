package models

import "strings"

type Pedestrian struct {
	ImageResult    ImageResult
	PedestrianID   string
	PedestrianReID string
	Timestamp      int64
	SensorID       string
	SensorName     string
	SensorType     int
	SensorURL      string
	Longitude      float64
	Latitude       float64
	AgeID          int32
	GenderID       int32
	NationID       int32
	UpperColor     int32
	LowerColor     int32
	HairStyle      int32
	Favorite       *Favorite
}

type PedestrianIndex struct {
	ImageResult       *CutboardImage
	PedestrianID      string
	PedestrianReID    string
	Timestamp         int64
	SensorID          string
	SensorIntID       int32
	OrgCode           int64
	SensorName        string
	Longitude         float64
	Latitude          float64
	AgeID             int
	GenderID          int
	NationID          int
	UpperColor        int
	LowerColor        int
	HairStyle         int
	Favorite          *Favorite
	RelationTypeSlice []int
}

func (pedestrianCapture *Pedestrian) AddFavorite(favorite *Favorite) *Pedestrian {
	if favorite != nil {
		pedestrianCapture.Favorite = favorite
	} else {
		pedestrianCapture.Favorite = &Favorite{FavoriteTargetID: pedestrianCapture.PedestrianID}
	}
	return pedestrianCapture
}

func (p *PedestrianIndex) AddFavorite(favorite *Favorite) *PedestrianIndex {
	if favorite != nil {
		p.Favorite = favorite
	} else {
		p.Favorite = &Favorite{FavoriteTargetID: p.PedestrianID}
	}
	return p
}

type PedestrianCaptureDetail struct {
	Pedestrian
	FaceID          string
	Speed           int32
	Direction       int32
	DescEye         int32
	DescHead        int32
	DescMouth       int32
	WithBackpack    int32
	WithShoulderBag int32
	WithHandbag     int32
	WithHandCarry   int32
	WithPram        int32
	WithLuggage     int32
	WithTrolley     int32
	WithUmbrella    int32
	WithHoldBaby    int32
	WithScarf       int32
	UpperColor      int32
	UpperTexture    int32
	UpperStyle      int32
	LowerColor      int32
	LowerStyle      int32
	ShoesColor      int32
	ShoesStyle      int32
	HairStyle       int32
	RelationTypes   []int
	RelationIDs     []string
	RelationURLs    []string
}

type PedestrianCaptureConfidence struct {
	Pedestrian
	Confidence       float32
	OriginImageIndex int
}

func (pedestrianCapture *Pedestrian) AddConfidence(featureResult *FeatureResult) *PedestrianCaptureConfidence {
	return &PedestrianCaptureConfidence{
		Pedestrian:       *pedestrianCapture,
		Confidence:       featureResult.Confidence,
		OriginImageIndex: featureResult.OriginImageIndex,
	}
}

type PedestrianConditionRequest struct {
	HairStyle       []int32
	UpColor         []int32
	UpperTexture    []int32
	UpperStyle      []int32
	LowerColor      []int32
	LowerStyle      []int32
	ShoesColor      []int32
	ShoesStyle      []int32
	AgeID           []int32
	GenderID        []int32
	NationID        []int32
	Direction       []int32
	Speed           []int32
	HasFace         []int
	SensorID        []string
	OrgCodes        []int64
	WithShoulderBag int32
	WithBackpack    int32
	WithHandbag     int32
	WithHandCarry   int32
	WithPram        int32
	WithLuggage     int32
	WithTrolley     int32
	WithUmbrella    int32
	WithHoldBaby    int32
	WithScarf       int32
	DescEye         int32
	DescHead        []int32
	DescMouth       int32
	Time            StartAndEndTimestamp
	Total           int
	OrderAsc        bool   // default false
	OrderBy         string // default "ts"
	FromTailPage    bool
	Pagination
	TryLimit        int             // loki inner: limit + 1, for use of nextPage for fe
	AscWithFromPage bool            // loki inner: OrderAsc + FromTailPage
	SensorIntIDs    []int32         `json:"-"`
	OrgCodeRanges   []*OrgCodeRange `json:"-"` // inner
}

func (req *PedestrianConditionRequest) GetLimit() int {
	return req.Limit
}

func (req *PedestrianConditionRequest) GetOffset() int {
	return req.Offset
}

func (req *PedestrianConditionRequest) SensorIDsVisibleOptimizerV1(userVisibleSensorIDs, allSensorIDsInPg []string) {
	req.SensorID = SensorIDsVisibleOptimizerV1(req.SensorID, userVisibleSensorIDs, allSensorIDsInPg)
}

func (req *PedestrianConditionRequest) SensorIDsVisibleOptimizerV2(userVisibleSensorIDs []string) {
	req.SensorID = SensorIDsVisibleOptimizerV2(req.SensorID, userVisibleSensorIDs)
}

func (req *PedestrianConditionRequest) PreProcess() {
	req.fix()
	req.TryLimit = req.Limit + 1
	req.AscWithFromPage = !req.OrderAsc && req.FromTailPage
}

func (req *PedestrianConditionRequest) fix() {
	if strings.EqualFold(req.OrderBy, "") {
		req.OrderBy = "ts"
		req.OrderAsc = false
	}
}
