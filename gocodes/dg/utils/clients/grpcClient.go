package clients

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	addr    string
	timeout time.Duration
	conn    *grpc.ClientConn
}

func NewGRPCClient(addr string, timeout time.Duration) *GRPCClient {
	return &GRPCClient{
		addr:    addr,
		timeout: timeout,
	}
}

func (this *GRPCClient) GetAddr() string {
	return this.addr
}

func (this *GRPCClient) Close() (err error) {
	if this.conn != nil {
		err = this.conn.Close()
	}
	this.conn = nil
	return errors.WithStack(err)
}

func (this *GRPCClient) EnsureConn() (*grpc.ClientConn, error) {
	if this.conn != nil {
		return this.conn, nil
	}
	conn, err := grpc.Dial(this.addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	this.conn = conn
	return conn, nil
}

func (this *GRPCClient) GetContext() context.Context {
	ctx := context.Background()
	context.WithTimeout(ctx, this.timeout)
	return ctx
}
