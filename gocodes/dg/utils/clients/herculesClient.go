package clients

import (
	"fmt"
	"time"

	"codes/gocodes/dg/utils/json"

	"github.com/pkg/errors"
)

type HerculesClient struct {
	*HTTPClient
}

func NewHerculesClient(baseAddr string, timeout time.Duration) *HerculesClient {
	return &HerculesClient{
		HTTPClient: NewHTTPClient(baseAddr, timeout, nil),
	}
}

type HerculesStatus struct {
	SrcRepoID        string
	TargetRepoID     string
	DownloadID       string
	Threshold        float32
	MaxCandidatesNum int
	CurrentState     string
	GetDBDataStatus  *GetDBDataStatus
	ComparisonStatus *ComparisonStatus
	InnerErrorStatus *InnerErrorStatus
	RollBackStaus    *InputRollBackStatus
	InputPath        string
}

type GetDBDataStatus struct {
	InputTotalNum uint
	RightNum      uint
	ErrorNum      uint
}

type ComparisonStatus struct {
	InputTotalNum uint
	RightNum      uint
	ErrorNum      uint
}

type InnerErrorStatus struct {
	Errors []string
}

type InputRollBackStatus struct {
	RollBackTransformType string
	RollBackMethodTypes   []string
}

type HerculesTask struct {
	SrcRepoId    string
	TargetRepoId string
	TaskId       string
	ResultId     string
	Threshold    float32
}

type HerculesStatusMap struct {
	TaskStatusMap map[string]*HerculesStatus
}

func (this *HerculesClient) GetStatus() (map[string]*HerculesStatus, error) {
	resp, err := this.Get("/status")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(HerculesStatusMap)
	result.TaskStatusMap = make(map[string]*HerculesStatus)
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result.TaskStatusMap, nil
}

func (this *HerculesClient) CreateTask(param *HerculesTask) error {
	_, err := this.PostJSON("/create", param)

	return errors.WithStack(err)
}

func (this *HerculesClient) Reset(resourceId string) error {
	url := fmt.Sprintf("/reset?taskid=%v", resourceId)
	_, err := this.Get(url)

	return errors.WithStack(err)
}

func (this *HerculesClient) Download(resultID string) string {
	base := this.baseAddr
	return fmt.Sprintf("%v/download?resultid=%v", base, resultID)
}

func (this *HerculesClient) DeleteTask(resultID string) error {
	url := fmt.Sprintf("/deleteresult?resultid=%v", resultID)

	_, err := this.Get(url)
	return errors.WithStack(err)
}

func (this *HerculesClient) CancelTask(resourceId string) error {
	url := fmt.Sprintf("/cancel?taskid=%v", resourceId)
	_, err := this.Get(url)
	return errors.WithStack(err)
}
