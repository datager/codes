package models

type MonitorRuleExtend struct {
	MonitorRule                 *MonitorRule
	MonitorRuleSensorRelation   *MonitorRuleSensorRelation
	MonitorRuleFaceRepoRelation *MonitorRuleFaceRepoRelation
}

func (e *MonitorRuleExtend) ToMonitorRuleQueryRespListModel() *MonitorRuleQueryRespList {
	respList := &MonitorRuleQueryRespList{
		MonitorRuleId:   e.MonitorRule.MonitorRuleId,
		MonitorRuleName: e.MonitorRule.MonitorRuleName,
		//Type: mr.Type,
		//RepoTypeCount   []RepoTypeCount
		//SensorCount:
		TimeIsLong: e.MonitorRule.TimeIsLong,
		//CreateUserName  string,
		Status: e.MonitorRule.Status,
		//Offset int
		//Limit    int
	}
	return respList
}
