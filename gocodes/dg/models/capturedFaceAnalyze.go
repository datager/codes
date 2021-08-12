package models

import "sync"

type CapturedFaceAnalyzeParam struct {
	StartTimestamp int64
	EndTimestamp   int64
	SpanSecond     int
	LeastDays      int
	Vid            string
	Type           int `required:"true" description:"[1]人脸落脚点 [2]车辆落脚点"`
}

type FootholdAnalyze struct {
	SensorID   string `json:"Id"`
	SensorName string `json:"Name"`
	Longitude  float64
	Latitude   float64

	Detail []PerSensorData
}

type PerSensorData struct {
	Start        int
	End          int
	TimeKey      string
	NumberOfDays int
	Times        int
}

type PerSensorDataList []PerSensorData

func (p PerSensorDataList) Len() int {
	return len(p)
}

func (p PerSensorDataList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PerSensorDataList) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}

type AccompanyAnalyzeParam struct {
	Vid               string
	StartTimestamp    int64
	EndTimestamp      int64
	LeastTimes        int
	LeastSensorCount  int
	AccompanyDuration int
}

type AccompanyMeta struct {
	TargetReIdMap       map[string]string
	SensorMap           map[string]string
	AccompanyVid        string
	AccompanyExtendList []*FacesIndexExtend
	AccompanyReIdMap    map[string]string
}

// 车辆伴随结果
type VehicleAccompanyMeta struct {
	TargetReIDMap       map[string]string `json:"TargetReIdMap"`
	SensorMap           map[string]string
	AccompanyVid        string
	AccompanyExtendList []*VehicleAnalyzeData
	AccompanyReIDMap    map[string]string `json:"AccompanyReIdMap"`
}

type VehicleAccompanyMap struct {
	TargetReIDMap       *sync.Map
	SensorMap           *sync.Map
	AccompanyVid        string
	AccompanyExtendList []*VehicleAnalyzeData
	AccompanyReIDMap    *sync.Map
}

func (v *VehicleAccompanyMap) ToMeta() *VehicleAccompanyMeta {
	if v == nil {
		return nil
	}

	reIDMap := make(map[string]string)
	v.TargetReIDMap.Range(func(k, v interface{}) bool {
		reIDMap[k.(string)] = v.(string)
		return true
	})

	sensorMap := make(map[string]string)
	v.SensorMap.Range(func(k, v interface{}) bool {
		sensorMap[k.(string)] = v.(string)
		return true
	})

	accReIDMap := make(map[string]string)
	v.AccompanyReIDMap.Range(func(k, v interface{}) bool {
		accReIDMap[k.(string)] = v.(string)
		return true
	})

	return &VehicleAccompanyMeta{
		TargetReIDMap:       reIDMap,
		SensorMap:           sensorMap,
		AccompanyVid:        v.AccompanyVid,
		AccompanyExtendList: v.AccompanyExtendList,
		AccompanyReIDMap:    accReIDMap,
	}
}

type VehicleAccompanyResult struct {
	Times                int
	SensorCount          int
	AccompanyVid         string
	AccompanyVehicleFile *VehicleFile
	TargetReidList       []string
	AccompanyReidList    []string
}

type VehicleAccompanyResultList []VehicleAccompanyResult

func (a VehicleAccompanyResultList) Len() int {
	return len(a)
}

func (a VehicleAccompanyResultList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a VehicleAccompanyResultList) Less(i, j int) bool {
	return a[i].Times > a[j].Times
}

type VehicleAccompany struct {
	Rets  []VehicleAccompanyResult
	Total int
}

type VehicleAccompaynDetail struct {
	Target    *VehiclesIndexExtend
	Accompany *VehiclesIndexExtend
}

// 人脸伴随结果
type AccompanyResult struct {
	Times               int
	SensorCount         int
	AccompanyVid        string
	AccompanyPersonFile *PersonFile
	TargetReidList      []string
	AccompanyReidList   []string
}

type AccompanyResultList []AccompanyResult

func (a AccompanyResultList) Len() int {
	return len(a)
}

func (a AccompanyResultList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a AccompanyResultList) Less(i, j int) bool {
	return a[i].Times > a[j].Times
}

type AccompanyDetailParam struct {
	TargetReidList    []string
	AccompanyReidList []string
}

type AccompanyDetail struct {
	Target    *FacesIndexExtend
	Accompany *FacesIndexExtend
}
