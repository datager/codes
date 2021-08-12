package models

const (
	ImageTypeUnknown = iota
	ImageTypeID
	ImageTypeCaptured
	ImageTypeUserPosted
)

const (
	CivilStatusUnknown = iota
	CivilStatusCreated
)

const (
	CivilGenderMale = iota + 1
	CivilGenderFemale
	CivilGenderOther
)

const (
	CivilSaveStatusSuccess = iota + 1
	CivilSaveStatusFailed
)

type Civil struct {
	CivilAttr    *CivilAttr
	Confidence   float32
	FaceRepoId   string
	FaceRepoName string
	ImageResults []*ImageResult
	NameListAttr int
	Status       TaskStatus
	Timestamp    int64
	HadEvents    bool
}

type CivilSort []*Civil

func (c CivilSort) Len() int {
	return len(c)
}

func (c CivilSort) Less(i, j int) bool {
	return c[i].Confidence > c[j].Confidence
}

func (c CivilSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
