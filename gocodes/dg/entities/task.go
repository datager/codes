package entities

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/redis"
)

const (
	cacheKeyTasks = "tasks-cache-key-task-id:%v:v2"
)

type Tasks struct {
	Uts        time.Time         `xorm:"not null default 'now()' DATETIME"`
	TaskID     string            `xorm:"'task_id' not null pk VARCHAR(1024)"`
	SensorName string            `xorm:"not null"`
	Ts         int64             `xorm:"not null"`
	SensorID   string            `xorm:"'sensor_id' not null"`
	LinkID     string            `xorm:"'link_id' not null"`
	Node       string            `xorm:"'node' not null"`
	DetTypes   int               `xorm:"not null"`
	Progress   int               `xorm:"not null"`
	RenderURL  string            `xorm:"'render_url' not null"`
	Config     models.TaskConfig `xorm:"json not null"`
	Status     int               `xorm:"not null"`
	OrgID      string            `xorm:"'org_id' not null"`
	Channel    int               `xorm:"'channel' not null"`
	Type       int               `xorm:"not null"`
}

func (t *Tasks) TableName() string {
	return TableNameTask
}

func (t *Tasks) CacheKey() string {
	return fmt.Sprintf(cacheKeyTasks, t.TaskID)
}

func (t *Tasks) CacheExpireTime() int64 {
	return redis.OneMonth
}

func (t *Tasks) ToModel() *models.Tasks {
	m := &models.Tasks{
		SimpleTask: models.SimpleTask{
			TaskID:   t.TaskID,
			SensorID: t.SensorID,
			Config:   t.Config,
			Status:   t.Status,
			Type:     t.Type,
			Channel:  t.Channel,
		},
		LinkID:     t.LinkID,
		Node:       t.Node,
		Progress:   t.Progress,
		RenderURL:  t.RenderURL,
		Uts:        t.Uts,
		SensorName: t.SensorName,
		Ts:         t.Ts,
		DetTypes:   t.DetTypes,
	}
	m.SetDetTypeInfo()
	return m
}

func (t *Tasks) ToInfo() *models.TaskInfo {
	if t == nil {
		return nil
	}

	return &models.TaskInfo{
		SensorID:   t.SensorID,
		DetType:    t.DetTypes,
		TaskStatus: t.Status,
		TaskID:     t.TaskID,
	}
}

func TasksModelToEntity(t models.Tasks) *Tasks {
	return &Tasks{
		Uts:        time.Now(),
		TaskID:     t.TaskID,
		SensorName: t.SensorName,
		Ts:         t.Ts,
		SensorID:   t.SensorID,
		LinkID:     t.LinkID,
		Node:       t.Node,
		Progress:   t.Progress,
		RenderURL:  t.RenderURL,
		DetTypes:   t.DetTypes,
		Config:     t.Config,
		Status:     t.Status,
		OrgID:      t.OrgID,
		Type:       t.Type,
		Channel:    t.Channel,
	}
}

func TasksToSensorIDsMap(tasks []*Tasks) map[string]*Tasks {
	taskMap := make(map[string]*Tasks)
	for _, t := range tasks {
		taskMap[t.SensorID] = t
	}
	return taskMap
}

func TasksHelper(ids []string) []redis.Entity {
	slice := make([]redis.Entity, len(ids))
	for i, v := range ids {
		slice[i] = &Tasks{TaskID: v}
	}

	return slice
}
