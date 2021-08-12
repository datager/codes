package db

import (
	"codes/gocodes/dg/configs"
	"codes/gocodes/dg/dbengine"

	"github.com/golang/glog"
)

// 用于测试pg或gauss等db是否能连接的小工具

func CheckDBConnectionAvailable() {
	glog.Infoln("start")
	//flag.Parse()

	glog.Infoln("---开始测试deepface_v5URL是否可连接, 是否可查询---")
	v5ConnStr := configs.GetInstance().DBServiceConfig.V5ConnStr
	glog.Infof("[deepface_v5] read from config is %v", v5ConnStr)
	engineV5 := dbengine.CreateEngine(v5ConnStr)
	sensorNames := make([]string, 0)
	if err := engineV5.SQL("select sensor_name from sensors limit 5").Find(&sensorNames); err != nil {
		panic(err)
	}
	glog.Infof("[deepface_v5]successfully query: result is %v", sensorNames)
	glog.Infoln("---测试完毕---")

	glog.Infoln()
	glog.Infoln()
	glog.Infoln()
	glog.Infoln()
	glog.Infoln()

	glog.Infoln("---开始测试deepface_v5URL是否可连接, 是否可查询---")
	specialConnStr := configs.GetInstance().DBServiceConfig.SpecialConnStr
	glog.Infof("[special] read from config is %v", v5ConnStr)
	engineSpecial := dbengine.CreateEngine(specialConnStr)
	taskIDs := make([]string, 0)
	if err := engineSpecial.SQL("select task_id from tasks").Find(&taskIDs); err != nil {
		panic(err)
	}
	glog.Infof("[special] successfully query: result is %v", taskIDs)
	glog.Infoln("---测试完毕---")
}
