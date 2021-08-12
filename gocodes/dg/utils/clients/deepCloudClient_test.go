package clients

import (
	"codes/gocodes/dg/utils/config"
)

var sensorID string
var LogicDeviceID string
var conf config.Config

// func TestReady(t *testing.T) {
// 	_, err := config.InitConfig("../../config")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	sensorID = uuid.GetUUID()
// 	conf = config.GetConfig()
// }

// func TestSensorRegister(t *testing.T) {
// 	client := NewDeepCloudClient(conf)

// 	add := models.SensorAddDeepCloudRequest{
// 		DeviceType: "nemo",
// 		DeviceName: "gotest测试摄像头" + sensorID,
// 		DeviceID:   sensorID,
// 		Comment:    "gotest测试摄像头",
// 	}

// 	m, err := client.SensorRegister(add)
// 	if err != nil {
// 		fmt.Println("add sensor failed:", err)
// 		t.Fail()
// 	} else {
// 		LogicDeviceID = m.LogicDeviceID
// 		fmt.Println("add sensor success")
// 	}
// }

// func TestSensorCancellation(t *testing.T) {
// 	client := NewDeepCloudClient(conf)

// 	del := models.SensorDeleteDeepCloudRequest{LogicDeviceID: LogicDeviceID}

// 	_, err := client.SensorCancellation(del)

// 	if err != nil {
// 		fmt.Println("del sensor failed:", err)
// 		t.Fail()
// 	} else {
// 		fmt.Println("del sensor success")
// 	}
// }

// func TestGetSimilar(t *testing.T) {
// 	client := NewDeepCloudClient(conf)

// 	request := models.ImageSimilarDeepCloudRequest{
// 		Image: models.ImageQuery{
// 			Feature: "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",
// 		},
// 		TopN:       10,
// 		Confidence: 0.7,
// 	}

// 	resp, err := client.GetSimilar(request)
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Fail()
// 	}

// 	if resp == nil {
// 		return
// 	}

// 	for _, v := range resp.Persons {
// 		fmt.Println(v.PersonID)
// 	}

// }
