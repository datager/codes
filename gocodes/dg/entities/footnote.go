package entities

import (
	"time"

	"codes/gocodes/dg/models"
)

type Footnote struct {
	ID           string    `xorm:"id not null pk index VARCHAR(1024)"`
	Vid          string    `xorm:"not null index VARCHAR(1024)"`
	NoteKind     int       `xorm:"not null default 0 SMALLINT"`
	Longitude    float32   `xorm:"default '-200'::real REAL"`
	Latitude     float32   `xorm:"default '-200'::real REAL"`
	Comment      string    `xorm:"default '''::character varying' VARCHAR"`
	Ts           int64     `xorm:"not null default 0 BIGINT"`
	Uts          time.Time `xorm:"not null default 'now() DATETIME"`
	FootnoteType int       `xorm:"footnote_type not null default 1"`
}

func (entity *Footnote) TableName() string {
	return TableNameFootnote
}

func (entity *Footnote) ToModel() *models.Footnote {
	if entity == nil {
		return nil
	}

	return &models.Footnote{
		ID:           entity.ID,
		Vid:          entity.Vid,
		NoteKind:     entity.NoteKind,
		Longitude:    entity.Longitude,
		Latitude:     entity.Latitude,
		Comment:      entity.Comment,
		Ts:           entity.Ts,
		Uts:          entity.Uts,
		NoteName:     entity.NoteName(),
		FootnoteType: entity.FootnoteType,
	}
}

func (entity *Footnote) NoteName() string {
	var ret string
	switch entity.NoteKind {
	case models.KindSettlements:
		ret = "居住点"
	case models.KindWorkshop:
		ret = "工作点"
	case models.KindTemporary:
		ret = "临时点"
	}

	return ret
}

func FootnoteModelToEntity(note *models.Footnote) *Footnote {
	return &Footnote{
		ID:           note.ID,
		Vid:          note.Vid,
		NoteKind:     note.NoteKind,
		Longitude:    note.Longitude,
		Latitude:     note.Latitude,
		Comment:      note.Comment,
		Ts:           note.Ts,
		Uts:          note.Uts,
		FootnoteType: note.FootnoteType,
	}
}
