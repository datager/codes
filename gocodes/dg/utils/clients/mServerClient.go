package clients

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/json"
	"codes/gocodes/dg/utils/log"

	"github.com/pkg/errors"
)

const (
	RtspPrefix       = "rtsp"
	NetposaVodPrefix = "netposa_vod"

	SUCCESS = "OK"
)

type MServerClient struct {
	*HTTPClient
	outAddr string
}

func NewMServerClient(baseAddr string, timeout time.Duration, outAddr string) *MServerClient {
	ret := &MServerClient{
		HTTPClient: NewHTTPClient(baseAddr, timeout, nil),
		outAddr:    outAddr,
	}

	ret.HTTPClient.defaultRespReader = ret.readResponse
	return ret
}

type MclusterBase struct {
	URI   string                 `json:"uri"`
	Param map[string]interface{} `json:"param"`
}

type MclusterRequest struct {
	Input      MclusterBase `json:"input"`
	Output     MclusterBase `json:"output"`
	StreamType string       `json:"stream_type"`
	Lazyload   bool         `json:"lazyload"`
	StartTime  int64        `json:"start_time,omitempty"`
	EndTime    int64        `json:"end_time,omitempyt"`
}

type RunRequest struct {
	Input        string
	Output       string
	Args         string
	Type         int
	Gb28181Param models.GB28181Param
}

type Stream struct {
	StreamID   string `json:"stream_id"`
	StreamType string `json:"stream_type"`
	ServerID   string `json:"server_id"`
	Time       string `json:"time"`
	Report     string `json:"report"`
	UserAgent  string `json:"user_agent"`
	Vps        string `json:"vps"`
	Sps        string `json:"sps"`
	Pps        string `json:"pps"`
	MserverIP  string `json:"mserver_ip"`
}

func getBaseRequest(req RunRequest) *MclusterRequest {
	request := &MclusterRequest{
		StreamType: "live",
		Lazyload:   true,
		Input: MclusterBase{
			Param: make(map[string]interface{}),
			URI:   req.Input,
		},
		Output: MclusterBase{
			Param: make(map[string]interface{}),
			URI:   req.Output,
		},
	}

	if req.Type == models.SensorType_Sensor_Type_Ipc || req.Type == models.SensorType_Sensor_Type_Face || req.Type == models.SensorType_Sensor_Type_Capture {
		if strings.HasPrefix(req.Input, RtspPrefix) {
			request.Input.Param["rtsp_transport"] = "tcp"
		}

		if strings.HasPrefix(req.Input, NetposaVodPrefix) {
			request.StreamType = "vod"

			urlInput := strings.Replace(req.Input, NetposaVodPrefix, "netposa", 1)
			urlStruct, err := url.Parse(urlInput)
			if err != nil {
				log.Errorln(err)
				return nil
			}

			query := urlStruct.Query()
			bstr := query.Get("beginTime")[:14]
			layout := "20060102150405"
			btime, err := time.ParseInLocation(layout, bstr, time.Local)
			if err != nil {
				log.Errorln(err)
			}

			request.StartTime = btime.Unix() * 1000

			endTimeStr := query.Get("endTime")[:14]
			endTime, err := time.ParseInLocation(layout, endTimeStr, time.Local)
			if err != nil {
				log.Errorln(err)
			}

			request.EndTime = endTime.Unix() * 1000

			urlStruct.RawQuery = ""
			request.Input.URI = urlStruct.String()
		}
	}

	if req.Type == models.SensorType_Sensor_Type_GB28181 {
		request.Input.Param["local_ip"] = req.Gb28181Param.LocalIp
		request.Input.Param["sip_ipproto"] = req.Gb28181Param.SipIpproto
		request.Input.Param["media_ipproto"] = req.Gb28181Param.MediaIpproto
		request.Input.Param["local_port"] = strconv.Itoa(req.Gb28181Param.LocalPort)

		if req.Gb28181Param.MediaId == "" {
			request.Input.Param["media_ip"] = "0.0.0.0"
		} else {
			request.Input.Param["media_ip"] = req.Gb28181Param.MediaIp
		}

		request.Input.Param["server_id"] = req.Gb28181Param.ServerId
		request.Input.URI = req.Input + "/" + req.Gb28181Param.MediaId
	}

	log.Debugln("mserver request:", *request)

	return request
}

