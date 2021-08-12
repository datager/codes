package models

import "time"

const (
	DeleteTaskTypeSensor      = "sensor"
	DeleteTaskTypeRepo        = "repo"
	DeleteTaskTypeOrg         = "org"
	DeleteTaskTypeVehicleRepo = "vehiclerepo"
)

const (
	DeleteTaskStatusCreated = iota + 1
)

type DeleteTask struct {
	TaskID    string
	TaskType  string
	Status    int
	DeleteID  string
	Ts        int64
	EndTime   int64
	Note      string
	UpdatedAt time.Time
}
