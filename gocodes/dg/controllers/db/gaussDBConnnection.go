package db

import (
	"codes/gocodes/dg/configs"
	"codes/gocodes/dg/dbengine"
	"fmt"
	"github.com/golang/glog"
)

func ConnToV5() {
	connStr := configs.GetInstance().DBServiceConfig.V5ConnStr
	xV5 := dbengine.CreateEngine(connStr)

	queryStr := fmt.Sprintf("select count(*) from %v where ts <= %v", configs.GetInstance().DBServiceConfig.GaussTestTableName, configs.GetInstance().DBServiceConfig.GaussTestTs)
	glog.Info(queryStr)
	ret, err := xV5.Query(queryStr)
	if err != nil {
		glog.Fatal(err)
	}

	if len(ret) == 0 {
		glog.Fatal("len(ret) = 0")
	}

	if len(ret[0]) == 0 {
		glog.Fatal("len(ret[0]) = 0")
	}

	glog.Infoln("ok", string(ret[0]["count"]))
}

/*
	v := entities.VehicleCaptureIndex{
		VehicleID:        "1",
		VehicleReID:      "2",
		VehicleVID:       "3",
		RelationTypes:    nil,
		Ts:               1585724696000,
		SensorID:         "4",
		SensorIntID:      5,
		OrgCode:          6,
		Speed:            7,
		Direction:        8,
		CutboardImageURI: "http://9",
		BrandID:          10,
		SubBrandID:       11,
		ModelYearID:      12,
		TypeID:           13,
		Side:             14,
		ColorID:          15,
		PlateText:        "沪A16",
		PlateTypeID:      17,
		PlateColorID:     18,
	}
	_ = v

	//ret, err := xV5.Table("vehicle_capture_index").Insert(&v)
	//if err != nil {
	//	glog.Fatal(err)
	//}
	//glog.Infoln("ok", ret)
*/
