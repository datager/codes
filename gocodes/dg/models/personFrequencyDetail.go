package models

import "time"

// 右侧纵向列表详情的响应体(同一 sensor 下最晚 ts 的 vid + reid数量)
type PersonFrequencyDetailWithReIDCountExtend struct {
	LatestTsOfCaptureVID *PersonFrequencyDetail // 每个 vid 选最晚ts的那个 reid 对应的 vid
	ReIDCount            int64                  // 同时返回 属于该 vid 的 reid 的数量
}

// 中部横向列表详情的响应体(同一 sensor + vid 的 各个 reid)
type PersonFrequencyDetail struct {
	Uts    time.Time
	Ts     int64
	TaskID string
	//MaxTsTag  int64
	//ReIDCount int64

	// 以下来自faces_index表
	CapturedFace
}
