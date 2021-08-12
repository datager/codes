package timecost

import (
	"sync"
	"time"

	"codes/gocodes/dg/utils/log"
)

var (
	mu      sync.Mutex
	timeMap = make(map[string]time.Time)
)

func Start(name string) {
	mu.Lock()
	defer mu.Unlock()

	timeMap[name] = time.Now()
	//log.Infoln("[TimeCost] ", name, " start")
}

func Stop(name string) {
	mu.Lock()
	defer mu.Unlock()

	now, ok := timeMap[name]
	if !ok {
		log.Errorln("[TimeCost] name not found")
		return
	}
	log.Infoln("[TimeCost] ", name, " stop took ", time.Since(now))
}
