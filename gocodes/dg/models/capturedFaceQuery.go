package models

type CapturedFaceQuery struct {
	UTCOffset       int
	Limit           int
	Offset          int
	Total           int
	SensorIds       []string
	StartTimestamp  int64
	EndTimestamp    int64
	Images          []*ImageQuery
	Confidence      float32
	TopX            int `json:"Topx"` // todo rename to TopX
	OrderAsc        bool
	OrderBy         string
	AgeIDStart      int `json:"AgeIdStart"`
	AgeIDEnd        int `json:"AgeIdEnd"`
	GlassIds        []int
	MaskIds         []int
	HatIds          []int
	GenderIds       []int
	NationIds       []int
	Vid             string
	LocatedSensor   bool
	TryLimit        int             // with from page
	AscWithFromPage bool            // with from page
	SensorIntIDs    []int32         `json:"-"`
	OrgCodeRanges   []*OrgCodeRange `json:"-"` // inner
}

func (query *CapturedFaceQuery) ShallowClone() *CapturedFaceQuery {
	return &CapturedFaceQuery{
		UTCOffset:       query.UTCOffset,
		Limit:           query.Limit,
		Offset:          query.Offset,
		Total:           query.Total,
		SensorIds:       query.SensorIds,
		StartTimestamp:  query.StartTimestamp,
		EndTimestamp:    query.EndTimestamp,
		Images:          query.Images,
		Confidence:      query.Confidence,
		TopX:            query.TopX,
		OrderAsc:        query.OrderAsc,
		OrderBy:         query.OrderBy,
		AgeIDStart:      query.AgeIDStart,
		AgeIDEnd:        query.AgeIDEnd,
		GlassIds:        query.GlassIds,
		MaskIds:         query.MaskIds,
		HatIds:          query.HatIds,
		GenderIds:       query.GenderIds,
		NationIds:       query.NationIds,
		Vid:             query.Vid,
		TryLimit:        query.TryLimit,
		AscWithFromPage: query.AscWithFromPage,
	}
}

func (query *CapturedFaceQuery) GetStartTimestamp() int64 {
	return query.StartTimestamp
}

func (query *CapturedFaceQuery) GetEndTimestamp() int64 {
	return query.EndTimestamp
}

func (query *CapturedFaceQuery) GetSensorIds() []string {
	return query.SensorIds
}

func (query *CapturedFaceQuery) GetSortBy() string {
	return query.OrderBy
}

func (query *CapturedFaceQuery) GetSortBy2nd() string {
	return "face_reid"
}

func (query *CapturedFaceQuery) GetSortAsc() bool {
	return query.OrderAsc
}

func (query *CapturedFaceQuery) GetSortAscWithFromPage() bool {
	return query.AscWithFromPage
}

func (query *CapturedFaceQuery) GetLimit() int {
	return query.Limit
}

func (query *CapturedFaceQuery) GetTryLimit() int {
	return query.TryLimit
}

func (query *CapturedFaceQuery) GetOffset() int {
	return query.Offset
}

func (query *CapturedFaceQuery) GetAgeIdStart() int {
	return query.AgeIDStart
}

func (query *CapturedFaceQuery) GetAgeIdEnd() int {
	return query.AgeIDEnd
}

func (query *CapturedFaceQuery) GetGlassIds() []int {
	return query.GlassIds
}

func (query *CapturedFaceQuery) GetHatIds() []int {
	return query.HatIds
}

func (query *CapturedFaceQuery) GetMaskIds() []int {
	return query.MaskIds
}

func (query *CapturedFaceQuery) GetGenderIds() []int {
	return query.GenderIds
}

func (query *CapturedFaceQuery) GetNationIds() []int {
	return query.NationIds
}

func (query *CapturedFaceQuery) GetImages() []*ImageQuery {
	return query.Images
}

func (query *CapturedFaceQuery) GetUTCOffset() int {
	return query.UTCOffset
}
