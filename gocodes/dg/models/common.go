package models

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"

	stackErr "github.com/pkg/errors"
)

const (
	TypeSearchCondition          = 1
	TypeSearchFeature            = 2
	TypeSearchConditionAndExport = 3
	TypeSearchFeatureAndExport   = 4
)

type SearchResult struct {
	Total      int64
	NextOffset int
	Rets       []interface{}
}

func (sr SearchResult) MarshalJSON() ([]byte, error) {
	if sr.Rets == nil {
		sr.Rets = make([]interface{}, 0)
	}
	return json.Marshal(struct {
		Total      int64
		NextOffset int
		Rets       []interface{}
	}{
		Total:      sr.Total,
		NextOffset: sr.NextOffset,
		Rets:       sr.Rets,
	})
}

type SearchListRequest struct {
	Type    int
	Payload json.RawMessage
}

// v1 search capture
func (m *SearchListRequest) ParseToPedestrianCondition() (*PedestrianConditionRequest, error) {
	result := new(PedestrianConditionRequest)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToNonmotorCondition() (*NonmotorConditionRequest, error) {
	result := new(NonmotorConditionRequest)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToFaceCondition() (*FaceConditionRequest, error) {
	result := new(FaceConditionRequest)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToVehicleCondition() (*VehicleConditionRequest, error) {
	result := new(VehicleConditionRequest)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToFeature() (*FeatureRequest, error) {
	result := new(FeatureRequest)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToCivilQuery() (*CivilQuery, error) {
	result := new(CivilQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, stackErr.WithStack(err)
}

func (m *SearchListRequest) ParseToFailedCivilQuery() (*FailedCivilQuery, error) {
	result := new(FailedCivilQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, stackErr.WithStack(err)
}

func (m *SearchListRequest) ParseToEventQuery() (*FaceEventQuery, error) {
	result := new(FaceEventQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, stackErr.WithStack(err)
}

func (m *SearchListRequest) ParseToFaceConditionExport() (*AthenaFaceCaptureExportQuery, error) {
	result := new(AthenaFaceCaptureExportQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToFeatureExport() (*AthenaExportQuery, error) {
	result := new(AthenaExportQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToPedestrianConditionExport() (*AthenaPedestrianCaptureExportQuery, error) {
	result := new(AthenaPedestrianCaptureExportQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToNonmotorConditionExport() (*AthenaNonmotorCaptureExportQuery, error) {
	result := new(AthenaNonmotorCaptureExportQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToVehicleConditionExport() (*AthenaVehicleCaptureExportQuery, error) {
	result := new(AthenaVehicleCaptureExportQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, err
}

func (m *SearchListRequest) ParseToEventQueryExport() (*AthenaFaceEventExportQuery, error) {
	result := new(AthenaFaceEventExportQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, stackErr.WithStack(err)
}

func (m *SearchListRequest) ParseToCivilsConditionAndExport() (*AthenaCivilsExportQuery, error) {
	result := new(AthenaCivilsExportQuery)
	err := json.Unmarshal(m.Payload, &result)
	return result, stackErr.WithStack(err)
}

type FeatureRequest struct {
	Images     []*ImageQuery
	Time       *StartAndEndTimestamp
	Confidence float32
	SensorIds  []string
	TopX       int
	Limit      int
	Offset     int
}

func (m *FeatureRequest) Check() error {
	if len(m.Images) == 0 {
		return errors.New("no images invalid")
	}
	errFlag := true
	for _, image := range m.Images {
		if !strings.EqualFold(image.Feature, "") || !strings.EqualFold(image.Url, "") {
			errFlag = false
		}
	}
	if errFlag {
		return errors.New("all image no feature or url")
	}
	if m.TopX == 0 {
		return errors.New("TopX shouldn't be zero")
	}
	return nil
}

type FeatureResult struct {
	ID               string
	Confidence       float32
	OriginImageIndex int
}

type FeatureResultSlice []*FeatureResult

func (featureResultSlice FeatureResultSlice) ToMap() FeatureResultMap {
	result := make(FeatureResultMap)
	for _, featureResult := range featureResultSlice {
		result[featureResult.ID] = featureResult
	}
	return result
}

func (featureResultSlice FeatureResultSlice) TopXConfidence(topX int) FeatureResultSlice {
	if len(featureResultSlice) <= topX {
		return featureResultSlice
	}
	sort.SliceStable(featureResultSlice, func(i, j int) bool {
		return featureResultSlice[i].Confidence > featureResultSlice[j].Confidence
	})
	return featureResultSlice[:topX]
}

type FeatureResultMap map[string]*FeatureResult

func (featureResultMap FeatureResultMap) ToSlice() FeatureResultSlice {
	var result FeatureResultSlice
	for _, featureResult := range featureResultMap {
		result = append(result, featureResult)
	}
	return result
}

func (featureResultMap FeatureResultMap) GetKeys() []string {
	var result []string
	for key := range featureResultMap {
		result = append(result, key)
	}
	return result
}

type SpecialCaptureCount struct {
	CaptureType int
	StartAndEndTimestamp
}

type StartAndEndTimestamp struct {
	StartTimestamp int64 // 开始时间
	EndTimestamp   int64 // 结束时间
}

type Pagination struct {
	Limit  int // 页面条数
	Offset int // 偏移量
}

func (pagination Pagination) NeedPagination() bool {
	return pagination.Limit != 0
}

type SensorTs struct {
	SensorID string
	Ts       int64
}

type RelationData struct {
	Type        int
	RelationID  string
	RelationURL string
}
