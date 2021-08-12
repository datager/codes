package models

type FaceRepo struct {
	RepoID        string `json:"Id"`
	RepoName      string `json:"Name"`
	IssuedRepo    *IssuedRepo
	IssuedTo      []*IssuedRepo
	OrgId         string
	OrgName       string
	Timestamp     int64
	Comment       string
	PicCount      int64
	ErrorPicCount int64
	NameListAttr  int
	Type          int64
	Status        TaskStatus
	SystemTag     int
}

type SearchFaceReposRequest struct {
	NameListAttr int      `form:"NameListAttr"` //黑白名单
	Type         int64    `form:"Type"`         //比对库类型：全部，人脸，车辆，未知
	SystemTag    int      `form:"SystemTag"`    //库类型：系统库和用户库
	OrgID        string   `form:"OrgID"`        // nolint: golint
	OrgIDs       []string `form:"-" json:"-"`   //组织ID列表，不从前端对接
	RepoIDs      []string `form:"-" json:"-"`   //比对库ID列表，不从前端对接
	RepoName     string   `form:"RepoName"`     //比对库名称
	Offset       int      `form:"Offset"`       //分页起始
	Limit        int      `form:"Limit"`        //分页数量
	OrderBy      string   `form:"OrderBy"`      //排序规则
}

type FaceReposAndCount struct {
	Total int64
	Rets  []*FaceRepo
}

type IssuedRepo struct {
	Ts     int64
	Id     string
	RepoId string
	OrgId  string
	UserId string
	Type   int64
}

type RepoInfo struct {
	RepoID       string
	RepoName     string
	NameListAttr int
	Type         int
}

func IssuedRepoToIssuedRepoIDs(issuedRepos []*IssuedRepo) []string {
	issuedRepoIDs := make([]string, 0)
	for _, r := range issuedRepos {
		issuedRepoIDs = append(issuedRepoIDs, r.RepoId)
	}
	return issuedRepoIDs
}

// 用于: 下发库 + 比对库 检索接口
type RepoResult struct {
	IssuedRepos []*IssuedRepoSearchResult // 下发库: 来自issued_repo表
	FaceRepos   []*FaceRepo               // 比对库: 来自face_repo表
}

type IssuedRepoSearchResult struct {
	FaceRepo      *FaceRepo // 其中的OrgID意为: 为被下发的库, 的来源组织(从哪个组织下发) (从 face_repos 表得到)
	OrgIDIssuedTo string    // 此字段意为: 被下发的库, 的目的组织(下发到哪个组织) (从 issued_repo 表得到)
}

type IssuedRepoRequest struct {
	RepoID string
	UserID string
	OrgIDs []string
}

type SetFaceRepoRequest struct {
	RepoID string `json:"Id"`
	// 组织id
	OrgID string `json:"OrgId"`
	// 库名称
	RepoName string `json:"Name"`
	// 备注信息
	Comment string
	// 1 黑名单，2 白名单
	NameListAttr int
	// 1 系统默认库，2 用户自建的库
	SystemTag int
	// 1 车辆库 ，2 人脸库
	Type int64
}

type AddToRepoRequest struct {
	Attr   *CivilAttr
	Imgs   []*ImageQuery
	RepoID string
}

const (
	SystemDefault   = iota + 1 // 系统默认库
	SystemUserBuild            // 用户建立的库
)

const (
	_                             = iota
	FaceRepoNameListAttrBlackList //黑名单
	FaceRepoNameListAttrWhiteList //白名单
)

type FaceRepoOrder []*FaceRepo

func (f FaceRepoOrder) Len() int {
	return len(f)
}

func (f FaceRepoOrder) Less(i, j int) bool {
	return f[i].Timestamp > f[j].Timestamp
}

func (f FaceRepoOrder) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (r *FaceRepo) ToRepoInRule() *RepoInRule {
	rir := &RepoInRule{
		//RuleId:
		RepoId:   r.RepoID,
		RepoName: r.RepoName,
		//Confidence: r
		//AlarmLevel int     // 报警等级
		//AlarmVoice string  // 报警声音
		//CountLimit int
		Type:   r.Type,
		Status: r.Status,
	}
	return rir
}
