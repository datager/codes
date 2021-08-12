package models

import (
	"time"

	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/json"
	"codes/gocodes/dg/utils/pool"
)

//任务目标类型常量
const (
	DetTypePedestrian = 1 << iota //任务行人类型
	DetTypeVehicle                //任务机动车类型
	DetTypeNonmotor               //任务非机动车类型
	DetTypeFace                   //任务人脸类型
	DetTypeAllObj                 // 全目标
)

//任务状态常量
const (
	TaskStatusCreating     = 1  //创建中
	TaskStatusPending      = 2  //排队中
	TaskStatusRunning      = 3  //运行中
	TaskStatusFailed       = 4  //失败
	TaskStatusEnd          = 5  //已结束
	TaskStatusDeleting     = 6  //删除中
	TaskStatusDeleted      = 7  //已删除
	TaskStatusDeleteFailed = 8  //删除失败
	TaskStatusModifying    = 9  //修改中
	TaskStatusPausing      = 10 //暂停中
	TaskStatusPaused       = 11 //已暂停
	TaskStatusStarting     = 12 //开始中
	TaskStatusRetrying     = 13 // 重试
	TaskStatusLimited      = 14 // 路数受限
)

var (
	TaskStatusStable            = []int{TaskStatusRunning, TaskStatusFailed, TaskStatusEnd, TaskStatusDeleted, TaskStatusDeleteFailed, TaskStatusPaused, TaskStatusRetrying, TaskStatusLimited, TaskStatusPending} // 稳定任务状态
	TaskStatusStableMap         map[int]struct{}
	TaskStatusOccupiedResources = []int{TaskStatusRunning, TaskStatusPausing, TaskStatusRetrying} // 正在占用资源的任务状态
)

func init() {
	//根据TaskStatusStable动态生成Map
	TaskStatusStableMap = make(map[int]struct{})
	for _, v := range TaskStatusStable {
		TaskStatusStableMap[v] = struct{}{}
	}
}

//任务类型常量
const (
	_                      = iota
	TaskTypePicture        //图片
	TaskTypeVideo          //视频
	TaskTypeRealTimeStream //实时流
	TaskTypeSnapShot       //抓拍机
	TaskTypeHistoryStream  //历史流
)

//抽帧/全帧分类
const (
	_                    = iota
	FramingStrategyFull  //全帧
	FramingStrategyDrawe //抽帧
)

//返回数据类型
const (
	_                         = iota
	ResultTypeDetail          //详情
	ResultTypeIDList          //只返回ID列表
	ResultTypeIDListAndDetail //返回详情和ID列表
)

//返回任务查询列表
type ResultTaskList struct {
	Total      int64      //筛选结果总数
	Rets       []*Tasks   //任务列表
	TaskIDList TaskIDList //ID列表
}

//ID列表
type TaskIDList []string

//返回批量任务处理结果
type ResultMapBool map[string]bool

//批量传递简单任务
type SimpleTaskList []SimpleTask

//简单任务信息
type SimpleTask struct {
	TaskID       string     //任务ID
	SensorID     string     //设备ID
	SensorType   int        //设备类型
	DetTypesInfo []int      //任务目标类型未聚合数组
	Config       TaskConfig //任务配置
	Status       int        //任务状态
	Type         int        //任务类型
	Channel      int        //任务占用路数
}

type UpdateSimpleTask struct {
	TaskID       string      //任务ID
	SensorID     string      //设备ID
	DetTypesInfo []int       //任务目标类型未聚合数组
	Config       *TaskConfig //任务配置
	Status       int         //任务状态
	Type         int         //任务类型
}

//返回的数据详情内容
type Tasks struct {
	SimpleTask
	LinkID     string    //linkID
	Node       string    //记录Link相关的Node信息
	Uts        time.Time //最后一次更新时间：毫秒单位
	OrgID      string    //OrgID
	OrgName    string    //组织名称
	SensorName string    //设备名称
	Progress   int       //进度
	RenderURL  string    //渲染流地址
	Ts         int64     //创建时间：毫秒单位
	DetTypes   int       `json:"-"` //任务目标类型聚合内容
}

type TaskInfo struct {
	SensorID   string
	DetType    int
	TaskStatus int
	TaskID     string
}

func (task Tasks) Pk() string {
	return task.TaskID
}
func (task Tasks) CurrentStatus() int {
	return task.Status
}
func (task Tasks) Instance() interface{} {
	return task
}
func (task Tasks) LockerKey() string {
	return task.SensorID
}
func (task Tasks) LinkNode() string {
	return task.Node
}
func (task Tasks) VideoProgress() int {
	return task.Progress
}
func (task Tasks) RenderURI() string {
	return task.RenderURL
}

func (task Tasks) SchedulerType() int {
	return pool.SchedulerTypeTask
}

func (task Tasks) CurrentChannel() int {
	return task.Channel
}

