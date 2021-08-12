package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/config"
)

type ArchimedesClient struct {
	*HTTPClient
}

func NewArchimedesClient(conf config.Config) *ArchimedesClient {
	timeoutSecond := conf.GetIntOrDefault("services.archimedes.timeout_second", 60)
	client := ArchimedesClient{
		HTTPClient: NewHTTPClient(conf.GetString("services.archimedes.addr"), time.Duration(timeoutSecond)*time.Second, nil),
	}
	return &client
}

func (archimedesClient ArchimedesClient) RepoDayCount(request models.RepoStatisticsRequest) ([]byte, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	return archimedesClient.HTTPClient.Post("/repo", r)
}

func (archimedesClient ArchimedesClient) VehicleCaptureAttributesCount(request models.CaptureAttributeCountStatisticsToArchimedesRequest) ([]byte, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	return archimedesClient.HTTPClient.Post("/vehicle/capture/attributes", r)
}

func (archimedesClient ArchimedesClient) FaceCaptureAttributesCount(request models.FaceCaptureAttrCountStatisticsToArchimedesReq) ([]byte, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	return archimedesClient.HTTPClient.Post("/face/capture/attributes", r)
}

func (archimedesClient ArchimedesClient) FaceCaptureCountForInterval(request models.CaptureAttributeCountStatisticsRequest) ([]byte, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	return archimedesClient.HTTPClient.Post("/face/capture/count/interval", r)
}

func (archimedesClient ArchimedesClient) CaptureCountForTotal(req models.CaptureAttributeCountStatisticsRequest, captureType int) ([]byte, error) {
	rs, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)

	var url string
	switch captureType {
	case models.DetTypeFace:
		url = "/face/capture/count/total"
	case models.DetTypeVehicle:
		url = "/vehicle/capture/count/total"
	case models.DetTypeNonmotor:
		url = "/nonmotor/capture/count/total"
	case models.DetTypePedestrian:
		url = "/pedestrian/capture/count/total"
	default:
		return nil, errors.New("captureType not support in ArchimedesClient.CaptureCountForTotal")
	}
	return archimedesClient.HTTPClient.Post(url, r)
}

func (archimedesClient ArchimedesClient) SpecialCount(request models.SpecialCount) (*models.OnlyCount, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/special/count", r)
	if err != nil {
		return nil, err
	}
	count := &models.OnlyCount{}
	err = json.Unmarshal(body, count)
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (archimedesClient ArchimedesClient) SpecialCountAtCertainTime(request models.SpecialCount) (*models.SpecialAtCertainTime, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/special/countatcertaintime", r)
	if err != nil {
		return nil, err
	}
	result := &models.SpecialAtCertainTime{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) SpecialCountBySensorTop(request models.SpecialCount) ([]*models.KeyValueCommon, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/special/countbysensortop", r)
	if err != nil {
		return nil, err
	}
	var result []*models.KeyValueCommon
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) SpecialCountByPlateTop(request models.SpecialCount) ([]*models.KeyValueCommon, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/special/countbyplatetop", r)
	if err != nil {
		return nil, err
	}
	var result []*models.KeyValueCommon
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) SpecialCountByIllegalType(request models.SpecialCount) ([]*models.KeyValueIllegal, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/special/illegalcount", r)
	if err != nil {
		return nil, err
	}
	var result []*models.KeyValueIllegal
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//报警按时段统计

func (archimedesClient ArchimedesClient) VehicleEventCountByTime(request models.EventCountStatisticsRequest) ([]*models.KeyValueCommon, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/vehicle/events/durations", r)
	if err != nil {
		return nil, err
	}
	var result []*models.KeyValueCommon
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) VehicleEventCountBySensor(request models.EventCountStatisticsRequest) ([]*models.EventCountStatisticsResp, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/vehicle/events/sensors", r)
	if err != nil {
		return nil, err
	}
	var result []*models.EventCountStatisticsResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) VehicleEventCountByRepo(request models.EventCountStatisticsRequest) ([]*models.EventCountStatisticsResp, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/vehicle/events/civil", r)
	if err != nil {
		return nil, err
	}
	var result []*models.EventCountStatisticsResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) FaceEventCountByTime(request models.EventCountStatisticsRequest) ([]*models.KeyValueCommon, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/face/events/durations", r)
	if err != nil {
		return nil, err
	}
	var result []*models.KeyValueCommon
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) FaceEventCountBySensor(request models.EventCountStatisticsRequest) ([]*models.EventCountStatisticsResp, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/face/events/sensors", r)
	if err != nil {
		return nil, err
	}
	var result []*models.EventCountStatisticsResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (archimedesClient ArchimedesClient) FaceEventCountByRepo(request models.EventCountStatisticsRequest) ([]*models.EventCountStatisticsResp, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	body, err := archimedesClient.HTTPClient.Post("/face/events/civil", r)
	if err != nil {
		return nil, err
	}
	var result []*models.EventCountStatisticsResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
