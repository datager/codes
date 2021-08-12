package clients

import (
	"io"
	"codes/gocodes/dg/models"

	dg_model "codes/gocodes/dg/dg.model"
)

type VseMatrixClient interface {
	io.Closer
	RecognizeAllObject(img *models.ImageQuery, recognizeType dg_model.RecognizeType, detectMode int, recognizeFunction []int32) (*models.AllObject, error)
	GetIndexTxt(indexType dg_model.IndexType) (*dg_model.IndexTxtResponse, error)
}
