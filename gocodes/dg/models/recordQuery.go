package models

type RecordQuery struct {
	Limit      int
	Offset     int
	StartTime  int64
	EndTime    int64
	UserIds    []string
	OperFields []string
	OperTypes  []string
	Results    []string
	KeyContent string
}

func (query *RecordQuery) GetStartTimestamp() int64 {
	return query.StartTime
}

func (query *RecordQuery) GetEndTimestamp() int64 {
	return query.EndTime
}

func (query *RecordQuery) GetLimit() int {
	return query.Limit
}

func (query *RecordQuery) GetOffset() int {
	return query.Offset
}

func (query *RecordQuery) GetSortBy() string {
	return "ts"
}

func (query *RecordQuery) GetSortAsc() bool {
	return false
}
