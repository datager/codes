package models

import (
	"strings"
)

type FaceCaptureConfidence struct {
	CapturedFaceDetial
	Confidence       float32
	OriginImageIndex int
}

func (capturedFace *CapturedFaceDetial) AddConfidence(featureResult *FeatureResult) *FaceCaptureConfidence {
	return &FaceCaptureConfidence{
		CapturedFaceDetial: *capturedFace,
		Confidence:         featureResult.Confidence,
		OriginImageIndex:   featureResult.OriginImageIndex,
	}
}

type FaceConditionRequest struct {
	GenderID      []int // 男/女/其他
	NationID      []int // 全部/其他/维吾尔族
	HatID         []int // 全部/无帽子/有帽子/有头巾
	GlassID       []int // 全部/无眼镜/有眼睛/有太阳镜
	MaskID        []int // 全部/无口罩/有口罩
	SensorID      []string
	OrgCodes      []int64
	Time          StartAndEndTimestamp
	Total         int    // use total + offset + limit: to deside whether to query by ts (asc or desc)
	OrderAsc      bool   // eg: false
	OrderBy       string // eg: "ts"
	FromTailPage  bool   // true: fromTali; false: from Head
	Vid           string
	LocatedSensor bool
	Pagination
	TryLimit        int             // loki inner: limit + 1, for use of nextPage for fe
	AscWithFromPage bool            // loki inner: OrderAsc + FromTailPage
	OrgCodeRanges   []*OrgCodeRange `json:"-"` // inner
}

func (req *FaceConditionRequest) GetLimit() int {
	return req.Limit
}

func (req *FaceConditionRequest) GetOffset() int {
	return req.Offset
}

func (req *FaceConditionRequest) Fix() {
	if strings.EqualFold(req.OrderBy, "") {
		req.OrderBy = "ts"
		req.OrderAsc = false
	}
}

func (req *FaceConditionRequest) SensorIDsVisibleOptimizerV1(userVisibleSensorIDs, allSensorIDsInPg []string) {
	req.SensorID = SensorIDsVisibleOptimizerV1(req.SensorID, userVisibleSensorIDs, allSensorIDsInPg)
}

func (req *FaceConditionRequest) SensorIDsVisibleOptimizerV2(userVisibleSensorIDs []string) {
	req.SensorID = SensorIDsVisibleOptimizerV2(req.SensorID, userVisibleSensorIDs)
}

func (req *FaceConditionRequest) PreProcess() {
	req.Fix()
	req.TryLimit = req.Limit + 1
	req.AscWithFromPage = !req.OrderAsc && req.FromTailPage
}

func (req *FaceConditionRequest) ToCapturedFaceQuery() *CapturedFaceQuery {
	return &CapturedFaceQuery{
		UTCOffset:      0,
		Limit:          req.Limit,
		Offset:         req.Offset,
		Total:          req.Total,
		SensorIds:      req.SensorID,
		StartTimestamp: req.Time.StartTimestamp,
		EndTimestamp:   req.Time.EndTimestamp,
		//Images        []*ImageQuery not in use
		//Confidence    float32
		//TopX          int `json:"Topx"`
		OrderAsc: req.OrderAsc,
		OrderBy:  req.OrderBy,
		//AgeIdStart    int
		//AgeIdEnd      int
		GlassIds:        req.GlassID,
		MaskIds:         req.MaskID,
		HatIds:          req.HatID,
		GenderIds:       req.GenderID,
		NationIds:       req.NationID,
		Vid:             req.Vid,
		LocatedSensor:   req.LocatedSensor,
		TryLimit:        req.TryLimit,
		AscWithFromPage: req.AscWithFromPage,
		OrgCodeRanges:   req.OrgCodeRanges,
	}
}
