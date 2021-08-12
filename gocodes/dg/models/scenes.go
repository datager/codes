package models

const (
	CreateTypeDefault  = 1 //默认场景
	CreateTypeByPerson = 2 //用户自建场景

)
const (
	ScenesTypePerson  = 1 //人员场景
	ScenesTypeVechile = 2 //车场景
)

type PersonScene struct {
	Timestamp    int64
	SceneID      string
	SceneName    string
	SceneType    int
	CreateType   int
	SceneComment string
	Tactics      []*PeopleTactics
}

type AddScenesRequest struct {
	SceneName    string //场景名称
	CreateType   int
	SceneComment string //场景描述
	ScenesType   int    //场景类型
}
type UpdateScenesRequest struct {
	SceneName    string //场景名称
	SceneComment string //场景描述
}

type AddScenesResponse struct {
	SceneID      string //场景id
	SceneName    string //场景名称
	SceneComment string //场景描述
	ScenesType   int    //场景类型
}
type ListResponse struct {
	Total int
	Rets  []interface{}
}
type GetPeopleScenesResponse struct {
	SceneID      string //场景id
	SceneName    string //场景名称
	SceneComment string //场景描述
	CreatType    int    //是否为默认场景
	Tactics      []*PeopleTactics
}
