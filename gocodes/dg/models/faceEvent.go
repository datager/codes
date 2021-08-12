package models

const (
	FaceEventIsChecked_Pending int = iota
	FaceEventIsChecked_Valid
	FaceEventIsChecked_Invalid
	FaceEventIsChecked_Unknown
)

var FaceEventIsCheckedNames map[int]string

func init() {
	FaceEventIsCheckedNames = make(map[int]string)
	FaceEventIsCheckedNames[FaceEventIsChecked_Pending] = "未审核"
	FaceEventIsCheckedNames[FaceEventIsChecked_Valid] = "审核有效"
	FaceEventIsCheckedNames[FaceEventIsChecked_Invalid] = "审核无效"
	FaceEventIsCheckedNames[FaceEventIsChecked_Unknown] = "未知"
}

type FaceEvent struct {
	EventId         string
	EventReId       string
	Timestamp       int64
	SensorId        string
	SensorName      string
	FaceRepoId      string
	FaceRepoName    string
	NameListAttr    int
	FaceId          string
	FaceReId        string
	CapturedImage   *ImageResult
	ImageId         string
	CivilAttr       *CivilAttr
	RegisteredImage *ImageResult
	RuleId          string
	Rule            string
	Confidence      float32
	IsChecked       int
	Comment         string
	UserIds         string
	EventIsChecked  int
	Status          TaskStatus
	RuleName        string
	OrgCode         int64
	SensorIntID     int32
}

func GetIsCheckedName(id int) string {
	if name, exist := FaceEventIsCheckedNames[id]; exist {
		return name
	}
	return "未知"
}
