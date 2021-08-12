package models

type PeopleTactics struct {
	Timestamp     int64
	SceneID       string
	TacticID      string
	TacticName    string
	RepoTagIDs    []int64
	TacticComment string
	PersonComment string
	AnalysisType  int
	AnalysisItem  []int64
	PageStyle     string
}

type AddPeopleTacticsRequest struct {
	SceneID       string  //场景id
	TacticName    string  //技战法名称
	TacticComment string  //技战法描述
	RepoTagIDs    []int64 //库标签
	PersonComment string  //人员描述
	AnalysisType  int     //分析类型枚举
	AnalysisItem  []int64 //配置项
	PageStyle     string  //图标
}

type AddPeopleTacticsResponse struct {
	TacticID      string  //技战法id
	SceneID       string  //场景id
	TacticName    string  //技战法名称
	TacticComment string  //技战法描述
	RepoTagIDs    []int64 //库标签
	PersonComment string  //人员描述
	AnalysisType  int     //分析类型枚举
	AnalysisItem  []int64 //配置项
	PageStyle     string  //图标
}
type GetPeopleTacticsResponse struct {
	Timestamp     int64
	TacticID      string  //技战法id
	SceneID       string  //场景id
	TacticName    string  //技战法名称
	TacticComment string  //技战法描述
	RepoTagIDs    []int64 //库标签
	PersonComment string  //人员描述
	AnalysisType  int     //分析类型枚举
	AnalysisItem  []int64 //配置项
	PageStyle     string  //图标
}

type UpdatePeopleTacticsRequest struct {
	TacticName    string //技战法名称
	TacticComment string //技战法描述
	PersonComment string //人员描述
	PageStyle     string //图标
}
