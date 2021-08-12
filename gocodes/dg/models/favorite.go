package models

import (
	"encoding/json"
	"errors"
)

type Favorite struct {
	FavoriteTags      map[int64]int64 // 收藏状态
	FavoriteTimestamp int64           // 收藏时间
	FavoriteTargetID  string          // 目标ID
	FavoriteStatus    bool            // 收藏状态
}

type CaptureFavoriteRequest struct {
	DetType  int64   // 类型(行人:1,机动车:2,非机动车:4,人脸:8)
	Tags     []int64 // 标签
	TargetID string  // 目标ID
}

func (m *CaptureFavoriteRequest) Check() error {
	if m.DetType != DetTypePedestrian &&
		m.DetType != DetTypeVehicle &&
		m.DetType != DetTypeNonmotor &&
		m.DetType != DetTypeFace {
		return errors.New("error request")
	}
	return nil
}

type CaptureFavoriteListRequest struct {
	DetType int64 // 类型(行人:1,机动车:2,非机动车:4,人脸:8)
	Tag     int64 // 标签
	Pagination
}

func (m *CaptureFavoriteListRequest) Check() error {
	if m.DetType != DetTypePedestrian &&
		m.DetType != DetTypeVehicle &&
		m.DetType != DetTypeNonmotor &&
		m.DetType != DetTypeFace {
		return errors.New("error request")
	}
	return nil
}

type FavoriteCaptureListResponse struct {
	Total       int64
	Pedestrians []*Pedestrian
	Vehicles    []*VehicleCapture
	Nonmotors   []*Nonmotor
	Faces       []*CapturedFaceDetial
}

func (m *FavoriteCaptureListResponse) MarshalJSON() ([]byte, error) {
	var rets interface{}
	if m.Pedestrians != nil {
		rets = m.Pedestrians
	}
	if m.Vehicles != nil {
		rets = m.Vehicles
	}
	if m.Nonmotors != nil {
		rets = m.Nonmotors
	}
	if m.Faces != nil {
		rets = m.Faces
	}
	if rets == nil {
		rets = []interface{}{}
	}
	return json.Marshal(struct {
		Total int64
		Rets  interface{}
	}{
		Total: m.Total,
		Rets:  rets,
	})
}
