package pool

import (
	"fmt"
	"sync"

	"codes/gocodes/dg/utils/log"
)

type Locker struct {
	locks   map[string]*sync.RWMutex
	locksRW *sync.RWMutex
}

var locker *Locker

func init() {
	locker = NewLocker()
}

// u can use name to tell the kind of key
// eg: pool.Lock(sensorID, "sensor")
// eg: pool.Lock(taskID, "task")
func Lock(key string, name ...string) {
	// st := time.Now()
	// log.Debugf("尝试获取%s写锁%s", name, key)
	// defer log.Debugf("成功获取%s写锁%s, ---took: %s", name, key, time.Since(st))

	locker.Lock(key)
}

func Unlock(key string, name ...string) {
	// st := time.Now()
	// log.Debugf("尝试释放%s写锁%s", name, key)
	// defer log.Debugf("成功释放%s写锁%s, ---took: %s", name, key, time.Since(st))

	locker.UnLock(key)
}

func RLock(key string, name ...string) {
	// st := time.Now()
	// log.Debugf("尝试获取%s读锁%s", name, key)
	// defer log.Debugf("成功获取%s读锁%s, ---took: %s", name, key, time.Since(st))

	locker.RLock(key)
}

func RUnlock(key string, name ...string) {
	// st := time.Now()
	// log.Debugf("尝试释放%s读锁%s", name, key)
	// defer log.Debugf("成功释放%s读锁%s, ---took: %s", name, key, time.Since(st))

	locker.RUnlock(key)
}

func DeleteKey(key string) {
	locker.deleteLock(key)
}

func NewLocker() *Locker {
	return &Locker{
		locks:   make(map[string]*sync.RWMutex),
		locksRW: new(sync.RWMutex),
	}
}

func (lker *Locker) Lock(key string) {
	lk, ok := lker.getLock(key)
	if !ok {
		lk = lker.newLock(key)
	}

	lk.Lock()
}

func (lker *Locker) UnLock(key string) {
	lk, ok := lker.getLock(key)
	if !ok {
		// panic(fmt.Errorf("BUG: Lock for key '%v' not initialized", key))
		log.Debugf("get lock by key:%v but not found")
		return
	}

	lk.Unlock()
}

func (lker *Locker) RLock(key string) {
	lk, ok := lker.getLock(key)
	if !ok {
		lk = lker.newLock(key)
	}

	lk.RLock()
}

func (lker *Locker) RUnlock(key string) {
	lk, ok := lker.getLock(key)
	if !ok {
		panic(fmt.Errorf("BUG: Lock for key '%v' not initialized", key))
	}

	lk.RUnlock()
}

func (lker *Locker) newLock(key string) *sync.RWMutex {
	lker.locksRW.Lock()
	defer lker.locksRW.Unlock()

	if lk, ok := lker.locks[key]; ok {
		return lk
	}

	lk := new(sync.RWMutex)
	lker.locks[key] = lk
	return lk
}

func (lker *Locker) getLock(key string) (*sync.RWMutex, bool) {
	lker.locksRW.RLock()
	defer lker.locksRW.RUnlock()

	lock, ok := lker.locks[key]
	return lock, ok
}

func (lker *Locker) deleteLock(key string) {
	lker.locksRW.Lock()
	defer lker.locksRW.Unlock()

	if _, ok := lker.locks[key]; ok {
		delete(lker.locks, key)
	}
}
