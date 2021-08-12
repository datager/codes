package models

type RepoDailyPicCount struct {
	Id       string
	RepoId   string
	PicCount int64
	Ts       int64
}

type QueryRepoDailyPicCountParam struct {
	RepoIds        []string
	StartTimestamp int64
	EndTimestamp   int64
}

type RepoDailyPicCountData struct {
	RepoId     string
	RepoName   string
	TsSlice    []int64
	CountSlice []int64
}
