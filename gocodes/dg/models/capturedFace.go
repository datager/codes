package models

type CutboardImage struct {
	CutboardImageURI string `json:"CutboardImageUri"`
}

type CapturedFace struct {
	FaceID            string
	FaceReID          string
	Timestamp         int64
	SensorID          string
	SensorIntID       int32
	OrgCode           int64
	SensorName        string
	Longitude         float64
	Latitude          float64
	AgeID             int `json:"AgeId"`
	GenderID          int `json:"GenderId"`
	GlassID           int `json:"GlassId"`
	ImageResult       *CutboardImage
	FaceVid           string
	Favorite          *Favorite
	RelationTypeSlice []int
}

type CapturedFaceDetial struct {
	FaceID        string
	FaceReID      string
	Timestamp     int64
	SensorID      string
	OrgCode       int64
	SensorName    string
	SensorType    int
	SensorURL     string
	Longitude     float64
	Latitude      float64
	Confidence    float32
	IsWarned      int
	AgeID         int `json:"AgeId"`
	GenderID      int `json:"GenderId"`
	GlassID       int `json:"GlassId"`
	CivilAttr     *CivilAttr
	ImageResult   *ImageResult
	Status        TaskStatus
	FaceVid       string
	Favorite      *Favorite
	RelationTypes []int
	RelationIDs   []string
	RelationURLs  []string
}

func (capturedFace *CapturedFace) AddFavorite(favorite *Favorite) *CapturedFace {
	if favorite != nil {
		capturedFace.Favorite = favorite
	} else {
		capturedFace.Favorite = &Favorite{FavoriteTargetID: capturedFace.FaceID}
	}
	return capturedFace
}

func (capturedFace *CapturedFaceDetial) AddFavorite(favorite *Favorite) *CapturedFaceDetial {
	if favorite != nil {
		capturedFace.Favorite = favorite
	} else {
		capturedFace.Favorite = &Favorite{FavoriteTargetID: capturedFace.FaceID}
	}
	return capturedFace
}

type FacesIndexExtend struct {
	CapturedFace
	Sensor Sensor
}

type FacesExtend struct {
	CapturedFaceDetial
	Sensor Sensor
}

type ExtendData interface {
	GetSensor() *Sensor
	GetCapturedTimestamp() int64
}

func (fe FacesIndexExtend) GetSensor() *Sensor {
	return &fe.Sensor
}

func (fe FacesIndexExtend) GetCapturedTimestamp() int64 {
	return fe.CapturedFace.Timestamp
}

type CountCaptured struct {
	Date           string
	Count          int
	StartTimestamp int64
	EndTimestamp   int64
}