func (mc *MServerClient) CheckAndStartStream(sensor *models.Sensor) (url string, err error) {
	if should := mc.ShouldRegister(sensor.Type, sensor.Url); should && sensor.FlatformSensor != models.Flatform_Sensor_type_Outside {
		var gbParam models.GB28181Param
		if sensor.Gb28181Param == nil {
			gbParam = models.GB28181Param{}
		} else {
			gbParam = *sensor.Gb28181Param
		}

		req := RunRequest{
			Input:        sensor.Url,
			Output:       "dmi://" + mc.GetOutAddr() + "/live/" + sensor.SensorID,
			Args:         "",
			Type:         sensor.Type,
			Gb28181Param: gbParam,
		}

		mserverIP, err := mc.StartStream(req)
		if err != nil || mserverIP == "" {
			return "", errors.WithStack(err)
		}

		sensor.MServerAddr = mserverIP
		return fmt.Sprintf("rtsp://%s/live/%s", mserverIP, sensor.SensorID), nil
	}
	return "", models.NewErrorf(models.ErrorCodeSensorPreviewNotSupport, fmt.Sprintf("sensor cannot start stream because should not register, sensor_id: %s, sensor_type: %d, sensor_url: %s, flatformSensor: %d", sensor.SensorID, sensor.Type, sensor.Url, sensor.FlatformSensor))
}

func (mc *MServerClient) StartStream(req RunRequest) (string, error) {

	request := getBaseRequest(req)

	resp, err := mc.PostJSON("/run", *request)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(resp), nil
}

func (mc *MServerClient) CheckAndStopStream(sensor *models.Sensor, forceStopStream bool) error {
	if forceStopStream {
		if err := mc.StopStream(sensor.SensorID); err != nil {
			return errors.WithStack(err)
		}
	} else if mc.ShouldRegister(sensor.Type, sensor.Url) && sensor.FlatformSensor != models.Flatform_Sensor_type_Outside {
		if err := mc.StopStream(sensor.SensorID); err != nil {
			return errors.WithStack(err)
		}
	}
	log.Infof("sensor cannot stop stream because should not register, sensor_id: %s, sensor_type: %d, sensor_url: %s", sensor.SensorID, sensor.Type, sensor.Url)
	return nil // no need to stop, just skip
}

func (mc *MServerClient) StopStream(streamID string) error {

	exist, err := mc.CheckExist(streamID)
	if err != nil {
		return errors.WithStack(err)
	}

	if !exist {
		return nil
	}

	url := fmt.Sprintf("/stop?stream_id=%v", streamID)
	resp, err := mc.Get(url)
	if err != nil {
		return errors.WithStack(err)
	}

	if string(resp) != SUCCESS {
		return fmt.Errorf("stop stream id:%v err:%v", streamID, string(resp))
	}

	return nil
}

func (mc *MServerClient) CheckExist(id string) (bool, error) {
	stream, err := mc.GetStreamByID(id)
	if err != nil && strings.Contains(err.Error(), "404") {
		return false, nil
	} else if err != nil {
		return false, errors.WithStack(err)
	}

	if stream.StreamID != "" {
		return true, nil
	}

	return false, nil
}

func (mc *MServerClient) GetStreamByID(id string) (*Stream, error) {
	stream := new(Stream)
	url := fmt.Sprintf("/status?stream_id=%v", id)

	resp, err := mc.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = json.Unmarshal(resp, stream)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stream, nil
}

func (mc *MServerClient) ShouldRegister(sensorType int, url string) bool {
	if sensorType == models.SensorType_Sensor_Type_Ipc {
		return true
	}

	if sensorType == models.SensorType_Sensor_Type_GB28181 || sensorType == models.SensorType_Sensor_Type_Netposa_PVG {
		return false
	}

	if url == "" {
		return false
	}

	if sensorType == models.SensorType_Sensor_Type_Face || sensorType == models.SensorType_Sensor_Type_Capture ||
		sensorType == models.SensorType_Sensor_Type_WithID_Device || sensorType == models.SensorType_Sensor_Type_WithoutID_Device ||
		sensorType == models.SensorType_Sensor_Type_HuiMu || sensorType == models.SensorType_Sensor_Type_HaiKang ||
		sensorType == models.SensorType_Sensor_Type_HuaWei || sensorType == models.SensorTypeDC84 {
		return true
	}

	return false
}

func (mc *MServerClient) GetOutAddr() string {
	return mc.outAddr
}
