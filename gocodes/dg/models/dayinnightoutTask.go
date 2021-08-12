package models

const (
	MaxDayInNightOutTaskNumInAthena                               = 10  // 所有 athena 里 最多10个 任务
	MaxDayInNightOutTaskNumOfOneUser                              = 10  // 单人最多 10个 任务
	VehicleDayInNightOutCreateTaskMaxSensorCount                  = 500 // 最多 500 sensor
	VehicleDayInNightOutTaskMaxDurationDaysBetweenStartTsAndEndTs = 31  // 最多 30 天
	DayInNightOutCreateTaskMaxTaskNameLen                         = 20  // 任务名必须<=20
)

const (
	DayInNightOutTaskErr   = iota
	DayInNightOutTaskStop  //停止碰撞任务
	DayInNightOutTaskStart //启动碰撞任务
)

type BatchDeleteReqDayInNightOutModel struct {
	TaskIDs []string //删除的任务IDs
}

type VehicleDayInNightOutTaskConfigInfo struct {
	TaskID        string
	TaskName      string
	SensorIDs     []string
	DayNightPairs []*DayNightPair
	PlateText     string // not required
}

// 一对儿昼夜时段 (一天一个)
// note: all time must be utc param
type DayNightPair struct {
	SharpDate int64       // 整点日期 eg: 1573401600000(11月1日0点0分0秒)
	Day       *TimePeriod // 昼时段 eg: (7-9点=>11月1日7点-11月1日9点)
	Night     *TimePeriod // 夜时段 eg: (23-1点=>11月1日23点-11月2日1点)
}

// 时间段
type TimePeriod struct {
	St int64
	Et int64
}

type VehicleDayInNightOutTaskStatusInfo struct {
	//任务状态类型
	StateType int32

	// 任务排队信息
	PosInQueue  int
	QueueLength int

	//任务进度信息
	ProcessBar int // 0~100

	Err string
}

//雅典娜 Task 任务
type AthenaUnmarshalVehicleDayInNightOutTaskStatusResp struct {
	ID         string // 任务ID
	Type       int    // 任务类别
	ConfigInfo *VehicleDayInNightOutTaskConfigInfo
	StatusInfo *VehicleDayInNightOutTaskStatusInfo
}

//用于查询列表
type QueryDayInNightOutTaskListReq struct {
	AthenaTaskStatus []int // 任务状态(排队中、进行中、已完成、暂停中)
	AthenaTaskType   []int
	CurrentUser      *User // loki内部使用
}

type QueryDayInNightOutTaskListResp struct {
	TaskID                string //任务ID
	TaskName              string // 任务名称
	ErrMsg                string
	ProcessPercent        float64 // 处理进度
	CreateTs              int64   // 任务创建时间
	AthenaTaskStatus      int     //任务状态
	ResultCount           int     //结果数量
	DayInNightOutTaskType int     //任务类型
	StartAndEndTimestamp  []int64
}

type QueryDayInNightOutTaskListResps []*QueryDayInNightOutTaskListResp

func NewDayInNightOutTaskListQueryResps() QueryDayInNightOutTaskListResps {
	return make(QueryDayInNightOutTaskListResps, 0)
}

func (resps QueryDayInNightOutTaskListResps) Len() int {
	return len(resps)
}

// CreateTs 降序
func (resps QueryDayInNightOutTaskListResps) Less(i, j int) bool {
	return resps[i].CreateTs > resps[j].CreateTs
}

func (resps QueryDayInNightOutTaskListResps) Swap(i, j int) {
	resps[i], resps[j] = resps[j], resps[i]
}

func (query *QueryDayInNightOutTaskListReq) FilterByquencyType(t int) *QueryDayInNightOutTaskListReq {
	filteredTypes := make([]int, 0)
	for _, tp := range query.AthenaTaskType {
		if tp == t {
			filteredTypes = append(filteredTypes, t)
		}
	}

	filtered := &QueryDayInNightOutTaskListReq{
		AthenaTaskStatus: query.AthenaTaskStatus,
		AthenaTaskType:   filteredTypes,
		CurrentUser:      query.CurrentUser,
	}
	return filtered
}

type BatchDelDayInNightOutTasksResp struct {
	IsAllTasksDeleteSucceed bool           // 是否全部task删除成功
	DetailInfo              ResultMapError // 详情, key 为 taskID, 值为 err 原因
}

func NewBatchDelDayInNightOutTasksResp() *BatchDelDayInNightOutTasksResp {
	return &BatchDelDayInNightOutTasksResp{
		IsAllTasksDeleteSucceed: true,
		DetailInfo:              make(map[string]error),
	}
}
