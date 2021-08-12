package models

type RepoUpdates struct {
	Id     string
	RepoId string
	Kind   int
	Value  int
	Ts     int64
}

type QueryRepoUpdatesParam struct {
	StartTimestamp int64
	EndTimestamp   int64
}

type RepoUpdatesStatisticsData struct {
	RepoId      string
	RepoName    string
	AddValue    int
	DeleteValue int
}
