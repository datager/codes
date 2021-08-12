package models

type PersonBangQuery struct {
	UTCOffset      int
	Limit          int
	Offset         int
	OrderAsc       bool
	OrderBy        string
	StartTimestamp int64
	EndTimestamp   int64
	RepoIds        []string
	Confidence     float32
	Vids           []string
}

func (query *PersonBangQuery) GetStartTimestamp() int64 {
	return query.StartTimestamp
}

func (query *PersonBangQuery) GetEndTimestamp() int64 {
	return query.EndTimestamp
}

func (query *PersonBangQuery) GetSortBy() string {
	return query.OrderBy
}

func (query *PersonBangQuery) GetSortAsc() bool {
	return query.OrderAsc
}

func (query *PersonBangQuery) GetLimit() int {
	return query.Limit
}

func (query *PersonBangQuery) GetOffset() int {
	return query.Offset
}

func (query *PersonBangQuery) GetRepoIds() []string {
	return query.RepoIds
}

func (query *PersonBangQuery) GetUTCOffset() int {
	return query.UTCOffset
}
