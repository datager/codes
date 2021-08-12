package models

import (
	"errors"
	"time"
)

const (
	// face capture age statistics: enum
	FaceCaptureAttrStatisticsAge0To14   = 0 // [0, 14]
	FaceCaptureAttrStatisticsAge15To29  = 1 // [15, 29]
	FaceCaptureAttrStatisticsAge30To49  = 2 // [30, 49]
	FaceCaptureAttrStatisticsAge50To100 = 3 // [50, 100]
)

type CaptureStatistics struct {
	Last24H      []*CaptureStatisticsItem
	Accumulation *CaptureStatisticsItem
}

type CaptureStatisticsItem struct {
	Ts         int64
	Face       int64
	Pedestrian int64
	Vehicle    int64
	Nonmotor   int64
}

type CaptureStatisticsQuery struct {
	Types               []int
	NeedGroupBySensorID bool
	NeedGroupByTs       bool
	StartAndEndTimestamp
}

func (csi *CaptureStatisticsItem) Merge(new *CaptureStatisticsItem) {
	csi.Face = new.Face
	csi.Pedestrian = new.Pedestrian
	csi.Vehicle = new.Vehicle
	csi.Nonmotor = new.Nonmotor
}

type CaptureCountStatistics struct {
	ID       int64
	Uts      time.Time
	Ts       int64
	SensorID string
	Type     int
	Count    int64
}

type TasksCountStatistics struct {
	// 图片流任务数量
	TaskTypePicture int64
	// 离线视频任务数量
	TaskTypeVideo int64
	//实时视频任务数量
	TaskTypeRealTimeStream int64
	//抓拍机任务数量
	TaskTypeSnapShot int64
	//历史流任务数量
	TaskTypeHistoryStream int64
}

type SensorsCountStatistics struct {
	// ftp 抓拍机
	SensorTypeCaptureFTP int64
	// 网络摄像机
	SensorTypeIPC int64
	// 离线视频摄像机
	SensorTypeVideo int64
	// 图片流摄像机
	SensorTypePicture int64
	// pvg摄像机
	SensorTypePVG int64
	// 国标28181摄像机
	SensorTypeGB28181 int64
	// 慧目抓拍机
	SensorTypeHuiMu int64
	// 海康抓拍机
	SensorTypeHaiKang int64
	// 大华抓拍机
	SensorTypeDaHua int64
	// 华为抓拍机
	SensorTypeHuaWei int64
	// dc84抓拍机
	SensorTypeDC84 int64
}

type EventsCountStatisticsParam struct {
	StartTimestamp int64 `form:"start"`
	BackwardDays   int   `form:"backward"`
}

type EventsCountStatistics struct {
	// 起始日时间戳
	Ts int64
	// 人脸报警数量
	Face int64
	// 车辆报警数量
	Vehicle int64
}

type RepoCountStatistics struct {
	//总库数量
	RepositoryCount int
	//关注的库人员数量
	RepoCivilCount []RepoCivilCount
}

type RepoCivilCount struct {
	Name  string
	Count int
}

type PersonFileCountStatistics struct {
	// 今日档案数
	TodayCount int64
	// 实名档案数
	RealNameCount int64
	// 未实名档案数
	AnonymousCount int64
}

type VehicleFileCountStatistics struct {
	// 今日档案数
	TodayCount int64
	// 档案总数
	FileCount int64
}

type CaptureAttributeCountStatisticsRequest struct {
	SensorIDs []string // 设备列表
	StartAndEndTimestamp
}

type CaptureAttributeCountStatisticsToArchimedesRequest struct {
	SensorIDs     []string // 设备列表
	AttributeKeys []int64
	StartAndEndTimestamp
}

type FaceCaptureAttrCountStatisticsToArchimedesReq struct {
	SensorIDs     []string // 设备列表
	AttributeKeys []int64
	StartAndEndTimestamp
}

func (request *CaptureAttributeCountStatisticsRequest) Check() error {
	if request.StartTimestamp == 0 || request.EndTimestamp == 0 {
		return errors.New("timestamp shouldn't be empty")
	}
	if request.StartTimestamp > request.EndTimestamp {
		return errors.New("start_timestamp shouldn't larger than end_timestamp")
	}
	return nil
}

type CaptureAttributeCountStatisticsResponse struct {
	Key   int64 // 标识(对应VSE字典的valueID)
	Value int64 // 统计数量
}

//作为一个通用型的返回结构存在
type KeyValueCommon struct {
	Key   string
	Value int64
}

type KeyValueIllegal struct {
	Key   int64
	Value int64
}

type KeyValueCommonInt64 struct {
	Key   int64
	Value int64
}

type RepoDetailCountStatisticsRequest struct {
	RepoIDs []string // 比对库列表
	StartAndEndTimestamp
}

func (request *RepoDetailCountStatisticsRequest) Check() error {
	if request.StartTimestamp == 0 || request.EndTimestamp == 0 {
		return errors.New("timestamp shouldn't be empty")
	}
	return nil
}

type RepoDetailCountStatisticsResponse struct {
	Name  string // 名称
	Value int64  // 统计数量
}

type RepoDetailTimeSeriesCountStatisticsResponse struct {
	Timestamp int64 // 时间戳
	Value     int64 // 统计数量
}

type RepoStatisticsRequest struct {
	StartTime int64
	EndTime   int64
	RepoIDs   []string
	Type      int64
}

type CountResult struct {
	Count int
}

//报警统计req
type EventCountStatisticsRequest struct {
	SensorIds []string // 设备列表
	RepoIds   []string //库列表
	StartAndEndTimestamp
}

type EventCountStatisticsResp struct {
	SensorID   string
	SensorName string
	RepoID     string
	RepoName   string
	CountResult
}

func (request *EventCountStatisticsRequest) Check() error {
	if request.StartTimestamp == 0 || request.EndTimestamp == 0 {
		return errors.New("timestamp shouldn't be empty")
	}
	if request.StartTimestamp > request.EndTimestamp {
		return errors.New("start_timestamp shouldn't larger than end_timestamp")
	}
	return nil
}
