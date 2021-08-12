package models

import "time"

const (
	FileTypePersonRealName = iota + 1
	FileTypePersonAnonymous
	FileTypeVehicle
)

type FileCountStatistics struct {
	ID    int64
	Ts    int64
	Uts   time.Time
	Type  int
	Count int64
}
