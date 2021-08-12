package entities

import (
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils"
)

type PathAnalyzeTask struct {
	Uts            time.Time                `xorm:"uts not null default 'now()' DATETIME"`
	Ts             int64                    `xorm:"ts not null default 0 BIGINT"`
	TaskID         string                   `xorm:"task_id pk not null VARCHAR(1024) default ''"`
	TaskName       string                   `xorm:"task_name"`
	Sensors        []models.MiniSensorSlice `xorm:"sensors json"`
	DayTime        []int64                  `xorm:"day_time"`
	PathConfidence int64                    `xorm:"path_confidence"`
	Type           int                      `xorm:"type"`
}

func (t *PathAnalyzeTask) TableName() string {
	return TableNamePathAnalyzeTask
}

func PathAnalyzeTaskModelToEntity(model *models.PathAnalyzeTaskUpdateStatusRequest, minSensors []models.MiniSensorSlice) *PathAnalyzeTask {
	return &PathAnalyzeTask{
		Uts:            time.Now(),
		Ts:             utils.GetNowTs(),
		TaskID:         model.TaskID,
		TaskName:       model.TaskName,
		Sensors:        minSensors,
		PathConfidence: model.PathConfidence,
		DayTime:        model.StartAndEndTimestamp,
		Type:           model.TaskType,
	}
}

func (t *PathAnalyzeTask) ToModel() *models.PathAnalyzeTaskUpdateStatusRequest {
	return &models.PathAnalyzeTaskUpdateStatusRequest{
		TaskID:               t.TaskID,
		TaskName:             t.TaskName,
		Sensors:              t.Sensors,
		StartAndEndTimestamp: t.DayTime,
		TaskType:             t.Type,
		PathConfidence:       t.PathConfidence,
	}
}
