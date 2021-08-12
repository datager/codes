package models

type AddPeopleTacticRulesRequest struct {
	TacticID       string   //技术战法id
	UserID         string   //user_id
	TacticRuleName string   //战法规则名称
	AnalysisType   int      //分析类型枚举
	AnalysisItem   []int    //配置项
	IsOpen         bool     //规则是否开启（true 开启，false 关闭）
	WarningLevel   int      //预警级别
	RepoIDs        []string //比对库列表
	Parameters     *Parameters
}
type Parameters struct {
	TimeFrequencyTraceBackDuration int64               //时间频次回溯时长
	TimeFrequencyNumber            int                 //时间频次次数
	AppearNumber                   int                 //出现次数
	AreaBreadth                    int                 //区域个数
	AreaDepth                      int                 //区域频次
	SensorIDs                      []string            //设备列表
	AreaRange                      map[string][]string //区域列表
	TimeRange                      *TimeRange          //时间范围
}

type TimeRange struct {
	Offset int64 // 偏移量
	Step   int64 // 步长
}
type AddPeopleTacticRulesResponse struct {
	TacticRuleID string //战法规则id
	*AddPeopleTacticRulesRequest
}

type UpdatePeopleTacticRulesRequest struct {
	*AddPeopleTacticRulesRequest
}

type DeletePeopleTacticRulesReques struct {
	TacticRuleID string
	UserID       string
}

type GetPeopleTacticRulesListRequest struct {
	TacticID string //技术战法id
	UserID   string
}

type PeopleTacticRules struct {
	Timestamp      int64
	TacticID       string
	UserID         string   //user_id
	PersonComment  string   // 人员描述
	TacticRuleID   string   //战法规则id
	TacticRuleName string   //战法规则名称
	AnalysisType   int      //分析类型枚举
	AnalysisItem   []int64  //配置项
	IsOpen         bool     //是否开启状态
	WarningLevel   int      //预警级别
	RepoIDs        []string //比对库列表
	Parameters     *Parameters
}
type GetPeopleTacticRulesListResponse struct {
	Total int
	Rets  []*PeopleTacticRules
}

//为规则引擎提供的接口返回体
type GetPeopleTacticRulesListResponse2 struct {
	Total int
	Rets  map[string]*TacticRuleConfig
}

func (r *GetPeopleTacticRulesListResponse) ListResponseToListResponse2() *GetPeopleTacticRulesListResponse2 {
	m := make(map[string]*TacticRuleConfig)
	for _, ret := range r.Rets {
		if ret.Parameters == nil {
			continue
		}
		if len(ret.Parameters.SensorIDs) != 0 {
			areange := make(map[string][]string)
			for _, sensorid := range ret.Parameters.SensorIDs {
				areange[sensorid] = []string{sensorid}
			}
			ret.Parameters.AreaRange = areange
		}
		c := &TacticRuleConfig{
			ID:          ret.TacticRuleID,
			Timestamp:   ret.Timestamp,
			Valid:       ret.IsOpen,
			RepoRange:   ret.RepoIDs,
			AreaRange:   ret.Parameters.AreaRange,
			AreaDepth:   ret.Parameters.AreaDepth,
			TimeRange:   ret.Parameters.TimeRange,
			AreaBreadth: ret.Parameters.AreaBreadth,
			TotalNumber: ret.Parameters.AppearNumber,
		}
		if ret.Parameters.TimeFrequencyTraceBackDuration != 0 && ret.Parameters.TimeFrequencyNumber != 0 {
			tf := &TimeFrequency{
				TracebackDuration: ret.Parameters.TimeFrequencyTraceBackDuration * 24 * 60 * 60 * 1000,
				Number:            ret.Parameters.TimeFrequencyNumber,
			}
			c.TimeFrequency = tf
		}
		m[c.ID] = c
	}
	return &GetPeopleTacticRulesListResponse2{
		Total: len(m),
		Rets:  m,
	}
}

type TacticRuleConfig struct {
	ID            string              `json:",omitempty"` // 规则标识
	Timestamp     int64               `json:",omitempty"` // 创建时间
	Valid         bool                `json:",omitempty"` // 是否有效
	RepoRange     []string            `json:",omitempty"` // 库范围
	AreaRange     map[string][]string `json:",omitempty"` // 设备范围
	TimeRange     *TimeRange          `json:",omitempty"` // 时间范围
	AreaBreadth   int                 `json:",omitempty"` // 区域个数
	AreaDepth     int                 `json:",omitempty"` // 区域频次
	TotalNumber   int                 `json:",omitempty"` // 总数
	TimeFrequency *TimeFrequency      `json:",omitempty"` // 时间频次
}

type TimeFrequency struct {
	TracebackDuration int64 `json:",omitempty"` // 回溯时长
	Number            int   `json:",omitempty"` // 次数
}
