package clients

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/config"

	"github.com/emirpasic/gods/maps/treebidimap"
	godsutils "github.com/emirpasic/gods/utils"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type DeepCloudClient struct {
	*HTTPClient
	on bool
}

func NewDeepCloudClient(conf config.Config) *DeepCloudClient {
	timeoutSecond := conf.GetIntOrDefault("services.deepcloud.timeout_second", 60)
	onOff := conf.GetBoolOrDefault("services.deepcloud.switch", true)
	client := DeepCloudClient{
		HTTPClient: NewHTTPClient(conf.GetString("services.deepcloud.addr"), time.Duration(timeoutSecond)*time.Second, nil),
		on:         onOff,
	}

	header := http.Header{}
	header["access_key"] = []string{conf.GetString("services.deepcloud.access_key")}
	header["secret_key"] = []string{conf.GetString("services.deepcloud.secret_key")}
	header["authkey"] = []string{conf.GetString("services.deepcloud.authkey")}

	client.HTTPClient.SetHeader(header)
	return &client
}

func (deepCloudClient DeepCloudClient) OnOff() bool {
	return deepCloudClient.on
}

func (deepCloudClient DeepCloudClient) SensorRegister(request models.SensorAddDeepCloudRequest) (*models.SensorDeepCloud, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	rs, err = deepCloudClient.HTTPClient.Post("/device/add", r)
	if err != nil {
		return nil, err
	}
	resp := models.DeepCloudResponse{}
	err = json.Unmarshal(rs, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		switch resp.Code {
		case 1203:
			err = nil
		default:
			err = errors.New("register sensor failed:" + resp.Msg)
		}
		if err != nil {
			return nil, err
		}
	}

	m := &models.SensorDeepCloud{}
	err = json.Unmarshal(resp.Data, m)
	return m, err
}

func (deepCloudClient DeepCloudClient) SensorCancellation(request models.SensorDeleteDeepCloudRequest) (*models.DeepCloudResponse, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	rs, err = deepCloudClient.HTTPClient.Post("/device/delete", r)
	if err != nil {
		return nil, err
	}
	resp := models.DeepCloudResponse{}
	err = json.Unmarshal(rs, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		switch resp.Code {
		case 9:
			err = nil
		default:
			err = errors.New("cancellation sensor failed:" + resp.Msg)
		}
	}
	return &resp, err
}

func (deepCloudClient DeepCloudClient) GetSimilar(request models.ImageSimilarDeepCloudRequest) (*models.SimilarDeepCloudResponse, error) {
	//Confidence值不为空值
	if request.Confidence == 0 {
		request.Confidence = 0.01
	}
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	rs, err = deepCloudClient.HTTPClient.Post("/data/person/similar", r)
	if err != nil {
		return nil, err
	}
	resp := models.DeepCloudResponse{}
	err = json.Unmarshal(rs, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		return nil, errors.New("get similar by image failed:" + resp.Msg)
	}

	similarResp := models.SimilarDeepCloudResponse{}
	err = json.Unmarshal(resp.Data, &similarResp)
	if err != nil {
		return nil, err
	}

	return &similarResp, nil
}

func (deepCloudClient DeepCloudClient) BatchSearch(queries []*SaxQuery) (*treebidimap.Map, error) {

	list := make([]*models.SimilarDeepCloudResponse, 0, len(queries))

	var g errgroup.Group
	var lock sync.Mutex
	for _, v := range queries {
		v := v
		g.Go(func() error {
			request := models.ImageSimilarDeepCloudRequest{
				Image:      models.ImageQuery{Feature: v.Feature},
				TopN:       v.TopX,
				Confidence: v.Confidence,
			}
			sdr, err := deepCloudClient.GetSimilar(request)
			if err != nil {
				return errors.WithStack(err)
			}

			lock.Lock()
			list = append(list, sdr)
			lock.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, errors.WithStack(err)
	}

	// sort value by score
	m := treebidimap.NewWith(godsutils.StringComparator, func(a, b interface{}) int {
		s1 := a.(*SaxCandidate).SimilarityDegree
		s2 := b.(*SaxCandidate).SimilarityDegree
		if s1 > s2 {
			return -1
		} else if s1 < s2 {
			return 1
		} else {
			return godsutils.StringComparator(a.(*SaxCandidate).Vid, b.(*SaxCandidate).Vid)
		}
	})

	for _, r := range list {
		for _, item := range r.Persons {
			if val, exist := m.Get(item.PersonID); exist {
				if item.Confidence > val.(*SaxCandidate).SimilarityDegree {
					t := &SaxCandidate{
						Vid:              item.PersonID,
						SimilarityDegree: item.Confidence,
					}
					m.Put(item.PersonID, t)
				}
			} else {
				t := &SaxCandidate{
					Vid:              item.PersonID,
					SimilarityDegree: item.Confidence,
				}
				m.Put(item.PersonID, t)
			}
		}
	}
	return m, nil
}

func (deepCloudClient DeepCloudClient) PersonChangeList(request models.PersonChangeListRequest) ([]*models.ChangedPerson, error) {
	rs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(rs)
	rs, err = deepCloudClient.HTTPClient.Post("/data/person/changelist", r)
	if err != nil {
		return nil, err
	}
	resp := models.DeepCloudResponse{}
	err = json.Unmarshal(rs, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		return nil, errors.New("get person changelist failed:" + resp.Msg)
	}

	if resp.Data == nil {
		return nil, nil
	}

	var list struct {
		Persons []*models.ChangedPerson
	}
	err = json.Unmarshal(resp.Data, &list)
	if err != nil {
		return nil, err
	}

	return list.Persons, nil
}

func (deepCloudClient *DeepCloudClient) SearchAllVidCombinesInGivenTime(startTime int64, endTimestamp int64) ([]*SaxVidCombine, error) {
	list := make([]*SaxVidCombine, 0, 1000)

	for {
		request := models.PersonChangeListRequest{
			StartTime: startTime,
			Limit:     1000,
		}

		persons, err := deepCloudClient.PersonChangeList(request)
		if err != nil {
			return list, err
		}

		if len(persons) == 0 {
			break
		}

		for _, v := range persons {
			t := &SaxVidCombine{
				SourceVid: v.OriginPersonID,
				TargetVid: v.MergedPersonID,
			}
			list = append(list, t)
			startTime = v.Uts
		}
		if len(persons) == 1 {
			break
		}
	}

	return list, nil
}
