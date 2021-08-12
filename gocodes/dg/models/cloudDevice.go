package models

type CachedCloudDevice struct {
	ID   string `json:"id"`   // from /etc/bigtoe/flags/device_id
	Name string `json:"name"` // from http://api.demo.nemoface.com/api/device/list by device_id, ak, sk
}

type CloudDeviceReq struct {
	Offset         int
	Limit          int
	LogicDeviceIDs []string
	DeviceNames    []string
	DeviceIDs      []string
}

type CloudDeviceResp struct {
	Count   int
	Devices []*Device
}

type Device struct {
	GroupID       string
	DeviceType    string
	LogicDeviceID string
	DeviceID      string
	DeviceName    string
}
