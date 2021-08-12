package models

const (
	TaskQueueUp = iota + 1
	TaskExecutory
	TaskStop
	TaskFinish
)

type TaskCollision struct {
	TaskId     string
	TaskName   string
	TargetRepo string
	SrcRepo    string
	Threshold  float32
}

type TaskCollisionModel struct {
	TaskCollision
	Timestamp int64
	Status    int
	Progress  float32
}
type TaskCollisionQuery struct {
	Limit          int
	Offset         int
	Total          int
	StartTimestamp int64
	EndTimestamp   int64
	Status         int
	RepoId         string
	TaskIds        []string
}
