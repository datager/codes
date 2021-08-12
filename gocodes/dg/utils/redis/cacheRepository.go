package redis

import (
	"reflect"

	"codes/gocodes/dg/utils/config"
	"codes/gocodes/dg/utils/log"

	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

type CacheRepository struct {
	ConnPool
	engine *xorm.Engine
}

type Entity interface {
	TableName() string
	CacheKey() string
	CacheExpireTime() int64
}

type entityHelper func([]string) []Entity
type entityMoreHelper func([]string, ...string) []Entity

func NewCacheRepository(conf config.Config, engine *xorm.Engine) *CacheRepository {
	addr := conf.GetString("services.redis.cache.addr")
	if addr == "" {
		panic("cache server addr is empty")
	}

	password := conf.GetString("services.redis.cache.password")

	conn := initConnPool(addr, password, cachePrefix)

	log.Debugln("NewCacheRepository...")
	return &CacheRepository{
		ConnPool: *conn,
		engine:   engine,
	}
}

func (c *CacheRepository) Close() error {
	return c.ConnPool.Close()
}

func (c *CacheRepository) Get(instance Entity) error {
	cacheKey := instance.CacheKey()
	err := c.GetCachedStruct(cacheKey, instance)
	if err == nil {
		return nil
	}

	session := c.engine.NewSession()
	defer session.Close()

	err = session.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	got, err := session.Get(instance)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = session.Commit(); err != nil {
		return errors.WithStack(err)
	}

	if !got {
		value := reflect.ValueOf(instance).Elem()
		dateType := value.Type()
		newData := reflect.New(dateType).Elem()
		value.Set(newData)
		return nil
	}

	_, err = c.CacheStruct(cacheKey, instance)
	if err != nil {
		return errors.WithStack(err)
	}

	expireTime := instance.CacheExpireTime()
	_, err = c.Expire(cacheKey, expireTime)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *CacheRepository) Gets(ids []string, helper entityHelper, instanceSlicePtr interface{}) error {
	instances := helper(ids)
	return c.gets(instances, instanceSlicePtr)
}

func (c *CacheRepository) GetsWithMoreColumns(ids []string, helper entityMoreHelper, instanceSlicePtr interface{}, extra ...string) error {
	instances := helper(ids, extra...)
	return c.gets(instances, instanceSlicePtr)
}

func (c *CacheRepository) gets(instances []Entity, instanceSlicePtr interface{}) error {
	instanceInterfaces := make([]interface{}, len(instances))
	for i, v := range instances {
		instanceInterfaces[i] = v
	}

	keys := make([]string, len(instances))
	for i, v := range instances {
		keys[i] = v.CacheKey()
	}

	unCachedIndex, err := c.MGetCachedStructs(keys, instanceInterfaces)
	if err != nil {
		return errors.WithStack(err)
	}

	if len(unCachedIndex) > 0 {
		session := c.engine.NewSession()
		defer session.Close()

		err = session.Begin()
		if err != nil {
			return errors.WithStack(err)
		}

		for _, idx := range unCachedIndex {
			unCachedInstance := instances[idx]
			got, err := session.Get(unCachedInstance)
			if err != nil {
				return errors.WithStack(err)
			}

			if got {
				unCachedInstanceKey := unCachedInstance.CacheKey()
				instanceInterfaces[idx] = unCachedInstance
				_, _ = c.CacheStruct(unCachedInstanceKey, unCachedInstance)
				_, _ = c.Expire(unCachedInstanceKey, unCachedInstance.CacheExpireTime())
			} else {
				instanceInterfaces[idx] = nil
			}
		}

		if err := session.Commit(); err != nil {
			return errors.WithStack(err)
		}
	}

	for _, v := range instanceInterfaces {
		if v != nil {
			slicePointerValue := reflect.ValueOf(instanceSlicePtr).Elem()
			slicePointerValue.Set(reflect.Append(slicePointerValue, reflect.ValueOf(v)))
		}
	}

	return nil
}

func (c *CacheRepository) GetsFlexibly(ids []string, helper entityHelper, queryFunction func(*xorm.Session, []string, int64, int64) ([]Entity, error), start, end int64, instanceSlicePtr interface{}) error {
	instances := helper(ids)
	instanceInterfaces := make([]interface{}, len(instances))
	for i, v := range instances {
		instanceInterfaces[i] = v
	}

	keys := make([]string, len(instances))
	for i, v := range instances {
		keys[i] = v.CacheKey()
	}

	unCachedIndex, err := c.MGetCachedStructs(keys, instanceInterfaces)
	if err != nil {
		return errors.WithStack(err)
	}

	if len(unCachedIndex) > 0 {
		unCachedIDs := make([]string, len(unCachedIndex))
		postitionMap := make(map[string]int)
		for i, idx := range unCachedIndex {
			unCachedIDs[i] = ids[idx]
			unCachedInstance := instances[idx]
			key := unCachedInstance.CacheKey()
			postitionMap[key] = idx

		}

		session := c.engine.NewSession()
		defer session.Close()

		ret, err := queryFunction(session, unCachedIDs, start, end)
		if err != nil {
			return errors.WithStack(err)
		}

		for _, v := range ret {
			key := v.CacheKey()
			pos := postitionMap[key]
			instanceInterfaces[pos] = v

			_, _ = c.CacheStruct(key, v)
			_, _ = c.Expire(key, v.CacheExpireTime())

			delete(postitionMap, key)
		}

		for _, idx := range postitionMap {
			instanceInterfaces[idx] = nil
		}

	}

	for _, v := range instanceInterfaces {
		if v != nil {
			slicePointerValue := reflect.ValueOf(instanceSlicePtr).Elem()
			slicePointerValue.Set(reflect.Append(slicePointerValue, reflect.ValueOf(v)))
		}
	}

	return nil
}
