package clients

const (
	// -未知厂商
	SAX_MANUFACTURER_UNKNOWN = 0
	// -深网
	SAX_MANUFACTURER_SENSENETS = 1
	// -格灵深瞳
	SAX_MANUFACTURER_DEEPGLINT = 2
	// -商汤
	SAX_MANUFACTURER_SENSETIME = 3
	// -旷世
	SAX_MANUFACTURER_FACEPLUSE = 4
	// -依图
	SAX_MANUFACTURER_YITU = 5

	SaxResponseContextStatusOK = "200"
)

type SaxQuery struct {
	Feature    string
	Confidence float32
	TopX       int
}

type SaxCandidate struct {
	Vid              string
	SimilarityDegree float32
}

type SaxRequestContext struct {
	SessionId string
	Params    map[string]string
}

type SaxSearchRequest struct {
	Context          *SaxRequestContext
	FaceImageBindata string
	FaceImageUrl     string
	Feature          string
	Manufacturer     int
}

type SaxResponseContext struct {
	SessionId  string
	Status     string
	Message    string
	RequestTs  int64
	ResponseTs int64
}

type SaxSearchResponse struct {
	Context    *SaxResponseContext
	Candidates []*SaxCandidate
}

// -Vid合并条目
type SaxVidCombine struct {
	SourceVid string
	TargetVid string
}

// -Vid变化条目检索接口请求
type SaxSearchVidCombinesRequest struct {
	// -请求上下文
	Context *SaxRequestContext
	// -开始时间戳(秒)
	StartTimestamp int64
	// -结束时间戳(秒）
	EndTimestamp int64
	// -分页limit
	Limit int64
	// -分页offset
	Offset int64
}

// -Vid变化条目检索接口返回
type SaxSearchVidCombinesResponse struct {
	// -返回上下文
	Context *SaxResponseContext
	// -相似返回候选集合
	VidCombines []*SaxVidCombine
}
