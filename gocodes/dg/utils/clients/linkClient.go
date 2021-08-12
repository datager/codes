package clients

import (
	"fmt"
	"strings"
	"time"

	"codes/gocodes/dg/utils/json"

	"github.com/pkg/errors"
)

type LinkClient struct {
	*HTTPClient
}

const (
	LinkTypeVideo          = "vse_video"    // vse全帧结构化
	LinkTypeKeyFrame       = "vse_keyframe" // vse超压结构化，只分析1帧
	LinkTypeImage          = "vse_image"    // vse图片流任务
	LinkTypeImporterFtp    = "importer_ftp"
	LinkTypeImporterLibraf = "importer_libraf"
)

const (
	EngineTypeFace    = "face"
	EngineTypeVehicle = "vehicle"
	EngineTypeAllObj  = "allobj" // 暂时不支持
)

const (
	SuccessCode = 0
)

const (
	QUEUEING = "QUEUEING" // 排队中
	RUNNING  = "RUNNING"  // 运行中
	OOQ      = "OOQ"      // 路数限制
	RETRYING = "RETRYING" // 重试中
	STOPPED  = "STOPPED"  // 已停止,一般是前端操作停止任务
	FAILED   = "FAILED"   // 任务失败
	EXITED   = "EXITED"   // 已完成，任务退出并且正常退出
	DELETED  = "DELETED"  // 已删除
)

func NewLinkClient(baseAddr string, timeout time.Duration) *LinkClient {
	return &LinkClient{
		HTTPClient: NewHTTPClient(baseAddr, timeout, nil),
	}
}

type LinkMeta struct {
	State        string
	Node         string
	Workers      int
	Progress     int
	ExitCode     string
	ErrorMessage string
	RenderURI    string `json:"render_uri"`
	Type         string
	Engine       string
	InstanceID   string `json:"instance_id"`
	Channel      int
	ConfigJSON   string `json:"config_json"`
	Resource     int
}

type LinkInstanceRet struct {
	Code int
	Data LinkMeta
}

type LinkInstanceRetList struct {
	Code int
	Data []*LinkMeta
}

type NodeLimit struct {
	Alive int
	Limit int
}

type LinkNodeLimit struct {
	Code int
	Data map[string]NodeLimit
}

type SummaryMeta struct {
	Function   string
	DetectType string `json:"detect_type"`
	Used       int
	Total      int
}

type NodeResource struct {
	ExpiredDate  int64          `json:"expired_date"`
	ResourceInfo []*SummaryMeta `json:"resource_info"`
}

type LinkResource struct {
	Summary []*SummaryMeta
	Details map[string]*NodeResource
}

func (lc *LinkClient) CreateLinkTask(linkType string, worker int, engineType string, configJSON string) (*LinkMeta, error) {
	linkURL := "/olympus/v1/instance?type=%v"

	linkURL = fmt.Sprintf(linkURL, linkType)
	if worker > 0 {
		linkURL += "&workers=%v"
		linkURL = fmt.Sprintf(linkURL, worker)
	}

	linkURL += "&engine=%v"
	linkURL = fmt.Sprintf(linkURL, engineType)

	rep, err := lc.PostForm(linkURL, "config_json", configJSON)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(LinkInstanceRet)

	err = json.Unmarshal(rep, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result.Code != SuccessCode {
		err := fmt.Errorf("create link task by type:%v worker:%v engine type:%v config json:%v error msg:%v", linkType, worker, engineType, configJSON, result.Data.ErrorMessage)
		return nil, err
	}

	return &result.Data, nil
}

func (lc *LinkClient) DeleteLinkTask(linkID string) (*LinkMeta, error) {
	linkURL := fmt.Sprintf("/olympus/v1/instance/delete?iid=%v", linkID)

	resp, err := lc.PostJSON(linkURL, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(LinkInstanceRet)

	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result.Code != SuccessCode {
		err := fmt.Errorf("delete link task by id:%v error msg:%v", linkID, result.Data.ErrorMessage)
		return nil, err
	}

	return &result.Data, nil
}

func (lc *LinkClient) GetLinkTaskByID(linkID string) (*LinkMeta, error) {
	linkURL := fmt.Sprintf("/olympus/v1/instance?iid=%v", linkID)

	resp, err := lc.Get(linkURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(LinkInstanceRet)

	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result.Code != SuccessCode {
		err := fmt.Errorf("get link task by id:%v error msg:%v", linkID, result.Data.ErrorMessage)
		return nil, errors.WithStack(err)
	}

	return &result.Data, nil
}

func (lc *LinkClient) GetLinkTasks() ([]*LinkMeta, error) {
	linkURL := "/olympus/v1/instance"

	resp, err := lc.Get(linkURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(LinkInstanceRetList)
	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result.Code != SuccessCode {
		err := fmt.Errorf("get link task list error msg:%v", result.Data[0].ErrorMessage)
		return nil, err
	}

	return result.Data, nil
}

func (lc *LinkClient) StartLinkTask(linkID string) (*LinkMeta, error) {
	linkURL := fmt.Sprintf("/olympus/v1/instance/start?iid=%v", linkID)

	resp, err := lc.PostJSON(linkURL, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(LinkInstanceRet)
	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result.Code != SuccessCode {
		err := fmt.Errorf("start link task by id:%v error msg:%v", linkID, result.Data.ErrorMessage)
		return nil, err
	}

	return &result.Data, nil
}

func (lc *LinkClient) StopLinkTask(linkID string) (*LinkMeta, error) {
	linkURL := fmt.Sprintf("/olympus/v1/instance/stop?iid=%v", linkID)

	resp, err := lc.PostJSON(linkURL, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(LinkInstanceRet)
	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result.Code != SuccessCode {
		err := fmt.Errorf("stop link task by id:%v error msg:%v", linkID, result.Data.ErrorMessage)
		return nil, err
	}

	return &result.Data, nil

}

func (lc *LinkClient) GetNodeLimit() (map[string]NodeLimit, error) {
	linkURL := fmt.Sprintf("/olympus/v1/instance/limit")

	resp, err := lc.Get(linkURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(LinkNodeLimit)
	result.Data = make(map[string]NodeLimit)
	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result.Code != SuccessCode {
		err := fmt.Errorf("get link node limit error code:%v", result.Code)
		return nil, err
	}

	return result.Data, nil
}

func (lc *LinkClient) UpdateLink(linkID, linkType, engineType, configJSON string, worker int) (*LinkMeta, error) {
	_, err := lc.DeleteLinkTask(linkID)
	if err != nil && !strings.Contains(err.Error(), "404") {
		return nil, err
	}

	return lc.CreateLinkTask(linkType, worker, engineType, configJSON)
}

func (lc *LinkClient) GetResources() (*LinkResource, error) {
	linkURL := fmt.Sprintf("/x-api/v1/resources")

	resp, err := lc.Get(linkURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := LinkResource{}
	result.Details = make(map[string]*NodeResource)
	result.Summary = make([]*SummaryMeta, 0)

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &result, nil
}
