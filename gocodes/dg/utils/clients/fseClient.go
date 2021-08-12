package clients

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"sync"
	"time"

	"codes/gocodes/dg/utils/json"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type FseCli struct {
	httpClient *HTTPClient
}

func NewFseClient(addr string, timeout time.Duration) (*FseCli, error) {
	if addr == "" {
		return nil, errors.New("addr is empty")
	}

	cli := &FseCli{
		NewHTTPClient(addr, timeout, nil),
	}

	return cli, nil
}

func (fc *FseCli) AddRepo(repo RepoInfo) error {
	if repo.ID == "" {
		return errors.New("repo id is empty")
	}
	if !IsFeatureType(repo.Type) {
		return errors.New(fmt.Sprintf("repo type is invalid, [repo Type] 【%v】", repo.Type))
	}
	if !IsIndexType(repo.IndexType) {
		return errors.New(fmt.Sprintf("repo index type is invalid, [repo index type] 【%v】", repo.IndexType))
	}
	if !IsRepoLevel(repo.Level) {
		return errors.New(fmt.Sprintf("repo level is invalid, [repo level] 【%v】", repo.Level))
	}

	return fc.addRepo(repo)
}

func (fc *FseCli) DeleteRepo(repoID string) error {
	if repoID == "" {
		return errors.New("repo id is empty")
	}

	return fc.deleteRepo(repoID)
}

func (fc *FseCli) QueryRepo(repoID string) bool {
	if repoID == "" {
		return false
	}

	return fc.queryRepo(repoID)
}

func (fc *FseCli) addRepo(repo RepoInfo) error {
	var URL = "/repositories"
	_, err := fc.httpClient.PostJSON(URL, repo)

	return err
}

func (fc *FseCli) deleteRepo(repoID string) error {
	var URL = path.Join("/repositories", repoID)
	_, err := fc.httpClient.Delete(URL, bytes.NewReader([]byte{}))

	return err
}

func (fc *FseCli) queryRepo(repoID string) bool {
	var URL = path.Join("/repositories", repoID)
	_, err := fc.httpClient.Get(URL)

	return err == nil
}

func (fc *FseCli) AddFeature(repoID string, feature EntityObject) error {
	if repoID == "" {
		return errors.New("repo id is empty")
	}
	if feature.ID == "" {
		return errors.New("feature id is empty")
	}
	if !IsEntityDataType(feature.Data.Type) {
		return errors.New(fmt.Sprintf("feature data type is invalid, [feature data Type] 【%v】", feature.Data.Type))
	}
	if feature.Data.Value == "" {
		return errors.New("feature data value is empty")
	}
	if feature.LocationID == "" {
		return errors.New("location id id empty")
	}

	return fc.addFeature(repoID, feature)
}

func (fc *FseCli) DeleteFeature(repoID string, featureID string) error {
	if repoID == "" {
		return errors.New("repo id is empty")
	}
	if featureID == "" {
		return errors.New("feature id is empty")
	}

	return fc.deleteFeature(repoID, featureID)
}

func (fc *FseCli) RankFeature(query RepoSearchCriteria) ([]SearchResultItem, error) {
	if !IsFeatureType(query.Type) {
		return nil, errors.New(fmt.Sprintf("repo type is invalid, [repo Type] 【%v】", query.Type))
	}
	if len(query.Include) == 0 {
		return nil, errors.New("query include is empty")
	}
	for i := 0; i < len(query.Include); i++ {
		if !IsEntityDataType(query.Include[i].Data.Type) {
			return nil, errors.New(fmt.Sprintf("include feature data type is invalid, [feature data Type] 【%v】", query.Include[i].Data.Type))
		}
		if query.Include[i].Data.Value == "" {
			return nil, errors.New("include feature data value is empty")
		}
	}
	if len(query.Repositories) == 0 {
		return nil, errors.New("query repos is empty")
	}

	return fc.rankFeature(query)
}

