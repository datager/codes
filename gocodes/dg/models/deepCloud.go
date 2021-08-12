package models

import "encoding/json"

type SensorAddDeepCloudRequest struct {
	DeviceType string
	DeviceName string
	DeviceID   string
	GroupID    string
	Comment    string
}

type SensorDeleteDeepCloudRequest struct {
	DeviceID string
}

type SensorDeepCloud struct {
	Cts           int64
	Uts           int64
	GroupID       string
	DeviceType    string
	LogicDeviceID string
	DeviceID      string
	DeviceName    string
	Comment       string
}

type DeepCloudResponse struct {
	Code int
	Msg  string
	Data json.RawMessage
}

type ImageSimilarDeepCloudRequest struct {
	Image      ImageQuery
	TopN       int
	Confidence float32
}

type SimilarDeepCloudResponse struct {
	Request struct {
		Image struct {
			ImageID string
			URL     string
			BinData string
			Feature string
		}
		TopN       int
		Confidence float32
	}
	Persons []PersonDeepCloud
}

type PersonDeepCloud struct {
	Age        int
	Cts        int64
	Uts        int64
	Confidence float32
	PersonID   string
	Gender     string
	ImageURL   string
}

type PersonChangeListRequest struct {
	StartTime      int64
	Limit          int
	OriginPersonID string
}

type ChangedPerson struct {
	Uts            int64
	MergedAge      int
	MergedGender   int
	OriginPersonID string
	MergedPersonID string
	OriginImageURL string
	MergedImageURL string
}
