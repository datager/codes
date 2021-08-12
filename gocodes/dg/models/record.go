package models

const (
	RECORD_FIELD_FACE_CAPTURE      = "face_capture"
	RECORD_FIELD_FACE_CAPTURE_TRAJ = "face_capture_traj"
	RECORD_FIELD_FACE_COMPARE      = "face_compare"
	RECORD_FIELD_FACE_EVENT        = "face_event"
	RECORD_FIELD_FACE_PERSON_FILE  = "face_person_file"
	RECORD_FIELD_FACE_WHITE_EVENT  = "face_white_event"
	RECORD_FIELD_FACE_RULE         = "face_rule"
	RECORD_FIELD_DEVICE_GEO        = "device_geo"
	RECORD_FIELD_GIS               = "gis"
	RECORD_FIELD_FACE_REPO         = "face_repo"
	RECORD_FIELD_CIVIL             = "civil"
	RECORD_FIELD_DEVICE            = "device"
	RECORD_FIELD_DEVICE_PARAM      = "device_param"
	RECORD_FIELD_DEVICE_ROLE       = "device_role"
	RECORD_FIELD_FUNC_ROLE         = "func_role"
	RECORD_FIELD_ORG               = "org"
	RECORD_FIELD_PERSON_FILE       = "person_file"
	RECORD_FIELD_USER              = "user"
	RECORD_FIELD_RECORD            = "record"
	RECORD_FIELD_USER_AUTH         = "user_auth"

	RECORD_TYPE_LOGIN        = "login"
	RECORD_TYPE_LOGOUT       = "logout"
	RECORD_TYPE_ADD          = "add"
	RECORD_TYPE_UPDATE       = "update"
	RECORD_TYPE_DELETE       = "delete"
	RECORD_TYPE_QUERY        = "query"
	RECORD_TYPE_RESET        = "reset"
	RECORD_TYPE_BATCH_ADD    = "batch_add"
	RECORD_TYPE_BATCH_UPDATE = "batch_update"
	RECORD_TYPE_BATCH_DELETE = "batch_delete"
)

var RecordFieldNames map[string]string

func init() {
	RecordFieldNames = map[string]string{
		RECORD_FIELD_FACE_CAPTURE:      "人脸抓拍检索",
		RECORD_FIELD_FACE_CAPTURE_TRAJ: "轨迹检索",
		RECORD_FIELD_FACE_COMPARE:      "1v1比对",
		RECORD_FIELD_FACE_EVENT:        "报警管理",
		RECORD_FIELD_FACE_PERSON_FILE:  "人员档案",
		RECORD_FIELD_FACE_WHITE_EVENT:  "陌生人报警管理",
		RECORD_FIELD_FACE_RULE:         "布控管理",
		RECORD_FIELD_DEVICE_GEO:        "地图标定",
		RECORD_FIELD_GIS:               "GIS接入",
		RECORD_FIELD_FACE_REPO:         "比对库管理",
		RECORD_FIELD_CIVIL:             "人员信息管理",
		// RECORD_FIELD_FAILED_CIVIL:      "入库失败人员",
		RECORD_FIELD_DEVICE:       "接入管理",
		RECORD_FIELD_DEVICE_PARAM: "高级参数",
		RECORD_FIELD_ORG:          "组织管理",
		RECORD_FIELD_PERSON_FILE:  "人员档案",
		RECORD_FIELD_DEVICE_ROLE:  "设备角色管理",
		RECORD_FIELD_FUNC_ROLE:    "功能角色管理",
		RECORD_FIELD_USER:         "用户管理",
		RECORD_FIELD_RECORD:       "日志管理",
		RECORD_FIELD_USER_AUTH:    "用户认证",
	}
}

type Record struct {
	Timestamp     int64
	RecordId      string
	UserName      string
	UserId        string
	OperField     string
	OperType      string
	OperContent   string
	Images        []*ImageQuery
	Url           []string
	IpAddr        string
	Result        string
	ResultContent string
	Status        int
}
