package redis

import (
	"codes/gocodes/dg/utils/config"
	"codes/gocodes/dg/utils/log"
)

const (
	cachePrefix = "loki-"
)

const (
	HalfMinute = 30
	OneMinute  = 60
	HalfHour   = 1800
	OneHour    = 3600
	HalfDay    = OneHour * 12
	OneDay     = OneHour * 24
	OneWeek    = OneDay * 7
	OneMonth   = OneDay * 30
	OneYear    = OneDay * 365
)

type CacheClient struct {
	ConnPool
}

func NewCacheClient(conf config.Config) *CacheClient {
	addr := conf.GetString("services.redis.cache.addr")
	if addr == "" {
		panic("cache server addr is empty")
	}

	password := conf.GetString("services.redis.cache.password")

	conn := initConnPool(addr, password, cachePrefix)
	log.Debug("NewCacheClient...")

	return &CacheClient{
		ConnPool: *conn,
	}
}
