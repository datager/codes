package models

import "time"

type VehicleAttrs struct {
	ID           string
	RepoID       string `binding:"required"`
	OrgID        string `binding:"required"`
	PlateText    string `binding:"required"`
	PlateColorID string
	ColorID      string
	BrandID      string
	SubBrandID   string
	ModelYearID  string
	OwnerName    string
	PhoneNumber  string
	IDType       int    `json:"IdType"`
	IDNo         string `json:"IdNo"`
	Comment      string
	Ts           int64
	Uts          time.Time
}

type VehicleAttrsExtend struct {
	Attr *VehicleAttrs
	Repo *FaceRepo
}

type VehicleAttrsPayload struct {
	Payload *VehicleAttrsQuery
}

type VehicleAttrsQuery struct {
	RepoIDs       []string
	VehicleStatus int
	PlateText     string
	Limit         int  `binding:"required"`
	Offset        *int `binding:"required"`
	Total         int64
	OrderAsc      bool
	OrderBy       string
	FromTailPage  bool
}

type VehicleAttrsList struct {
	Rets       []*VehicleAttrsExtend
	NextOffset int
	Total      int64
}

type CheckPlateParam struct {
	RepoID    string `form:"repoid" binding:"required"`
	PlateText string `form:"platetext" binding:"required"`
}

type CheckPlateRet struct {
	Exist bool
}
