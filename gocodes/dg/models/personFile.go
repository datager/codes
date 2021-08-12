package models

type PersonFile struct {
	Ts               int64
	Vid              string
	Name             string
	Birthday         int64
	IdType           int
	IdNo             string
	Address          string
	GenderId         int
	NationId         int
	Comment          string
	ImageUrl         string
	MarkState        int
	IsNeedAutoUpdate bool
	Status           int
	Confidence       float32
	CapturedFace     *CapturedFace
	IsImpacted       bool

	SystemTags []string
	CustomTags []string

	ExistInRepos      []*ReposData
	CapturedFaceCount int
}
type Tags struct {
	SystemTags []string
	CustomTags []string
}

type ReposData struct {
	RepoID   string
	RepoName string
}

type CommentJSON struct {
	Repos map[string]string
	Tags  string
}
