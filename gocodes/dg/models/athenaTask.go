package models

import (
	"fmt"
	"math"
	"strings"
	"time"

	"codes/gocodes/dg/utils/config"

	"github.com/json-iterator/go"
)

const (
	AthenaTaskTypeUnknown = iota // 0
	AthenaTaskTypeExport
	AthenaTaskTypePlatformSyncGB
	AthenaTaskTypeVehicleFrequencyAnalysis
	AthenaTaskTypePersonFrequencyAnalysis // 人员频次分析
	AthenaTaskTypeExportIncludeScore
	AthenaTaskTypePlatformSyncPVG
	AthenaTaskTypeVehicleAreaCollision
	AthenaTaskTypePersonAreaCollision
	AthenaTaskTypeVehicleDayInNightOut
	AthenaTaskTypeVehicleFirstAppearance
	AthenaTaskTypeVehicleFakePlate
	AthenaTaskTypeVehicleHidden
	AthenaTaskTypeVehiclePathAnalyze
	AthenaTaskTypePersonPathAnalyze

	AthenaTaskTypeFaceCaptureExport
	AthenaTaskTypeVehicleCaptureExport
	AthenaTaskTypeNonmotorCaptureExport
	AthenaTaskTypePedestrianCaptureExport

	AthenaTaskTypeFaceCaptureFeatureExport
	AthenaTaskTypeVehicleCaptureFeatureExport
	AthenaTaskTypeNonmotorCaptureFeatureExport
	AthenaTaskTypePedestrianCaptureFeatureExport

	AthenaTaskTypeDuck = -1
)

const (
	AthenaTaskInnerTypeUnknown = iota
	AthenaTaskInnerTypeFaceCapture
	AthenaTaskInnerTypeVehicleCapture
	AthenaTaskInnerTypeNonmotorCapture
	AthenaTaskInnerTypePedestrianCapture
	AthenaTaskInnerTypeFaceEvent
	AthenaTaskInnerTypeCivils

	AthenaTaskInnerTypeFeatureCapture = 100
)

const (
	AthenaTaskTypeExportIncludeScoreUniqueID = "唯一标识符"
)

const (
	AthenaTaskStateTypeEnumRunning  = "Running"  // 运行中
	AthenaTaskStateTypeEnumError    = "Error "   // 错误
	AthenaTaskStateTypeEnumFinish   = "Finish"   // 完成
	AthenaTaskStateTypeEnumRollback = "Rollback" // 回滚
	AthenaTaskStateTypeEnumQueue    = "Queue"    // 排队
	AthenaTaskStateTypeEnumPause    = "Pause"    // 暂停
)

var AthenaTaskStateTypeMap = map[int]map[string]int{
	AthenaTaskTypeExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypePlatformSyncGB: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
	},
	AthenaTaskTypeVehicleFrequencyAnalysis: {
		AthenaTaskStateTypeEnumRunning:  0,
		AthenaTaskStateTypeEnumError:    1,
		AthenaTaskStateTypeEnumFinish:   2,
		AthenaTaskStateTypeEnumRollback: 3,
		AthenaTaskStateTypeEnumQueue:    4,
		AthenaTaskStateTypeEnumPause:    5,
	},
	AthenaTaskTypeVehicleAreaCollision: {
		AthenaTaskStateTypeEnumRunning:  0,
		AthenaTaskStateTypeEnumError:    1,
		AthenaTaskStateTypeEnumFinish:   2,
		AthenaTaskStateTypeEnumRollback: 3,
		AthenaTaskStateTypeEnumQueue:    4,
		AthenaTaskStateTypeEnumPause:    5,
	},
	AthenaTaskTypeVehicleFirstAppearance: {
		AthenaTaskStateTypeEnumRunning:  0,
		AthenaTaskStateTypeEnumError:    1,
		AthenaTaskStateTypeEnumFinish:   2,
		AthenaTaskStateTypeEnumRollback: 3,
		AthenaTaskStateTypeEnumQueue:    4,
		AthenaTaskStateTypeEnumPause:    5,
	},
	AthenaTaskTypeVehicleFakePlate: {
		AthenaTaskStateTypeEnumRunning:  0,
		AthenaTaskStateTypeEnumError:    1,
		AthenaTaskStateTypeEnumFinish:   2,
		AthenaTaskStateTypeEnumRollback: 3,
		AthenaTaskStateTypeEnumQueue:    4,
		AthenaTaskStateTypeEnumPause:    5,
	},
	AthenaTaskTypeVehicleHidden: {
		AthenaTaskStateTypeEnumRunning:  0,
		AthenaTaskStateTypeEnumError:    1,
		AthenaTaskStateTypeEnumFinish:   2,
		AthenaTaskStateTypeEnumRollback: 3,
		AthenaTaskStateTypeEnumQueue:    4,
		AthenaTaskStateTypeEnumPause:    5,
	},
	AthenaTaskTypeVehiclePathAnalyze: {
		AthenaTaskStateTypeEnumRunning:  0,
		AthenaTaskStateTypeEnumError:    1,
		AthenaTaskStateTypeEnumFinish:   2,
		AthenaTaskStateTypeEnumRollback: 3,
		AthenaTaskStateTypeEnumQueue:    4,
		AthenaTaskStateTypeEnumPause:    5,
	},
	AthenaTaskTypeFaceCaptureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypeVehicleCaptureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypeNonmotorCaptureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypePedestrianCaptureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypeFaceCaptureFeatureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypeVehicleCaptureFeatureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypeNonmotorCaptureFeatureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
	AthenaTaskTypePedestrianCaptureFeatureExport: {
		AthenaTaskStateTypeEnumRunning: 0,
		AthenaTaskStateTypeEnumError:   1,
		AthenaTaskStateTypeEnumFinish:  2,
		AthenaTaskStateTypeEnumQueue:   4,
	},
}

func GetAthenaStatusEnumByTaskType(athenaTaskType int) []int {
	var result []int
	for _, enum := range AthenaTaskStateTypeMap[athenaTaskType] {
		result = append(result, enum)
	}
	return result
}

