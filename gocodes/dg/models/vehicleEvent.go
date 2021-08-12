package models

type VehicleEvent struct {
	EventID              string
	EventReID            string
	Timestamp            int64
	SensorID             string
	SensorName           string
	VehicleRepoID        string //车辆比对库ID
	VehicleRepoName      string //车辆比对库名
	VehicleID            string //抓拍ID
	VehicleReID          string
	CapturedImage        *ImageResult  //图片
	VehicleAttr          *VehicleAttrs //对象在比对库信息
	RuleID               string
	IsChecked            int    //审核状态
	Comment              string //备注
	UserID               string //审核人
	Status               TaskStatus
	RuleName             string
	VehicleCaptureDetail *VehicleCaptureDetail
	SensorIntID          int32
	OrgCode              int64
}

type VehicleEventList struct {
	Rets  []*VehicleEvent
	Total int64
}
