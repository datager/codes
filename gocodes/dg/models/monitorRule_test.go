package models

import (
	"codes/gocodes/dg/utils/uuid"
	"fmt"
	"time"

	"codes/gocodes/dg/utils/json"
	"codes/gocodes/dg/utils/log"
)

var (
	sensors []*Sensor
	repos   []*RepoInRule
)

func init() {
	sensors = make([]*Sensor, 0)
	sensors = append(sensors,
		&Sensor{
			SensorID: uuid.NewV4().String(),
		},
		&Sensor{
			SensorID: uuid.NewV4().String(),
		},
	)

	repos = make([]*RepoInRule, 0)
	repos = append(repos,
		&RepoInRule{
			RepoId:     uuid.NewV4().String(),
			Confidence: 0.7000000000000001,
			AlarmLevel: 1,
			AlarmVoice: "在逃库",
		},
		&RepoInRule{
			RepoId:     uuid.NewV4().String(),
			Confidence: 0.89000000000000001,
			AlarmLevel: 3,
			AlarmVoice: "吸毒库",
		},
	)

	//ExampleMockcreateMonitorRule()
	//ExampleMockQueryMonitorRuleList()
	//ExampleMockDeleteMonitorRule()
	//ExampleMockEditMonitorRule()
	//ExampleGetMonitorRuleDetail()
	ExampleBatchEditMonitorRuleStatus()
}

// 1. 添加布控: POST http://{{loki}}/monitorrule

// 请求
// body 中
// Model: MonitorRuleSet
// 示例 见 req
/*
{
    "MonitorRuleName":"逃犯吸毒布控",
    "SensorIDs":[
        "36878bc3-612a-4573-b917-8130626e7975",
        "ad96a203-ddbe-4727-85d0-32637cdc6c37"
    ],
    "RepoInRules":[
        {
            "RuleId":"",
            "RepoId":"",
            "RepoName":"",
            "Confidence":0.7,
            "AlarmLevel":1,
            "AlarmVoice":"在逃库",
            "CountLimit":0,
            "Status":0
        },
        {
            "RuleId":"",
            "RepoId":"",
            "RepoName":"",
            "Confidence":0.89,
            "AlarmLevel":3,
            "AlarmVoice":"吸毒库",
            "CountLimit":0,
            "Status":0
        }
    ],
    "TimeIsLong":true,
    "Comment":"这是一个布控, 有2个设备, 2个库"
}
*/

// 响应
// Model: 无
// 示例 见 resp
// 若成功: 返回 204, 无 body
// 若失败: 返回 500, 无 body
func ExampleMockcreateMonitorRule() {
	req := &MonitorRuleSet{
		MonitorRuleName: "逃犯吸毒布控",
		SensorIds:       []string{uuid.NewV4().String(), uuid.NewV4().String()},
		RepoInRules:     repos,
		TimeIsLong:      true,
		Comment:         "这是一个布控, 有2个设备, 2个库",
	}
	reqBody, _ := json.Marshal(req)
	log.Infoln("req: ", string(reqBody))

	log.Infof("resp: Code: 204/500")
}

// 2. 查询布控列表: POST http://{{loki}}/monitorruleslist

// 请求
// body 中
// Model: MonitorRuleQueryReq
// 示例 见 req
/*
{
    "MonitorRuleID":"e0bcf699-7bb0-467b-95a7-b3e95fb21b30",
    "MonitorRuleName":"逃犯吸毒布控",
    "SensorIDs":[
        "4bb4afbb-7016-4551-8af6-884cdd97d3b3",
        "d4817efe-439e-4663-966b-6b1749a9daeb"
    ],
    "FaceRepoIds":[
        "a15ebb33-5239-43fb-af3a-1cf180d9433b",
        "a2d4e424-5f70-4e90-be52-57fa28142ade"
    ],
    "Type":[
        2,
        1
    ],
    "Status":[
        1,
        2
    ],
    "Offset":10,
    "Limit":10,
}
*/

