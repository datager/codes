package clients

import (
	"io"

	globalmodels "codes/gocodes/dg/dg.model"
)

type MatrixClient interface {
	io.Closer
	BatchRecognize(req *globalmodels.WitnessRequest) (*globalmodels.WitnessResponse, error)
}
