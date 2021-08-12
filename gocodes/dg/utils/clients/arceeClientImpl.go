package clients

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/log"

	"github.com/pkg/errors"
	globalmodels "codes/gocodes/dg/dg.model"
)

type ArceeClientImpl struct {
	*GRPCClient
}

func NewArceeClient(addr string, timeout time.Duration) ArceeClient {
	return &ArceeClientImpl{
		NewGRPCClient(addr, timeout),
	}
}

func (aci *ArceeClientImpl) BatchAddImages(imgs []*models.ImageQuery) ([]string, error) {
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
	resp, err := aci.BatchAddBase64(req)
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

func (aci *ArceeClientImpl) BatchDeleteImages(imageURLs []string) error {
	return nil
}

func (aci *ArceeClientImpl) BatchAddBase64(imgs map[string]string) (map[string]string, error) {
	if len(imgs) == 0 {
		return nil, nil
	}
	client, err := aci.ensureClient()
	if err != nil {
		return nil, errors.WithStack(err)
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
	req := &globalmodels.WeedFileMapMessage{
		MapData: m,
	}
	start := time.Now()
	log.Debugf("Sending request to rpc://%v/PostFileContent, len of images = %v", aci.GetAddr(), totalLength)
	resp, err := client.PostFileContent(aci.GetContext(), req)
	if err != nil {
		log.Errorf("Failed to send request to rpc://%v/PostFileContent\n%+v", aci.GetAddr(), err)
		return nil, errors.WithStack(err)
	}
	if resp == nil {
		log.Errorf("Got empty response from rpc://%v/PostFileContent", aci.GetAddr())
		return nil, fmt.Errorf("Got invalid response from arcee")
	}
	log.Debugf("Got response from rpc://%v/PostFileContent, resp = %v, took: %v", aci.GetAddr(), resp, time.Since(start))
	return resp.MapIds, nil
}

func (aci *ArceeClientImpl) ensureClient() (globalmodels.GrpcServiceClient, error) {
	conn, err := aci.EnsureConn()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return globalmodels.NewGrpcServiceClient(conn), nil
}