// 响应
// Model: QueryAndCountResult 内包MonitorRuleQueryRespList
// 示例 见 resp
// 若成功: 返回 200, + body
// 若失败: 返回 500, + body
/*
{
    "Rets":{
        "MonitorRuleID":"b66de794-b742-4994-a1b6-8ca164d5c3b5",
        "MonitorRuleName":"逃犯吸毒布控",
        "Type":2,
        "RepoTypeCount":[
            {
                "Type":2,
                "Count":9
            },
            {
                "Type":1,
                "Count":9
            }
        ],
        "SensorCount":12234,
        "TimeIsLong":true,
        "CreateUserName":"admin",
        "Status":1,
        "Offset":0,
        "Limit":0
    },
    "Total":10
}
*/

func ExampleMockQueryMonitorRuleList() {
	req := &MonitorRuleQueryReq{
		MonitorRuleId:   uuid.NewV4().String(),
		MonitorRuleName: "逃犯吸毒布控",
		SensorIds:       []string{uuid.NewV4().String(), uuid.NewV4().String()},
		FaceRepoIds:     []string{uuid.NewV4().String(), uuid.NewV4().String()},
		Type:            MonitorRuleTypeOnlyFaceRepo,
		Status:          MonitorRuleStatusOn,
		Offset:          10,
		Limit:           10,
	}
	reqBody, _ := json.Marshal(req)
	log.Infoln("req: ", string(reqBody))

	repoTypeCounts := make([]*RepoTypeCount, 0)
	repoTypeCounts = append(repoTypeCounts,
		&RepoTypeCount{
			Type:  RepoTypeFace,
			Count: 9,
		},
		&RepoTypeCount{
			Type:  RepoTypeVehicle,
			Count: 9,
		},
	)
	var resp = &QueryAndCountResult{
		Rets: &MonitorRuleQueryRespList{
			MonitorRuleId:   uuid.NewV4().String(),
			MonitorRuleName: "逃犯吸毒布控",
			Type:            []RepoType{RepoTypeFace},
			RepoTypeCount:   repoTypeCounts,
			SensorCount:     12234,
			TimeIsLong:      true,
			CreateUserName:  "admin",
			Status:          MonitorRuleStatusOn,
		},
		Total: 10,
	}
	respBody, _ := json.Marshal(resp)
	log.Infoln("resp: ", string(respBody))
}

// 3. 批量删除布控: DELETE http://{{loki}}/monitorrules

// 请求
// body 中
// Model: []string, 每一项为 MonitorRuleID
// 示例 见 req
/*
[
    "5bec9ac5-a0f4-4089-aa72-1ce152cf54c3",
    "a73290d8-5820-4a5a-943a-b789f4befb2a"
]
*/

// 响应
// Model: map[string]bool
// 示例 见 resp
// key为 RuleId, value 若为true则此项成功, 若false则此项失败
/*
{
    "5bec9ac5-a0f4-4089-aa72-1ce152cf54c3":true,
    "a73290d8-5820-4a5a-943a-b789f4befb2a":false
}
*/
func ExampleMockDeleteMonitorRule() {
	req := []string{
		uuid.NewV4().String(),
		uuid.NewV4().String(),
	}
	reqBody, _ := json.Marshal(req)
	log.Infoln("req: ", string(reqBody))

	resp := &map[string]error{
		uuid.NewV4().String(): fmt.Errorf("某原因导致此布控删除失败"),
		uuid.NewV4().String(): nil,
	}
	respBody, _ := json.Marshal(resp)
	log.Infoln("resp: ", string(respBody))
}

// 4. 编辑布控: PATCH http://{{loki}}/monitorrule/:monitorruleid
// 注意: 使用PATCH方式: 若只需要改部分字段, 有两种传参方式: 1. 只传部分字段; 2. 传全量字段, 其中不改变的字段传原值(而非默认空值)

