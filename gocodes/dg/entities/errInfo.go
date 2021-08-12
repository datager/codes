package entities

import (
	"time"

	"codes/gocodes/dg/utils/json"

	"codes/gocodes/dg/models"
)

type ErrInfo struct {
	Uts       time.Time `xorm:"not null default 'now()' DATETIME"`
	Ts        int64     `xorm:"not null BIGINT"`
	ErrId     string    `xorm:"not null pk index VARCHAR(1024)"`
	ErrType   string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	ErrMsg    string    `xorm:"not null TEXT"`
	ErrReason string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	RepoId    string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Status    int       `xorm:"not null default 1 SMALLINT"`
}

type ErrInfoExtend struct {
	ErrInfo  `xorm:"extends"`
	FaceRepo FaceRepos `xorm:"extends"`
}

func (ErrInfoExtend) TableName() string {
	return TableNameErrInfo
}

type croatiaObj struct {
	Name  string
	ID    string
	Image *croatiaImage
}

type croatiaImage struct {
	Uri string
}

func (entity *ErrInfoExtend) ToModel() *models.Civil {
	var obj croatiaObj
	json.Unmarshal([]byte(entity.ErrMsg), &obj) // todo error ignored...
	return &models.Civil{
		CivilAttr: &models.CivilAttr{
			Name:       obj.Name,
			IdType:     models.IdType_Id_Type_IdNo,
			IdNo:       obj.ID,
			SaveStatus: models.CivilSaveStatusFailed,
		},
		FaceRepoId:   entity.RepoId,
		FaceRepoName: entity.FaceRepo.RepoName,
		ImageResults: []*models.ImageResult{
			&models.ImageResult{
				ImageUri: obj.Image.Uri,
				IsRanked: 0,
			},
		},
		NameListAttr: entity.FaceRepo.NameListAttr,
		Status:       models.TaskStatus(entity.Status),
		Timestamp:    entity.Ts,
	}
}
