package timecost

import (
	"codes/gocodes/dg/utils/uuid"
	"fmt"
	"sync"
	"time"

	"codes/gocodes/dg/utils/log"
)

var (
	muUUID      sync.Mutex
	timeMapUUID = make(map[string]time.Time)
)

func StartWithUUID(name string) (uid string) {
	muUUID.Lock()
	defer muUUID.Unlock()

	uid = uuid.NewV4().String()
	key := fmt.Sprintf("%s-%s", name, uid)

	timeMapUUID[key] = time.Now()
	log.Infoln("[TimeCost-UUID] ", key, " start")

	return uid
}

func StopWithUUID(name, uuid string) {
	muUUID.Lock()
	defer muUUID.Unlock()

	key := fmt.Sprintf("%s-%s", name, uuid)
	st, ok := timeMapUUID[key] // start time
	if !ok {
		log.Errorln("[TimeCost-UUID] name not found")
		return
	}
	log.Infoln("[TimeCost-UUID] ", key, " stop took ", time.Since(st))
	delete(timeMapUUID, key)
}