// 请求
// body 中
// Model: MonitorRuleSet
// 示例 见 req
/*
{
    "MonitorRuleName":"逃犯吸毒布控",
    "SensorIDs":[
        "36878bc3-612a-4573-b917-8130626e7975",
        "ad96a203-ddbe-4727-85d0-32637cdc6c37"
    ],
    "RepoInRules":[
        {
            "RuleId":"",
            "RepoId":"b9055210-b94b-4727-971e-62eadf66cbfb",
            "RepoName":"",
            "Confidence":0.7,
            "AlarmLevel":1,
            "AlarmVoice":"在逃库",
            "CountLimit":0,
            "Status":0
        },
        {
            "RuleId":"",
            "RepoId":"bddfada2-e55d-41e4-beb3-c9bdb16c1954",
            "RepoName":"",
            "Confidence":0.89,
            "AlarmLevel":3,
            "AlarmVoice":"吸毒库",
            "CountLimit":0,
            "Status":0
        }
    ],
    "TimeIsLong":true,
    "Comment":"这是一个布控, 有2个设备, 2个库"
}
*/

// 响应
// Model: 无
// 示例 见 resp
// 若成功: 返回 204, 无 body
// 若失败: 返回 500, 无 body
func ExampleMockEditMonitorRule() {
	req := &MonitorRuleSet{
		MonitorRuleName: "逃犯吸毒布控",
		SensorIds:       []string{uuid.NewV4().String(), uuid.NewV4().String()},
		RepoInRules:     repos,
		TimeIsLong:      true,
		Comment:         "这是一个布控, 有2个设备, 2个库",
	}
	reqBody, _ := json.Marshal(req)
	log.Infoln("req: ", string(reqBody))

	log.Infoln("resp: Code: 204/500")
}

// 5. 布控详情查询 GET http://{{loki}}/monitorrule/:monitorruleid

// 请求
// param 中
// Model: monitorruleid
// 示例 见 req
/*
http://{{loki}}/monitorrule/e85c4c39-6bd5-4416-a56a-e4c2fcd505cb
*/

