package wolf

import (
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/uuid"
	"fmt"
	"github.com/golang/glog"
	"time"
)

func InsertIntoMonitorRuleSensorRelation() {
	mrsrsToBeInsert := make([]*entities.MonitorRuleSensorRelation, 0)
	for i := 0; i < 1; i++ {
		mrsrsToBeInsert = append(mrsrsToBeInsert,
			&entities.MonitorRuleSensorRelation{
				Uts:           time.Now(),
				Ts:            utils.GetNowTs(),
				MonitorRuleID: uuid.NewV4().String(),
				SensorID:      uuid.NewV4().String(),
			},
		)
	}

	//st := time.Now()
	//cnt, err := dbengine.GetV5Instance().Table(entities.TableNameMonitorRuleSensorRelation).InsertMulti(mrsrsToBeInsert)
	//cost := time.Since(st)
	//glog.Info(cost)
	//if err != nil {
	//	panic(err)
	//}
	//glog.Info(cnt)

	dataSQL := ""
	for i := range mrsrsToBeInsert {
		dataSQL += fmt.Sprintf("(%v,'%v','%v')", mrsrsToBeInsert[i].Ts, mrsrsToBeInsert[i].MonitorRuleID, mrsrsToBeInsert[i].SensorID)
		if len(mrsrsToBeInsert) > 1 && i < len(mrsrsToBeInsert)-1 {
			dataSQL += ","
		}
	}

	sql := fmt.Sprintf("INSERT INTO monitor_rule_sensor_relation (ts, monitor_rule_id, sensor_id) VALUES %v", dataSQL)

	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()

	st := time.Now()
	sqlRet, err := ss.Exec(sql)
	cost := time.Since(st)
	glog.Info(cost)
	if err != nil {
		panic(err)
	}

	rows, err := sqlRet.RowsAffected()
	if err != nil {
		panic(err)
	}
	glog.Info("aaa", rows)
}
