package pool

import (
	"fmt"
	"sync"

	"codes/gocodes/dg/utils/log"
)

const (
	SchedulerTypeTask = iota + 1
	SchedulerTypeSensor
	SchedulerTypeCalculate
)

type Scheduler struct {
	locker       *sync.RWMutex
	queue        Link
	indexMap     *sync.Map
	statusMap    *sync.Map
	linkNodeMap  *sync.Map
	progressMap  *sync.Map
	renderMap    *sync.Map
	channelMap   *sync.Map
	taskCountMap *sync.Map

	// link资源及状态
	faceUsed     uint32
	vehicleUsed  uint32
	faceTotal    uint32
	vehicleTotal uint32
	allTotal     uint32
	allUsed      uint32
	status       int
}

type Object interface {
	Pk() string
	CurrentStatus() int
	Instance() interface{}
	LockerKey() string
	LinkNode() string
	VideoProgress() int
	RenderURI() string
	SchedulerType() int
	CurrentChannel() int
}

var (
	scheduler *Scheduler
	linkOnce  sync.Once
)

func InitScheduler() *Scheduler {
	linkOnce.Do(func() {
		scheduler = &Scheduler{
			queue:        Link{},
			indexMap:     &sync.Map{},
			locker:       &sync.RWMutex{},
			statusMap:    &sync.Map{},
			linkNodeMap:  &sync.Map{},
			progressMap:  &sync.Map{},
			renderMap:    &sync.Map{},
			channelMap:   &sync.Map{},
			taskCountMap: &sync.Map{},
		}

	})

	return scheduler
}

func (s *Scheduler) LoadFromDb(initData []Object) {
	for _, v := range initData {
		s.Push(v)
	}
}

func (s *Scheduler) readLock() {
	s.locker.RLock()
}

func (s *Scheduler) readUnLock() {
	s.locker.RUnlock()
}

func (s *Scheduler) writeLock() {
	s.locker.Lock()
}

func (s *Scheduler) writeUnLock() {
	s.locker.Unlock()
}

func (s *Scheduler) Push(obj Object) {
	s.writeLock()
	defer s.writeUnLock()
	// log.Debugf("scheduler Push, PK:%v status:%v", obj.Pk(), obj.CurrentStatus())

	kind := obj.SchedulerType()

	node, ok := s.indexMap.Load(obj.Pk())
	if ok {
		s.statusMap.Store(obj.Pk(), obj.CurrentStatus())
		if kind == SchedulerTypeTask {
			s.linkNodeMap.Store(obj.Pk(), obj.LinkNode())
			s.progressMap.Store(obj.Pk(), obj.VideoProgress())
			s.renderMap.Store(obj.Pk(), obj.RenderURI())
			s.channelMap.Store(obj.Pk(), obj.CurrentChannel())
		}

		s.queue.Update(node.(*Node), obj.Instance())
		return
	}

	var count int
	value, exist := s.taskCountMap.Load(kind)
	if exist {
		count = value.(int) + 1
	} else {
		count++
	}

	s.taskCountMap.Store(kind, count)

	node = s.queue.Append(obj.Instance())
	s.indexMap.Store(obj.Pk(), node)
	s.statusMap.Store(obj.Pk(), obj.CurrentStatus())

	if kind == SchedulerTypeTask {
		s.linkNodeMap.Store(obj.Pk(), obj.LinkNode())
		s.progressMap.Store(obj.Pk(), obj.VideoProgress())
		s.renderMap.Store(obj.Pk(), obj.RenderURI())
		s.channelMap.Store(obj.Pk(), obj.CurrentChannel())
	}
}