// 响应
// Model: MonitorRule
// 示例 见 resp
// 若成功: 返回 200 + body
// 若失败: 返回 500 + body
/*
{
    "MonitorRuleID":"fa8a6593-2e0c-44e1-bf99-ddebb92caf1b",
    "MonitorRuleName":"逃犯吸毒布控",
    "Timestamp":1559725252780,
    "Sensors":[
        {
            "Uts":"0001-01-01T00:00:00Z",
            "Timestamp":0,
            "OrgId":"",
            "OrgName":"",
            "Id":"5263da76-64f6-43bc-8a93-891d63be88ba",
            "Name":"",
            "SerialID":"",
            "Type":0,
            "Status":0,
            "Longitude":0,
            "Latitude":0,
            "Ip":"",
            "Port":"",
            "Url":"",
            "RenderedUrl":"",
            "RtmpUrl":"",
            "FtpAddr":"",
            "FtpDir":"",
            "GateIP":"",
            "Comment":"",
            "ConfigJson":"",
            "OlympusId":"",
            "GateThreshold":0,
            "MserverAddr":"",
            "AlgType":"",
            "EurThreshold":0,
            "CosThreshold":0,
            "ProcessCapturedReid":0,
            "ProcessWarnedReid":0,
            "WindowSize":0,
            "Properties":null,
            "Gb28181Param":null,
            "Speed":0,
            "NvrAddr":"",
            "FlatformSensor":0
        },
        {
            "Uts":"0001-01-01T00:00:00Z",
            "Timestamp":0,
            "OrgId":"",
            "OrgName":"",
            "Id":"b9eedd4a-b151-41ef-a0f9-9fe9814f3d0a",
            "Name":"",
            "SerialID":"",
            "Type":0,
            "Status":0,
            "Longitude":0,
            "Latitude":0,
            "Ip":"",
            "Port":"",
            "Url":"",
            "RenderedUrl":"",
            "RtmpUrl":"",
            "FtpAddr":"",
            "FtpDir":"",
            "GateIP":"",
            "Comment":"",
            "ConfigJson":"",
            "OlympusId":"",
            "GateThreshold":0,
            "MserverAddr":"",
            "AlgType":"",
            "EurThreshold":0,
            "CosThreshold":0,
            "ProcessCapturedReid":0,
            "ProcessWarnedReid":0,
            "WindowSize":0,
            "Properties":null,
            "Gb28181Param":null,
            "Speed":0,
            "NvrAddr":"",
            "FlatformSensor":0
        }
    ],
    "FaceRepos":[
        {
            "RuleId":"",
            "RepoId":"550b1cbe-38c9-4636-a94b-7d8ef8d2622b",
            "RepoName":"",
            "Confidence":0.7,
            "AlarmLevel":1,
            "AlarmVoice":"在逃库",
            "CountLimit":0,
            "Status":0
        },
        {
            "RuleId":"",
            "RepoId":"89f30d25-7afa-402d-9e37-06382781ef78",
            "RepoName":"",
            "Confidence":0.89,
            "AlarmLevel":3,
            "AlarmVoice":"吸毒库",
            "CountLimit":0,
            "Status":0
        }
    ],
    "TimeIsLong":true,
    "Comment":"这是一个布控, 有2个设备, 2个库",
    "Type":2,
    "Status":1,
    "UserID":"b638f7a5-745a-4ccc-ae49-95f7588b3d41",
    "OrgId":"1eaa8790-c40d-460b-87f9-d0c12a7be632"
}
*/
func ExampleGetMonitorRuleDetail() {
	log.Infoln("req: ", "http://{{loki}}/monitorrule/:monitorruleid/e85c4c39-6bd5-4416-a56a-e4c2fcd505cb")

	resp := &MonitorRule{
		MonitorRuleId:   uuid.NewV4().String(),
		MonitorRuleName: "逃犯吸毒布控",
		Timestamp:       time.Now().UnixNano() / 1e6,
		Sensors:         sensors,
		FaceRepos:       repos,
		TimeIsLong:      true,
		Comment:         "这是一个布控, 有2个设备, 2个库",
		Type:            MonitorRuleTypeOnlyFaceRepo,
		Status:          MonitorRuleStatusOn,
		UserID:          uuid.NewV4().String(),
		OrgID:           uuid.NewV4().String(),
	}
	respBody, _ := json.Marshal(resp)
	log.Infoln("resp: ", string(respBody))
}

// 6. 批量编辑 布控状态: PATCH http://{{loki}}/monitorrules/status

// 请求
// body 中
// Model: []*MonitorRuleBatchEditStatusReq
// 示例 见 req
/*
[
    {
        "MonitorRuleID":"1c662012-6cb5-44f6-9136-f484f053a8f8",
        "DstEditStatus":1
    },
    {
        "MonitorRuleID":"8d7da8fa-7537-4544-9200-9d4608ee96c6",
        "DstEditStatus":1
    }
]
*/

// 响应
// Model: map[string]bool
// 示例 见 resp
// key为 MonitorRuleId, value 若为true则此项成功, 若false则此项失败
/*
{
    "5bec9ac5-a0f4-4089-aa72-1ce152cf54c3":true,
    "a73290d8-5820-4a5a-943a-b789f4befb2a":false
}
*/
func ExampleBatchEditMonitorRuleStatus() {
	req := make([]*MonitorRuleBatchEditStatusReq, 0)
	req = append(req,
		&MonitorRuleBatchEditStatusReq{
			MonitorRuleID: uuid.NewV4().String(),
			DstEditStatus: MonitorRuleStatusOn,
		},
		&MonitorRuleBatchEditStatusReq{
			MonitorRuleID: uuid.NewV4().String(),
			DstEditStatus: MonitorRuleStatusOn,
		})
	reqBody, _ := json.Marshal(req)
	log.Infoln("req: ", string(reqBody))
}
