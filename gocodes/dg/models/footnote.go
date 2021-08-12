package models

import "time"

const (
	KindSettlements = iota + 1
	KindWorkshop
	KindTemporary
)

const (
	FootnoteTypeFace = iota + 1
	FootnoteTypeVehicle
)

type Footnote struct {
	ID           string `json:"Id"`
	Vid          string
	NoteKind     int
	Longitude    float32
	Latitude     float32
	Comment      string
	Ts           int64
	Uts          time.Time
	NoteName     string
	FootnoteType int
}
