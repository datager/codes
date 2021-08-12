package models

import "strings"

// 布控状态
type MonitorRuleStatus int64

const (
	MonitorRuleStatusAll MonitorRuleStatus = 0
	MonitorRuleStatusOn  MonitorRuleStatus = 1
	MonitorRuleStatusOff MonitorRuleStatus = 2
)

// 布控类型
type MonitorRuleType int64

const (
	MonitorRuleTypeAll                    MonitorRuleType = 0 // 不可写入, 只可检索: 全部
	MonitorRuleTypeOnlyVehicleRepo        MonitorRuleType = 1 // 可写入, 可检索:     库仅有车辆的,
	MonitorRuleTypeOnlyFaceRepo           MonitorRuleType = 2 // 可写入, 可检索: 	   库仅有人脸的,
	MonitorRuleTypeBothVehicleAndFaceRepo MonitorRuleType = 3 // 可写入, 可检索:     库既有车辆又有人脸的
	MonitorRuleTypeHasAnyVehicleRepo      MonitorRuleType = 4 // 不可写入, 只可检索: 库包含车辆的(仅有车辆 or 即有人脸又有车辆),
	MonitorRuleTypeHasAnyFaceRepo         MonitorRuleType = 5 // 不可写入, 只可检索: 库包含人脸的(仅有人脸 or 即有人脸又有车辆)
	MonitorRuleTypeUnknown                MonitorRuleType = 9
)

// 布控库类型
type RepoType int64

const (
	RepoTypeAll     RepoType = 0
	RepoTypeVehicle RepoType = 1
	RepoTypeFace    RepoType = 2
	RepoTypeUnknown RepoType = 9 // 在 sql 中, 默认写入 2 为人脸类型, 9 没用到
)

// 获取布控详情 响应体
type MonitorRule struct {
	MonitorRuleId   string
	MonitorRuleName string
	Timestamp       int64
	Sensors         []*Sensor
	FaceRepos       []*RepoInRule // todo: 和前端沟通, 该字段类型不变, 但是名称改为 RepoInRules
	TimeIsLong      bool          //Times *TimeInRule		// 布控时段
	// Receivers []Users									// 创建人
	// MonitorAlarms []*MonitorAlarms 						// 报警数量
	Comment                           string
	Type                              MonitorRuleType   // 类型:人/车
	Status                            MonitorRuleStatus // 开关状态
	OrgID                             string
	UserID                            string
	LastUpdateOrgID                   string
	LastUpdateUserID                  string
	CanDeleteOrEditNonAlarmLevelVoice bool // 是否可删除, 是否可编辑"除报警声音等级之外的项"(如布控名称, 增减设备, 增减库, 改库阈值)
	//CurrentUser *User 					// 用于鉴权
}

// 创建/编辑 布控 请求体
type MonitorRuleSet struct {
	MonitorRuleID   string `uri:"monitorruleid"`
	MonitorRuleName string
	SensorIds       []string
	RepoInRules     []*RepoInRule
	TimeIsLong      bool
	Comment         string
	Type            MonitorRuleType
	Status          MonitorRuleStatus // 开关状态
	CurrentUser     *User             // 用于鉴权
}

// 查询 布控 请求体
type MonitorRuleQueryReq struct {
	ResultType      int64 // 返回数据类型, 如果值为 3 需要返回 MonitorRuleIdList的[]string, 和 task 接口相同
	MonitorRuleId   string
	MonitorRuleName string
	SensorIds       []string
	FaceRepoIds     []string
	Type            MonitorRuleType   // 类型:人/车
	Status          MonitorRuleStatus // 开关状态
	Offset          int
	Limit           int
	CurrentUser     *User
}

// 布控库 类型+数量
type RepoTypeCount struct {
	Type  RepoType
	Count int64
}

// 返回布控查询列表
type ResultMonitorRuleList struct {
	Total             int64                       //筛选结果总数
	Rets              []*MonitorRuleQueryRespList //布控列表
	MonitorRuleIDList []string                    //ID列表
}

