package entities

import (
	"codes/gocodes/dg/models"
	"time"
)

type VehicleEvents struct {
	Uts                       time.Time `xorm:"uts"`
	Ts                        int64     `xorm:"ts"`
	EventID                   string    `xorm:"event_id"`
	EventReID                 string    `xorm:"event_reid"`
	SensorID                  string    `xorm:"sensor_id"`
	SensorName                string    `xorm:"sensor_name"`
	RuleID                    string    `xorm:"rule_id"`
	RuleName                  string    `xorm:"rule_name"`
	VehicleID                 string    `xorm:"vehicle_id"`
	RepoID                    string    `xorm:"repo_id"`
	RepoName                  string    `xorm:"repo_name"`
	UserID                    string    `xorm:"user_id"`
	EventType                 int       `xorm:"event_type"`
	IsChecked                 int       `xorm:"is_checked"`
	Comment                   string    `xorm:"comment"`
	Status                    int       `xorm:"status"`
	CapturedTypeID            int32     `xorm:"captured_type_id"`
	CapturedColorID           int32     `xorm:"captured_color_id"`
	CapturedPlateText         string    `xorm:"captured_plate_text"`
	CapturedPlateTypeID       int32     `xorm:"captured_plate_type_id"`
	CapturedPlateColorID      int32     `xorm:"captured_plate_color_id"`
	CapturedImageURI          string    `xorm:"captured_image_uri"`
	CapturedThumbnailImageURI string    `xorm:"captured_thumbnail_image_uri"`
	CapturedCutboardImageURI  string    `xorm:"captured_cutboard_image_uri"`
	CapturedCutboardX         int       `xorm:"captured_cutboard_x"`
	CapturedCutboardY         int       `xorm:"captured_cutboard_y"`
	CapturedCutboardWidth     int       `xorm:"captured_cutboard_width"`
	CapturedCutboardHeight    int       `xorm:"captured_cutboard_height"`
	CapturedCutboardResWidth  int       `xorm:"captured_cutboard_res_width"`
	CapturedCutboardResHeight int       `xorm:"captured_cutboard_res_height"`
	CapturedBrandID           int32     `xorm:"captured_brand_id"`
	CapturedSubBrandID        int32     `xorm:"captured_sub_brand_id"`
	CapturedModelYearID       int32     `xorm:"captured_model_year_id"`
	VehicleAttrID             string    `xorm:"vehicle_attr_id"`
	VehiclePlateText          string    `xorm:"vehicle_plate_text"`
	VehiclePlateColorID       string    `xorm:"vehicle_plate_color_id"`
	VehicleColorID            string    `xorm:"vehicle_color_id"`
	VehicleOwnerName          string    `xorm:"vehicle_owner_name"`
	VehiclePhoneNumber        string    `xorm:"vehicle_phone_number"`
	VehicleIDType             int       `xorm:"vehicle_id_type"`
	VehicleIDNo               string    `xorm:"vehicle_id_no"`
	VehicleComment            string    `xorm:"vehicle_comment"`
	VehicleBrandID            string    `xorm:"vehicle_brand_id"`
	VehicleSubBrandID         string    `xorm:"vehicle_sub_brand_id"`
	VehicleModelYearID        string    `xorm:"vehicle_model_year_id"`
	SensorIntID               int32     `xorm:"sensor_int_id"`
	OrgCode                   int64
}

func ToModels(data []*VehicleEvents) []*models.VehicleEvent {
	models := make([]*models.VehicleEvent, len(data))
	for i := 0; i < len(data); i++ {
		models[i] = data[i].Tomodel()
	}
	return models
}

func (entity *VehicleEvents) Tomodel() *models.VehicleEvent {
	model := new(models.VehicleEvent)
	vehicleCaptureDetail := new(models.VehicleCaptureDetail)
	vehicleAttr := new(models.VehicleAttrs)
	capturedImage := new(models.ImageResult)
	model.Timestamp = entity.Ts
	model.EventID = entity.EventID
	model.EventReID = entity.EventReID
	model.SensorID = entity.SensorID
	model.SensorName = entity.SensorName
	model.VehicleRepoID = entity.RepoID
	model.VehicleRepoName = entity.RepoName
	model.VehicleID = entity.VehicleID
	model.RuleID = entity.RuleID
	model.RuleName = entity.RuleName
	model.IsChecked = entity.IsChecked
	model.Comment = entity.Comment
	model.UserID = entity.UserID
	model.Comment = entity.Comment
	model.Status = models.TaskStatus(entity.Status)

	vehicleCaptureDetail.BrandID = entity.CapturedBrandID
	vehicleCaptureDetail.SubBrandID = entity.CapturedSubBrandID
	vehicleCaptureDetail.ModelYearID = entity.CapturedModelYearID
	vehicleCaptureDetail.TypeID = entity.CapturedTypeID
	vehicleCaptureDetail.PlateText = entity.CapturedPlateText
	vehicleCaptureDetail.ColorID = entity.CapturedColorID
	vehicleCaptureDetail.PlateTypeID = entity.CapturedPlateTypeID
	vehicleCaptureDetail.PlateColorID = entity.CapturedPlateColorID

	vehicleAttr.ID = entity.VehicleAttrID
	vehicleAttr.RepoID = entity.RepoID
	vehicleAttr.PlateText = entity.VehiclePlateText
	vehicleAttr.PlateColorID = entity.VehiclePlateColorID
	vehicleAttr.ColorID = entity.VehicleColorID
	vehicleAttr.BrandID = entity.VehicleBrandID
	vehicleAttr.SubBrandID = entity.VehicleSubBrandID
	vehicleAttr.ModelYearID = entity.VehicleModelYearID
	vehicleAttr.OwnerName = entity.VehicleOwnerName
	vehicleAttr.PhoneNumber = entity.VehiclePhoneNumber
	vehicleAttr.Comment = entity.Comment

	capturedImage.ImageUri = entity.CapturedImageURI
	capturedImage.ThumbnailImageUri = entity.CapturedThumbnailImageURI
	capturedImage.CutboardImageUri = entity.CapturedCutboardImageURI
	capturedImage.CutboardX = entity.CapturedCutboardX
	capturedImage.CutboardY = entity.CapturedCutboardY
	capturedImage.CutboardWidth = entity.CapturedCutboardWidth
	capturedImage.CutboardHeight = entity.CapturedCutboardHeight
	capturedImage.CutboardResWidth = entity.CapturedCutboardResWidth
	capturedImage.CutboardResHeight = entity.CapturedCutboardResHeight

	model.CapturedImage = capturedImage
	model.VehicleAttr = vehicleAttr
	model.VehicleCaptureDetail = vehicleCaptureDetail
	model.OrgCode = entity.OrgCode
	model.SensorIntID = entity.SensorIntID

	return model
}
