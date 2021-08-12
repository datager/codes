package models

import (
	"fmt"
)

type PersonFileQuery struct {
	UTCOffset      int
	Limit          int
	Offset         int
	Total          int
	StartTimestamp int64
	EndTimestamp   int64
	Addr           string
	OrderAsc       bool
	OrderBy        string
	IDNo           string // todo rename to IdNo
	IDType         int    // todo rename to IdType
	Name           string
	Autonymous     int
	MarkStates     []int
	Images         []*ImageQuery
	Confidence     float32
	TopX           int           `json:"Topx"` // todo rename to TopX
	Vids           []interface{} `json:"-"`    // json ignore
	Status         int
}

func (query *PersonFileQuery) Validate() error {
	if query.Name != "" {
		autonymous := NewNullableBoolByInt(query.Autonymous)
		if !autonymous.Null() && !autonymous.True() {
			return fmt.Errorf("Invalid model")
		}
	}
	return nil
}

type PersonFileCapturedCount struct {
	StartTimestamp int64 `form:"starttime"`
	EndTimestamp   int64 `form:"endtime"`
}

func (query *PersonFileQuery) ShallowClone() *PersonFileQuery {
	return &PersonFileQuery{
		UTCOffset:      query.UTCOffset,
		Limit:          query.Limit,
		Offset:         query.Offset,
		Total:          query.Total,
		StartTimestamp: query.StartTimestamp,
		EndTimestamp:   query.EndTimestamp,
		Addr:           query.Addr,
		OrderAsc:       query.OrderAsc,
		OrderBy:        query.OrderBy,
		IDNo:           query.IDNo,
		IDType:         query.IDType,
		Name:           query.Name,
		Autonymous:     query.Autonymous,
		MarkStates:     query.MarkStates,
		Images:         query.Images,
		Confidence:     query.Confidence,
		TopX:           query.TopX,
		Vids:           query.Vids,
		Status:         query.Status,
	}
}

func (query *PersonFileQuery) GetStartTimestamp() int64 {
	return query.StartTimestamp
}

func (query *PersonFileQuery) GetEndTimestamp() int64 {
	return query.EndTimestamp
}

func (query *PersonFileQuery) GetName() string {
	return query.Name
}

func (query *PersonFileQuery) GetAddr() string {
	return query.Addr
}

func (query *PersonFileQuery) GetIdNo() string {
	return query.IDNo
}

func (query *PersonFileQuery) GetIdType() int {
	return query.IDType
}

func (query *PersonFileQuery) GetSortBy() string {
	return query.OrderBy
}

func (query *PersonFileQuery) GetSortAsc() bool {
	return query.OrderAsc
}

func (query *PersonFileQuery) GetLimit() int {
	return query.Limit
}

func (query *PersonFileQuery) GetOffset() int {
	return query.Offset
}

func (query *PersonFileQuery) GetImages() []*ImageQuery {
	return query.Images
}

func (query *PersonFileQuery) GetUTCOffset() int {
	return query.UTCOffset
}

func (query *PersonFileQuery) GetGenderID() int {
	return 0
}

func (query *PersonFileQuery) GetSearchID() string {
	return ""
}
