package entities

const PeopleTacticRulesResultTableName = "people_tactic_rules_result"

func (PeopleTacticRulesResultEntity) TableName() string {
	return PeopleTacticRulesResultTableName
}

type PeopleTacticRulesResultEntity struct {
	ID                 int64   `xorm:"id serial"`
	EventID            string  `xorm:"event_id"`
	EventReID          string  `xorm:"event_reid"`
	TacticRuleID       string  `xorm:"tactic_rule_id"`
	Timestamp          int64   `xorm:"ts"`
	HitSignCode        string  `xorm:"hit_sign_code"`
	UserID             string  `xorm:"user_id"`
	SensorID           string  `xorm:"sensor_id"`
	SensorName         string  `xorm:"sensor_name"`
	FaceRepoID         string  `xorm:"face_repo_id"`
	FaceRepoName       string  `xorm:"face_repo_name"`
	FaceID             string  `xorm:"face_id"`
	FaceReID           string  `xorm:"face_reid"`
	CivilAttrName      string  `xorm:"civil_attr_name"`
	CivilAttrIdType    int32   `xorm:"civil_attr_id_type"`
	CivilAttrIdNo      string  `xorm:"civil_attr_id_no"`
	CivilAttrId        string  `xorm:"civil_attr_id"`
	CapturedImageUri   string  `xorm:"captured_image_uri"`
	ThumbnailImageUri  string  `xorm:"thumbnail_image_uri"`
	CutboardImageUri   string  `xorm:"cutboard_image_uri"`
	CutboardX          int32   `xorm:"cutboard_x"`
	CutboardY          int32   `xorm:"cutboard_y"`
	CutboardWidth      int32   `xorm:"cutboard_width"`
	CutboardHeight     int32   `xorm:"cutboard_height"`
	CutboardResWidth   int32   `xorm:"cutboard_res_width"`
	CutboardResHeight  int32   `xorm:"cutboard_res_height"`
	RegisteredImageUri string  `xorm:"registered_image_uri"`
	Confidence         float32 `xorm:"confidence"`
}
type PeopleTacticRulesResultSensorCountEntity struct {
	SensorId string `xorm:"sensor_id"`
	Count    int64  `xorm:"count"`
}

//func (e *PeopleTacticRulesResultEntity) ToModel() *models.TacticRulesResult {
//	civilAttr := &models.CivilAttr{
//		Name:        e.CivilAttrName,
//		IdType:      e.CivilAttrIdType,
//		IdNo:        e.CivilAttrIdNo,
//		CivilAttrId: e.CivilAttrIdNo,
//	}
//	capturedImage := &models.ImageResult{
//		ImageUri:          e.CapturedImageUri,
//		ThumbnailImageUri: e.ThumbnailImageUri,
//		CutboardImageUri:  e.CutboardImageUri,
//		CutboardX:         e.CutboardX,
//		CutboardY:         e.CutboardY,
//		CutboardWidth:     e.CutboardWidth,
//		CutboardHeight:    e.CutboardHeight,
//		CutboardResHeight: e.CutboardResHeight,
//		CutboardResWidth:  e.CutboardResWidth,
//	}
//	registeredImage := &models.ImageResult{
//		ImageUri: e.RegisteredImageUri,
//	}
//
//	return &models.TacticRulesResult{
//		EventID:         e.EventID,
//		EventReID:       e.EventReID,
//		Timestamp:       e.Timestamp,
//		TacticRuleID:    e.TacticRuleID,
//		SensorID:        e.SensorID,
//		SensorName:      e.SensorName,
//		FaceRepoID:      e.FaceRepoID,
//		FaceRepoName:    e.FaceRepoName,
//		FaceID:          e.FaceID,
//		Confidence:      e.Confidence,
//		UserID:          e.UserID,
//		CivilAttr:       civilAttr,
//		CapturedImage:   capturedImage,
//		RegisteredImage: registeredImage,
//	}
//
//}
//func TacticRulesResultKafkaInputDataToEntity(data *models.TacticRulesResultKafkaInputData) *PeopleTacticRulesResultEntity {
//	m := data.Data
//	entity := &PeopleTacticRulesResultEntity{
//		EventID:            m.EventId,
//		EventReID:          m.EventReId,
//		TacticRuleID:       data.RuleID,
//		HitSignCode:        data.HitSignCode,
//		Timestamp:          m.Timestamp,
//		SensorID:           m.SensorId,
//		SensorName:         m.SensorName,
//		FaceID:             m.FaceId,
//		FaceRepoName:       m.FaceRepoName,
//		FaceRepoID:         m.FaceReId,
//		FaceReID:           m.FaceReId,
//		CivilAttrId:        m.CivilAttr.CivilAttrId,
//		CivilAttrName:      m.CivilAttr.Name,
//		CivilAttrIdNo:      m.CivilAttr.IdNo,
//		CivilAttrIdType:    m.CivilAttr.IdType,
//		CapturedImageUri:   m.CapturedImage.ImageUri,
//		ThumbnailImageUri:  m.CapturedImage.ThumbnailImageUri,
//		CutboardImageUri:   m.CapturedImage.CutboardImageUri,
//		CutboardX:          m.CapturedImage.CutboardX,
//		CutboardY:          m.CapturedImage.CutboardY,
//		CutboardWidth:      m.CapturedImage.CutboardWidth,
//		CutboardHeight:     m.CapturedImage.CutboardHeight,
//		CutboardResWidth:   m.CapturedImage.CutboardResWidth,
//		CutboardResHeight:  m.CapturedImage.CutboardResHeight,
//		RegisteredImageUri: m.RegisteredImage.ImageUri,
//		Confidence:         m.Confidence,
//	}
//	return entity
//}
