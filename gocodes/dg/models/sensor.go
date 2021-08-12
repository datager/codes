package models

import (
	"fmt"
	"time"

	"codes/gocodes/dg/utils/pool"

	"github.com/pkg/errors"
)

//设备状态常量
const (
	SensorStatusCreated   = 1 //已创建
	SensorStatusDeleting  = 2 //删除中
	SensorStatusModifying = 3 //修改中
	// 4枚举占位弃用: 之前很多repository 层查 sensor 时, 都过滤了 where status != 4的条件
	SensorStatusDeleted = 5 //已删除(GB平台专用)
)

const (
	SensorType_Sensor_Type_Unknown = iota
	SensorType_Sensor_Type_Face
	SensorType_Sensor_Type_Capture
	SensorType_Sensor_Type_Ipc
	SensorType_Sensor_Type_Video
	SensorType_Sensor_Type_Picture
	SensorType_Sensor_Type_Gate
	SensorType_Sensor_Type_WithID_Device
	SensorType_Sensor_Type_WithoutID_Device
	SensorType_Sensor_Type_Car
	SensorType_Sensor_Type_Netposa_PVG
	SensorType_Sensor_Type_GB28181
	SensorType_Sensor_Type_HuiMu
	SensorType_Sensor_Type_HaiKang
	SensorType_Sensor_Type_DaHua
	SensorType_Sensor_Type_HuaWei

	SensorTypeDC84 // dc84设备，变量名使用CamelCase
	SensorTypeDC83 // dc83设备，变量名使用CamelCase
)

const (
	Flatform_Sensor_type_Unknow = iota
	Flatform_Sensor_type_Local
	Flatform_Sensor_type_Outside
)

const (
	OuterPlatformStatusNormal  = 1 //正常
	OuterPlatformStatusDeleted = 2 //已删除
)

const (
	AlgType             string  = "EUR"
	EurThreshold        float32 = 5.5
	CosThreshold        float32 = 0.9
	ProcessCapturedReid int32   = 1
	ProcessWarnedReid   int32   = 1
	WindowSize          int32   = 20
)

type sensorValidator func(*Sensor) error

var sensorValidators map[int]sensorValidator

func init() {
	sensorValidators = make(map[int]sensorValidator)
	sensorValidators[SensorType_Sensor_Type_Face] = validate_LibraF
	sensorValidators[SensorType_Sensor_Type_Capture] = validate_Ftp
	sensorValidators[SensorType_Sensor_Type_Ipc] = validate_Vsd
	sensorValidators[SensorType_Sensor_Type_Video] = validate_Vsd
	sensorValidators[SensorType_Sensor_Type_Picture] = validate_Ftp
	sensorValidators[SensorType_Sensor_Type_WithID_Device] = validate_Ftp
	sensorValidators[SensorType_Sensor_Type_GB28181] = validate_GB28181
	sensorValidators[SensorType_Sensor_Type_HuiMu] = validate_Importer
	sensorValidators[SensorType_Sensor_Type_HaiKang] = validate_Importer
	sensorValidators[SensorType_Sensor_Type_HuaWei] = validate_Importer
	sensorValidators[SensorType_Sensor_Type_DaHua] = validate_Importer
	sensorValidators[SensorTypeDC84] = validate_Importer
}

type Sensor struct {
	Uts                 time.Time
	Timestamp           int64
	OrgId               string
	OrgName             string
	OrgCode             int64
	SensorID            string `json:"Id"`
	SensorName          string `json:"Name"`
	SerialID            string `json:"SerialId"`
	Type                int
	Status              int
	Longitude           float64
	Latitude            float64
	Ip                  string
	Port                string
	Url                 string
	RenderedUrl         string
	RtmpUrl             string
	FtpAddr             string
	FtpDir              string
	GateIP              string
	Comment             string
	ConfigJson          string `json:"-"`
	OlympusId           string
	GateThreshold       float32
	MServerAddr         string `json:"MserverAddr"` // thor compatibility
	AlgType             string
	EurThreshold        float32
	CosThreshold        float32
	ProcessCapturedReid int
	ProcessWarnedReid   int
	WindowSize          int
	Properties          map[string]interface{}
	Gb28181Param        *GB28181Param
	NvrAddr             string
	FlatformSensor      int
	OuterPlatformStatus int
	UnBindTask          bool
	SensorIntID         int32 `json:"-"`
	AdditionalInfo      string
	UniqueID            string
}

