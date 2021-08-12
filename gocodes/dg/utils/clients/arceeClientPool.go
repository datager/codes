package clients

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"sync"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/log"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	dg_model "codes/gocodes/dg/dg.model"

	"google.golang.org/grpc"
)

type arceeConnection struct {
	client dg_model.GrpcServiceClient
	conn   *grpc.ClientConn
}

type ArceePoolClient struct {
	addr string

	connectNubmer int

	timeoutSecond time.Duration
	tryTimes      int

	connectionPool []*arceeConnection

	serviceOperationLock sync.RWMutex
	connectIndex         int
	connectLock          sync.Mutex
	isServiceOpen        bool
}

func NewArceePoolClient(addr string, connectNumber int, timeoutSecond time.Duration, tryTimes int) *ArceePoolClient {
	if addr == "" {
		panic("addr is empty")
	}

	service := &ArceePoolClient{
		addr:          addr,
		connectNubmer: connectNumber,
		timeoutSecond: timeoutSecond,
		tryTimes:      tryTimes,
	}

	service.init()

	return service
}

func (aps *ArceePoolClient) init() {
	aps.serviceOperationLock.Lock()
	defer aps.serviceOperationLock.Unlock()

	if aps.isServiceOpen {
		return
	}

	for {

		if len(aps.connectionPool) < aps.connectNubmer {
			conn, err := grpc.Dial(aps.addr, grpc.WithInsecure())
			if err != nil {
				log.Errorln("arcee connect err:", err)
				<-time.After(time.Second)
				continue
			}

			connClient := dg_model.NewGrpcServiceClient(conn)
			aps.connectionPool = append(aps.connectionPool, &arceeConnection{
				conn:   conn,
				client: connClient,
			})

			// log.Infoln("create a new connect to fse black, total connect num is ", len(aps.blackConnectionPool))
			continue
		}

		break
	}

	aps.isServiceOpen = true
}

func (aps *ArceePoolClient) BatchAddImages(imgs []*models.ImageQuery) ([]string, error) {
	if len(imgs) == 0 {
		return nil, nil
	}
	results := make([]string, len(imgs))
	req := make(map[string]string)
	for i, img := range imgs {
		if img.Url != "" {
			results[i] = img.Url
			continue
		}
		if img.BinData == "" {
			return nil, fmt.Errorf("Invalid bindata")
		}
		req[strconv.Itoa(i)] = img.BinData
	}
	if len(req) == 0 {
		return results, nil
	}
	resp, err := aps.BatchAddBase64(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for k, v := range resp {
		i, err := strconv.Atoi(k)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if i >= len(results) || results[i] != "" {
			return nil, fmt.Errorf("Invalid key %d", i)
		}
		results[i] = v
	}
	return results, nil
}

func (aps *ArceePoolClient) BatchDeleteImages(imageURLs []string) error {
	aps.serviceOperationLock.Lock()
	defer aps.serviceOperationLock.Unlock()

	var tryTimes int
	conn := aps.connectionPool[aps.getConnectIndex()]

	if len(imageURLs) == 0 {
		return nil
	}

	req := &dg_model.WeedFileArrayMessage{
		ArrayData: imageURLs,
	}

	start := time.Now()

	for {
		tryTimes++
		ctx, cancel := context.WithTimeout(context.Background(), aps.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/BatchDeleteImages, req = %v", aps.addr, req)

		resp, err := conn.client.DeleteFile(ctx, req)
		if err != nil || len(resp.ArrayData) != len(imageURLs) {
			if tryTimes <= aps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/BatchDeleteImages\n%+v", aps.addr, err)
			return errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/BatchDeleteImages, resp = %v, took: %v", aps.addr, resp, time.Since(start))

		break
	}

	return nil
}

func (aps *ArceePoolClient) BatchAddBase64(imgs map[string]string) (map[string]string, error) {
	aps.serviceOperationLock.Lock()
	defer aps.serviceOperationLock.Unlock()

	var tryTimes int
	var conn *arceeConnection
	var addr string

	conn = aps.connectionPool[aps.getConnectIndex()]
	addr = aps.addr

	if len(imgs) == 0 {
		return nil, nil
	}

	m := make(map[string][]byte)
	totalLength := 0

	for key, bindata := range imgs {
		body, err := base64.StdEncoding.DecodeString(bindata)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		m[key] = body
		totalLength += len(body)
	}

	req := &dg_model.WeedFileMapMessage{
		MapData: m,
	}

	start := time.Now()

	var resp *dg_model.WeedFileMapId
	var err error

	for {
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), aps.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/BatchAddBase64, length = %v", addr, totalLength)

		resp, err = conn.client.PostFileContent(ctx, req)
		if err != nil {
			if tryTimes <= aps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/BatchAddBase64\n%+v", addr, err)
			return nil, errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/BatchAddBase64, resp = %v, took: %v", addr, resp, time.Since(start))

		break
	}

	return resp.MapIds, nil
}

func (aps *ArceePoolClient) Close() error {
	aps.serviceOperationLock.Lock()
	defer aps.serviceOperationLock.Unlock()

	if !aps.isServiceOpen {
		return nil
	}

	for _, connection := range aps.connectionPool {
		connection.client = nil
		connection.conn.Close()
	}

	aps.connectionPool = []*arceeConnection{}

	aps.isServiceOpen = false
	log.Infoln("arcee service closed")

	return nil
}

func (aps *ArceePoolClient) getConnectIndex() int {
	aps.connectLock.Lock()
	defer aps.connectLock.Unlock()

	index := aps.connectIndex
	aps.connectIndex = (aps.connectIndex + 1) % aps.connectNubmer

	return index
}

func (aps *ArceePoolClient) newContext() (*dg_model.RankRequestContext, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &dg_model.RankRequestContext{
		SessionId: id.String(),
	}, nil
}
