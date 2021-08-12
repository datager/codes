package dbengine

import (
	"sync"
	"time"

	"codes/gocodes/dg/configs"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var (
	engineOnceV5         sync.Once
	engineOnceStatistics sync.Once
	engineOnceSpecial    sync.Once

	engineV5         *xorm.Engine
	engineStatistics *xorm.Engine
	EngineSpecial    *xorm.Engine
)

func GetV5Instance() *xorm.Engine {
	engineOnceV5.Do(func() {
		engineV5 = CreateEngine(configs.GetInstance().DBServiceConfig.V5ConnStr)
	})
	return engineV5
}

func GetStatisticsInstance() *xorm.Engine {
	engineOnceStatistics.Do(func() {
		engineStatistics = CreateEngine(configs.GetInstance().DBServiceConfig.StatisticsConnStr)
	})
	return engineStatistics
}

func GetSpecialInstance() *xorm.Engine {
	engineOnceSpecial.Do(func() {
		EngineSpecial = CreateEngine(configs.GetInstance().DBServiceConfig.SpecialConnStr)
	})
	return EngineSpecial
}

func CreateEngine(connStr string) *xorm.Engine {
	e, err := xorm.NewEngine("postgres", connStr) // v5/statistics/special
	if err != nil {
		glog.Fatal(err)
	}

	// conn
	e.SetMaxOpenConns(configs.GetInstance().DBServiceConfig.MaxOpenConnsPerClient)
	e.SetMaxIdleConns(configs.GetInstance().DBServiceConfig.MaxIdleConnsPerClient)
	if configs.GetInstance().DBServiceConfig.MaxLifeTimeSeconds != 0 {
		e.SetConnMaxLifetime(time.Duration(configs.GetInstance().DBServiceConfig.MaxLifeTimeSeconds) * time.Second)
	}

	// log
	//e.ShowSQL(configs.GetInstance().DBServiceConfig.IsShowSQL)
	e.ShowSQL(true)
	//e.ShowExecTime(true)
	//e.SetLogLevel(core.LOG_INFO)

	glog.Info("ping start")
	err = e.Ping()
	if err != nil {
		glog.Error("ping error")
		glog.Fatal(err)
	}
	glog.Info("ping ok")
	return e
}
