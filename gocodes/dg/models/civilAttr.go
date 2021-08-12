package models

type CivilAttr struct {
	Name             string
	IdType           IdType
	IdNo             string
	GenderId         int
	GenderConfidence float32
	AgeId            int
	AgeConfidence    float32
	NationId         int
	NationConfidence float32
	GlassId          int
	GlassConfidence  float32
	MaskId           int
	MaskConfidence   float32
	HatId            int
	HatConfidence    float32
	HalmetId         int
	HalmetConfidence float32
	Addr             string
	Birthday         *int64
	Comment          string
	Status           TaskStatus
	AttrId           string
	SaveStatus       int
	SearchID         string
}
