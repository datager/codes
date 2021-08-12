package redis

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type ConnPool struct {
	conn            *redis.Pool
	server          string
	maxIdleConn     int
	connIdelTimeout time.Duration
	prefix          string
}

func initConnPool(server, password string, prefix string) *ConnPool {
	dialFunc := func() (redis.Conn, error) {
		var conn redis.Conn
		var err error
		if password != "" {
			passwordOpt := redis.DialPassword(password)
			conn, err = redis.Dial("tcp", server, passwordOpt)
		} else {
			conn, err = redis.Dial("tcp", server)
		}

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return conn, nil
	}

	pool := &redis.Pool{
		MaxIdle:     int(math.Pow(2, 30)),
		IdleTimeout: 1 * time.Second,
		Dial:        dialFunc,
	}

	connTest, err := dialFunc()
	if err != nil {
		panic(err)
	}
	defer connTest.Close()

	cp := &ConnPool{
		prefix: prefix,
	}

	cp.conn = pool
	cp.server = server
	cp.maxIdleConn = pool.MaxIdle
	cp.connIdelTimeout = pool.IdleTimeout

	return cp
}

func (r *ConnPool) Close() error {
	return r.conn.Close()
}

func (r *ConnPool) getKeys(key string) ([]string, error) {
	v, err := redis.Strings(r.do("KEYS", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetInt(key string) (int, error) {
	v, err := redis.Int(r.do("GET", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetInt64(key string) (int64, error) {
	v, err := redis.Int64(r.do("GET", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetString(key string) (string, error) {
	v, err := redis.String(r.do("GET", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetBytes(key string) ([]byte, error) {
	v, err := redis.Bytes(r.do("GET", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetBool(key string) (bool, error) {
	v, err := redis.Bool(r.do("GET", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) Set(key string, value interface{}, expireIn int64) (bool, error) {
	var args []interface{}
	if expireIn == 0 {
		args = append(args, r.compositeKey(key), value)
	} else {
		args = append(args, r.compositeKey(key), value, "EX", expireIn)
	}

	v, err := redis.String(r.do("SET", args...))
	if err != nil {
		return false, errors.WithStack(err)
	}

	var ret bool
	if v == "OK" {
		ret = true
	} else {
		ret = false
	}

	return ret, nil
}

func (r *ConnPool) Expire(key string, seconds int64) (bool, error) {
	v, err := redis.Bool(r.do("EXPIRE", r.compositeKey(key), seconds))
	return v, errors.WithStack(err)
}

func (r *ConnPool) TTL(key string) (int, error) {
	v, err := redis.Int(r.do("TTL", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) MExist(keys []string) (map[string]bool, error) {
	conn := r.conn.Get()
	defer conn.Close()

	ret := make(map[string]bool)
	for _, key := range keys {
		err := conn.Send("EXISTS", r.compositeKey(key))
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if err := conn.Flush(); err != nil {
		return nil, errors.WithStack(err)
	}

	for _, key := range keys {
		v, err := redis.Bool(conn.Receive())
		if err != nil {
			return nil, errors.WithStack(err)
		}

		ret[key] = v
	}

	return ret, nil
}

func (r *ConnPool) MGetInt64Slice(keys []string) (map[string]int64, error) {
	conn := r.conn.Get()
	defer conn.Close()

	rv := make(map[string]int64)
	var err error

	for _, key := range keys {
		err = conn.Send("GET", r.compositeKey(key))
		if err != nil {
			return rv, errors.WithStack(err)
		}
	}

	if err = conn.Flush(); err != nil {
		return rv, errors.WithStack(err)
	}

	for _, key := range keys {
		if v, err := redis.Int64(conn.Receive()); err == nil {
			rv[key] = v
		}
	}

	return rv, nil
}

func (r *ConnPool) Incr(key string, increment int) (int, error) {
	v, err := redis.Int(r.do("INCRBY", r.compositeKey(key), increment))
	return v, errors.WithStack(err)
}

func (r *ConnPool) Decr(key string, decrement int) (int, error) {
	if decrement > 0 {
		decrement = -decrement
	}

	v, err := redis.Int(r.do("INCRBY", r.compositeKey(key), decrement))
	return v, errors.WithStack(err)
}

func (r *ConnPool) Delete(key string) (int, error) {
	v, err := redis.Int(r.do("DEL", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) MDelete(keys []string) (int, error) {
	interfaceKeys := make([]interface{}, len(keys))
	for i, k := range keys {
		interfaceKeys[i] = r.compositeKey(k)
	}

	v, err := redis.Int(r.do("DEL", interfaceKeys...))
	return v, errors.WithStack(err)
}

func (r *ConnPool) FuzzyDelete(fuzzyKey string) (int, error) {
	keys, err := r.getKeys(fuzzyKey)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	interfaceKeys := make([]interface{}, len(keys))
	for i, k := range keys {
		interfaceKeys[i] = k
	}

	if len(interfaceKeys) == 0 {
		return 0, nil
	}

	v, err := redis.Int(r.do("DEL", interfaceKeys...))

	return v, errors.WithStack(err)
}

func (r *ConnPool) Exists(key string) (bool, error) {
	v, err := redis.Bool(r.do("EXISTS", r.compositeKey(key)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) CacheStruct(key string, valueStruct interface{}) (bool, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(valueStruct)
	if err != nil {
		return false, errors.WithStack(err)
	}

	v, err := redis.Int64(r.do("SETNX", r.compositeKey(key), buffer.String()))
	if err != nil {
		return false, errors.WithStack(err)
	}

	var ret bool
	if v == 1 {
		ret = true
	} else {
		return false, fmt.Errorf("key:%v is exist", key)
	}

	return ret, nil
}

func (r *ConnPool) GetCachedStruct(key string, retStruct interface{}) error {
	v, err := redis.Bytes(r.do("GET", r.compositeKey(key)))
	if err != nil {
		return errors.WithStack(err)
	}

	decoder := gob.NewDecoder(bytes.NewReader(v))
	if derr := decoder.Decode(retStruct); derr != nil {
		return errors.WithStack(derr)
	}

	return nil
}

func (r *ConnPool) CacheIDsWithExpire(key string, ids []string, expireIn int64) (bool, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ids)
	if err != nil {
		return false, errors.WithStack(err)
	}

	v, err := redis.Int64(r.do("SETNX", r.compositeKey(key), buffer.String()))
	if err != nil {
		return false, errors.WithStack(err)
	}

	_, err = r.Expire(key, expireIn)
	if err != nil {
		return false, errors.WithStack(err)
	}

	var ret bool
	if v == 1 {
		ret = true
	} else {
		return false, fmt.Errorf("Key:%v is exist", key)
	}

	return ret, nil
}

func (r *ConnPool) CacheIDs(key string, ids []string) (bool, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ids)
	if err != nil {
		return false, errors.WithStack(err)
	}

	v, err := redis.Int64(r.do("SETNX", r.compositeKey(key), buffer.String()))
	if err != nil {
		return false, errors.WithStack(err)
	}

	var ret bool
	if v == 1 {
		ret = true
	} else {
		return false, fmt.Errorf("key:%v is exist", key)
	}

	return ret, nil
}

func (r *ConnPool) GetCachedIDs(key string) ([]string, error) {
	var ids []string
	v, err := redis.Bytes(r.do("GET", r.compositeKey(key)))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	decoder := gob.NewDecoder(bytes.NewReader([]byte(v)))
	if err := decoder.Decode(&ids); err != nil {
		return nil, errors.WithStack(err)
	}

	return ids, nil
}

func (r *ConnPool) MGetCachedStructs(keys []string, valueStructs []interface{}) ([]int, error) {
	conn := r.conn.Get()
	defer conn.Close()

	size := len(keys)
	for _, key := range keys {
		err := conn.Send("GET", r.compositeKey(key))
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if err := conn.Flush(); err != nil {
		return nil, errors.WithStack(err)
	}

	unCachedIndex := make([]int, 0)

	for i := 0; i < size; i++ {
		s := valueStructs[i]
		v, err := redis.Bytes(conn.Receive())
		if err != nil {
			valueStructs[i] = nil
			unCachedIndex = append(unCachedIndex, i)
		} else {
			decoder := gob.NewDecoder(bytes.NewReader(v))
			err := decoder.Decode(s)
			if err != nil {
				valueStructs[i] = nil
				unCachedIndex = append(unCachedIndex, i)
			}
		}
	}

	return unCachedIndex, nil
}

// -----------------
// Set commands
// -----------------
// 向set中添加一个成员member
func (r *ConnPool) SAdd(setName string, member interface{}) (bool, error) {
	v, err := redis.Bool(r.do("SADD", r.compositeKey(setName), member))
	return v, errors.WithStack(err)
}

// 向set中添加一个或者多个成员
func (r *ConnPool) SAddMembers(setName string, members ...interface{}) (int, error) {
	conn := r.conn.Get()
	defer conn.Close()

	for _, v := range members {
		err := conn.Send("SADD", r.compositeKey(setName), v)
		if err != nil {
			return 0, errors.WithStack(err)
		}
	}

	if err := conn.Flush(); err != nil {
		return 0, errors.WithStack(err)
	}

	var ret int
	for i := 0; i < len(members); i++ {
		v, err := redis.Int(conn.Receive())
		if err != nil {
			return ret, errors.WithStack(err)
		}

		if v == 1 {
			ret++
		}
	}

	return ret, nil
}

// 获取set中成员的个数
func (r *ConnPool) SCard(setName string) (int, error) {
	v, err := redis.Int(r.do("SCARD", r.compositeKey(setName)))
	return v, errors.WithStack(err)
}

// 移除set中的元素member
func (r *ConnPool) SRem(setName string, member interface{}) (int, error) {
	v, err := redis.Int(r.do("SREM", r.compositeKey(setName), member))
	return v, errors.WithStack(err)
}

// 返回set中所有的成员
func (r *ConnPool) SMembersIntSlice(setName string) ([]int, error) {
	v, err := redis.Ints(r.do("SMEMBERS", r.compositeKey(setName)))
	return v, errors.WithStack(err)
}

// 返回set中所有的元素
func (r *ConnPool) SMembersStringSlice(setName string) ([]string, error) {
	v, err := redis.Strings(r.do("SMEMBERS", r.compositeKey(setName)))
	return v, errors.WithStack(err)
}

// 判断member是否在set中
func (r *ConnPool) SIsMember(setName string, member interface{}) (bool, error) {
	v, err := redis.Bool(r.do("SISMEMBER", r.compositeKey(setName), member))
	return v, errors.WithStack(err)
}

func (r *ConnPool) sRandMember(setName string, amount int) (interface{}, error) {
	v, err := r.do("SRANDMEMBER", setName, amount)
	return v, errors.WithStack(err)
}

// 返回set中随机的一个integer
func (r *ConnPool) SRandMemberInt(setName string) (int, error) {
	v, err := redis.Int(r.sRandMember(r.compositeKey(setName), 1))
	return v, errors.WithStack(err)
}

// 返回set中随机的多个integer
func (r *ConnPool) SRandMemberInts(setName string, amount int) ([]int, error) {
	v, err := redis.Ints(r.sRandMember(r.compositeKey(setName), amount))
	return v, errors.WithStack(err)
}

func (r *ConnPool) SRandMemberString(setName string) (string, error) {
	v, err := redis.String(r.sRandMember(r.compositeKey(setName), 1))
	return v, errors.WithStack(err)
}

func (r *ConnPool) SRandMemberStrings(setName string, amount int) ([]string, error) {
	v, err := redis.Strings(r.sRandMember(r.compositeKey(setName), amount))
	return v, errors.WithStack(err)
}

// -----------------
// Hash commands
// -----------------
func (r *ConnPool) HSet(hashName string, fieldName string, value interface{}) (int, error) {
	v, err := redis.Int(r.do("HSET", r.compositeKey(hashName), fieldName, value))
	return v, errors.WithStack(err)
}

func (r *ConnPool) HGetInt(hashName string, fieldName string) (int, error) {
	v, err := redis.Int(r.do("HGET", r.compositeKey(hashName), fieldName))
	return v, errors.WithStack(err)
}

func (r *ConnPool) HGetString(hashName, fielldName string) (string, error) {
	v, err := redis.String(r.do("HGET", r.compositeKey(hashName), fielldName))
	return v, errors.WithStack(err)
}

func (r *ConnPool) HDel(hashName, fieldName string) (int, error) {
	v, err := redis.Int(r.do("HDEL", r.compositeKey(hashName), fieldName))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetCachedIntMap(hashName string) (map[string]int, error) {
	v, err := redis.IntMap(r.do("HGETALL", r.compositeKey(hashName)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetCachedInt64Map(hashName string) (map[string]int64, error) {
	v, err := redis.Int64Map(r.do("HGETALL", r.compositeKey(hashName)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) GetCachedStringMap(hashName string) (map[string]string, error) {
	v, err := redis.StringMap(r.do("HGETALL", r.compositeKey(hashName)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) HExists(hashName string, fieldName string) (bool, error) {
	v, err := redis.Bool(r.do("HEXISTS", r.compositeKey(hashName), fieldName))
	return v, errors.WithStack(err)
}

func (r *ConnPool) HIncrBy(hashName, fieldName string, increment int) (int, error) {
	v, err := redis.Int(r.do("HINCRBY", r.compositeKey(hashName), fieldName, increment))
	return v, errors.WithStack(err)
}

func (r *ConnPool) HKeys(hashName string) ([]interface{}, error) {
	v, err := redis.Values(r.do("HKEYS", r.compositeKey(hashName)))
	return v, errors.WithStack(err)
}

// --------------------
// Sorted set commands
// --------------------

// 返回有序集中元素的个数
func (r *ConnPool) ZCount(setName string) (int, error) {
	v, err := redis.Int(r.do("ZCOUNT", r.compositeKey(setName), "-inf", "+inf"))
	return v, errors.WithStack(err)
}

// 向有序集中添加元素或者更新已存在成员的分数
func (r *ConnPool) ZAdd(setName string, score interface{}, member interface{}) (bool, error) {
	v, err := redis.Bool(r.do("ZADD", r.compositeKey(setName), score, member))
	return v, errors.WithStack(err)
}

// 移除成员member
func (r *ConnPool) ZRem(setName string, member interface{}) (int, error) {
	v, err := redis.Int(r.do("ZREM", r.compositeKey(setName), member))
	return v, errors.WithStack(err)
}

// 移除指定排名区间的内的元素
func (r *ConnPool) ZRemRangeByRank(setName string, start, stop int) (int, error) {
	v, err := redis.Int(r.do("ZREMRANGEBYRANK", r.compositeKey(setName), start, stop))
	return v, errors.WithStack(err)
}

func (r *ConnPool) ZAddMembers(setName string, memberScoreMap map[string]int) (int, error) {
	conn := r.conn.Get()
	defer conn.Close()

	for member, score := range memberScoreMap {
		err := conn.Send("ZADD", r.compositeKey(setName), score, member)
		if err != nil {
			return 0, errors.WithStack(err)
		}
	}

	if err := conn.Flush(); err != nil {
		return 0, errors.WithStack(err)
	}

	var ret int
	for i := 0; i < len(memberScoreMap); i++ {
		v, err := conn.Receive()
		if err != nil {
			return 0, errors.WithStack(err)
		}

		if v == 1 {
			ret++
		}
	}

	return ret, nil
}

// 检查元素是否存在
func (r *ConnPool) ZCheckMember(setName string, member interface{}) (bool, error) {
	ret, err := r.do("ZSCORE", r.compositeKey(setName), member)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if ret == nil {
		return false, nil
	}

	return true, nil
}

// 返回有序集中元素的分数
func (r *ConnPool) ZScore(setName string, member string) (float64, error) {
	v, err := redis.Float64(r.do("ZSCORE", r.compositeKey(setName), member))
	return v, errors.WithStack(err)
}

// 返回有序集中区间内的元素，并按照分数排序
func (r *ConnPool) ZRangeInts(setName string, start int, end int, desc bool) ([]int, error) {
	var v []int
	var err error
	if !desc {
		v, err = redis.Ints(r.do("ZRANGE", r.compositeKey(setName), start, end))
	} else {
		v, err = redis.Ints(r.do("ZREVRANGE", r.compositeKey(setName), start, end))
	}
	return v, errors.WithStack(err)
}

// 返回有序集中区间内的元素，并按照分数排序
func (r *ConnPool) ZRangeStrings(setName string, start int, end int, desc bool) ([]string, error) {
	var v []string
	var err error
	if !desc {
		v, err = redis.Strings(r.do("ZRANGE", r.compositeKey(setName), start, end))
	} else {
		v, err = redis.Strings(r.do("ZREVRANGE", r.compositeKey(setName), start, end))
	}
	return v, errors.WithStack(err)
}

// 返回有序集中区间内的元素和分数，并按照分数排序
func (r *ConnPool) ZRangeWithScoreStringInt(setName string, start int, end int, desc bool) (map[string]int, error) {
	var vs map[string]int
	var err error
	if !desc {
		vs, err = redis.IntMap(r.do("ZRANGE", r.compositeKey(setName), start, end, "WITHSCORES"))
	} else {
		vs, err = redis.IntMap(r.do("ZREVRANGE", r.compositeKey(setName), start, end, "WITHSCORES"))
	}
	return vs, errors.WithStack(err)
}

func (r *ConnPool) ZRangeWithScoreStringInt64(setName string, start int, end int, desc bool) (map[string]int64, error) {
	var vs map[string]int64
	var err error
	if !desc {
		vs, err = redis.Int64Map(r.do("ZRANGE", r.compositeKey(setName), start, end, "WITHSCORES"))
	} else {
		vs, err = redis.Int64Map(r.do("ZREVRANGE", r.compositeKey(setName), start, end, "WITHSCORES"))
	}
	return vs, errors.WithStack(err)
}

func (r *ConnPool) ZRangeWithScoreIntInt(setName string, start int, end int, desc bool) (map[int]int, error) {
	vs, err := r.ZRangeWithScoreStringInt(r.compositeKey(setName), start, end, desc)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret = make(map[int]int)
	for k, v := range vs {
		ik, err := toInt(k)
		if err != nil {
			return nil, err
		}
		ret[ik] = v
	}
	return ret, errors.WithStack(err)
}

func (r *ConnPool) ZRangeWithScoreIntInt64(setName string, start int, end int, desc bool) (map[int]int64, error) {
	vs, err := r.ZRangeWithScoreStringInt64(r.compositeKey(setName), start, end, desc)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret = make(map[int]int64)
	for k, v := range vs {
		ik, err := toInt(k)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret[ik] = v
	}
	return ret, errors.WithStack(err)
}

func (r *ConnPool) ZIncrBy(setName string, count interface{}, member interface{}) (float64, error) {
	v, err := redis.Float64(r.do("ZINCRBY", r.compositeKey(setName), count, member))
	return v, errors.WithStack(err)
}

// ZRank return member's rank
func (r *ConnPool) ZRank(setName string, member interface{}) (int, error) {
	v, err := redis.Int(r.do("ZRANK", r.compositeKey(setName), member))
	return v, errors.WithStack(err)
}

// ZRevRank return member's rank, desc
func (r *ConnPool) ZRevRank(setName string, member interface{}) (int, error) {
	v, err := redis.Int(r.do("ZREVRANK", r.compositeKey(setName), member))
	return v, errors.WithStack(err)
}

// ------------------------
// List Commands
// ------------------------

func (r *ConnPool) LPush(listName string, members ...interface{}) (int, error) {
	var len int
	for _, member := range members {
		v, err := redis.Int(r.do("LPUSH", r.compositeKey(listName), member))
		if err != nil {
			return v, errors.WithStack(err)
		}
		len = v
	}
	return len, nil
}

func (r *ConnPool) LPop(listName string) (string, error) {
	v, err := redis.String(r.do("LPOP", r.compositeKey(listName)))
	return v, errors.WithStack(err)
}

// RPush return list length
func (r *ConnPool) RPush(listName string, members ...interface{}) (int, error) {
	var len int
	for _, member := range members {
		v, err := redis.Int(r.do("RPUSH", r.compositeKey(listName), member))
		if err != nil {
			return v, errors.WithStack(err)
		}
		len = v
	}
	return len, nil
}

func (r *ConnPool) RPop(listName string) (string, error) {
	v, err := redis.String(r.do("RPOP", r.compositeKey(listName)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) LBeforeInsert(listName string, originValue, targetValue interface{}) (int, error) {
	v, err := redis.Int(r.do("LINSERT", r.compositeKey(listName), "BEFORE", originValue, targetValue))
	return v, errors.WithStack(err)
}

func (r *ConnPool) LAfterInsert(listName string, originValue, targetValue interface{}) (int, error) {
	v, err := redis.Int(r.do("LINSERT", r.compositeKey(listName), "AFTER", originValue, targetValue))
	return v, errors.WithStack(err)
}

func (r *ConnPool) RLen(listName string) (int, error) {
	v, err := redis.Int(r.do("LLEN", r.compositeKey(listName)))
	return v, errors.WithStack(err)
}

func (r *ConnPool) LIdex(listName string, index int) (string, error) {
	v, err := redis.String(r.do("LINDEX", r.compositeKey(listName), index))
	return v, errors.WithStack(err)
}

func (r *ConnPool) LRem(listName string, count int, member interface{}) (int, error) {
	v, err := redis.Int(r.do("LREM", r.compositeKey(listName), count, member))
	return v, errors.WithStack(err)
}

func (r *ConnPool) LRange(listName string, start, stop int) ([]string, error) {
	v, err := redis.Strings(r.do("LRANGE", r.compositeKey(listName), start, stop))
	return v, errors.WithStack(err)
}

func (r *ConnPool) LRangeInts(listName string, start, stop int) ([]int, error) {
	v, err := redis.Ints(r.do("LRANGE", r.compositeKey(listName), start, stop))
	return v, errors.WithStack(err)
}

func (r *ConnPool) do(command string, args ...interface{}) (reply interface{}, err error) {
	conn := r.conn.Get()
	defer conn.Close()

	reply, err = conn.Do(command, args...)
	return
}

func (r *ConnPool) compositeKey(key string) string {
	return fmt.Sprintf("%v%v", r.prefix, key)
}

func toInt(origin interface{}) (int, error) {
	switch typed := origin.(type) {
	case int64:
		x := int(typed)
		if int64(x) != typed {
			return 0, strconv.ErrRange
		}
		return x, nil
	case []byte:
		x, err := strconv.Atoi(string(typed))
		return x, err
	case string:
		return strconv.Atoi(typed)
	case nil:
		return 0, errors.New("Type Converter: nil returned")
	}
	return 0, fmt.Errorf("Type Converter: unexpected type for int, got type %T", origin)
}
