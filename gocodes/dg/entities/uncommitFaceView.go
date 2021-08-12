package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type FacesNotCommitView struct {
	Uts time.Time `xorm:"uts"`
	Ts  int64     `xorm:"ts"`

	PairId            string  `xorm:"pair_id"`
	SensorId          string  `xorm:"sensor_id"`
	SensorName        string  `xorm:"sensor_name"`
	FaceId            string  `xorm:"face_id"`
	Feature           string  `xorm:"feature"`
	Confidence        float64 `xorm:"confidence"`
	NationId          int64   `xorm:"nation_id"`
	ImageType         int     `xorm:"image_type"`
	ImageUri          string  `xorm:"image_uri"`
	ThumbnailImageUri string  `xorm:"thumbnail_image_uri"`
	CutboardImageUri  string  `xorm:"cutboard_image_uri"`
	CutboardX         int     `xorm:"cutboard_x"`
	CutboardY         int     `xorm:"cutboard_y"`
	CutboardWidth     int     `xorm:"cutboard_width"`
	CutboardHeight    int     `xorm:"cutboard_height"`
	CutboardResWidth  int     `xorm:"cutboard_res_width"`
	CutboardResHeight int     `xorm:"cutboard_res_height"`

	GateSensorId          string  `xorm:"gate_sensor_id" `
	GateFaceId            string  `xorm:"gate_face_id" `
	GateFeature           string  `xorm:"gate_feature" `
	GateConfidence        float64 `xorm:"gate_confidence" `
	GateNationId          int64   `xorm:"gate_nation_id" `
	GateImageType         int     `xorm:"gate_image_type" `
	GateImageUri          string  `xorm:"gate_image_uri" `
	GateThumbnailImageUri string  `xorm:"gate_thumbnail_image_uri" `
	GateCutboardImageUri  string  `xorm:"gate_cutboard_image_uri" `
	GateCutboardX         int     `xorm:"gate_cutboard_x" `
	GateCutboardY         int     `xorm:"gate_cutboard_y" `
	GateCutboardWidth     int     `xorm:"gate_cutboard_width" `
	GateCutboardHeight    int     `xorm:"gate_cutboard_height" `
	GateCutboardResWidth  int     `xorm:"gate_cutboard_res_width" `
	GateCutboardResHeight int     `xorm:"gate_cutboard_res_height" `

	Name          string  `xorm:"name" `
	IdType        int64   `xorm:"id_type" ` //证件类型
	IDNo          string  `xorm:"id_no" `   //身份证号
	Address       string  `xorm:"address" ` //家庭住址
	CmpConfidence float64 `xorm:"cmp_confidence" `
	CheckValid    int     `xorm:"check_valid" ` // 0：未审核；1：审核有效；2：审核无效；3：无法确定
	Comment       string  `xorm:"comment" `
}

func (FacesNotCommitView) TableName() string {
	return TableNameUnCommitFaceView
}

func (entity *FacesNotCommitView) ToModel() *models.UncommitFace {
	uncommitface := &models.UncommitFace{
		PairId:     entity.PairId,
		Name:       entity.Name,
		Timestamp:  entity.Ts,
		SensorId:   entity.SensorId,
		SensorName: entity.SensorName,
		Confidence: entity.CmpConfidence,
		IsChecked:  entity.CheckValid,
		Comment:    entity.Comment,
	}
	captuerImage := &models.UncommitFaceImage{
		FaceId:            entity.GateFaceId,
		ImageUri:          entity.GateImageUri,
		ImageType:         int64(entity.GateImageType),
		CutboardImageUri:  entity.GateCutboardImageUri,
		CutboardX:         int64(entity.GateCutboardX),
		CutboardY:         int64(entity.GateCutboardY),
		CutboardWidth:     int64(entity.GateCutboardWidth),
		CutboardHeight:    int64(entity.GateCutboardHeight),
		CutboardResWidth:  int64(entity.GateCutboardResWidth),
		CutboardResHeight: int64(entity.GateCutboardResHeight),
		Feature:           entity.GateFeature,
	}
	registeredImage := &models.UncommitFaceImage{
		FaceId:            entity.FaceId,
		ImageUri:          entity.ImageUri,
		ImageType:         int64(entity.ImageType),
		CutboardImageUri:  entity.CutboardImageUri,
		CutboardX:         int64(entity.CutboardX),
		CutboardY:         int64(entity.CutboardY),
		CutboardWidth:     int64(entity.CutboardWidth),
		CutboardHeight:    int64(entity.CutboardHeight),
		CutboardResWidth:  int64(entity.CutboardResWidth),
		CutboardResHeight: int64(entity.CutboardResHeight),
		Feature:           entity.Feature,
	}
	civilattr := &models.UncommitFaceCivilAttr{
		Name:     entity.Name,
		IDType:   int(entity.IdType),
		IDNo:     entity.IDNo,
		Addr:     entity.Address,
		NationId: int(entity.NationId),
	}
	uncommitface.CapturedImage = captuerImage
	uncommitface.RegisteredImage = registeredImage
	uncommitface.CivilAttr = civilattr
	return uncommitface
}
