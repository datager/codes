package entities

import (
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils"
)

type UserCaptureFavorite struct {
	ID       int64  `xorm:"id serial pk not null"`
	TS       int64  `xorm:"ts"`
	UserID   string `xorm:"user_id"`
	Tag      int64  `xorm:"tag"`
	Type     int64  `xorm:"type"`
	TargetID string `xorm:"target_id"`
}

func (*UserCaptureFavorite) TableName() string {
	return TableNameUserCaptureFavorite
}

func NewUserCaptureFavorites(user *models.User, request *models.CaptureFavoriteRequest) []*UserCaptureFavorite {
	var result []*UserCaptureFavorite
	for _, tag := range request.Tags {
		result = append(result, &UserCaptureFavorite{
			TS:       utils.GetNowTs(),
			UserID:   user.UserId,
			Tag:      tag,
			Type:     request.DetType,
			TargetID: request.TargetID,
		})
	}
	return result
}