func (fc *FseCli) MutliRankFeature(query RepoSearchCriteria) ([]SearchResultItem, error) {
	var g errgroup.Group
	var lock sync.Mutex
	ret := make([]SearchResultItem, 0)
	for index, v := range query.Include {
		v := v
		index := index
		g.Go(func() error {
			param := RepoSearchCriteria{
				Type:             query.Type,
				Include:          []SearchItem{v},
				Exclude:          query.Exclude,
				IncludeThreshold: query.IncludeThreshold,
				ExcludeThreshold: query.ExcludeThreshold,
				Repositories:     query.Repositories,
				Location:         query.Location,
				StartTime:        query.StartTime,
				EndTime:          query.EndTime,
				MaxCandidates:    query.MaxCandidates,
				Options:          query.Options,
			}

			r, err := fc.RankFeature(param)
			if err != nil {
				return errors.WithStack(err)
			}

			lock.Lock()
			for _, v := range r {
				v.Index = index
				ret = append(ret, v)
			}
			lock.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, errors.WithStack(err)
	}

	return ret, nil
}

func (fc *FseCli) addFeature(repoID string, feature EntityObject) error {
	var URL = path.Join("/repositories", repoID, "entities")
	_, err := fc.httpClient.PostJSON(URL, feature)

	return err
}

func (fc *FseCli) deleteFeature(repoID string, featureID string) error {
	var URL = path.Join("/repositories", repoID, "/entities", featureID)
	_, err := fc.httpClient.Delete(URL, bytes.NewReader([]byte{}))

	return err
}

func (fc *FseCli) rankFeature(query RepoSearchCriteria) ([]SearchResultItem, error) {
	var URL = path.Join("/repositories/search")
	respByte, err := fc.httpClient.PostJSON(URL, query)
	if err != nil {
		return nil, err
	}
	var ret RepoSearchResult
	if err = json.Unmarshal(respByte, &ret); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unmarshal resp byte to repo search result err [resp byte] 【%v】", respByte))
	}

	return ret.Results, nil
}

// 特征类型枚举
const (
	FeatureTypeFace     = "face"
	FeatureTypeMotor    = "motor"
	FeatureTypeNonmotor = "nonmotor"
	FeatureTypePerson   = "person"
)

func IsFeatureType(t string) bool {
	return strings.EqualFold(t, FeatureTypeFace) ||
		strings.EqualFold(t, FeatureTypeMotor) ||
		strings.EqualFold(t, FeatureTypeNonmotor) ||
		strings.EqualFold(t, FeatureTypePerson)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////// 库CRUD操作的model ///////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 检索类型枚举
const (
	IndexTypeInt8  = "int8"
	IndexTypeFloat = "float"
)

func IsIndexType(t string) bool {
	return strings.EqualFold(t, IndexTypeInt8) || strings.EqualFold(t, IndexTypeFloat)
}

// 库级别枚举
const (
	RepoLevelRAM  = "ram"
	RepoLevelGPU  = "gpu"
	RepoLevelDisk = "disk"
)

func IsRepoLevel(l string) bool {
	return strings.EqualFold(l, RepoLevelRAM) || strings.EqualFold(l, RepoLevelGPU) || strings.EqualFold(l, RepoLevelDisk)
}

type RepoVersion struct {
	Name          string `json:"name"`           // 特征版本名称
	FeatureLength int32  `json:"feature_length"` // 特征长度
	Size          int64  `json:"size"`           // 当前版本特征数量，仅在查询时作为返回值使用，其它请求中会被忽略
}

type RepoInfo struct {
	ID             string      `json:"id"`              // 比对库ID
	Type           string      `json:"type"`            // 比对库存储的特征类型，即人脸，机动车，非机动车和人体
	IndexType      string      `json:"index_type"`      // 检索类型，用来指定以何种类型组织特征数据。建议使用int8，int8占用的内存只有float的1/4并且比对精度几乎不会损失
	Level          string      `json:"level"`           // 库级别。disk模式表示数据存储在磁盘上，在计算过程中加载到内存上，支持最近n天数据在内存上缓存（暂不建议使用）；ram和gpu模式分别表示数据存储在内存或者GPU显存上，在计算过程中直接使用，不需要发生IO操作
	Replications   int         `json:"replications"`    // default: 1 此参数指明为当前库生成多少个副本，有副本的库比对性能会得到提升，安全性也会增强。此参数默认为1，即不生成副本，如若需要一个副本可将其改为2。若要修改此参数，只能向上修改，不能向下修改，即可以将1改为2或3，不能将2改为1
	Size           int64       `json:"size"`            // 库当前大小，仅在查询时作为返回值使用，其它请求中会被忽略。
	Capacity       int64       `json:"capacity"`        // 容量
	DefaultVersion RepoVersion `json:"default_version"` // 当前库默认的特征版本，支持多版本特征后，添加特征数据时若不指定版本号，将按默认特征版本进行入库。（目前暂不支持多版本特征，创建比对库时不需要指定此参数）
	// 一个string类型的键值对集合，用来存储一些高级参数，通常情况下这些参数可以为默认值。下面是目前支持的高级参数：
	// hashingNum，值为字符串形式的正整数，表示集群FSE中为当前库进行raft分组时分组的个数，通常情况下集群会根据当前库的配置自动设置该参数。
	// UseFeatureIDMap，值为字符串形式的bool值，指定是否使用特征ID映射表，即是否将映射表存储在内存上。启用后可避免比对和入库时的一些读库操作，但是会造成额外的内存占用（1亿特征约24G）, 默认为"false"。要求黑名单库设置这个参数为"true"，否则会影响比对QPS。针对抓拍库，当特征库保存在SSD时，建议设置该参数为"false"，可以减少内存占用，并且对性能影响很小；当特征库保存在Sata盘时，建议设置该参数为"true"，否则对性能影响较大。
	// TimerDeleteDays, 值为字符串形式的整数，根据距离今天的天数删除旧特征，默认为"0", 表示不删除。比对库创建完成后，可以随时修改该值。如需关闭该功能，改为"0"即可。
	// PreFilter，值为字符串形式的布尔值，默认为"true"。表示是否开启前过滤，即根据时间和地点来进行计算的筛选，开启后1亿特征大概多占用2G内存
	Options map[string]string `json:"options"`
}

type Repositories struct {
	TotalCount int64      `json:"total_count"` // 数据集总规模
	Repos      []RepoInfo `json:"repos"`       // 比对库列表
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////// 特征CRUD操作以及对比操作的model ////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 比对库中的实例中的数据类型枚举
const (
	EntityDataTypeURI       = "uri"
	EntityDataTypeFile      = "file"
	EntityDataTypeFeature   = "feature"
	EntityDataTypeFeatureID = "feature_id"
)

func IsEntityDataType(t string) bool {
	return strings.EqualFold(t, EntityDataTypeURI) || strings.EqualFold(t, EntityDataTypeFile) ||
		strings.EqualFold(t, EntityDataTypeFeature) || strings.EqualFold(t, EntityDataTypeFeatureID)
}

// 比对库中的实例数据
type EntityData struct {
	RepoID  string `json:"repo_id"` // 仅当type为"feature_id"时生效，表示该特征ID从库ID为repo_id的比对库中获取，value表示此库中的特征ID
	Value   string `json:"value"`   // 用户传入的数据，可以是base64编码的特征、特征ID、图片的uri或二进制数据，根据type确定类型。当type为"feature_id"时，表示特征ID，非base64编码
	Type    string `json:"type"`    // 描述value字段传入数据的类型，"uri"表示传入图片的uri地址，"file"表示传入图片的二进制数据，"feature"表示传入图片特征(base64编码)，"feature_id"表示传入指定库ID和指定特征ID。目前search接口仅支持"feature"和"feature_id"，当值为"feature_id"时，需要同时指定repo_id字段。compare接口仅支持"feature"
	Version string `json:"version"` // 特征版本名称，仅在compare接口中进行分数映射时使用。调用compare接口时所有特征的"version"字段必须相同，否则会报错。version可填写为分数版本（如"1.8.1.0"）或为空字符串，表示不进行分数映射
}

// 比对库中的实例
type EntityObject struct {
	ID         string     `json:"id"`          // 实例ID
	Data       EntityData `json:"data"`        // 实例数据
	LocationID string     `json:"location_id"` // 地点ID
	Time       int64      `json:"time"`        // 时间，以毫秒记, default: 0
}

type SearchItem struct {
	Data       EntityData `json:"data"`       // 搜索的数据
	Confidence float64    `json:"confidence"` // default: 1 表示当前特征的质量分，普通用户无需填写。比对多张图片时，将按照每张图片对应特征的质量分对各图片的检索结果进行合并（加权平均）；若不指定该项，则每张图片具有相同的权重
}

type RepoSearchCriteria struct {
	Type             string       `json:"type"`              // 进行比对的特征类型，分别表示人脸、机动车、非机动车和行人
	Include          []SearchItem `json:"include"`           // 检索目标图片集合
	Exclude          []SearchItem `json:"exclude"`           // 排除目标图片集合
	IncludeThreshold float64      `json:"include_threshold"` // default: 0 检索目标阈值，相似度小于该值的结果不返回；值在0~1之间，默认为0表示不过滤
	ExcludeThreshold float64      `json:"exclude_threshold"` // default: 0 排除目标阈值；值在0~1之间，默认为0表示不过滤
	Repositories     []string     `json:"repositories"`      // 待比对的特征库ID，可以指定多个repo ID，每个repo的特征长度应该与待比对的特征的长度相同。以图搜图引擎对每个repo分别进行比对，按照分数合并结果。
	Location         []string     `json:"locations"`         // 地点ID数组。当比对库为非GPU模式时，比对结果会根据该数组进行过滤，若该数组为空，则不对地点进行过滤。
	StartTime        int64        `json:"start_time"`        // default: 0 与end_time配合使用，指定搜索的时间范围起点，使用unix timestamp毫秒，仅当比对库是非GPU模式时生效。
	EndTime          int64        `json:"end_time"`          // default: 0 与start_time配合使用，指定搜索的时间范围终点，使用unix timestamp毫秒，仅当比对库是非GPU模式时生效。若为0，则不限制时间终点，搜索从时间起点开始的全部特征
	MaxCandidates    int32        `json:"max_candidates"`    // default: 3 返回top K，默认为3
	//一个string类型的键值对集合，用来指定一些扩展参数，普通用户不需要使用。目前支持的扩展参数有：
	//normalization，是否对返回的相似度作分数映射，默认对人脸作分数映射("true")，对其它类型特征不作分数映射("false")。
	//return_feature，在比对结果中是否返回特征值，默认为"false"。
	Options map[string]string `json:"options"`
}

type SearchResultItem struct {
	ID         string     `json:"id"`          // 特征ID
	Similarity float64    `json:"similarity"`  // 当前特征与比对特征的相似度，介于0与1之间的小数
	RepoID     string     `json:"repo_id"`     // 特征所属库的ID
	Data       EntityData `json:"data"`        // 数据
	LocationID string     `json:"location_id"` // 地点ID
	Time       int64      `json:"time"`        // 时间，以毫秒记
	Index      int        `json:"-"`
}

type RepoSearchResult struct {
	Results []SearchResultItem `json:"results"` // 查询到的结果集合
	Message string             `json:"message"` // 分布式中为了保证部分节点掉线之后还能返回结果，增加此字段返回那些没有检索的节点
}

type CompareItem struct {
	ID   string     `json:"id"`   // 特征ID
	Data EntityData `json:"data"` // 数据
}

type CompareCriteria struct {
	Type      string        `json:"type"`      // 进行比对的特征类型
	MObjects  []CompareItem `json:"m_objects"` // 特征集合一
	NObjects  []CompareItem `json:"n_objects"` // 特征集合二
	Threshold float64       `json:"threshold"` // 比对目标阈值，相似度小于该值不返回；值在0~1之间，默认为0表示不过滤
}

type CompareResultItem struct {
	MID        string  `json:"m_id"`       // 特征集合一中特征的ID
	NID        string  `json:"n_id"`       // 特征集合二中特征的ID
	Similarity float64 `json:"similarity"` // 特征相似度分数，介于0与1之间
}

type CompareResult struct {
	Results []CompareResultItem `json:"results"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////// 特征检测操作的model ///////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DetectModeAuto            = "auto"             // 自动模式（暂不支持)
	DetectModeCapturedFace    = "captured_face"    // 证件照或者人脸抓拍图中人脸检测
	DetectModeCapturedVehicle = "captured_vehicle" // 机动车抓拍图中乘客人脸检测
	DetectModeCapturedPerson  = "captured_person"  // 行人抓拍图中人脸检测
	DetectModeScene           = "scene"            // 场景大图中人脸检测 (默认值)
)

type Rectangle struct {
	X int64 `json:"x"` // 对象左上角的x座标
	Y int64 `json:"y"` // 对象左上角的y座标
	W int64 `json:"w"` // 对象宽度
	H int64 `json:"h"` // 对象高度
}

type BoundingBox struct {
	Type string    `json:"type"` // 目标类型，即人脸、机动车、非机动车和人体
	Rect Rectangle `json:"rect"` // 目标检测框
}

type Pos struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ROIPolygon struct {
	PolygonArea     []Pos   `json:"polygon_area"`     // 不规则多边形感兴趣。通过一个(类似)二维数组来表示，列长度为2，比如 [[0.1, 0.2], [0.2, 0.3], [0.3, 0.1]]被用来表示一个三角形区域
	FilterThreshold float64 `json:"filter_threshold"` // 感兴趣或屏蔽区域阈值，正数为感兴趣区域，负数为屏蔽区域，范围:[-1.00, 1.00]，目标在polygon_area多边形中的面积占目标整体面积的比例小于此阀值时，目标将被过滤掉（filter_threshold > 0 时）或者输出（filter_threshold < 0 时）。缺省值0.9
}

type Image struct {
	File          string        `json:"file"`           // 图片base64；file和url至少一个不为空
	URL           string        `json:"url"`            // 图片url
	BoundingBoxes []BoundingBox `json:"bounding_boxes"` // 多个目标检测框；当指定此参数时，跳过目标检测(暂不支持)
	ROIsPolygon   []ROIPolygon  `json:"rois_polygon"`   // 多个感兴趣区域
}

// 一个感兴趣目标，目前仅支持此目标是输入场景图image一部分的图片
type InterestedObj struct {
	File        string    `json:"file"`         // 暂不支持；图片base64；file、url和bounding_box至少一个不为空
	URL         string    `json:"url"`          // 暂不支持；图片URL；file、url和bounding_box至少一个不为空
	BoundingBox Rectangle `json:"bounding_box"` // 感兴趣目标在原图中的位置座标；file、url和bounding_box至少一个不为空
}

type InterestedObject struct {
	Type             string        `json:"type"`              // 当前目标类型；目前只支持face！
	InterestedObject InterestedObj `json:"interested_object"` // 一个感兴趣目标，目前仅支持此目标是输入场景图image一部分的图片
}

type DetectCriteria struct {
	// 仅在输入参数detect_type是face或者all时生效；因为算法局限,不同场景的图片,需要使用不同的算法进行检测，才能达到比较大的QPS和准确度。用户可以自己选择不同的算法。目前引擎提供如下模式：
	// auto：自动模式（暂不支持）
	// captured_face: 证件照或者人脸抓拍图中人脸检测
	// captured_vehicle: 机动车抓拍图中乘客人脸检测
	// captured_person: 行人抓拍图中人脸检测
	// scene: 场景大图中人脸检测 (默认值)
	DetectMode        string             `json:"detect_mode"`
	Image             Image              `json:"image"`              // 一张图片
	InterestedObjects []InterestedObject `json:"interested_objects"` // 多个感兴趣目标，当指定此参数时，只返回在输入图片image中与感兴趣目标interested_objects关联的全目标结构化信息
	Extra             map[string]string  `json:"extra"`              // 是个字典，透传的客户数据，请求处理后的结果会携带此数据，格灵深瞳不做任何处理。 示例：{"additionalProp1": "user_data1", "additionalProp2": "user_data2"}
}

const (
	ObjectTypeVehicle       = "vehicle"
	ObjectTypeNonVehicle    = "nonvehicle"
	ObjectTypePedestrian    = "pedestrian"
	ObjectTypeFace          = "face"
	ObjectTypeLicensePlate  = "licenseplate"
	ObjectTypePassenger     = "passenger"
	ObjectTypeMisc          = "misc"
	ObjectTypeVehicleSymbol = "vehiclesymbal" // 此处为api文档中提供的枚举（有错别字symbal）
)

type ObjBoundingBox struct {
	X     int `json:"x"`      // 对象左上角的x座标
	Y     int `json:"y"`      // 对象左上角的y座标
	W     int `json:"w"`      // 对象宽度
	H     int `json:"h"`      // 对象高度
	OrigW int `json:"orig_w"` // 原图宽度
	OrigH int `json:"orig_h"` // 原图高度
}

const (
	SpeedFast    = "fast"
	SpeedMedium  = "medium"
	SpeedSlow    = "slow"
	SpeedUnknown = "unknown"
)

const (
	DirectionUp        = "up"
	DirectionDown      = "down"
	DirectionLeft      = "left"
	DirectionRight     = "right"
	DirectionLeftUp    = "left-up"
	DirectionLeftDown  = "left-down"
	DirectionRightUp   = "right-up"
	DirectionRightDown = "right-down"
)

type ObjSnapShot struct {
	Timestamp   int64          `json:"timestamp"`    // 对象所在的帧的时间戳 符合ISO8601格式要求的UTC时间，如2018-05-25T03:27:18Z
	BoundingBox ObjBoundingBox `json:"bounding_box"` // 对象在原图中的位置座标
}

type ObjMove struct {
	Speed      string        `json:"speed"`      // 对象运动速度，取值快 中 慢 未知
	Direction  string        `json:"direction"`  // 对象在画面中的相对运动方向
	Trajectory []ObjSnapShot `json:"trajectory"` // 数组，对象的运动轨迹
}

type ObjAttribute struct {
	AttributeID   int     `json:"attribute_id"`   // 该属性在字典中的id
	AttributeName string  `json:"attribute_name"` // 该属性的中文显示名
	ValueID       int     `json:"value_id"`       // 属性在字典中的id
	ValueName     string  `json:"value_name"`     // 属性的值
	Confidence    float64 `json:"confidence"`     // 置信度
}

type Obj struct {
	Type     string `json:"type"`      // 类型
	ObjectID string `json:"object_id"` // 结构化对象ID
}

// 结构化对象
type Object struct {
	ID            string            `json:"id"`              // 结构化对象ID
	Time          int64             `json:"time"`            // 对象所在的帧的时间戳 符合ISO8601格式要求的UTC时间，如2018-05-25T03:27:18Z
	FrameID       string            `json:"frame_id"`        // 对象所在帧号
	Type          string            `json:"type"`            // 可能值为： vehicle: 机动车 nonvehicle: 非机动车 pedestrian: 行人/人体 face: 人脸 licenseplate: 车牌 passenger: 乘客 misc: 其他
	Subtype       string            `json:"subtype"`         // 对象的子类型，如非机动车中可能有两轮车(bicycle)，三轮车(tricycle)等
	SceneImageURL string            `json:"scene_image_url"` // 对象所在图片的存储访问地址
	ObjImageURL   string            `json:"obj_image_url"`   // 对象本身的存储访问地址
	File          string            `json:"file"`            // 图片二进制
	Confidence    float64           `json:"confidence"`      // 置信度
	BoundingBox   ObjBoundingBox    `json:"bounding_box"`    // 对象在原图中的位置座标
	Movement      ObjMove           `json:"movement"`        // 视频中对象的运动轨迹，只在视频流分析任务中输出
	Feature       string            `json:"feature"`         // 如果对象可以提取特征，这里是特征的base64字符串
	Attributes    []ObjAttribute    `json:"attributes"`      // 对象的结构化属性
	RelatedTo     []Obj             `json:"related_to"`      // 数组，表示一组关系
	Has           []string          `json:"has"`             // 如果该对象包含其他被检测出的对象，这里会包括第一级的children对象的id
	Level         int               `json:"level"`           // 相应Object的层级
	BelongsTo     string            `json:"belongs_to"`      // 如果该对象属于某个其他对象，在这里会表示出来
	Extra         map[string]string `json:"extra"`           // 是个字典，透传的客户数据，请求处理后的结果会携带此数据，格灵深瞳不做任何处理。
}

type Objects struct {
	TotalCount int64    `json:"total_count"` // 数据集总规模
	Data       []Object `json:"data"`        // 对象列表
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////// 违法行为识别的model ///////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	ViolationType1208 = "1208"
	ViolationType1301 = "1301"
	ViolationType1345 = "1345"
	ViolationType1357 = "1357"
	ViolationType1625 = "1625"
	ViolationTypeDDCS = "ddcs"
	ViolationTypeQJCS = "qjcs"
)

type Img struct {
	File string `json:"file"` // 图片二进制；file和url两者至少一个不为空
	URL  string `json:"url"`  // 图片URL；file和url两者至少一个不为空
}

type IncidentDetectCriteria struct {
	// 违法行为类型：
	// 1208: 不按导向车道行驶
	// 1301: 逆向行驶
	// 1345: 违反禁止标线指示
	// 1357: 未礼让行人
	// 1625: 违法交通信号
	// ddcs: 单点测速
	// qjcs: 区间测速
	ViolationType string            `json:"violation_type"`
	PlateNumber   string            `json:"plate_number"` // 车牌号码
	Images        []Img             `json:"images"`       // 识别的图像数据，目前只支持一张图片
	Extra         map[string]string `json:"extra"`        // 是个字典，透传的客户数据，请求处理后的结果会携带此数据，格灵深瞳不做任何处理
}

type ViolationObject struct {
	IsValid               bool              `json:"is_valid"`                // default: false 识别结果是否有效：true(违法可删除或者违法可找回)；false(无效)
	Reason                string            `json:"reason"`                  // 仅当is_valid为true时有效；违法判断依据和原因，例如E00001表示车牌未识别：无一张图片能检测到号牌
	PlateNumberMatched    int               `json:"plate_number_matched"`    // default: 0, 0-未识别；1-一致；2-部分号牌不匹配（两位及以下）；3-未追踪到车辆；
	PlateType             int               `json:"plate_type"`              // 识别出的车牌类型（例如普通蓝牌等，参考格灵深瞳车牌属性字典）
	RecognizedPlateNumber string            `json:"recognized_plate_number"` // 识别出的车牌号码
	VehicleType           int               `json:"vehicle_type"`            // 目标车辆类型（例如面包车等，参考格灵深瞳机动车算法属性字典中车型部分）
	Direction             int               `json:"direction"`               // default: 0, 行车方向：0-未识别；1-直行；2-左转；3-右转；4-掉头
	Angle                 int               `json:"angle"`                   // default: 0, 车辆角度：0-未识别；1-车头；2-车尾
	IsComplete            int               `json:"is_complete"`             // default: 0, 目标是否完整：0-未识别；1-完整；2-不完整但号牌悬挂处完整；3-不完整且号牌悬挂处不完整
	Position              int               `json:"position"`                // default: 0, 车辆位置：0-未识别；1-中心分隔线左侧；2-中心分隔线右侧
	TrafficLight          string            `json:"traffic_light"`           // 行车方向信号灯：1R2Y3B，表示行车方向信号灯第一张为红灯、第二张为黄灯、第三张未识别
	HasPoliceman          int               `json:"has_policeman"`           // default: 0, 是否存在交警：0-未识别；1-存在；2-不存在
	FastenSeatBelt        int               `json:"fasten_seatbelt"`         // default: 0, 是否不系安全带：0-未识别；1-已系安全带；2-未系安全带
	CallPhone             int               `json:"call_phone"`              // default: 0, 是否打手机：0-未识别；1-打手机；2-未打手机
	OverLine              int               `json:"over_line"`               // default: 0, 行车过程是否压线：0-未识别；1-压线；2-未压线
	UnderConstruction     int               `json:"under_construction"`      // default: 0, 是否存在施工占道：0-未识别；1-存在；2-不存在
	Extra                 map[string]string `json:"extra"`                   // 是个字典，透传的客户数据，请求处理后的结果会携带此数据，格灵深瞳不做任何处理
}

type ResponseError struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error"`
}

var _ error = &ResponseError{}

func (e ResponseError) Error() string {
	return "fse response 【error】, error code 【" + e.ErrorCode + "】, error message 【" + e.ErrorMessage + "】"
}
