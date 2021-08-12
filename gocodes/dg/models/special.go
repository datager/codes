package models

type CoverFaceAtCertainTime struct {
	DailyAverage []int
	Yestday      []int
	Today        []int
}

type TopNeed struct {
	Top int
}

type VehicleSpecialRequest struct {
	Plate        string
	SensorID     []string
	Time         StartAndEndTimestamp
	OrderAsc     bool
	OrderBy      string
	FromTailPage bool
	Pagination
	TryLimit        int
	AscWithFromPage bool
}

type VehicleSpecialWithDriverAndPassengerRequest struct {
	*VehicleSpecialRequest
	DriverType    []int // 遮挡面部:(墨镜:3,遮阳板:4);没系安全带:1
	PassengerType []int // 遮挡面部:(墨镜:3,遮阳板:4);没系安全带:1
}

type NonmotorSpecialRequest struct {
	SensorID     []string
	Time         StartAndEndTimestamp
	OrderAsc     bool
	OrderBy      string
	FromTailPage bool
	Pagination
	TryLimit        int
	AscWithFromPage bool
}

type NonmotorWithPeople struct {
	*NonmotorSpecialRequest
	Type []int // 自行车:1,摩托车:2,三轮车:3
}

type ThemisItem struct {
	Timestamp int64
	CaptureID string
	SensorID  string
	PlateText string
}

type ThemisResponse struct {
	Total      int64
	NextOffset int
	Rets       []ThemisItem
}

type NonmotorThemisItem struct {
	Timestamp              int64
	CaptureID              string
	SensorID               string
	NonmotorWithPeopleType int
}

type NonmotorThemisResponse struct {
	Total      int64
	NextOffset int
	Rets       []NonmotorThemisItem
}

type SpecialCount struct {
	StartAndEndTimestamp
	TopN      int
	SpecialID string
	SensorIDs []string
}

type OnlyCount struct {
	Count int64
}

//按时段进行统计
type SpecialAtCertainTime struct {
	DailyAverage []int
	Yesterday    []int
	Today        []int
}
