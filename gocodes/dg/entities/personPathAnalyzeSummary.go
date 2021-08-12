package entities

import "codes/gocodes/dg/models"

type PersonPathAnalyzeSummary struct {
	TaskID    string `xorm:"task_id  not null VARCHAR(1024) default ''"`
	FaceVid   string `xorm:"face_vid" not null VARCHAR(36) default ''`
	ReidCount int    `xorm:"reid_count" bigint not null default 0`
}

func (t *PersonPathAnalyzeSummary) TableName() string {
	return TableNamePersonPathAnalyzeSummary
}

func (t *PersonPathAnalyzeSummary) ToModel() *models.PathAnalyzeSummaryItem {

	return &models.PathAnalyzeSummaryItem{
		Count:  int64(t.ReidCount),
		VID:    t.FaceVid,
		TaskID: t.TaskID,
	}
}