type AthenaTask struct {
	Uts       time.Time
	Ts        int64
	TaskID    string
	Name      string
	UserID    string
	OrgID     string
	Type      int
	InnerType int
	StartTS   int64
	EndTS     int64
}

type AthenaTaskSlice []AthenaTask

func (ats AthenaTaskSlice) GetTaskIDs() []string {
	result := make([]string, 0)
	for _, athenaTask := range ats {
		result = append(result, athenaTask.TaskID)
	}
	return result
}

func AthenaTaskModelsToIDs(models []*AthenaTask) []string {
	result := make([]string, 0)
	for _, athenaTask := range models {
		result = append(result, athenaTask.TaskID)
	}
	return result
}

type AthenaTaskQuery struct {
	Pagination
	UserID string
	OrgID  string
	Type   uint
	Status int // TODO 暂时未使用
}

// 平台同步
type PlatformSyncConfig struct {
	IP       string `json:"Ip"`
	Port     int
	Username string
	Password string
}

func (psc *PlatformSyncConfig) ToString() string {
	result, _ := jsoniter.MarshalToString(&psc)
	return result
}

func (psc *PlatformSyncConfig) FromString(str string) {
	_ = jsoniter.UnmarshalFromString(str, &psc)
}

func (psc *PlatformSyncConfig) ToGBURL() string {
	return fmt.Sprintf(config.GetConfig().GetString("services.athena.platform_sync.gb.url"), psc.IP, psc.Port)
}

func (psc *PlatformSyncConfig) ToPVGURL() string {
	return fmt.Sprintf(config.GetConfig().GetString("services.athena.platform_sync.pvg.url"), psc.Username, psc.Password, psc.IP, psc.Port)
}

type PlatformSyncQuery struct {
	Name         string
	IsDistribute int
	Pagination
}

func (platformSyncGBQuery *PlatformSyncQuery) Normalized() {
	platformSyncGBQuery.Name = strings.TrimSpace(platformSyncGBQuery.Name)
}

func (platformSyncGBQuery *PlatformSyncQuery) CheckFuzzyMatching() bool {
	return !strings.EqualFold(platformSyncGBQuery.Name, "")
}

type PlatformSyncGB struct {
	Uts          time.Time
	Ts           int64
	DeviceID     string
	DeviceName   string
	SensorID     string
	Online       int
	ParentID     string
	Name         string
	Longitude    float64
	Latitude     float64
	Manufacturer string
	Model        string
	Owner        string
	CivilCode    string
	Address      string
	IPAddress    string
	Port         int
	Parental     int
	Secrecy      int
	IsDistribute int
}

func (pst *PlatformSyncGB) IsDistributed() bool {
	return pst.IsDistribute != PlatformSyncNoDistribute
}

func (pst *PlatformSyncGB) ToPlatformSyncSensor() *PlatformSyncDevice {
	return &PlatformSyncDevice{
		DeviceID:   pst.DeviceID,
		DeviceName: pst.DeviceName,
		OrgID:      pst.CivilCode,
		Online:     pst.Online,
	}
}

type PlatformSyncPVG struct {
	Uts          time.Time
	Ts           int64
	DeviceID     string
	DeviceName   string
	SensorID     string
	URL          string
	Name         string
	AVType       int64
	Host         string
	HostTitle    string
	HostType     int64
	IsCurNode    int64
	Level        int64
	Org          string
	OrgPath      string
	Path         string
	Title        string
	Longitude    float64
	Latitude     float64
	IsDistribute int
}

func (pst *PlatformSyncPVG) IsDistributed() bool {
	return pst.IsDistribute != PlatformSyncNoDistribute
}

type PlatformSyncStatusResult struct {
	Status          int    // -1:未配置状态;0:未知状态;1:忙碌状态;2:空闲状态
	LastUpdatedTime int64  `json:",omitempty"` // 上次更新时间(未更新过会不传该字段)
	LastUpdatedInfo string // 上次更新错误信息(未发生错误则为空)
}

type PlatformSyncStatus int

const (
	PlatformSyncStatusNotConfig = -1
	PlatformSyncStatusUnknown   = 0
	PlatformSyncStatusBusy      = 1
	PlatformSyncStatusIdle      = 2
)

func (pss PlatformSyncStatus) IsAvailable() bool {
	return pss == PlatformSyncStatusIdle
}

func (pss PlatformSyncStatus) IsConfig() bool {
	return pss != PlatformSyncStatusNotConfig
}

func (pss PlatformSyncStatus) IsBusy() bool {
	return pss == PlatformSyncStatusBusy
}

// 国标组织树
type PlatformSyncOrgRequest struct {
	PlatformType int // 1:国标,2:PVG
}

type PlatformSyncDevice struct {
	DeviceID   string `json:"Id"`
	DeviceName string `json:"Name"`
	OrgID      string
	Online     int
}

type PlatformSyncOrgResponse struct {
	OrgID         string                     `json:"Id"`
	SuperiorOrgID string                     `json:"SuperiorOrgId"`
	OrgName       string                     `json:"Name"`
	Orgs          []*PlatformSyncOrgResponse `json:"Orgs"`
	Devices       []*PlatformSyncDevice      `json:"Sensors"`
}

const (
	PlatformSyncGBTaskID          = "platformSyncGB"
	PlatformSyncPVGTaskID         = "platformSyncPVG"
	PlatformSyncGBLastUpdatedKey  = "platformSyncGBLastUpdatedKey"
	PlatformSyncPVGLastUpdatedKey = "platformSyncPVGLastUpdatedKey"

	PlatformSyncNoDistribute       = 1
	PlatformSyncAlreadyDistributed = 2
)

// 分配设备
type DistributeQuery []DistributeOrg

type DistributeOrg struct {
	OrgID                 string
	PlatformSyncDeviceIDs []string
}

func (dq DistributeQuery) GetCount() int {
	var result int
	for _, distributeOrg := range dq {
		result += len(distributeOrg.PlatformSyncDeviceIDs)
	}
	return result
}