func (sensor *Sensor) IsNonStableStatus() bool {
	return sensor.Status == SensorStatusDeleting || sensor.Status == SensorStatusModifying
}

type SensorInfo struct {
	SensorID       string
	SensorName     string
	SensorIntID    int32
	SensorStatus   int
	AdditionalInfo string
	Longitude      float64
	Latitude       float64
}

func SensorIDsToSensors(sensorIDs []string) []*Sensor {
	sensors := make([]*Sensor, 0)
	for _, id := range sensorIDs {
		sensors = append(sensors, &Sensor{
			SensorID: id,
		})
	}
	return sensors
}

//为组织结构提供的Sensor结构，会包括相关的设备，布控以及将来可添加的设备相关信息
type OrgSensor struct {
	*Sensor
	Task        *Tasks
	MonitorRule *MonitorRule
}

type GB28181Param struct {
	ServerId     string
	MediaId      string
	LocalIp      string
	LocalPort    int
	MediaIp      string
	SipIpproto   string
	MediaIpproto string
}

func (sensor *Sensor) Validate() error {
	// todo validate other fields
	if validator, exist := sensorValidators[sensor.Type]; exist {
		return validator(sensor)
	}
	return nil
}

func validate_Ftp(sensor *Sensor) error {
	if err := validatePropertyString(sensor, "IP"); err != nil {
		return errors.WithStack(err)
	}
	if err := validatePropertyFloat64(sensor, "Port"); err != nil {
		return errors.WithStack(err)
	}
	if err := validatePropertyString(sensor, "UserName"); err != nil {
		return errors.WithStack(err)
	}
	if err := validatePropertyString(sensor, "Path"); err != nil {
		return errors.WithStack(err)
	}
	if err := validatePropertyFloat64(sensor, "ImgRate"); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func validate_LibraF(sensor *Sensor) error {
	if err := validatePropertyString(sensor, "Port"); err != nil {
		return errors.WithStack(err)
	}
	if err := validatePropertyString(sensor, "PublicKey"); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func validate_Vsd(sensor *Sensor) error {
	if sensor.Url != "" {
		// not a video file
		return nil
	}
	if err := validatePropertyString(sensor, "FilePath"); err != nil {
		return errors.WithStack(err)
	}
	if err := validatePropertyString(sensor, "DateTime"); err != nil {
		return errors.WithStack(err)
	}
	// 设备与倍速无关，无需验证
	// if sensor.Speed < 1 {
	// 	return fmt.Errorf("Invalid speed")
	// }
	return nil
}

func validate_GB28181(sensor *Sensor) error {
	// todo gb28181
	return nil
}

func validate_Importer(sensor *Sensor) error {
	if err := validatePropertyString(sensor, "IP"); err != nil {
		return errors.WithStack(err)
	}
	if err := validatePropertyFloat64(sensor, "Port"); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func validatePropertyString(sensor *Sensor, property string) error {
	if val, ok := sensor.Properties[property]; ok {
		if str, ok := val.(string); !ok || str == "" {
			return newPropertyError(property)
		}
	}
	return nil
}

func validatePropertyFloat64(sensor *Sensor, property string) error {
	if val, ok := sensor.Properties[property]; ok {
		if _, ok := val.(float64); !ok {
			return newPropertyError(property)
		}
	}
	return nil
}

func newPropertyError(property string) error {
	return fmt.Errorf("Invalid property %v", property)
}

func SensorToSensorIDs(sensors []*Sensor) []string {
	sIDs := make([]string, 0)
	for _, s := range sensors {
		sIDs = append(sIDs, s.SensorID)
	}
	return sIDs
}

// 实现任务调度interface
func (sensor Sensor) Pk() string {
	return sensor.SensorID
}

func (sensor Sensor) CurrentStatus() int {
	return int(sensor.Status)
}

func (sensor Sensor) Instance() interface{} {
	return sensor
}

func (sensor Sensor) LockerKey() string {
	return sensor.SensorID
}

func (sensor Sensor) LinkNode() string {
	return ""
}

func (sensor Sensor) VideoProgress() int {
	return 0
}

func (sensor Sensor) RenderURI() string {
	return ""
}

func (sensor Sensor) SchedulerType() int {
	return pool.SchedulerTypeSensor
}

func (sensor Sensor) CurrentChannel() int {
	return 0
}
