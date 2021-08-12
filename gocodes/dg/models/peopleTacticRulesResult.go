package models

const (
	StatusUnconfirmed = 1 //待确认
	StatusRemove      = 2 //排除
	StatusConfirmed   = 3 //已确认
)
const (
	IsFollowUnknow = 0 //未知
	IsFollow       = 1 //关注
	IsFollowNot    = 2 //未关注
)

type UpdateRulesResultStatusRequest struct {
	IsFollow int
	Status   int //状态
}

type CountStatusResponse struct {
	IsFollowNum    int64 //关注数量
	UnconfirmedNum int64 //待确认数量
	RemoveNum      int64 //排除数量
}

type TacticRulesResultRequest struct {
	TacticRuleIDs  []string //规则
	UserID         string
	KeyWord        string   // 关键词
	StartTimestamp int64    //起始时间
	EndTimestamp   int64    //结束时间
	SensorIDs      []string //设备列表
	Status         int      //状态
	Limit          int      // 页面条数
	Offset         int      // 偏移量
	IsFollow       int      //是否关注
}
type TacticRulesResult struct {
	ResultID        string
	CreateTime      int64
	EventID         string
	EventReID       string
	UserID          string
	Timestamp       int64
	TacticRuleID    string
	SensorID        string
	SensorName      string
	FaceRepoID      string
	FaceRepoName    string
	FaceID          string
	FaceReID        string
	Confidence      float32
	CivilAttr       *CivilAttr
	CapturedImage   *ImageResult
	RegisteredImage *ImageResult
	Status          int   //状态
	IsFollow        int   //是否关注
	EventCount      int64 //报警数量
}

type GetResultSensorsResponse struct {
	Total int64
	Rets  []*ResultSensor
}
type ResultSensor struct {
	SensorID string
	Count    int64
}

type TacticRulesResultResponse struct {
	Total int64
	Rets  []*TacticRulesResult
}

type TacticRulesResultKafkaInputData struct {
	RuleID      string
	HitSignCode string
	Data        *FaceWarnIndex
}

// 人脸报警索引
type FaceWarnIndex struct {
	*FaceWarn
	EnterTimeMilliseconds int64
	LeaveTimeMilSeconds   int64
	DataIndexType         int32
}

// 人脸报警
type FaceWarn struct {
	EventId         string       `json:",omitempty"`
	EventReId       string       `json:",omitempty"`
	EventType       int32        `json:",omitempty"`
	Timestamp       int64        `json:",omitempty"`
	SensorId        string       `json:",omitempty"`
	SensorName      string       `json:",omitempty"`
	FaceRepoId      string       `json:",omitempty"`
	FaceRepoName    string       `json:",omitempty"`
	NameListAttr    int32        `json:",omitempty"`
	FaceId          string       `json:",omitempty"`
	FaceReId        string       `json:",omitempty"`
	CapturedImage   *ImageResult `json:",omitempty"`
	ImageId         string       `json:",omitempty"`
	CaptureAttr     *CivilAttr   `json:",omitempty"`
	CivilAttr       *CivilAttr   `json:",omitempty"`
	RegisteredImage *ImageResult `json:",omitempty"`
	RuleId          string       `json:",omitempty"`
	Rule            string       `json:",omitempty"`
	Confidence      float32      `json:",omitempty"`
	IsChecked       int32        `json:",omitempty"`
	Comment         string       `json:",omitempty"`
	UserIds         string       `json:",omitempty"`
	MonitorRuleId   string       `json:",omitempty"`
	MonitorRuleName string       `json:",omitempty"`
	AlarmLevel      int          `json:",omitempty"`
	AlarmVoice      string       `json:",omitempty"`
	Status          int          `json:",omitempty"`
	Longitude       float64      `json:",omitempty"`
	Latitude        float64      `json:",omitempty"`
	SensorType      int          `json:",omitempty"`
}
