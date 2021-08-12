package models

type VehicleEventQueryRequest struct {
	Type    int
	Payload *VehicleEventQuery
}

type VehicleEventQuery struct {
	UTCOffset      int
	Total          int
	CheckStatus    []int    //审核状态 0:未审核 1：有效 2：无效
	StartTimestamp int64    //开始时间
	EndTimestamp   int64    //结束时间
	OrderAsc       bool     //排序
	OrderBy        string   //排序字段
	PlateText      string   //车牌
	OwnerName      string   //车主姓名
	RepoIds        []string //比对库IDs
	SensorIds      []string //设备IDs
	VehicleIds     []string
	AttrID         string `json:"AttrId"`
	RuleName       string //布控名
	PlateColorIds  []int  //车牌颜色
	ColorIds       []int  //车辆颜色
	Pagination            //limit offset
}

func (query *VehicleEventQuery) ShallowClone() *VehicleEventQuery {
	return &VehicleEventQuery{
		UTCOffset:      query.UTCOffset,
		Total:          query.Total,
		CheckStatus:    query.CheckStatus,
		StartTimestamp: query.StartTimestamp,
		EndTimestamp:   query.EndTimestamp,
		OrderAsc:       query.OrderAsc,
		OrderBy:        query.OrderBy,
		OwnerName:      query.OwnerName,
		RepoIds:        query.RepoIds,
		SensorIds:      query.SensorIds,
		RuleName:       query.RuleName,
		PlateText:      query.PlateText,
		ColorIds:       query.ColorIds,
		PlateColorIds:  query.PlateColorIds,
		VehicleIds:     query.VehicleIds,
		Pagination: Pagination{
			Limit:  query.Limit,
			Offset: query.Offset,
		},
	}
}

func (query *VehicleEventQuery) GetStartTimestamp() int64 {
	return query.StartTimestamp
}

func (query *VehicleEventQuery) GetEndTimestamp() int64 {
	return query.EndTimestamp
}

func (query *VehicleEventQuery) GetSensorIds() []string {
	return query.SensorIds
}

func (query *VehicleEventQuery) GetVehicleIds() []string {
	return query.VehicleIds
}

func (query *VehicleEventQuery) GetRepoIds() []string {
	return query.RepoIds
}

func (query *VehicleEventQuery) GetName() string {
	return query.OwnerName
}

func (query *VehicleEventQuery) GetPlateText() string {
	return query.PlateText
}

func (query *VehicleEventQuery) GetColorIDs() []int {
	return query.ColorIds
}

func (query *VehicleEventQuery) GetPlateColorIDs() []int {
	return query.PlateColorIds
}

func (query *VehicleEventQuery) GetSortBy() string {
	return query.OrderBy
}

func (query *VehicleEventQuery) GetSortAsc() bool {
	return query.OrderAsc
}

func (query *VehicleEventQuery) GetLimit() int {
	return query.Limit
}

func (query *VehicleEventQuery) GetOffset() int {
	return query.Offset
}

func (query *VehicleEventQuery) GetUTCOffset() int {
	return query.UTCOffset
}

func (query *VehicleEventQuery) GetGenderID() int {
	return 0
}

func (query *VehicleEventQuery) GetSearchID() string {
	return ""
}
