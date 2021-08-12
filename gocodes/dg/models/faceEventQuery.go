package models

type FaceEventQuery struct {
	UTCOffset      int
	Total          int
	Addr           string
	CheckStatus    []int
	StartTimestamp int64
	EndTimestamp   int64
	OrderAsc       bool
	OrderBy        string
	IDNo           string `json:"IdNo"`   // todo rename to IdNo
	IDType         int    `json:"IdType"` // todo rename to IdType
	Name           string
	RepoIds        []string
	SensorIds      []string
	Confidence     float32
	FaceIds        []string
	RuleName       string
	CivilAttrID    string `json:"AttrId"`
	//Limit          int
	//Offset         int
	Pagination
}

func (query *FaceEventQuery) ShallowClone() *FaceEventQuery {
	return &FaceEventQuery{
		UTCOffset:      query.UTCOffset,
		Total:          query.Total,
		Addr:           query.Addr,
		CheckStatus:    query.CheckStatus,
		StartTimestamp: query.StartTimestamp,
		EndTimestamp:   query.EndTimestamp,
		OrderAsc:       query.OrderAsc,
		OrderBy:        query.OrderBy,
		IDNo:           query.IDNo,
		IDType:         query.IDType,
		Name:           query.Name,
		RepoIds:        query.RepoIds,
		SensorIds:      query.SensorIds,
		Confidence:     query.Confidence,
		RuleName:       query.RuleName,
		CivilAttrID:    query.CivilAttrID,
		Pagination: Pagination{
			Limit:  query.Limit,
			Offset: query.Offset,
		},
	}
}

func (query *FaceEventQuery) GetStartTimestamp() int64 {
	return query.StartTimestamp
}

func (query *FaceEventQuery) GetEndTimestamp() int64 {
	return query.EndTimestamp
}

func (query *FaceEventQuery) GetSensorIds() []string {
	return query.SensorIds
}

func (query *FaceEventQuery) GetRepoIds() []string {
	return query.RepoIds
}

func (query *FaceEventQuery) GetName() string {
	return query.Name
}

func (query *FaceEventQuery) GetAddr() string {
	return query.Addr
}

func (query *FaceEventQuery) GetIdNo() string {
	return query.IDNo
}

func (query *FaceEventQuery) GetIdType() int {
	return query.IDType
}

func (query *FaceEventQuery) GetSortBy() string {
	return query.OrderBy
}

func (query *FaceEventQuery) GetSortAsc() bool {
	return query.OrderAsc
}

func (query *FaceEventQuery) GetLimit() int {
	return query.Limit
}

func (query *FaceEventQuery) GetOffset() int {
	return query.Offset
}

func (query *FaceEventQuery) GetUTCOffset() int {
	return query.UTCOffset
}

func (query *FaceEventQuery) GetGenderID() int {
	return 0
}

func (query *FaceEventQuery) GetSearchID() string {
	return ""
}
