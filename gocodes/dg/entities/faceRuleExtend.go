package entities

import (
	"codes/gocodes/dg/utils/uuid"
	"time"

	"codes/gocodes/dg/models"
)

type FaceRuleExtend struct {
	ID         string    `xorm:"id not null pk VARCHAR(1024)"`
	SensorID   string    `xorm:"sensor_id not null VARCHAR(1024)"`
	FaceRepoID string    `xorm:"face_repo_id not null VARCHAR(1024)"`
	UserID     string    `xorm:"user_id not null VARCHAR(1024)"`
	AlarmLevel int       `xorm:"alarm_level not null default 0 SMALLINT"`
	AlarmVoice string    `xorm:"alarm_voice not null VARCHAR(1024)"`
	Ts         int64     `xorm:"ts not null default 0 BIGINT"`
	OrgID      string    `xorm:"org_id not null VARCHAR(1024)"`
	RuleID     string    `xorm:"rule_id not null VARCHAR(1024)"`
	UpdatedAt  time.Time `xorm:"updated"`
}

func (fre *FaceRuleExtend) TableName() string {
	return TableNameFaceRuleExtend
}

func (fre *FaceRuleExtend) ToModel() *models.FaceRuleExtend {
	e := &models.FaceRuleExtend{
		ID:         fre.ID,
		SensorID:   fre.SensorID,
		FaceRepoID: fre.FaceRepoID,
		UserID:     fre.UserID,
		AlarmLevel: fre.AlarmLevel,
		AlarmVoice: fre.AlarmVoice,
		Ts:         fre.Ts,
		OrgID:      fre.OrgID,
		RuleID:     fre.RuleID,
		UpdatedAt:  fre.UpdatedAt,
	}
	return e
}

func FaceRuleExtendModelToEntity(fre *models.FaceRuleExtend) *FaceRuleExtend {
	e := &FaceRuleExtend{
		ID:         fre.ID,
		SensorID:   fre.SensorID,
		FaceRepoID: fre.FaceRepoID,
		UserID:     fre.UserID,
		AlarmLevel: fre.AlarmLevel,
		AlarmVoice: fre.AlarmVoice,
		Ts:         fre.Ts,
		OrgID:      fre.OrgID,
		RuleID:     fre.RuleID,
		UpdatedAt:  fre.UpdatedAt,
	}
	if e.ID == "" {
		e.ID = uuid.NewV4().String()
	}
	return e
}

func FaceRuleExtendsModelsToEntities(mdls []*models.FaceRuleExtend) []*FaceRuleExtend {
	ens := make([]*FaceRuleExtend, 0)
	for _, m := range mdls {
		ens = append(ens, FaceRuleExtendModelToEntity(m))
	}
	return ens
}
