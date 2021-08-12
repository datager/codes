package wolf

import (
	"codes/gocodes/dg/service/org"
	"codes/gocodes/dg/service/sensor"
)

func CreateOrgSensor() {
	// 最终111个org, 11000个sensor
	orgLevel := 3    // 3
	orgNum := 10     // 10
	sensorNum := 100 // 100

	// select min(sensor_int_id), max(sensor_int_id) from sensors;
	//x, err := dbengine.GetV5Instance().Exec(" ALTER SEQUENCE sensor_id_seq RESTART WITH 1; ") // 空库时可用
	//if err != nil {
	//	glog.Fatal(x, err)
	//}

	org.BatchCreate(orgLevel, orgNum)
	sensor.BatchCreate(orgLevel, orgNum, sensorNum)
}
