package models

import (
	"strings"
)

type SensorQuery struct {
	OrgIds      []string `form:"orgid"`
	Recursive   bool     `form:"recursive"`
	SensorTypes []int    `form:"sensortype"`
	SensorName  string   `form:"sensorname"`
	//RequireSensorsDetail bool     `form:"sensorsdetail"`
	Offset              int   `form:"offset"`
	Limit               int   `form:"limit"`
	OuterPlatformStatus []int `form:"outerplatformstatus"`
	CurrentUser         *User

	// 基于设备可见性查询设备(通过组织或设备角色):
	// 1. 如果是默认设备角色, 置空 SensorIDs 条件即可;
	// 2. 如果不是默认设备角色, 置 SensorIDs为 visibleSensorIDs 即可.
	SensorIDs []string `form:"sensorids"` // 前端也可以在 GET 请求中只传 sensorids
}

func (query *SensorQuery) VisibleFilter(currentUser *User, userVisibleSensorIDs, allSensorIDsInPg []string) {
	query.CurrentUser = currentUser
	query.SensorIDs = SensorIDsVisibleOptimizerV1(query.SensorIDs, userVisibleSensorIDs, allSensorIDsInPg)
}

func (query *SensorQuery) CheckFuzzyMatching() bool {
	return !strings.EqualFold(query.SensorName, "")
}

func (query *SensorQuery) GetSensorTypes() []int {
	return query.SensorTypes
}

func (query *SensorQuery) GetLimit() int {
	return query.Limit
}

func (query *SensorQuery) GetOffset() int {
	return query.Offset
}
