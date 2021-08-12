package entities

type PersonPathAnalyzeDetail struct {
	TaskID   string `xorm:"task_id  not null VARCHAR(1024) default ''"`
	Ts       int64  `xorm:"ts bigint not null default 0"`
	FaceReid string `xorm:"face_reid" varchar ( 36 ) not null default ''`
	FaceVid  string `xorm:"face_vid" varchar ( 36 ) not null default ''`
	SensorID string `xorm:"sensor_id"varchar ( 36 ) not null default ''`
}

func (t *PersonPathAnalyzeDetail) TableName() string {
	return TableNamePersonPathAnalyzeDetail
}
