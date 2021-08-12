package models

import (
	"codes/gocodes/dg/utils/slice"
)

// GetSensorIDs
// 1. ByUserVisibleFilter
// 2. BySQLOptimizer, 有 in sensor 优化; 但若设备被删, 搜抓拍会多搜出结果
func SensorIDsVisibleOptimizerV1(userQuerySensorIDs, userVisibleSensorIDs, allSensorIDsInPg []string) []string {
	visibleSensorIDsOfUserQuery := visibleSensorIDsFilter(userQuerySensorIDs, userVisibleSensorIDs)
	return sqlOptimizerSensorIDs(visibleSensorIDsOfUserQuery, allSensorIDsInPg)
}

// GetSensorIDs
// 1. ByUserVisibleFilter, 无 in sensor 优化; 但若设备被删, 搜抓拍结果正常
func SensorIDsVisibleOptimizerV2(userQuerySensorIDs, userVisibleSensorIDs []string) []string {
	return visibleSensorIDsFilter(userQuerySensorIDs, userVisibleSensorIDs)
}

// visibleFilterOfUserQuerySensorIDs
// userQuery, userVisible, allInPg: sensorIDs
func visibleSensorIDsFilter(userQuery, userVisible []string) (visibleOfUserQuery []string) {
	// log.Debugf("[before--visible--filter] 用户请求中的 len(sensor) %d", len(userQuery))
	// log.Debugf("[before--visible--filter] 用户可见  的 len(sensor) %d", len(userVisible))

	// 是否有 sensor 检索条件
	if len(userQuery) == 0 {
		visibleOfUserQuery = userVisible
	} else {
		visibleOfUserQuery = slice.GetIntersectionStrings(userQuery, userVisible)
	}
	// log.Debugf("[after--visible--filter] 可见过滤后, 用户请求的 len(sensor) %d", len(visibleOfUserQuery))
	return
}

// sqlOptimizerOfUserVisibleSensorIDs
// if == allSensorIDsInPg, do not use In(sensor_ids)... condition in sql
func sqlOptimizerSensorIDs(visibleOfUserQuery, allInPg []string) (SQLOptimizerSensorIDs []string) {
	// log.Debugf("[before--SQLOptimizer--filter] sql优化前, 数据库中所有的 len(sensor) %d", len(allInPg))
	if slice.IsEqualWithoutOrder(visibleOfUserQuery, allInPg) {
		SQLOptimizerSensorIDs = nil
	} else {
		SQLOptimizerSensorIDs = visibleOfUserQuery
	}
	// log.Debugf("[after--SQLOptimizer--filter] sql优化后, 用户请求的 len(sensor) %d", len(SQLOptimizerSensorIDs))
	return
}