// 查询布控列表 响应体
type MonitorRuleQueryRespList struct {
	MonitorRuleId                     string
	MonitorRuleName                   string
	Type                              []RepoType // 类型:人/车
	RepoTypeCount                     []*RepoTypeCount
	SensorCount                       int64
	TimeIsLong                        bool
	CreateUserName                    string
	Status                            MonitorRuleStatus // 开关状态
	CanDeleteOrEditNonAlarmLevelVoice bool              // 是否可删除, 是否可编辑"除报警声音等级之外的项"(如布控名称, 增减设备, 增减库, 改库阈值)
	Offset                            int
	Limit                             int
}

// 批量修改 布控状态 请求体
type MonitorRuleBatchEditStatusReq struct {
	MonitorRuleID string            // 某布控
	DstEditStatus MonitorRuleStatus // 该布控的 目标修改状态
}

type BatchEditMonitorRuleStatusRet struct {
	MonitorRuleID string
	DstEditStatus MonitorRuleStatus
	IsCanEdit     bool
	Err           error
}

type MonitorRuleInfo struct {
	MonitorRuleID   string
	MonitorRuleName string
}

func NewBatchEditMonitorRuleStatusRet(ruleID string, dstEditStatus MonitorRuleStatus, isCanEdit bool, err error) *BatchEditMonitorRuleStatusRet {
	return &BatchEditMonitorRuleStatusRet{
		MonitorRuleID: ruleID,
		DstEditStatus: dstEditStatus,
		IsCanEdit:     isCanEdit,
		Err:           err,
	}
}

func (mr *MonitorRule) ToMonitorRuleSetModel() *MonitorRuleSet {
	mrSet := &MonitorRuleSet{
		MonitorRuleID:   mr.MonitorRuleId,
		MonitorRuleName: mr.MonitorRuleName,
		SensorIds:       SensorToSensorIDs(mr.Sensors),
		RepoInRules:     mr.FaceRepos,
		TimeIsLong:      mr.TimeIsLong,
		Comment:         mr.Comment,
		Type:            mr.Type,
		Status:          mr.Status,
	}
	return mrSet
}

func (mrs *MonitorRuleSet) ToMonitorRuleModel() *MonitorRule {
	sensors := make([]*Sensor, 0)
	for _, sID := range mrs.SensorIds {
		sensors = append(sensors, &Sensor{
			SensorID: sID,
		})
	}

	mr := &MonitorRule{
		MonitorRuleId:    mrs.MonitorRuleID,
		MonitorRuleName:  mrs.MonitorRuleName,
		Sensors:          sensors,
		FaceRepos:        mrs.RepoInRules,
		TimeIsLong:       mrs.TimeIsLong,
		Comment:          mrs.Comment,
		LastUpdateOrgID:  mrs.CurrentUser.OrgId,
		LastUpdateUserID: mrs.CurrentUser.UserId,
		Type:             mrs.Type,
		Status:           mrs.Status,
	}
	return mr
}

// 查询 布控 请求体, 内部 repository 层使用
// 和MonitorRuleQueryReq的区别是:
// 1. MonitorRuleId 是数组
type MonitorRuleQueryReqForRepository struct {
	MonitorRuleIds  []string
	MonitorRuleName string
	SensorIds       []string          // 此条件已经在 MonitorRuleIds 中体现
	FaceRepoIds     []string          // 此条件已经在 MonitorRuleIds 中体现
	Type            MonitorRuleType   // 类型:人/车
	Status          MonitorRuleStatus // 开关状态
	Offset          int
	Limit           int
	CurrentUser     *User
}

func (query *MonitorRuleQueryReqForRepository) CheckFuzzyMatching() bool {
	return !strings.EqualFold(query.MonitorRuleName, "")
}

func (mrqr *MonitorRuleQueryReq) ToMonitorRuleQueryReqForRepositoryModel(monitorRuleIDs []string) *MonitorRuleQueryReqForRepository {
	queryForRepository := &MonitorRuleQueryReqForRepository{
		MonitorRuleIds:  monitorRuleIDs,
		MonitorRuleName: mrqr.MonitorRuleName,
		SensorIds:       mrqr.SensorIds,   // 此条件已经在 MonitorRuleIds 中体现
		FaceRepoIds:     mrqr.FaceRepoIds, // 此条件已经在 MonitorRuleIds 中体现
		Type:            mrqr.Type,        // 类型:人/车
		Status:          mrqr.Status,      // 开关状态
		Offset:          mrqr.Offset,
		Limit:           mrqr.Limit,
		CurrentUser:     mrqr.CurrentUser,
	}
	return queryForRepository
}
