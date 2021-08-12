package clients

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"codes/gocodes/dg/utils/json"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/log"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type ThorClient struct {
	*HTTPClient
}

type ThorBaseConds struct {
	QueryId      string
	Offset       int32
	Limit        int32
	SortBy       string
	SortOrderAsc bool
}

type ThorSensorResponse struct {
	AllSize int
	Sensors []*models.Sensor
}

type ThorFaceRuleQuery struct {
	BaseConds *ThorBaseConds
	OrgId     string `json:"orgid"`
}

type ThorFaceRepoResponse struct {
	AllSize   int
	FaceRepos []*models.FaceRepo
}

type ThorFaceRuleResponse struct {
	AllSize   int
	FaceRules []*models.FaceRule
}

func NewThorClient(baseAddr string, timeout time.Duration) *ThorClient {
	ret := &ThorClient{
		HTTPClient: NewHTTPClient(baseAddr, timeout, nil),
	}
	ret.HTTPClient.defaultRespReader = ret.readResponse
	return ret
}

func (this *ThorClient) DeleteOrg(orgId string) error {
	_, err := this.Delete(fmt.Sprintf("/api/account/org?org_id=%v", orgId), nil)
	return errors.WithStack(err)
}

func (this *ThorClient) GetFaceRepos() (*ThorFaceRepoResponse, error) {
	resp, err := this.Get("/api/face/repos")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result ThorFaceRepoResponse
	err = json.Unmarshal(resp, &result)
	return &result, errors.WithStack(err)
}

func (this *ThorClient) GetFaceRulesEx(orgId string) ([]*models.FaceRule, error) {
	query := &ThorFaceRuleQuery{
		BaseConds: &ThorBaseConds{
			Limit: math.MaxInt32,
		},
		OrgId: orgId,
	}
	resp, err := this.SearchFaceRules(query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp.FaceRules, nil
}

func (this *ThorClient) SearchFaceRules(query *ThorFaceRuleQuery) (*ThorFaceRuleResponse, error) {
	resp, err := this.PostJSON("/api/bingo/rules", query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result ThorFaceRuleResponse
	err = json.Unmarshal(resp, &result)
	return &result, errors.WithStack(err)
}

func (this *ThorClient) SetFaceRule(rule *models.FaceRule) error {
	_, err := this.PostJSON("/api/bingo/rule", rule)
	return errors.WithStack(err)
}

func (this *ThorClient) StartFaceRule(sensorId string) error {
	url := fmt.Sprintf("/api/bingo/startrule?sensor_id=%v", sensorId)
	_, err := this.Get(url)
	return errors.WithStack(err)
}

func (this *ThorClient) StopFaceRule(sensorId string) error {
	url := fmt.Sprintf("/api/bingo/stoprule?sensor_id=%v", sensorId)
	_, err := this.Get(url)
	return errors.WithStack(err)
}

func (this *ThorClient) DeleteFaceRule(rule *models.FaceRule) error {
	_, err := this.DeleteJSON("/api/bingo/rule", rule)
	return errors.WithStack(err)
}

func (this *ThorClient) GetAllSensorsEx() ([]*models.Sensor, error) {
	resp, err := this.GetSensors("")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp.Sensors, nil
}

func (this *ThorClient) GetSensors(orgId string) (*ThorSensorResponse, error) {
	url := "/api/device/sensors"
	if orgId != "" {
		url = fmt.Sprintf("%v?org_id=%v", url, orgId)
	}
	resp, err := this.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result ThorSensorResponse
	err = json.Unmarshal(resp, &result)
	return &result, errors.WithStack(err)
}

func (this *ThorClient) GetSensor(sensorId string) (*models.Sensor, error) {
	resp, err := this.Get(fmt.Sprintf("/api/device/sensor?sensor_id=%v", sensorId))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result models.Sensor
	err = json.Unmarshal(resp, &result)
	return &result, errors.WithStack(err)
}

func (this *ThorClient) AddSensor(sensor *models.Sensor) error {
	_, err := this.PostJSON("/api/device/sensor", sensor)
	return errors.WithStack(err)
}

func (this *ThorClient) UpdateSensor(sensor *models.Sensor) error {
	_, err := this.PutJSON("/api/device/sensor", sensor)
	return errors.WithStack(err)
}

func (this *ThorClient) DeleteSensor(sensorId string) error {
	_, err := this.Delete(fmt.Sprintf("/api/device/sensor?id=%v", sensorId), nil)
	return err
}

func (this *ThorClient) GetSensorDefaultConfigParam(sensorType int) (string, error) {
	url := fmt.Sprintf("/api/device/defaultconfig?type=%d", sensorType)
	body, err := this.Get(url)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return strconv.Unquote(string(body))
}

func (this *ThorClient) readResponse(resp *http.Response, req *http.Request) ([]byte, error) {
	bytes, err := this.HTTPClient.readResponse(resp, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(bytes) == 0 {
		log.Warnln("Thor response body empty, ignore parsing body")
		return bytes, nil
	}
	cr := gjson.GetBytes(bytes, "Code")
	if c := cr.Int(); c != 1 {
		return nil, fmt.Errorf("Thor response error with code = %d", c)
	}
	result := gjson.GetBytes(bytes, "Data")
	var raw []byte
	if result.Index > 0 {
		raw = bytes[result.Index : result.Index+len(result.Raw)]
	} else {
		raw = []byte(result.Raw)
	}
	return raw, nil
}
