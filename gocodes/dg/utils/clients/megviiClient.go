package clients

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"codes/gocodes/dg/utils/json"
	"codes/gocodes/dg/utils/log"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type MegviiCient struct {
	*HTTPClient
	needTestData bool
}

func NewMegviiClient(baseAddr string, timeout time.Duration, needTestData bool) *MegviiCient {
	return &MegviiCient{
		HTTPClient:   NewHTTPClient(baseAddr, timeout, nil),
		needTestData: needTestData,
	}
}

type FaceGroup struct {
	Name     string
	ReadOnly bool
	Remark   string
}

type GroupRet struct {
	Items []*FaceGroup
}

type Rect struct {
	Height int
	Left   int
	Top    int
	Width  int
}

type FindParam struct {
	Image  string
	Groups []string
	Limit  int
}

type Photo struct {
	Id   int
	Rect string
	Url  string
}

type Subject struct {
	Category    string
	CertId      string
	District    string
	Gender      string
	GroupId     string
	Id          int64
	Labels      []interface{}
	Minority    string
	Name        string
	Passport    string
	Remark      string
	Timestamp   int64
	Trusted     bool
	VehicleInfo string
}

type FaceGroupInfo struct {
	Group   string
	Photo   *Photo
	Score   float32
	Subject *Subject
}

type FindFaceResult struct {
	Items []*FaceGroupInfo
}

func (mg *MegviiCient) NeedTestData() bool {
	return mg.needTestData
}

func (mg *MegviiCient) GetGroups() ([]*FaceGroup, error) {
	resp, err := mg.Get("/api/group/")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(GroupRet)
	result.Items = make([]*FaceGroup, 0)

	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result.Items, nil
}

func (mg *MegviiCient) FindFaces(faceImagePath string, binData string, limit int) ([]*FaceGroupInfo, error) {
	var path string

	defer func() {
		if path != "" {
			err := os.Remove(path)
			if err != nil {
				log.Errorln(err)
			}
		}
	}()

	url := "/api/face-search/"
	var content []byte

	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	filename := id.String()
	path = fmt.Sprintf("./%v.jpg", filename)

	if binData != "" {
		fileBase64, err := base64.StdEncoding.DecodeString(binData)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		content = fileBase64

	} else if faceImagePath != "" {
		resp, err := http.Get(faceImagePath)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		rd, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		content = rd
	}

	err = ioutil.WriteFile(path, content, 0666)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	faceImagePath = path

	resp, err := mg.PostFile(url, fmt.Sprintf("@%v", faceImagePath), faceImagePath, limit)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := new(FindFaceResult)
	result.Items = make([]*FaceGroupInfo, 0)
	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result.Items, nil
}

func (mg *MegviiCient) GetImageURL(url string) string {
	return fmt.Sprintf("%v%v", mg.baseAddr, url)
}
