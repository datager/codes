package models

// (车+人)频次分析结果 详情 请求
type QueryFrequencyResultDetailReq struct {
	TaskID   string
	SensorID string
	VID      string
	Pagination
}