func (dq DistributeQuery) GetIDs() []string {
	var result []string
	for _, distributeOrg := range dq {
		result = append(result, distributeOrg.PlatformSyncDeviceIDs...)
	}
	return result
}

func (dq DistributeQuery) GetOrgID(ID string) string {
	for _, distributeOrg := range dq {
		for _, deviceID := range distributeOrg.PlatformSyncDeviceIDs {
			if strings.EqualFold(deviceID, ID) {
				return distributeOrg.OrgID
			}
		}
	}
	return ""
}

type AthenaPlatformSyncGBConfig struct {
	URI           string
	LastUpdateKey string
}

type AthenaPlatformSyncGBStatusInfo struct {
	StateType      int    // 任务状态
	Done           int    // 同步完成设备
	TotalSensors   int    // 需要同步的总设备
	LastUpdateTime string // 上次同步时间
	Err            string // 错误信息
}

func (athenaPlatformSyncGBStatusInfo *AthenaPlatformSyncGBStatusInfo) IsError() bool {
	taskStatusInfoMap, ok := AthenaTaskStateTypeMap[AthenaTaskTypePlatformSyncGB]
	if !ok {
		return false
	}
	return athenaPlatformSyncGBStatusInfo.StateType == taskStatusInfoMap[AthenaTaskStateTypeEnumError]
}

func (athenaPlatformSyncGBStatusInfo *AthenaPlatformSyncGBStatusInfo) IsAvailable() bool {
	taskStatusInfoMap, ok := AthenaTaskStateTypeMap[AthenaTaskTypePlatformSyncGB]
	if !ok {
		return false
	}
	return athenaPlatformSyncGBStatusInfo.StateType != taskStatusInfoMap[AthenaTaskStateTypeEnumRunning]
}

func (athenaPlatformSyncGBStatusInfo *AthenaPlatformSyncGBStatusInfo) NeedReset() bool {
	taskStatusInfoMap, ok := AthenaTaskStateTypeMap[AthenaTaskTypePlatformSyncGB]
	if !ok {
		return false
	}
	return athenaPlatformSyncGBStatusInfo.StateType == taskStatusInfoMap[AthenaTaskStateTypeEnumError] || athenaPlatformSyncGBStatusInfo.StateType == taskStatusInfoMap[AthenaTaskStateTypeEnumFinish]
}

// 导出
const AthenaExportNoPaginationPageSize = -1

type AthenaExportConfig struct {
	SQL                    string          // SQL语句
	TableHead              []string        // 表头
	ImageColsName          map[string]bool // 指定需要根据单元格中的uri加载图片的列
	PageSize               int64           // 获取数据库数据时的分页大小,默认100,<0则不分页
	MaxGoroutinesToQueryDB int64           // 查询数据库时的最大并发度,默认10
	MaxGoroutinesToLoadPic int64           // 查询数据库时的最大并发度,默认10
	ExportPath             string          // 指定导出的excel文件的位置,默认./athenaDir/tmp/result/${taskId}.${timestamp}.xlxs
	URI                    string          // 可以指定把结果的excel用post的方式上传至URL,默认不会上传

	IncludeCustomImages *AthenaExportFeature // 带有上传图片的导出任务
}

type AthenaExportStatusInfo struct {
	StateType       int    // 任务状态
	TextDone        int    // 文本填入数
	TextRowTotal    int    // 总共行数
	PicLoadDone     int    // 图片加载数
	PicSkipped      int    // 图片跳过数
	PicsToLoadTotal int    // 需要加载的图片总数
	ExportPath      string // 存放位置
	URI             string // 上传位置
	UploadedToURI   bool   // 是否上传
	Err             string // 错误信息

	TaskID string // 任务ID
}

func (athenaExportStatusInfo *AthenaExportStatusInfo) IsFinishOrError() bool {
	taskStatusInfoMap, ok := AthenaTaskStateTypeMap[AthenaTaskTypeExport]
	if !ok {
		return false
	}
	return athenaExportStatusInfo.StateType == taskStatusInfoMap[AthenaTaskStateTypeEnumFinish] || athenaExportStatusInfo.StateType == taskStatusInfoMap[AthenaTaskStateTypeEnumError]
}

type AthenaExportField struct {
	DBName        string            // 数据库字段名
	ExportName    string            // 导出标题名
	NeedLoadImage bool              // 是否需要加载图片
	Dict          *AthenaExportDict // 字典解析配置
	StaticDict    map[string]string // 静态字典解析配置
	IsCustomImage bool              // 用户图片
	IsScore       bool              // 分数
}

type AthenaExportDict struct {
	No     int               // 第几个字典,默认第一个
	Key    string            // 字典key
	Expect map[string]string // 例外转换
	IsBit  bool              // 按位存储字典
}

var AthenaExportSpeedDict = map[string]string{
	"0": "未知",
	"1": "快速",
	"2": "中速",
	"3": "慢速",
}

var AthenaExportDirectionDict = map[string]string{
	"0": "未知",
	"1": "向上",
	"2": "向下",
	"3": "向左",
	"4": "向右",
	"5": "左上",
	"6": "左下",
	"7": "右上",
	"8": "右下",
}

var AthenaExportIDTypeDict = map[string]string{
	"0": "未知",
	"1": "身份证",
	"2": "护照",
	"3": "警官证",
	"4": "军官证",
	"5": "临时身份证",
}

var AthenaFaceCaptureByConditionExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(faces_index.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"Age": {
		ExportName: "年龄",
		DBName:     "faces_index.age_id",
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "faces_index.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Glass": {
		ExportName: "眼镜",
		DBName:     "faces_index.glass_id",
		Dict:       &AthenaExportDict{Key: "3"},
	},
	"Hat": {
		ExportName: "帽子",
		DBName:     "faces_index.hat_id",
		Dict:       &AthenaExportDict{Key: "4"},
	},
	"Helmet": {
		ExportName: "安全帽",
		DBName:     "''",
	},
	"Mask": {
		ExportName: "口罩",
		DBName:     "''",
	},
	"OriginalImage": {
		ExportName:    "高清图",
		DBName:        "faces_index.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "特征图",
		DBName:        "faces_index.cutboard_image_uri",
		NeedLoadImage: true,
	},
}

var AthenaFaceCaptureByFeatureExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(faces_index.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"Age": {
		ExportName: "年龄",
		DBName:     "faces_index.age_id",
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "faces_index.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Glass": {
		ExportName: "眼镜",
		DBName:     "faces_index.glass_id",
		Dict:       &AthenaExportDict{Key: "3"},
	},
	"Hat": {
		ExportName: "帽子",
		DBName:     "faces_index.hat_id",
		Dict:       &AthenaExportDict{Key: "4"},
	},
	"Helmet": {
		ExportName: "安全帽",
		DBName:     "''",
	},
	"Mask": {
		ExportName: "口罩",
		DBName:     "''",
	},
	"OriginalImage": {
		ExportName:    "高清原图",
		DBName:        "faces_index.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "比对图片",
		DBName:        "faces_index.cutboard_image_uri",
		NeedLoadImage: true,
	},
	"UploadImage": {
		ExportName:    "上传图片",
		DBName:        "faces_index.face_id",
		IsCustomImage: true,
	},
	"Similarity": {
		ExportName: "相似度",
		DBName:     "faces_index.face_id",
		IsScore:    true,
	},
	"唯一标识符": {
		ExportName: "唯一标识符",
		DBName:     "faces_index.face_id",
	},
}

var AthenaPedestrianCaptureByConditionExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(pedestrian_capture.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"HeadWear": {
		ExportName: "头部特征",
		DBName:     "( (CASE WHEN pedestrian_capture.desc_eye = 1 THEN '眼镜 ' ELSE '' END) || (CASE WHEN pedestrian_capture.desc_head = 1 THEN '帽子 ' WHEN pedestrian_capture.desc_head = 2 THEN '头盔 ' ELSE '' END) || (CASE WHEN pedestrian_capture.desc_mouth = 1 THEN '口罩 ' ELSE '' END) )",
	},
	"Bag": {
		ExportName: "包",
		DBName:     "( (CASE WHEN pedestrian_capture.with_shoulder_bag = 1 THEN '单肩包或者斜挎包 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_backpack = 1 THEN '双肩包 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_handbag = 1 THEN '手提包 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_hand_carry = 1 THEN '拎物品 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_pram = 1 THEN '婴儿车 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_luggage = 1 THEN '行李箱 ' ELSE '' END) )",
	},
	"UpperColor": {
		ExportName: "上身颜色",
		DBName:     "pedestrian_capture.upper_color",
		Dict:       &AthenaExportDict{Key: "54"},
	},
	"UpperTexture": {
		ExportName: "上身纹理",
		DBName:     "pedestrian_capture.upper_texture",
		Dict:       &AthenaExportDict{Key: "9"},
	},
	"LowerColor": {
		ExportName: "下身颜色",
		DBName:     "pedestrian_capture.lower_color",
		Dict:       &AthenaExportDict{Key: "56"},
	},
	"LowerType": {
		ExportName: "下身类别",
		DBName:     "pedestrian_capture.lower_style",
		Dict:       &AthenaExportDict{Key: "10"},
	},
	"AttachedItems": {
		ExportName: "附属物品",
		DBName:     "( (CASE WHEN pedestrian_capture.with_trolley = 1 THEN '手拉车 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_umbrella = 1 THEN '雨伞 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_hold_baby = 1 THEN '抱小孩 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_scarf = 1 THEN '围巾 ' ELSE '' END) )",
	},
	"UpperType": {
		ExportName: "上衣款式",
		DBName:     "pedestrian_capture.upper_style",
		Dict:       &AthenaExportDict{Key: "22"},
	},
	"ShoesType": {
		ExportName: "鞋子款式",
		DBName:     "pedestrian_capture.shoes_style",
		Dict:       &AthenaExportDict{Key: "23"},
	},
	"ShoesColor": {
		ExportName: "鞋子颜色",
		DBName:     "pedestrian_capture.shoes_color",
		Dict:       &AthenaExportDict{Key: "24"},
	},
	"HairStyle": {
		ExportName: "发型",
		DBName:     "pedestrian_capture.hair_style",
		Dict:       &AthenaExportDict{Key: "21"},
	},
	"Age": {
		ExportName: "年龄",
		DBName:     "pedestrian_capture.age_id",
		Dict:       &AthenaExportDict{Key: "8"},
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "pedestrian_capture.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Direction": {
		ExportName: "目标方向",
		DBName:     "pedestrian_capture.direction",
		StaticDict: AthenaExportDirectionDict,
	},
	"Speed": {
		ExportName: "目标速度",
		DBName:     "pedestrian_capture.speed",
		StaticDict: AthenaExportSpeedDict,
	},
	"HasFace": {
		ExportName: "是否有人脸",
		DBName:     "(CASE WHEN pedestrian_capture.has_face THEN '有' ELSE '无' END)",
	},
	"OriginalImage": {
		ExportName:    "高清图",
		DBName:        "pedestrian_capture.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "特征图",
		DBName:        "pedestrian_capture.cutboard_image_uri",
		NeedLoadImage: true,
	},
}

var AthenaPedestrianCaptureByFeatureExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(pedestrian_capture.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"HeadWear": {
		ExportName: "头部特征",
		DBName:     "( (CASE WHEN pedestrian_capture.desc_eye = 1 THEN '眼镜 ' ELSE '' END) || (CASE WHEN pedestrian_capture.desc_head = 1 THEN '帽子 ' WHEN pedestrian_capture.desc_head = 2 THEN '头盔 ' ELSE '' END) || (CASE WHEN pedestrian_capture.desc_mouth = 1 THEN '口罩 ' ELSE '' END) )",
	},
	"Bag": {
		ExportName: "包",
		DBName:     "( (CASE WHEN pedestrian_capture.with_shoulder_bag = 1 THEN '单肩包或者斜挎包 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_backpack = 1 THEN '双肩包 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_handbag = 1 THEN '手提包 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_hand_carry = 1 THEN '拎物品 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_pram = 1 THEN '婴儿车 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_luggage = 1 THEN '行李箱 ' ELSE '' END) )",
	},
	"UpperColor": {
		ExportName: "上身颜色",
		DBName:     "pedestrian_capture.upper_color",
		Dict:       &AthenaExportDict{Key: "54"},
	},
	"UpperTexture": {
		ExportName: "上身纹理",
		DBName:     "pedestrian_capture.upper_texture",
		Dict:       &AthenaExportDict{Key: "9"},
	},
	"LowerColor": {
		ExportName: "下身颜色",
		DBName:     "pedestrian_capture.lower_color",
		Dict:       &AthenaExportDict{Key: "56"},
	},
	"LowerType": {
		ExportName: "下身类别",
		DBName:     "pedestrian_capture.lower_style",
		Dict:       &AthenaExportDict{Key: "10"},
	},
	"AttachedItems": {
		ExportName: "附属物品",
		DBName:     "( (CASE WHEN pedestrian_capture.with_trolley = 1 THEN '手拉车 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_umbrella = 1 THEN '雨伞 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_hold_baby = 1 THEN '抱小孩 ' ELSE '' END) || (CASE WHEN pedestrian_capture.with_scarf = 1 THEN '围巾 ' ELSE '' END) )",
	},
	"UpperType": {
		ExportName: "上衣款式",
		DBName:     "pedestrian_capture.upper_style",
		Dict:       &AthenaExportDict{Key: "22"},
	},
	"ShoesType": {
		ExportName: "鞋子款式",
		DBName:     "pedestrian_capture.shoes_style",
		Dict:       &AthenaExportDict{Key: "23"},
	},
	"ShoesColor": {
		ExportName: "鞋子颜色",
		DBName:     "pedestrian_capture.shoes_color",
		Dict:       &AthenaExportDict{Key: "24"},
	},
	"HairStyle": {
		ExportName: "发型",
		DBName:     "pedestrian_capture.hair_style",
		Dict:       &AthenaExportDict{Key: "21"},
	},
	"Age": {
		ExportName: "年龄",
		DBName:     "pedestrian_capture.age_id",
		Dict:       &AthenaExportDict{Key: "8"},
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "pedestrian_capture.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Direction": {
		ExportName: "目标方向",
		DBName:     "pedestrian_capture.direction",
		StaticDict: AthenaExportDirectionDict,
	},
	"Speed": {
		ExportName: "目标速度",
		DBName:     "pedestrian_capture.speed",
		StaticDict: AthenaExportSpeedDict,
	},
	"HasFace": {
		ExportName: "是否有人脸",
		DBName:     "(CASE WHEN pedestrian_capture.has_face THEN '有' ELSE '无' END)",
	},
	"OriginalImage": {
		ExportName:    "高清原图",
		DBName:        "pedestrian_capture.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "比对图片",
		DBName:        "pedestrian_capture.cutboard_image_uri",
		NeedLoadImage: true,
	},
	"UploadImage": {
		ExportName:    "上传图片",
		DBName:        "pedestrian_capture.pedestrian_id",
		IsCustomImage: true,
	},
	"Similarity": {
		ExportName: "相似度",
		DBName:     "pedestrian_capture.pedestrian_id",
		IsScore:    true,
	},
	"唯一标识符": {
		ExportName: "唯一标识符",
		DBName:     "pedestrian_capture.pedestrian_id",
	},
}

var AthenaNonmotorCaptureByConditionExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(nonmotor_capture.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"Type": {
		ExportName: "车型",
		DBName:     "nonmotor_capture.nonmotor_type",
		Dict:       &AthenaExportDict{Key: "4"},
	},
	"LicenseNumber": {
		ExportName: "车辆号牌",
		DBName:     "nonmotor_capture.plate_text",
	},
	"LicenseColor": {
		ExportName: "车牌颜色",
		DBName:     "nonmotor_capture.plate_color_id",
		Dict:       &AthenaExportDict{Key: "21", No: 1},
	},
	"Attitude": {
		ExportName: "车辆角度",
		DBName:     "nonmotor_capture.nonmotor_gesture",
		Dict:       &AthenaExportDict{Key: "1"},
	},
	"Color": {
		ExportName: "车身颜色",
		DBName:     "nonmotor_capture.nonmotor_color_id",
		Dict:       &AthenaExportDict{Key: "8"},
	},
	"UpperColor": {
		ExportName: "上衣颜色",
		DBName:     "nonmotor_capture.upper_color",
		Dict:       &AthenaExportDict{Key: "5"},
	},
	"UpperStyle": {
		ExportName: "上衣样式",
		DBName:     "nonmotor_capture.upper_style",
		Dict:       &AthenaExportDict{Key: "6"},
	},
	"HeadWear": {
		ExportName: "头部特征",
		DBName:     "( (CASE WHEN nonmotor_capture.desc_eye = 1 THEN '眼镜 ' ELSE '' END) || (CASE WHEN nonmotor_capture.desc_head = 1 THEN '帽子 ' WHEN nonmotor_capture.desc_head = 2 THEN '头盔 ' ELSE '' END) || (CASE WHEN nonmotor_capture.desc_mouth = 1 THEN '口罩 ' ELSE '' END) )",
	},
	"AttachedItems": {
		ExportName: "附属物品",
		DBName:     "( (CASE WHEN nonmotor_capture.with_shoulder_bag = 1 THEN '单肩包或斜挎包 ' ELSE '' END) || (CASE WHEN nonmotor_capture.with_backpack = 1 THEN '双肩包 ' ELSE '' END) )",
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "nonmotor_capture.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Direction": {
		ExportName: "目标方向",
		DBName:     "nonmotor_capture.direction",
		StaticDict: AthenaExportDirectionDict,
	},
	"Speed": {
		ExportName: "目标速度",
		DBName:     "nonmotor_capture.speed",
		StaticDict: AthenaExportSpeedDict,
	},
	"HasFace": {
		ExportName: "是否有人脸",
		DBName:     "(CASE WHEN nonmotor_capture.has_face THEN '有' ELSE '无' END)",
	},
	"OriginalImage": {
		ExportName:    "高清图",
		DBName:        "nonmotor_capture.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "特征图",
		DBName:        "nonmotor_capture.cutboard_image_uri",
		NeedLoadImage: true,
	},
}

var AthenaNonmotorCaptureByFeatureExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(nonmotor_capture.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"Type": {
		ExportName: "车型",
		DBName:     "nonmotor_capture.nonmotor_type",
		Dict:       &AthenaExportDict{Key: "4"},
	},
	"LicenseNumber": {
		ExportName: "车辆号牌",
		DBName:     "nonmotor_capture.plate_text",
	},
	"LicenseColor": {
		ExportName: "车牌颜色",
		DBName:     "nonmotor_capture.plate_color_id",
		Dict:       &AthenaExportDict{Key: "21", No: 1},
	},
	"Attitude": {
		ExportName: "车辆角度",
		DBName:     "nonmotor_capture.nonmotor_gesture",
		Dict:       &AthenaExportDict{Key: "1"},
	},
	"Color": {
		ExportName: "车身颜色",
		DBName:     "nonmotor_capture.nonmotor_color_id",
		Dict:       &AthenaExportDict{Key: "8"},
	},
	"UpperColor": {
		ExportName: "上衣颜色",
		DBName:     "nonmotor_capture.upper_color",
		Dict:       &AthenaExportDict{Key: "5"},
	},
	"UpperStyle": {
		ExportName: "上衣样式",
		DBName:     "nonmotor_capture.upper_style",
		Dict:       &AthenaExportDict{Key: "6"},
	},
	"HeadWear": {
		ExportName: "头部特征",
		DBName:     "( (CASE WHEN nonmotor_capture.desc_eye = 1 THEN '眼镜 ' ELSE '' END) || (CASE WHEN nonmotor_capture.desc_head = 1 THEN '帽子 ' WHEN nonmotor_capture.desc_head = 2 THEN '头盔 ' ELSE '' END) || (CASE WHEN nonmotor_capture.desc_mouth = 1 THEN '口罩 ' ELSE '' END) )",
	},
	"AttachedItems": {
		ExportName: "附属物品",
		DBName:     "( (CASE WHEN nonmotor_capture.with_shoulder_bag = 1 THEN '单肩包或斜挎包 ' ELSE '' END) || (CASE WHEN nonmotor_capture.with_backpack = 1 THEN '双肩包 ' ELSE '' END) )",
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "nonmotor_capture.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Direction": {
		ExportName: "目标方向",
		DBName:     "nonmotor_capture.direction",
		StaticDict: AthenaExportDirectionDict,
	},
	"Speed": {
		ExportName: "目标速度",
		DBName:     "nonmotor_capture.speed",
		StaticDict: AthenaExportSpeedDict,
	},
	"HasFace": {
		ExportName: "是否有人脸",
		DBName:     "(CASE WHEN nonmotor_capture.has_face THEN '有' ELSE '无' END)",
	},
	"OriginalImage": {
		ExportName:    "高清原图",
		DBName:        "nonmotor_capture.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "比对图片",
		DBName:        "nonmotor_capture.cutboard_image_uri",
		NeedLoadImage: true,
	},
	"UploadImage": {
		ExportName:    "上传图片",
		DBName:        "nonmotor_capture.nonmotor_id",
		IsCustomImage: true,
	},
	"Similarity": {
		ExportName: "相似度",
		DBName:     "nonmotor_capture.nonmotor_id",
		IsScore:    true,
	},
	"唯一标识符": {
		ExportName: "唯一标识符",
		DBName:     "nonmotor_capture.nonmotor_id",
	},
}

var AthenaVehicleCaptureByConditionExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(vehicle_capture_index.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"LicenseNumber": {
		ExportName: "车辆号牌",
		DBName:     "vehicle_capture_index.plate_text",
	},
	"Type": {
		ExportName: "车辆类型",
		DBName:     "vehicle_capture_index.type_id",
		Dict:       &AthenaExportDict{Key: "34"},
	},
	"BrandModel": {
		ExportName: "品牌型号",
		DBName:     "vehicle_brand_dict.desc",
	},
	"Color": {
		ExportName: "车辆颜色",
		DBName:     "vehicle_capture_index.color_id",
		Dict:       &AthenaExportDict{Key: "39"},
	},
	"LicenseType": {
		ExportName: "车牌类型",
		DBName:     "vehicle_capture_index.plate_type_id",
		Dict:       &AthenaExportDict{Key: "20", No: 1},
	},
	"LicenseColor": {
		ExportName: "车牌颜色",
		DBName:     "vehicle_capture_index.plate_color_id",
		Dict:       &AthenaExportDict{Key: "21", No: 1},
	},
	"Special": {
		ExportName: "特殊车辆",
		DBName:     "vehicle_capture_index.special_str",
		Dict:       &AthenaExportDict{Key: "58", IsBit: true},
	},
	"Symbol": {
		ExportName: "标志物",
		DBName:     "vehicle_capture_index.symbols_desc",
	},
	"Direction": {
		ExportName: "目标方向",
		DBName:     "vehicle_capture_index.direction",
		StaticDict: AthenaExportDirectionDict,
	},
	"Speed": {
		ExportName: "目标速度",
		DBName:     "vehicle_capture_index.speed",
		StaticDict: AthenaExportSpeedDict,
	},
	"HasFace": {
		ExportName: "是否有人脸",
		DBName:     "(CASE WHEN vehicle_capture_index.has_face THEN '有' ELSE '无' END)",
	},
	"OriginalImage": {
		ExportName:    "高清图",
		DBName:        "vehicle_capture_index.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "特征图",
		DBName:        "vehicle_capture_index.cutboard_image_uri",
		NeedLoadImage: true,
	},
}

var AthenaVehicleCaptureByFeatureExportMap = map[string]AthenaExportField{
	"Time": {
		ExportName: "经过时间",
		DBName:     "to_char(to_timestamp(vehicle_capture_index.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "设备名称",
		DBName:     "sensors.sensor_name",
	},
	"LicenseNumber": {
		ExportName: "车辆号牌",
		DBName:     "vehicle_capture_index.plate_text",
	},
	"Type": {
		ExportName: "车辆类型",
		DBName:     "vehicle_capture_index.type_id",
		Dict:       &AthenaExportDict{Key: "34"},
	},
	"BrandModel": {
		ExportName: "品牌型号",
		DBName:     "vehicle_brand_dict.desc",
	},
	"Color": {
		ExportName: "车辆颜色",
		DBName:     "vehicle_capture_index.color_id",
		Dict:       &AthenaExportDict{Key: "39"},
	},
	"LicenseType": {
		ExportName: "车牌类型",
		DBName:     "vehicle_capture_index.plate_type_id",
		Dict:       &AthenaExportDict{Key: "20", No: 1},
	},
	"LicenseColor": {
		ExportName: "车牌颜色",
		DBName:     "vehicle_capture_index.plate_color_id",
		Dict:       &AthenaExportDict{Key: "21", No: 1},
	},
	"Special": {
		ExportName: "特殊车辆",
		DBName:     "vehicle_capture_index.special_str",
		Dict:       &AthenaExportDict{Key: "58", IsBit: true},
	},
	"Symbol": {
		ExportName: "标志物",
		DBName:     "vehicle_capture_index.symbols_desc",
	},
	"Direction": {
		ExportName: "目标方向",
		DBName:     "vehicle_capture_index.direction",
		StaticDict: AthenaExportDirectionDict,
	},
	"Speed": {
		ExportName: "目标速度",
		DBName:     "vehicle_capture_index.speed",
		StaticDict: AthenaExportSpeedDict,
	},
	"HasFace": {
		ExportName: "是否有人脸",
		DBName:     "(CASE WHEN vehicle_capture_index.has_face THEN '有' ELSE '无' END)",
	},
	"OriginalImage": {
		ExportName:    "高清原图",
		DBName:        "vehicle_capture_index.image_uri",
		NeedLoadImage: true,
	},
	"FeatureImage": {
		ExportName:    "比对图片",
		DBName:        "vehicle_capture_index.cutboard_image_uri",
		NeedLoadImage: true,
	},
	"UploadImage": {
		ExportName:    "上传图片",
		DBName:        "vehicle_capture_index.vehicle_id",
		IsCustomImage: true,
	},
	"Similarity": {
		ExportName: "相似度",
		DBName:     "vehicle_capture_index.vehicle_id",
		IsScore:    true,
	},
	"唯一标识符": {
		ExportName: "唯一标识符",
		DBName:     "vehicle_capture_index.vehicle_id",
	},
}

var AthenaFaceEventByConditionExportMap = map[string]AthenaExportField{
	"EventImage": {
		ExportName:    "抓拍人脸",
		DBName:        "face_events_index.captured_cutboard_image_uri",
		NeedLoadImage: true,
	},
	"MonitorRuleImage": {
		ExportName:    "对比人脸",
		DBName:        "face_events_index.civil_image_uri",
		NeedLoadImage: true,
	},
	"Confidence": {
		ExportName: "比对分值",
		DBName:     "to_char( TRUNC( face_events_index.confidence * 10000 ) / 100.0, '99D99' )|| '%'",
	},
	"Time": {
		ExportName: "报警时间",
		DBName:     "to_char(to_timestamp(face_events_index.ts/1000),'YYYY-MM-DD HH24:MI:SS')",
	},
	"SensorName": {
		ExportName: "报警地点",
		DBName:     "face_events_index.sensor_name",
	},
	"RepoName": {
		ExportName: "所在库",
		DBName:     "face_events_index.repo_name",
	},
	"Name": {
		ExportName: "姓名",
		DBName:     "face_events_index.civil_name",
	},
	"IDNo": {
		ExportName: "身份证",
		DBName:     "face_events_index.civil_id_no",
	},
	"Status": {
		ExportName: "审核状态",
		DBName:     "(CASE WHEN face_events_index.is_checked = 0 THEN '待审核' WHEN face_events_index.is_checked = 1 THEN '有效' WHEN face_events_index.is_checked = 2 THEN '无效' WHEN face_events_index.is_checked = 3 THEN '未知' ELSE '' END)",
	},
	"CheckPerson": {
		ExportName: "审核人",
		DBName:     "face_events_index.user_ids",
	},
}

var AthenaCivilsByConditionExportMap = map[string]AthenaExportField{
	"Name": {
		ExportName: "姓名",
		DBName:     "civil_attrs.name",
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "civil_attrs.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Repo": {
		ExportName: "所在库",
		DBName:     "face_repos.repo_name",
	},
	"Birthday": {
		ExportName: "出生日期",
		DBName:     "(CASE WHEN civil_attrs.birthday != '' THEN to_char(to_timestamp(to_number(civil_attrs.birthday, '9999999999999')/1000),'YYYY-MM-DD HH24:MI:SS') ELSE '' END)",
	},
	"IdType": {
		ExportName: "证件类型",
		DBName:     "civil_attrs.id_type",
		StaticDict: AthenaExportIDTypeDict,
	},
	"IdNo": {
		ExportName: "证件号码",
		DBName:     "civil_attrs.id_no",
	},
	"Comment": {
		ExportName: "备注",
		DBName:     "civil_attrs.comment",
	},
	"TargetImage": {
		ExportName:    "比对人脸",
		DBName:        "civil_attrs.image_uris",
		NeedLoadImage: true,
	},
}

var AthenaCivilsByFeatureExportMap = map[string]AthenaExportField{
	"Name": {
		ExportName: "姓名",
		DBName:     "civil_attrs.name",
	},
	"Gender": {
		ExportName: "性别",
		DBName:     "civil_attrs.gender_id",
		Dict:       &AthenaExportDict{Key: "16"},
	},
	"Repo": {
		ExportName: "所在库",
		DBName:     "face_repos.repo_name",
	},
	"Birthday": {
		ExportName: "出生日期",
		DBName:     "(CASE WHEN civil_attrs.birthday != '' THEN to_char(to_timestamp(to_number(civil_attrs.birthday, '9999999999999')/1000),'YYYY-MM-DD HH24:MI:SS') ELSE '' END)",
	},
	"IdType": {
		ExportName: "证件类型",
		DBName:     "civil_attrs.id_type",
		StaticDict: AthenaExportIDTypeDict,
	},
	"IdNo": {
		ExportName: "证件号码",
		DBName:     "civil_attrs.id_no",
	},
	"Comment": {
		ExportName: "备注",
		DBName:     "civil_attrs.comment",
	},
	"SrcImage": {
		ExportName:    "上传人脸",
		DBName:        "civil_attrs.civil_attr_id",
		IsCustomImage: true,
	},
	"Score": {
		ExportName: "相似度",
		DBName:     "civil_attrs.civil_attr_id",
		IsScore:    true,
	},
	"TargetImage": {
		ExportName:    "比对人脸",
		DBName:        "civil_images.image_uri",
		NeedLoadImage: true,
	},
	"唯一标识符": {
		ExportName: "唯一标识符",
		DBName:     "civil_images.image_id",
	},
}

type AthenaExportCondition struct {
	SelectedIDs []string
}

type AthenaExportFeature struct {
	Images []ImageQuery
	IDMap  []AthenaExportIDMap

	LoadFromRequest map[string]string
	DefaultSortKey  string
}

type AthenaExportIDMap struct {
	ID    string
	Score float64
	NO    uint
}

type AthenaExportQuery struct {
	Field       []string
	Export      []AthenaExportField
	InlineImage bool
	FileName    string
	AthenaExportCondition
	IncludeCustomImages *AthenaExportFeature
}

func (athenaExportQuery *AthenaExportQuery) CheckAndParse(athenaExportMap map[string]AthenaExportField) error {
	athenaExportQuery.Export = []AthenaExportField{}
	var customImageColName string
	var scoreColName string
	for _, field := range athenaExportQuery.Field {
		athenaExportField, ok := athenaExportMap[field]
		if !ok {
			return NewErrorVerificationFailed(ErrorCodeExportModelFieldNoFound)
		}
		athenaExportQuery.Export = append(athenaExportQuery.Export, athenaExportField)
	}
	for _, item := range athenaExportQuery.Export {
		if item.IsCustomImage {
			customImageColName = item.ExportName
		}
		if item.IsScore {
			scoreColName = item.ExportName
		}
	}
	if athenaExportQuery.IncludeCustomImages != nil {
		athenaExportQuery.IncludeCustomImages.DefaultSortKey = AthenaTaskTypeExportIncludeScoreUniqueID
		athenaExportField, ok := athenaExportMap[AthenaTaskTypeExportIncludeScoreUniqueID]
		if !ok {
			return NewErrorVerificationFailed(ErrorCodeExportModelFieldNoFound)
		}
		athenaExportQuery.Export = append(athenaExportQuery.Export, athenaExportField)
		for i, item := range athenaExportQuery.IncludeCustomImages.IDMap {
			if item.NO > 0 && item.NO >= uint(len(athenaExportQuery.IncludeCustomImages.Images)) {
				return NewErrorVerificationFailed(ErrorCodeExportModelFeatureIDMapNOExceedLimit)
			}
			athenaExportQuery.SelectedIDs = append(athenaExportQuery.SelectedIDs, item.ID)
			// fix score
			item.Score = math.Floor(item.Score * 100)
			athenaExportQuery.IncludeCustomImages.IDMap[i] = item
		}
		athenaExportQuery.IncludeCustomImages.LoadFromRequest = make(map[string]string)
		if !strings.EqualFold(customImageColName, "") { // export custom image
			athenaExportQuery.IncludeCustomImages.LoadFromRequest[customImageColName] = "NO"
		}
		if !strings.EqualFold(scoreColName, "") { // export score
			athenaExportQuery.IncludeCustomImages.LoadFromRequest[scoreColName] = "Score"
		}
	}
	return nil
}

type AthenaFaceCaptureExportQuery struct {
	FaceConditionRequest
	AthenaExportQuery
	FileName string
}

type AthenaPedestrianCaptureExportQuery struct {
	PedestrianConditionRequest
	AthenaExportQuery
	FileName string
}

type AthenaNonmotorCaptureExportQuery struct {
	NonmotorConditionRequest
	AthenaExportQuery
	FileName string
}

type AthenaVehicleCaptureExportQuery struct {
	VehicleConditionRequest
	AthenaExportQuery
	FileName string
}

type AthenaFaceEventExportQuery struct {
	FaceEventQuery
	AthenaExportQuery
}

type AthenaCivilsExportQuery struct {
	CivilQuery
	AthenaExportQuery
}

type AthenaExportStatusQuery struct {
	Type int `form:"type"`
}

func (athenaExportStatusQuery *AthenaExportStatusQuery) Check() error {

	switch athenaExportStatusQuery.Type {
	case AthenaTaskInnerTypeFaceCapture:
		athenaExportStatusQuery.Type = AthenaTaskTypeFaceCaptureExport
		return nil
	case AthenaTaskInnerTypeVehicleCapture:
		athenaExportStatusQuery.Type = AthenaTaskTypeVehicleCaptureExport
		return nil
	case AthenaTaskInnerTypeNonmotorCapture:
		athenaExportStatusQuery.Type = AthenaTaskTypeNonmotorCaptureExport
		return nil
	case AthenaTaskInnerTypePedestrianCapture:
		athenaExportStatusQuery.Type = AthenaTaskTypePedestrianCaptureExport
		return nil
	case AthenaTaskInnerTypeFaceEvent:
		return nil
	case AthenaTaskInnerTypeCivils:
		return nil
	case AthenaTaskInnerTypeFeatureCapture:
		return nil
	}

	return NewErrorVerificationFailed(ErrorCodeExportModelTargetInnerTypeNotFound)
}

func (athenaExportStatusQuery *AthenaExportStatusQuery) UserFeatureExportType() bool {
	return athenaExportStatusQuery.Type == AthenaTaskInnerTypeFeatureCapture
}

type AthenaExportResetQuery struct {
	TaskID string
}
