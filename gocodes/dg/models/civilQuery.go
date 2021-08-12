package models

const (
	DefaultAlgorithmDeepglint = iota
	AlgorithmMegvii
)

type CivilQueryAndCountResult struct {
	Rets       interface{}
	NextOffset int
	Total      int
}

type CivilQuery struct {
	UTCOffset        int
	Limit            int
	Offset           int
	Total            int
	Addr             string
	GenderID         int `json:"GenderId"`
	OrderAsc         bool
	OrderBy          string
	IDNo             string `json:"IdNo"`   // todo rename to IdNo
	IDType           int    `json:"IdType"` // todo rename to IdType
	Name             string
	RepoIds          []string
	Images           []*ImageQuery
	Confidence       float32
	TopX             int           `json:"Topx"` // todo rename to TopX
	CivilImageIds    []interface{} `json:"-"`    // json ignore
	CivilAttrIdNotIn []interface{} `json:"-"`
	SearchID         string

	AlgorithmType int

	FromTailPage bool
	RepoImageAll bool
}

func (query *CivilQuery) ShallowClone() *CivilQuery {
	return &CivilQuery{
		UTCOffset:        query.UTCOffset,
		Limit:            query.Limit,
		Offset:           query.Offset,
		Total:            query.Total,
		Addr:             query.Addr,
		GenderID:         query.GenderID,
		OrderAsc:         query.OrderAsc,
		OrderBy:          query.OrderBy,
		IDNo:             query.IDNo,
		IDType:           query.IDType,
		Name:             query.Name,
		RepoIds:          query.RepoIds,
		Images:           query.Images,
		Confidence:       query.Confidence,
		TopX:             query.TopX,
		CivilImageIds:    query.CivilImageIds,
		CivilAttrIdNotIn: query.CivilAttrIdNotIn,
		SearchID:         query.SearchID,
		FromTailPage:     query.FromTailPage,
	}
}

func (query *CivilQuery) GetRepoIds() []string {
	return query.RepoIds
}

func (query *CivilQuery) GetName() string {
	return query.Name
}

func (query *CivilQuery) GetAddr() string {
	return query.Addr
}

func (query *CivilQuery) GetIdNo() string {
	return query.IDNo
}

func (query *CivilQuery) GetIdType() int {
	return query.IDType
}

func (query *CivilQuery) GetSortBy() string {
	return query.OrderBy
}

func (query *CivilQuery) GetSortAsc() bool {
	return query.OrderAsc
}

func (query *CivilQuery) GetLimit() int {
	return query.Limit
}

func (query *CivilQuery) GetOffset() int {
	return query.Offset
}

func (query *CivilQuery) GetUTCOffset() int {
	return query.UTCOffset
}

func (query *CivilQuery) GetGenderID() int {
	return query.GenderID
}

func (query *CivilQuery) GetSearchID() string {
	return query.SearchID
}

func (query *CivilQuery) GetTailPage() bool {
	return query.FromTailPage
}

type FailedCivilQuery struct {
	UTCOffset    int
	Limit        int
	Offset       int
	OrderAsc     bool
	OrderBy      string
	RepoIds      []string
	FromTailPage bool
}

func (query *FailedCivilQuery) GetRepoIds() []string {
	return query.RepoIds
}

func (query *FailedCivilQuery) GetSortBy() string {
	return query.OrderBy
}

func (query *FailedCivilQuery) GetSortAsc() bool {
	return query.OrderAsc
}

func (query *FailedCivilQuery) GetLimit() int {
	return query.Limit
}

func (query *FailedCivilQuery) GetOffset() int {
	return query.Offset
}

func (query *FailedCivilQuery) GetUTCOffset() int {
	return query.UTCOffset
}
