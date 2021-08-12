package models

import "strings"

type VehicleBaseConditionRequest struct {
	Plate        string
	SensorID     []string
	Time         StartAndEndTimestamp
	OrderAsc     bool   // default false
	OrderBy      string // default "ts"
	FromTailPage bool
	Pagination
	TryLimit        int  // loki inner: limit + 1, for use of nextPage for fe
	AscWithFromPage bool // loki inner: OrderAsc + FromTailPage
}

func (query *VehicleBaseConditionRequest) PreProcess() {
	query.fix()
	query.TryLimit = query.Limit + 1
	query.AscWithFromPage = !query.OrderAsc && query.FromTailPage
}

func (query *VehicleBaseConditionRequest) fix() {
	if strings.EqualFold(query.OrderBy, "") {
		query.OrderBy = "ts"
		query.OrderAsc = false
	}
}

type VehicleCommonData struct {
	Timestamp int64
	CaptureId string
	SensorId  string
	PlateText string
}

type CoverFaceRequest struct {
	*VehicleBaseConditionRequest
	DriverType    []int
	PassengerType []int
}

type UnbeltedRequest struct {
	*VehicleBaseConditionRequest
	DriverType    []int
	PassengerType []int
}

type NonmotorBaseConditionRequest struct {
	SensorID     []string
	Time         StartAndEndTimestamp
	OrderAsc     bool   // default false
	OrderBy      string // default "ts"
	FromTailPage bool
	Pagination
	TryLimit        int  // loki inner: limit + 1, for use of nextPage for fe
	AscWithFromPage bool // loki inner: OrderAsc + FromTailPage
}

func (query *NonmotorBaseConditionRequest) PreProcess() {
	query.fix()
	query.TryLimit = query.Limit + 1
	query.AscWithFromPage = !query.OrderAsc && query.FromTailPage
}

func (query *NonmotorBaseConditionRequest) fix() {
	if strings.EqualFold(query.OrderBy, "") {
		query.OrderBy = "ts"
		query.OrderAsc = false
	}
}

type NonmotorCommonData struct {
	Timestamp              int64
	CaptureId              string
	SensorId               string
	NonmotorWithPeopleType int
}
