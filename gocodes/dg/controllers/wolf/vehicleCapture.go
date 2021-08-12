package wolf

import (
	"codes/gocodes/dg/configs"
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/service/sensor"
	"codes/gocodes/dg/service/vehicle"
	"time"

	"github.com/golang/glog"
)

// kind: captureTableKind
func InsertCaptureIndexTable(kind int) {
	minSIntID := configs.GetInstance().CaptureConfig.Sensor.MinSensorIntID
	maxSIntID := configs.GetInstance().CaptureConfig.Sensor.MaxSensorIntID

	totalSt := time.Now()
	for sIntID := minSIntID; sIntID <= maxSIntID; sIntID++ {
		s, err := sensor.GetBySensorIntID(sIntID)
		if err != nil {
			glog.Error(err)
			continue
		}

		st := time.Now()

		// insert
		switch kind {
		case models.DetTypeVehicle:
			vehicle.InsertVehicleCaptureIndexOfOneSensor(s.SensorID, s.SensorIntID, s.OrgCode)
		default:
			glog.Error("captureTableKind wrong")
		}

		glog.Infof("sIntID:%v write captureNumToInsertOfOneSensor:%v done, took:%vms, totalTook:%vms",
			sIntID, configs.GetInstance().CaptureConfig.CaptureNumToInsertOfOneSensor, time.Now().Sub(st), time.Now().Sub(totalSt))
		st = time.Now()
	}
}
