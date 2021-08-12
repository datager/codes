package clients

import (
	"fmt"
	"codes/gocodes/dg/utils/json"
	"time"
)

type PseClient struct {
	*HTTPClient
}

type AddPlateParam struct {
	PlateText string
}

type PseResponse struct {
	Code       int64
	Message    string
	InnerError error
}

func NewPseClient(baseAddr string, timeout time.Duration) *PseClient {
	return &PseClient{
		HTTPClient: NewHTTPClient(baseAddr, timeout, nil),
	}
}

func (pc *PseClient) AddRepo(repoID string) error {
	url := fmt.Sprintf("/repos/%v", repoID)
	resp, err := pc.PostJSON(url, nil)
	if err != nil {
		ret := new(PseResponse)
		_ = json.Unmarshal(resp, ret)

		e := fmt.Errorf("err:%v code:%v message:%v", err, ret.Code, ret.Message)
		return e
	}

	return nil
}

func (pc *PseClient) DeleteRepo(repoID string) error {
	url := fmt.Sprintf("/repos/%v", repoID)
	resp, err := pc.Delete(url, nil)
	if err != nil {
		ret := new(PseResponse)
		_ = json.Unmarshal(resp, ret)

		e := fmt.Errorf("err:%v code:%v message:%v", err, ret.Code, ret.Message)
		return e
	}

	return nil

}

func (pc *PseClient) AddPlate(plateID, plateText, repoID string) error {
	url := fmt.Sprintf("/repos/%v/plates/%v", repoID, plateID)
	param := AddPlateParam{
		PlateText: plateText,
	}

	resp, err := pc.PostJSON(url, &param)
	if err != nil {
		ret := new(PseResponse)
		_ = json.Unmarshal(resp, ret)

		e := fmt.Errorf("err:%v code:%v message:%v", err, ret.Code, ret.Message)
		return e
	}

	return nil

}

func (pc *PseClient) DeletePlate(plateID, repoID string) error {
	url := fmt.Sprintf("/repos/%v/plates/%v", repoID, plateID)
	resp, err := pc.Delete(url, nil)
	if err != nil {
		ret := new(PseResponse)
		_ = json.Unmarshal(resp, ret)

		e := fmt.Errorf("err:%v code:%v message:%v", err, ret.Code, ret.Message)
		return e
	}

	return nil

}
