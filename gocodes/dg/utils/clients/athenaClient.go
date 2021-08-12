package clients

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"codes/gocodes/dg/utils/log"
)

type AthenaClient struct {
	httpClient *HTTPClient
	muGB       sync.Mutex // 国标同步锁
}

func NewAthenaClient(baseAddr string, timeout time.Duration) *AthenaClient {
	return &AthenaClient{
		httpClient: NewHTTPClient(baseAddr, timeout, nil),
	}
}

// 创建配置
type CreateConfig struct {
	Type       int
	ConfigInfo interface{}
}

func NewCreateConfig(taskType int, config interface{}) *CreateConfig {
	return &CreateConfig{
		Type:       taskType,
		ConfigInfo: config,
	}
}

// 操作
type OperateConfig struct {
	Type int
}

const (
	OperateTypeStart  = 0 // 开启操作
	OperateTypeStop   = 1 // 停止操作
	OperateTypeCancel = 2 // 取消操作
	OperateTypeReset  = 3 // 重置操作
)

// 状态
type StatusResponse struct {
	ID         string          // 任务ID
	Type       int             // 任务类别
	ConfigInfo json.RawMessage // 任务配置信息
	StatusInfo json.RawMessage // 任务状态信息
}

// 错误
type ErrorInfo interface {
	Error() string
	GetCode() int
	IsTaskNotExist() bool
	IsTaskStateChanged() bool
	IsMaxTaskNum() bool
}

type ErrorResponse struct {
	Code    int    // 错误码
	Message string // 错误描述
}

const (
	ErrorCodeTaskUnknown      = 0      // 未知错误
	ErrorCodeTaskMaxNum       = 100000 // 任务上限(默认 10)
	ErrorCodeTaskNotExist     = 100001 // 任务不存在
	ErrorCodeTaskStateChanged = 100009 // 任务状态变化
)

func NewErrorResponseUnknown(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    ErrorCodeTaskUnknown,
		Message: err.Error(),
	}
}

func (errorResponse *ErrorResponse) GetCode() int {
	return errorResponse.Code
}

func (errorResponse *ErrorResponse) Error() string {
	return fmt.Sprintf("code %d message %s", errorResponse.Code, errorResponse.Message)
}

func (errorResponse *ErrorResponse) IsTaskNotExist() bool {
	return errorResponse.Code == ErrorCodeTaskNotExist
}

func (errorResponse *ErrorResponse) IsTaskStateChanged() bool {
	return errorResponse.Code == ErrorCodeTaskStateChanged
}

func (errorResponse *ErrorResponse) IsMaxTaskNum() bool {
	return errorResponse.Code == ErrorCodeTaskMaxNum
}

func (ac *AthenaClient) CreateTask(taskID string, config *CreateConfig) ErrorInfo {
	return ac.parseErrorResponse(ac.httpClient.PostJSON(fmt.Sprintf("/tasks/%s", taskID), config))
}

func (ac *AthenaClient) GetStatus(taskIDs ...string) (map[string]*StatusResponse, error) {
	var err error
	// 获取状态
	resp, err := ac.httpClient.GetParams("/tasks", map[string][]string{"taskID": taskIDs})
	if err != nil { // 请求错误
		return nil, errors.WithStack(err)
	}
	// 解析body
	rawMessageMap := make(map[string]json.RawMessage)
	err = jsoniter.Unmarshal(resp, &rawMessageMap)
	if err != nil {
		log.Debugf("athenaClient parse response error, response is %s, error is %s", string(resp), err.Error())
		return nil, fmt.Errorf("athenaClient parse response error, response is %s, error is %s", string(resp), err.Error())
	}
	result := make(map[string]*StatusResponse)
	for key, rawMessage := range rawMessageMap {
		statusResponse := new(StatusResponse)
		err = jsoniter.Unmarshal(rawMessage, &statusResponse)
		if err != nil {
			log.Debugf("athenaClient parse response error, response is %s, error is %s", string(resp), err.Error())
			return nil, fmt.Errorf("athenaClient parse response error, response is %s, error is %s", string(resp), err.Error())
		}
		result[key] = statusResponse
	}
	return result, nil
}

func (ac *AthenaClient) StartTask(taskID string) ErrorInfo {
	return ac.operateTask(taskID, OperateTypeStart)
}

func (ac *AthenaClient) StopTask(taskID string) ErrorInfo {
	return ac.operateTask(taskID, OperateTypeStop)
}

func (ac *AthenaClient) CancelTask(taskID string) ErrorInfo {
	return ac.operateTask(taskID, OperateTypeCancel)
}

func (ac *AthenaClient) ResetTask(taskID string) ErrorInfo {
	return ac.operateTask(taskID, OperateTypeReset)
}

func (ac *AthenaClient) operateTask(taskID string, operateType int) ErrorInfo {
	return ac.parseErrorResponse(ac.httpClient.PutJSON(fmt.Sprintf("/tasks/%s", taskID), OperateConfig{Type: operateType}))
}

func (ac *AthenaClient) parseErrorResponse(body []byte, err error) ErrorInfo {
	if err != nil {
		// 无返回body
		if len(body) == 0 {
			log.Debugf("athenaClient http error, status not ok but body is null")
			return NewErrorResponseUnknown(err)
		}
		// 解析body
		errorResponse := new(ErrorResponse)
		err = jsoniter.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Debugf("athenaClient parse response error, response is %s, error is %s", string(body), err.Error())
			return NewErrorResponseUnknown(err)
		}
		return errorResponse
	}
	return nil
}
