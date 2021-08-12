package models

type ImageQuery struct {
	Url     string
	BinData string
	Feature string
}

type FeatureCompare struct {
	FeatureA string
	FeatureB string
}

type ImagesQuery interface {
	GetImages() []*ImageQuery
}

type DateTimeRangeQuery interface {
	GetStartTimestamp() int64
	GetEndTimestamp() int64
}

type SortableQuery interface {
	GetSortBy() string
	GetSortAsc() bool
}

type SortableQueryWithFromPage interface {
	GetSortBy() string
	GetSortBy2nd() string
	GetSortAscWithFromPage() bool // asc = orderAsc + withHeadPage
}

type PagingQuery interface {
	GetLimit() int
	GetOffset() int
}

type PagingQueryWithTryLimit interface {
	GetTryLimit() int
	GetOffset() int
}

type SensorIdRangeQuery interface {
	GetSensorIds() []string
}

type SensorTypeRangeQuery interface {
	GetSensorTypes() []int
}

type RepoIdRangeQuery interface {
	GetRepoIds() []string
}

type CivilAttrQuery interface {
	GetName() string
	GetAddr() string
	GetIdNo() string
	GetIdType() int
	GetGenderID() int
	GetSearchID() string
}

type FaceAttrQuery interface {
	GetAgeIdStart() int
	GetAgeIdEnd() int
	GetGlassIds() []int
	GetMaskIds() []int
	GetHatIds() []int
	GetGenderIds() []int
	GetNationIds() []int
	DateTimeRangeQuery
}

type UTCOffsetQuery interface {
	GetUTCOffset() int
}

type pagingQueryImpl struct {
	Limit  int
	Offset int
}

func NewPagingQuery(limit, offset int) PagingQuery {
	return &pagingQueryImpl{
		Limit:  limit,
		Offset: offset,
	}
}

func (query *pagingQueryImpl) GetLimit() int {
	return query.Limit
}

func (query *pagingQueryImpl) GetOffset() int {
	return query.Offset
}

type pagingQueryImplWithTryLimit struct {
	TryLimit int
	Offset   int
}

func (query *pagingQueryImplWithTryLimit) GetTryLimit() int {
	return query.TryLimit
}

func (query *pagingQueryImplWithTryLimit) GetOffset() int {
	return query.Offset
}

type ObjectQuery struct {
	ImageQuery
	// description:"[1]行人 [2]机动车 [4]非机动车 [8]人脸"
	InterestedType []int
}