//任务可配置项
type TaskConfig struct {
	Speed                int       //速度：视频任务可用
	FramingStrategy      int       //抽帧/全帧：视频任务可用
	VideoStartTime       time.Time `json:"-"` //任务相对视频的开始时间：视频任务可用
	VideoEndTime         time.Time `json:"-"` //任务相对视频的结束时间：视频任务可用
	BaseTime             time.Time `json:"-"` //离线视频任务的基准时间：任务开始跑了之后，以这个时间为进准对分析结果开始记录
	ConfigVideoStartTime int64     `json:"VideoStartTime"`
	ConfigVideoEndTime   int64     `json:"VideoEndTime"`
	ConfigBaseTime       int64     `json:"BaseTime"`
	Rois                 []Roi     //roi信息
}

func (tconfig *TaskConfig) ToDB() ([]byte, error) {
	return json.Marshal(tconfig)
}

func (tconfig *TaskConfig) FromDB(b []byte) error {
	err := json.Unmarshal(b, tconfig)
	tconfig.VideoStartTime = utils.TsToTime(tconfig.ConfigVideoStartTime)
	tconfig.VideoEndTime = utils.TsToTime(tconfig.ConfigVideoEndTime)
	tconfig.BaseTime = utils.TsToTime(tconfig.ConfigBaseTime)
	return err
}

func (tconfig *TaskConfig) SetToTime() {
	tconfig.VideoStartTime = utils.TsToTime(tconfig.ConfigVideoStartTime)
	tconfig.VideoEndTime = utils.TsToTime(tconfig.ConfigVideoEndTime)
	tconfig.BaseTime = utils.TsToTime(tconfig.ConfigBaseTime)
}

func (tconfig *TaskConfig) GetFromTime() {
	tconfig.ConfigVideoStartTime = utils.TimeToTs(tconfig.VideoStartTime)
	tconfig.ConfigVideoEndTime = utils.TimeToTs(tconfig.VideoEndTime)
	tconfig.ConfigBaseTime = utils.TimeToTs(tconfig.BaseTime)
}

//检索接口接收参数
type SearchTaskRequest struct {
	FramingStrategy int      `form:"FramingStrategy" json:"FramingStrategy"` //抽帧/全帧查询
	ResultType      int64    `form:"ResultType" json:"ResultType"`           //返回数据类型
	SensorName      string   `form:"SensorName" json:"SensorName"`           //设备名称
	TaskStatus      int64    `form:"TaskStatus" json:"TaskStatus"`           //任务状态
	DetType         int      `form:"DetType" json:"DetType"`                 //任务目标类型
	TaskType        int      `form:"TaskType" json:"TaskType"`               //任务类型
	Offset          int      `form:"Offset" json:"Offset"`                   //起始数
	Limit           int      `form:"Limit" json:"Limit"`                     //分页大小
	OrgIDs          []string `form:"-" json:"-"`                             //用户所属组织ID展开
	CurrentUser     *User    `form:"-" json:"-"`                             //当前用户信息
}

//统计结果
type TaskStatistics struct {
	ProcessingCount int64 //正在处理数量
	RemainCount     int64 //还可以处理数量
	LinkStatus      int   //link的状态，0为异常，1为正常
	Ability         struct {
		//结构化处理能力统计
		PvnSum    int64 //人车非
		Face      int64 //人脸
		AllObject int64 // 全目标
	}
	Task struct {
		//结构化任务类型统计
		Vehicle    int64 //车辆
		Pedestrian int64 //行人
		Nonmotor   int64 //非机动车
		Face       int64 //人脸
		AllObject  int64 // 全目标
	}
}

//设定任务目标类型的详细标记
func (t *Tasks) SetDetTypeInfo() {
	t.DetTypesInfo = TypeToTypeSlice(t.DetTypes)
}

//把任务目标类型的数组标记转换到聚合后的数据里面
func (t *Tasks) SetDetType() {
	t.DetTypes = TypeSliceToType(t.DetTypesInfo)
}

func TypeSliceToType(list []int) int {
	rs := 0
	for _, v := range list {
		if v == DetTypePedestrian {
			rs = rs | DetTypePedestrian
		}
		if v == DetTypeVehicle {
			rs = rs | DetTypeVehicle
		}
		if v == DetTypeNonmotor {
			rs = rs | DetTypeNonmotor
		}
		if v == DetTypeFace {
			rs = rs | DetTypeFace
		}
		if v == DetTypeAllObj {
			rs = rs | DetTypeAllObj
		}
	}
	return rs
}

func TypeToTypeSlice(t int) []int {
	rs := make([]int, 0, 4)

	if t&DetTypePedestrian == DetTypePedestrian {
		rs = append(rs, DetTypePedestrian)
	}

	if t&DetTypeVehicle == DetTypeVehicle {
		rs = append(rs, DetTypeVehicle)
	}

	if t&DetTypeNonmotor == DetTypeNonmotor {
		rs = append(rs, DetTypeNonmotor)
	}

	if t&DetTypeFace == DetTypeFace {
		rs = append(rs, DetTypeFace)
	}

	if t&DetTypeAllObj == DetTypeAllObj {
		rs = append(rs, DetTypeAllObj)
	}
	return rs
}
