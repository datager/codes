package models

const (
	AnalysisTypeFirstAppear         = 1001 // 首次出现
	AnalysisTypeMultipleAppear      = 1002 //多次出现
	AnalysisTypeMultiLocationAppear = 1003 //多地点出现
	AnalysisTypeFrequencyAppear     = 1004 //频次出现
	AnalysisTypeLeave               = 2001 //人员离开
	AnalysisTypAppearSameTime       = 3001 //同时出现
)
const (
	AnalysisItemSensor            = 1 //设备
	AnalysisItemArea              = 2 //区域
	AnalysisItemTimeInterval      = 3 //时间段
	AnalysisItemAppearNumber      = 4 //出现次数
	AnalysisItemLocationNumber    = 5 //地点数量
	AnalysisItemLocationFrequency = 6 //地点频次
	AnalysisItemTimeFrequency     = 7 //时间频次
)
const (
	AnalysisTypeAppear    = "出现类"
	AnalysisTypeDisAppear = "消失类"
	AnalysisTypeHideout   = "窝点类"
)

var AnalysisTypeMap map[string][]map[int]string

func init() {
	AnalysisTypeMap = make(map[string][]map[int]string)
	appearMaps := make([]map[int]string, 0)
	appearMap1 := make(map[int]string)
	appearMap2 := make(map[int]string)
	appearMap3 := make(map[int]string)
	appearMap4 := make(map[int]string)
	appearMap1[AnalysisTypeFirstAppear] = "首次出现"
	appearMap2[AnalysisTypeMultipleAppear] = "多次出现"
	appearMap3[AnalysisTypeMultiLocationAppear] = "多地点出现"
	appearMap4[AnalysisTypeFrequencyAppear] = "频次出现"
	appearMaps = append(appearMaps, appearMap1, appearMap2, appearMap3, appearMap4)
	AnalysisTypeMap[AnalysisTypeAppear] = appearMaps
	disAppearMaps := make([]map[int]string, 0)
	disAppearMap1 := make(map[int]string)
	disAppearMap1[AnalysisTypeLeave] = "人员离开"
	disAppearMaps = append(disAppearMaps, disAppearMap1)
	AnalysisTypeMap[AnalysisTypeDisAppear] = disAppearMaps
	hideoutMaps := make([]map[int]string, 0)
	hideoutMap1 := make(map[int]string)
	hideoutMap1[AnalysisTypAppearSameTime] = "同时出现"
	hideoutMaps = append(hideoutMaps, hideoutMap1)
	AnalysisTypeMap[AnalysisTypeHideout] = hideoutMaps
}
