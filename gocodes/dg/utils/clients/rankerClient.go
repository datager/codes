package clients

import (
	"io"

	"github.com/emirpasic/gods/maps/treebidimap"
	globalmodels "codes/gocodes/dg/dg.model"
)

type RankerClient interface {
	io.Closer
	AddRepo(repoId string, repoParams *RankerRepoParams) error
	RankFeature(query *RankFeatureQuery) ([]*globalmodels.RankItem, error)
	BatchRankFeature(queries []*RankFeatureQuery) (*treebidimap.Map, error)
	QueryFeature(repoId string, id string) (*globalmodels.RankFeatureOperation, error)
	QueryFeatureEx(repoId string, id string) (*globalmodels.ObjectProperty, error)
	BatchAddRankerFeature(params []*RankerFace, repoID, location string) error
	BatchDeleteRankerFeature(repoID string, imageIDs []string) error
	DeleteRepo(repoID string) error
	QueryRepo(repoID string) (bool, error)
}
