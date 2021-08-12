package entities

const PeopleTacticRulesResultIndexTableName = "people_tactic_rules_result_index"

func (PeopleTacticRulesResultIndexEntity) TableName() string {
	return PeopleTacticRulesResultIndexTableName
}

type PeopleTacticRulesResultIndexEntity struct {
	Id                 string  `xorm:"id  not null pk VARCHAR(1024)"`
	EventId            string  `xorm:"event_id"`
	EventReId          string  `xorm:"event_reid"`
	TacticRuleID       string  `xorm:"tactic_rule_id"`
	CreateTime         int64   `xorm:"create_time"`
	Timestamp          int64   `xorm:"ts"`
	HitSignCode        string  `xorm:"hit_sign_code"`
	UserID             string  `xorm:"user_id"`
	SensorId           string  `xorm:"sensor_id"`
	SensorName         string  `xorm:"sensor_name"`
	FaceRepoId         string  `xorm:"face_repo_id"`
	FaceRepoName       string  `xorm:"face_repo_name"`
	FaceId             string  `xorm:"face_id"`
	FaceReId           string  `xorm:"face_reid"`
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
	IsFollow           int     `xorm:"is_follow"` //是否关注
	Status             int     `xorm:"status"`    // 待确认，排除，已确认
	Count              int64   `xorm:"count"`     //计数
}

//func TacticRulesResultKafkaInputDataToIndexEntity(data *models.TacticRulesResultKafkaInputData) *PeopleTacticRulesResultIndexEntity {
//	m := data.Data
//	entity := &PeopleTacticRulesResultIndexEntity{
//		Id:                 utils.GetUUID(),
//		EventId:            m.EventId,
//		EventReId:          m.EventReId,
//		TacticRuleID:       data.RuleID,
//		HitSignCode:        data.HitSignCode,
//		Timestamp:          m.Timestamp,
//		SensorId:           m.SensorId,
//		SensorName:         m.SensorName,
//		FaceId:             m.FaceId,
//		FaceRepoName:       m.FaceRepoName,
//		FaceRepoId:         m.FaceReId,
//		FaceReId:           m.FaceReId,
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

//func (e *PeopleTacticRulesResultIndexEntity) ToModel() *models.TacticRulesResult {
//	if e == nil {
//		return nil
//	}
//	civilAttr := &models.CivilAttr{
//		Name:   e.CivilAttrName,
//		IdType: models.IdType(e.CivilAttrIdType),
//		IdNo:   e.CivilAttrIdNo,
//		AttrId: e.CivilAttrId,
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
//		ResultID:        e.Id,
//		CreateTime:      e.CreateTime,
//		EventID:         e.EventId,
//		EventReID:       e.EventReId,
//		Timestamp:       e.Timestamp,
//		TacticRuleID:    e.TacticRuleID,
//		SensorID:        e.SensorId,
//		SensorName:      e.SensorName,
//		FaceRepoID:      e.FaceRepoId,
//		FaceRepoName:    e.FaceRepoName,
//		FaceID:          e.FaceId,
//		FaceReID:        e.FaceReId,
//		Confidence:      e.Confidence,
//		UserID:          e.UserID,
//		//CivilAttr:       civilAttr,
//		//CapturedImage:   capturedImage,
//		//RegisteredImage: registeredImage,
//		Status:          e.Status,
//		IsFollow:        e.IsFollow,
//		EventCount:      e.Count,
//	}
//
//}
