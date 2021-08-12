package db

import (
	"fmt"
	"math"
	"sync"
	"time"

	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/log"
	"codes/gocodes/dg/utils/random"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"xorm.io/xorm"
)

func handleMySQL() {
	// 连接sdb
	//e, err := xorm.NewEngine("mysql", "root:12345678@tcp(127.0.0.1:3306)/db?charset=utf8")
	e, err := xorm.NewEngine("mysql", "root:@tcp(192.168.2.48:3306)/db?charset=utf8")
	if err != nil {
		panic(err)
	}
	e.SetMaxOpenConns(10)
	e.SetMaxIdleConns(10)
	e.ShowSQL(false)
	err = e.Ping()
	if err != nil {
		panic(err)
	}
	log.Info("connected")

	// 定义结构
	type TB struct {
		Id      int64
		Num     int64
		Namestr string
		Timestr string
		Decstr  string
	}
	tableName := "tb"
	//
	//// 造数据
	const lenPerMonth = 10000
	// 共12个月
	for month := 1; month <= 12; month++ {
		// 每个月lenPerMonth条数据
		records := make([]TB, lenPerMonth, lenPerMonth)
		for i := 0; i < lenPerMonth; i++ {
			// 生成闭区间内随机时间
			start := time.Date(2016, time.Month(month), 1, 0, 0, 0, 0, time.Local)
			end := start.AddDate(0, 1, 0)
			randomTimestamp := random.RandInt64(utils.TimeToTs(start), utils.TimeToTs(end))
			randomTimeStr := utils.TsToTime(randomTimestamp).Format("20160101")
			// 组装1个record
			id := random.RandInt64(0, math.MaxInt32)
			records[i] = TB{
				Id:      id,
				Num:     random.RandInt64(0, math.MaxInt32),
				Namestr: fmt.Sprintf("test%v", id),
				Timestr: randomTimeStr,
				Decstr:  random.RandomString(10),
			}
		}
		inserted, err := e.Table(tableName).Insert(records)
		if err != nil {
			panic(err)
		}
		log.Infof("month %v, inserted %v", month, inserted)
	}

	// 查数据
	const threads = 10
	wg := sync.WaitGroup{}
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		i := i
		go func() {
			session := e.NewSession()
			defer session.Close()
			defer wg.Done()

			st := time.Now()
			ret := TB{}
			_, err := session.Table(tableName).Select("*").Limit(1).Get(&ret)
			if err != nil {
				glog.Error(err)
			}
			glog.Infof("[thread%v] %+v, took%v", i, ret, time.Since(st))
		}()
	}

	glog.Info("[wait] start")
	wg.Wait()
	glog.Info("[wait] end")

	glog.Info("all done")
}
