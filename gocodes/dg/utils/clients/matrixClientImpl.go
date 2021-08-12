package clients

import (
	"time"

	"codes/gocodes/dg/utils/log"

	"github.com/pkg/errors"
	globalmodels "codes/gocodes/dg/dg.model"
)

type MatrixClientImpl struct {
	*GRPCClient
}

func NewMatrixClient(addr string, timeout time.Duration) MatrixClient {
	return &MatrixClientImpl{
		NewGRPCClient(addr, timeout),
	}
}

func (this *MatrixClientImpl) BatchRecognize(req *globalmodels.WitnessRequest) (*globalmodels.WitnessResponse, error) {
	client, err := this.ensureClient()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	start := time.Now()
	log.Debugf("Sending request to rpc://%v/BatchRecognize, req = %v", this.GetAddr(), req)
	// resp, err := client.BatchRecognize(this.GetContext(), req)
	resp, err := client.Recognize(this.GetContext(), req)
	if err != nil {
		log.Errorf("Failed to send request to rpc://%v/BatchRecognize\n%+v", this.GetAddr(), err)
		return nil, errors.WithStack(err)
	}
	log.Debugf("Got response from rpc://%v/BatchRecognize, resp = %v, took: %v", this.GetAddr(), resp, time.Now().Sub(start))
	return resp, nil
}

func (this *MatrixClientImpl) ensureClient() (globalmodels.WitnessServiceClient, error) {
	conn, err := this.EnsureConn()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return globalmodels.NewWitnessServiceClient(conn), nil
}
