package models

type Org struct {
	Ts            int64
	OrgID         string `json:"Id"`
	OrgName       string `json:"Name"`
	OrgLevel      int64
	OrgCode       int64
	SuperiorOrgId string
	// SuperiorOrgName string
	Comment      string
	Status       int
	SensorCount  int
	Orgs         []*Org
	DeviceRoles  []*DeviceRole
	FaceRepos    []*FaceRepo
	VehicleRepos []*FaceRepo
	FaceRules    []*FaceRule
	FuncRoles    []*FuncRole
	Sensors      []*OrgSensor
	Users        []*User
	TaskCount    int
}

type OrgOrder []*Org

func (o OrgOrder) Len() int {
	return len(o)
}

func (o OrgOrder) Less(i, j int) bool {
	return o[i].Ts > o[j].Ts
}

func (o OrgOrder) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type BatchAddOrgsRequest struct {
	Ts            int64
	SuperiorOrgId string
	OrgNames      []string
	Comment       string
}

type OrgQuery struct {
	Recursive              bool  `form:"recursive"`
	IncludeSelf            bool  `form:"self"`
	IncludeDeviceRoles     bool  `form:"deviceroles"`
	IncludeFaceRepos       bool  `form:"facerepos"`
	IncludeVehicleRepos    bool  `form:"vehiclerepos"`
	IncludeFaceRules       bool  `form:"facerules"`
	IncludeFuncRoles       bool  `form:"funcroles"`
	IncludeSensors         bool  `form:"sensors"`
	IncludeUsers           bool  `form:"users"`
	IncludeSensorCount     bool  `form:"sensorcount"`
	IncludeTasks           bool  `form:"tasks"`
	IncludeMonitorRules    bool  `form:"monitorrules"`
	SensorTypes            []int `form:"sensortype"`
	RequireSensorsDetail   bool  `form:"sensorsdetail"`
	RequireFaceReposDetail bool  `form:"facereposdetail"`
	RequireFaceRulesDetail bool  `form:"facerulesdetail"`
	CurrentUser            *User
}

// IncludeSensorCount 复用 IncludeSensors 的逻辑: 只是登陆者可见的设备/设备数量
// IncludeSensors = true,  IncludeSensorCount = true:  返回   Sensors, 返回 SensorCount
// IncludeSensors = true,  IncludeSensorCount = false: 返回   Sensors, 返回 SensorCount
// IncludeSensors = false, IncludeSensorCount = true:  返回   Sensors, 返回 SensorCount
// IncludeSensors = false, IncludeSensorCount = false: 不返回 Sensors, 不返回 SensorCount
func (query *OrgQuery) FixParam() {
	// 如果要IncludeSensorCount, 就置IncludeSensors为 true, 并其结果中的len(sensor)处拿
	if query.IncludeSensorCount {
		query.IncludeSensors = true
	}
}

func (query *OrgQuery) GetSensorTypeMap() map[int]bool {
	length := len(query.SensorTypes)
	if length == 0 {
		return nil
	}
	m := make(map[int]bool, length)
	for _, t := range query.SensorTypes {
		m[t] = true
	}
	return m
}

type OrgListParam struct {
	OrgID   string
	OrgName string
	Offset  int
	Limit   int
}

type OrgDict struct {
	OrgID         string
	OrgCode       int64
	OrgName       string
	OrgLevel      int64
	SuperiorOrgID string
	Comment       string
	Status        int
	Ts            int64
}

type OrgResources struct {
	Sensors []*Sensor
	Orgs    []*OrgDict
}

type Resources struct {
	Sensors []*SensorInfo
	Tasks   []*TaskInfo
	Repos   []*RepoInfo
	Rules   []*MonitorRuleInfo
}
