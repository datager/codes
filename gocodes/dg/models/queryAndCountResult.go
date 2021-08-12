package models

type QueryAndCountResult struct {
	Rets  interface{}
	Total int
}

type ResultWithPage struct {
	Rets       interface{}
	Total      int
	NextStart  int
	IsLastPage bool
}

type FaceRepoGetResult struct {
	Rets  []*FaceRepo
	Total int
}
