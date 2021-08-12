package clients

import (
	"bytes"
	"codes/gocodes/dg/configs"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"codes/gocodes/dg/utils/json"

	"codes/gocodes/dg/utils/log"

	"github.com/pkg/errors"
)

type HTTPClient struct {
	hc                http.Client
	baseAddr          string
	defaultRespReader HTTPResponseReader
	header            http.Header
	startTime         time.Time
}

type HTTPResponseReader func(*http.Response, *http.Request) ([]byte, error)

func NewHTTPClient(baseAddr string, timeout time.Duration, defaultRespReader HTTPResponseReader) *HTTPClient {
	ret := &HTTPClient{
		baseAddr: baseAddr,
		hc: http.Client{
			Timeout: timeout,
		},
	}
	if defaultRespReader == nil {
		ret.defaultRespReader = ret.readResponse
	} else {
		ret.defaultRespReader = defaultRespReader
	}
	return ret
}

func (hc *HTTPClient) SetHeader(header http.Header) {
	hc.header = header
}

func (hc *HTTPClient) Get(url string) ([]byte, error) {
	return hc.fetch("GET", url, nil, nil)
}

func (hc *HTTPClient) GetParams(url string, params map[string][]string) ([]byte, error) {
	if len(params) == 0 {
		return hc.fetch("GET", url, nil, nil)
	}
	var paramSlice []string
	for key, valSlice := range params {
		for _, val := range valSlice {
			paramSlice = append(paramSlice, key+"="+val)
		}
	}
	url = url + "?" + strings.Join(paramSlice, "&")
	return hc.fetch("GET", url, nil, nil)

}

func (hc *HTTPClient) PostJSON(url string, obj interface{}) ([]byte, error) {
	body, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return hc.Post(url, bytes.NewBuffer(body))
}

func (hc *HTTPClient) Post(url string, body io.Reader) ([]byte, error) {
	return hc.fetch("POST", url, body, nil)
}

func (hc *HTTPClient) PutJSON(url string, obj interface{}) ([]byte, error) {
	body, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return hc.Put(url, bytes.NewBuffer(body))
}

func (hc *HTTPClient) PostForm(url, key, value string) ([]byte, error) {
	return hc.form("POST", url, key, value, nil)
}

func (hc *HTTPClient) Put(url string, body io.Reader) ([]byte, error) {
	return hc.fetch("PUT", url, body, nil)
}

func (hc *HTTPClient) DeleteJSON(url string, obj interface{}) ([]byte, error) {
	body, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return hc.Delete(url, bytes.NewBuffer(body))
}

func (hc *HTTPClient) Delete(url string, body io.Reader) ([]byte, error) {
	return hc.fetch("DELETE", url, body, nil)
}

func (hc *HTTPClient) fetch(method, url string, body io.Reader, respReader HTTPResponseReader) ([]byte, error) {
	targetURL := fmt.Sprintf("%v%v", hc.baseAddr, url)

	if configs.GetInstance().Debug {
		hc.startTime = time.Now()
		log.Debugf("Sending %v request to %v", method, targetURL)
		log.Debugf("Body = %v", body)
	}

	req, err := http.NewRequest(method, targetURL, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header = hc.header
	resp, err := hc.hc.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	r := hc.getRespReader(respReader)
	return r(resp, req)
}

func (hc *HTTPClient) form(method, formURL, key, value string, respReader HTTPResponseReader) ([]byte, error) {
	targetURL := fmt.Sprintf("%v%v", hc.baseAddr, formURL)
	log.Debugf("Sending %v request to %v", method, targetURL)
	log.Debugf("form key =%v value=%v", key, value)

	data := url.Values{}
	data.Set(key, value)

	req, err := http.NewRequest(method, targetURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := hc.hc.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	r := hc.getRespReader(respReader)

	return r(resp, req)
}

func (hc *HTTPClient) fetchWithContentType(method, url, contentType string, body io.Reader, respReader HTTPResponseReader) ([]byte, error) {
	targetURL := fmt.Sprintf("%v%v", hc.baseAddr, url)

	if configs.GetInstance().Debug {
		log.Debugf("Sending %v requst to %v", method, targetURL)
		log.Debugf("Body = %v", body)
	}

	req, err := http.NewRequest(method, targetURL, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := hc.hc.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	r := hc.getRespReader(respReader)
	return r(resp, req)
}

func (hc *HTTPClient) PostFile(url, filename, filePath string, limit int) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()

	//创建一个模拟的form中的一个选项,这个form项现在是空的
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	//相当于现在还没选择文件, form项里选择文件的选项
	fileWriter, err := bodyWriter.CreateFormFile("image", filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if fileWriter, err = bodyWriter.CreateFormField("limit"); err != nil {
		return nil, errors.WithStack(err)
	}

	if _, err = fileWriter.Write([]byte(fmt.Sprintf("%v", limit))); err != nil {
		return nil, errors.WithStack(err)
	}

	//获取上传文件的类型,multipart/form-data; boundary=...
	contentType := bodyWriter.FormDataContentType()

	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
	bodyWriter.Close()

	//发送post请求到服务端
	resp, err := hc.fetchWithContentType("POST", url, contentType, bodyBuf, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return resp, nil
}

func (hc *HTTPClient) getRespReader(respReader HTTPResponseReader) HTTPResponseReader {
	if respReader != nil {
		return respReader
	}
	return hc.defaultRespReader
}

func (hc *HTTPClient) readResponse(resp *http.Response, req *http.Request) ([]byte, error) {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if configs.GetInstance().Debug {
		log.Debugf("Got response from %v %v, status code = %d, body = %v took = %v", req.Method, req.URL, resp.StatusCode, string(bodyBytes), time.Since(hc.startTime))
	}

	if resp.StatusCode/100 != 2 {
		return bodyBytes, fmt.Errorf("HTTP request to %v %v failed with status code %d", req.Method, req.URL, resp.StatusCode)
	}
	return bodyBytes, nil
}
