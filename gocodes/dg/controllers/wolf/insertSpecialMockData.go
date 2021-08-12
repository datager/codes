package wolf

import (
	"fmt"

	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/slice"

	"github.com/golang/glog"
)

type VehicleSpecialMockData struct {
	CutboardURL string // 自己找一个抽烟url
	PlateText   string // 自己找一个抽烟plateText
}

// 插入车辆专题模拟数据
// taskID: 0001~0006
func InsertVehicleSpecialMockData(taskID string) {
	sensorID := "10efab97-8355-4b46-9305-e461aa0a5216" // 自己找一个
	mockVehicle := make([]*VehicleSpecialMockData, 0)
	mockVehicle = append(mockVehicle,
		&VehicleSpecialMockData{"http://192.168.2.118:8501/api/v2/file/4/21de5ff5786d3d", "川A2XU75"},
		&VehicleSpecialMockData{"http://192.168.2.118:8501/api/v2/file/1/21de026bf047f7", "辽K75807"},
	)
	toChangeVehicleIDs := findVehicleIDs(taskID, sensorID, len(mockVehicle))
	insertCapIDsIntoSpecial(taskID, sensorID, toChangeVehicleIDs, mockVehicle)
	updateVehicleIDsFromV5(toChangeVehicleIDs, mockVehicle)
}

func findVehicleIDs(taskID, sensorID string, limit int) []string {
	existCaptureIDs := make([]string, 0)
	sqlExist := fmt.Sprintf("select capture_id from %v", GetVehicleSpecialRelationTableName(taskID))
	err := dbengine.GetSpecialInstance().SQL(sqlExist).Find(&existCaptureIDs)
	if err != nil {
		panic("find existCaptureIDs")
	}

	vehicleIDs := make([]string, 0) // deepface_v5库的抓拍表里有, special库里没有的vehicle_id
	sqlFindID := fmt.Sprintf("select vehicle_id from vehicle_capture_index where sensor_id = '%v' and (vehicle_id not in (%v)) limit '%v' order by ts desc", sensorID, slice.FormatIDsByAddComma(slice.FormatIDsByAddQuote(existCaptureIDs)), limit)
	err = dbengine.GetV5Instance().SQL(sqlFindID).Find(&vehicleIDs)
	if err != nil {
		panic("find vehicle_id")
	}
	glog.Info("ok")
	return vehicleIDs
}

func insertCapIDsIntoSpecial(taskID, sensorID string, capIDs []string, data []*VehicleSpecialMockData) {
	for i := range data {
		sqlCap := fmt.Sprintf("insert into %v (capture_id, plate_text, sensor_id, ts) values ('%v','%v','%v','%v')", GetVehicleSpecialCaptureTableName(taskID), capIDs[i], data[i].PlateText, sensorID, utils.GetNowTs())
		_, err := dbengine.GetSpecialInstance().Exec(sqlCap)
		if err != nil {
			panic("sqlCap")
		}

		sqlRelation := fmt.Sprintf("insert into %v (task_id, capture_id) values ('%v',' %v')", GetVehicleSpecialRelationTableName(taskID), taskID, capIDs[i])
		_, err = dbengine.GetSpecialInstance().Exec(sqlRelation)
		if err != nil {
			panic("sqlRelation")
		}

		sqlSpecialCount := fmt.Sprintf("update detail_type_special_count set count = count +1 where (detail_type = '10000' and special_task_id = '%v')", taskID)
		_, err = dbengine.GetStatisticsInstance().Exec(sqlSpecialCount)
		if err != nil {
			panic("sqlSpecialCount")
		}
	}
}

func updateVehicleIDsFromV5(capIDs []string, data []*VehicleSpecialMockData) {
	for i := range data {
		sqlCapIndex := fmt.Sprintf("update vehicle_capture_index set cutboard_image_uri = '%v', plate_text = '%v' where vehicle_id = '%v'", data[i].CutboardURL, data[i].PlateText, capIDs[i])
		_, err := dbengine.GetV5Instance().Exec(sqlCapIndex)
		if err != nil {
			panic("sqlCapIndex")
		}

		sqlCap := fmt.Sprintf("update vehicle_capture set cutboard_image_uri = '%v', image_uri = '%v', plate_text = '%v' where vehicle_id = '%v'", data[i].CutboardURL, data[i].CutboardURL, data[i].PlateText, capIDs[i])
		_, err = dbengine.GetV5Instance().Exec(sqlCap)
		if err != nil {
			panic("sqlSpecialCount")
		}
	}
}

func GetVehicleSpecialRelationTableName(taskID string) string {
	switch taskID {
	case "0001":
		return "task_unbelted_captures_relation"
	case "0002":
		return "task_driving_smoking_captures_relation"
	case "0003":
		return "task_cover_face_captures_relation"
	case "0004":
		return "task_driving_smoking_captures_relation"
	case "0005":
		return "task_chemicals_codriver_empty_captures_relation"
	case "0006":
		return "task_nonmotor_with_people_captures_relation"
	default:
		panic("GetVehicleSpecialRelationTableName bad case")
	}
}

func GetVehicleSpecialCaptureTableName(taskID string) string {
	switch taskID {
	case "0001":
		return entities.UnbeltedTableName
	case "0002":
		return entities.DriveCallTableName
	case "0003":
		return entities.CoverFaceTableName
	case "0004":
		return entities.DriveSmokingTableName
	case "0005":
		return entities.ChemicalsCodriverEmptyTableName
	case "0006":
		return entities.NonmotorWithPeopleTableName
	default:
		panic("GetVehicleSpecialCaptureTableName bad case")
	}
}
