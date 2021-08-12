package models

type FacesIndexDict struct {
	FaceID           string
	SensorID         string
	FaceVid          string
	FaceReID         string
	GenderID         int
	AgeID            int
	NationID         int
	GlassID          int
	MaskID           int
	HatID            int
	HalmetID         int
	CutboardImageURI string
	Ts               int64
}
