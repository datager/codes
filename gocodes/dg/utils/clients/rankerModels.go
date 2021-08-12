package clients

import globalmodels "codes/gocodes/dg/dg.model"

type RankFeatureQuery struct {
	RepoId         string
	Feature        string
	Confidence     float32
	TopX           int
	SensorIds      []string
	StartTimestamp int64
	EndTimestamp   int64
	RecognizeType  globalmodels.RecognizeType
}

type RankerResponse interface {
	GetContext() *globalmodels.RankResponseContext
}

type RankItem struct {
	Id    string
	Score float32
	index int
}

func (item *RankItem) GetIndex() int {
	return item.index
}

type RankerRepoParams struct {
	GPU             string
	FeatureLen      int32
	FeatureDataType int32
	RepoLevel       int32
	RepoCapacity    int32
	UseFeatureIDMap string
	PreFilter       string
}

type RankerFace struct {
	ImageID string
	Feature string
}