func (s *Scheduler) Pop(obj Object) {
	s.writeLock()
	defer s.writeUnLock()

	kind := obj.SchedulerType()
	var count int
	value, _ := s.taskCountMap.Load(kind)
	count = value.(int) - 1
	s.taskCountMap.Store(kind, count)

	log.Debugf("scheduler Pop, PK:%v status:%v", obj.Pk(), obj.CurrentStatus())

	value, ok := s.indexMap.Load(obj.Pk())
	if ok {
		s.queue.Remove(value.(*Node))

		s.indexMap.Delete(obj.Pk())
	}

	_, ok = s.statusMap.Load(obj.Pk())
	if ok {
		s.statusMap.Delete(obj.Pk())
	}

	if kind == SchedulerTypeTask {
		_, ok = s.linkNodeMap.Load(obj.Pk())
		if ok {
			s.linkNodeMap.Delete(obj.Pk())
		}

		_, ok = s.progressMap.Load(obj.Pk())
		if ok {
			s.progressMap.Delete(obj.Pk())
		}

		_, ok = s.renderMap.Load(obj.Pk())
		if ok {
			s.renderMap.Delete(obj.Pk())
		}

		_, ok = s.channelMap.Load(obj.Pk())
		if ok {
			s.channelMap.Delete(obj.Pk())
		}
	}
}

func (s *Scheduler) Len() int {
	s.readLock()
	defer s.readUnLock()

	return int(s.queue.Len())
}

func (s *Scheduler) Next() Object {
	s.readLock()
	defer s.readUnLock()

	content := s.queue.Next()
	if content == nil {
		return nil
	}

	return content.val.(Object)
}

func (s *Scheduler) GetStatusFromMap(key interface{}) int {
	s.readLock()
	defer func() {
		s.readUnLock()
		if r := recover(); r != nil {
			log.Errorln(r)
		}
	}()

	value, ok := s.statusMap.Load(key)
	if !ok {
		panic(fmt.Errorf("the key:%v not save in the status map", key))
	}

	return value.(int)
}

func (s *Scheduler) GetLinkNodeFromMap(key interface{}) string {
	s.readLock()
	defer func() {
		s.readUnLock()
		if r := recover(); r != nil {
			log.Errorln(r)
		}
	}()

	value, ok := s.linkNodeMap.Load(key)
	if !ok {
		panic(fmt.Errorf("the key:%v not save in the link node map", key))
	}

	return value.(string)
}

func (s *Scheduler) GetProgressFromMap(key interface{}) int {
	s.readLock()
	defer func() {
		s.readUnLock()
		if r := recover(); r != nil {
			log.Errorln(r)
		}
	}()

	value, ok := s.progressMap.Load(key)
	if !ok {
		panic(fmt.Errorf("the key:%v not save in the progress map", key))
	}

	return value.(int)
}

func (s *Scheduler) GetRenderURIFromMap(key interface{}) string {
	s.readLock()
	defer func() {
		s.readUnLock()
		if r := recover(); r != nil {
			log.Errorln(r)
		}
	}()

	value, ok := s.renderMap.Load(key)
	if !ok {
		panic(fmt.Errorf("the key:%v not save in the render map", key))
	}

	return value.(string)
}

func (s *Scheduler) GetChannelFromMap(key interface{}) int {
	s.readLock()
	defer s.readUnLock()

	value, ok := s.channelMap.Load(key)
	if !ok {
		panic(fmt.Errorf("the key:%v not save in the channel map", key))
	}

	return value.(int)
}

func (s *Scheduler) SetResource(status int, faceTotal, vehicleTotal, faceUsed, vehicleUsed, allTotal, allUsed uint32) {
	s.writeLock()
	defer s.writeUnLock()

	s.status = status
	s.faceTotal = faceTotal
	s.vehicleTotal = vehicleTotal
	s.faceUsed = faceUsed
	s.vehicleUsed = vehicleUsed
	s.allTotal = allTotal
	s.allUsed = allUsed
}

func (s *Scheduler) GetResource() (status int, faceTotal uint32, vehicleTotal uint32, faceUsed uint32, vehicleUsed, allTotal, allUsed uint32) {
	s.readLock()
	defer s.readUnLock()

	status = s.status
	faceTotal = s.faceTotal
	vehicleTotal = s.vehicleTotal
	faceUsed = s.faceUsed
	vehicleUsed = s.vehicleUsed
	allTotal = s.allTotal
	allUsed = s.allUsed

	return
}

func (s *Scheduler) GetTaskCount(kind int) int {
	s.readLock()
	defer s.readUnLock()

	v, ok := s.taskCountMap.Load(kind)
	if !ok {
		return 0
	}

	return v.(int)
}
