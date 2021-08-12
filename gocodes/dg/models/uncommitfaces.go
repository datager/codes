package models

type QueryListReq struct {
	Payload *UncommitFaceReq
	Type    int
}
type UncommitFaceReq struct {
	Limit          int64 //limit 和 offset 有一个为0，则视为不分页
	Offset         int64
	OrderAsc       bool
	OrderBy        string
	CheckStatus    []int    //审核状态数组：0：未审核；1：审核有效；2：审核无效；3：无法确定
	StartTimestamp int64    //起始记录时间戳
	EndTimestamp   int64    //结束时间时间戳
	SensorIds      []string //设备ID数组，用于匹配一批需要检索的设备
	Name           string   //姓名关键字，使用like 模糊匹配
	IdType         int64    //证件类型
	IdNo           string   //证件号码，必须全量匹配
	Addr           string   //使用like模糊匹配地址关键字
}
type UncommitFaceCheckReq struct {
	PairId     string
	CheckValid int
}
type UncommitFace struct {
	PairId          string
	Timestamp       int64
	Name            string
	SensorId        string
	SensorName      string
	Confidence      float64
	IsChecked       int
	Comment         string
	CapturedImage   *UncommitFaceImage
	RegisteredImage *UncommitFaceImage
	CivilAttr       *UncommitFaceCivilAttr
}

type UncommitFaceImage struct {
	FaceId            string
	ImageUri          string
	CutboardImageUri  string
	CutboardX         int64
	CutboardY         int64
	CutboardWidth     int64
	CutboardHeight    int64
	CutboardResWidth  int64
	CutboardResHeight int64
	Feature           string
	ImageType         int64
}
type UncommitFaceCivilAttr struct {
	Name     string
	IDType   int
	IDNo     string
	Addr     string
	NationId int
}
