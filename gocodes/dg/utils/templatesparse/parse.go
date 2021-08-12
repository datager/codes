package templatesparse

import (
	"fmt"
	"io/ioutil"
	"strings"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/config"

	"github.com/pkg/errors"
)

// TODO: 文件路径问题
const (
	ISDCONFIGPATH            = "isd.json"
	VSDCONFIGPATH            = "link_video.json"
	IMPORTERFTPCONFIGPATH    = "importer.json"
	IMPORTERLIBRAFCONFIGPATH = "importer_libraF.json"
	SDKHUIMUCONFIGPATH       = "huimu.json"
	SDKHAIKANGCONFIGPATH     = "haikangsdk.json"
	SDKHUAWEICONFIGPATH      = "huawei.json"
	SDKDAHUACONFIGPATH       = "dahua.json"
	SDKDC84PATH              = "dc84.json"
)

type ConfigType int

const (
	ConfigTypeUnkown         ConfigType = 0
	ConfigTypeVsd            ConfigType = 1
	ConfigTypeImportorLibraF ConfigType = 2
	ConfigTypeImportorFTP    ConfigType = 3
	ConfigTypeNeedNotOlympus ConfigType = 4
	ConfigTypeSDKHuiMu       ConfigType = 5
	ConfigTypeSDKHaiKang     ConfigType = 6
	ConfigTypeSDKDaHua       ConfigType = 7
	ConfigTypeSDKHuaWei      ConfigType = 8
	ConfigTypeHistoryVideo   ConfigType = 9
	ConfigTypeDC84           ConfigType = 10
	ConfigTypeDC83           ConfigType = 11
)

var sensorTypeToConfigTypeMap = map[int]ConfigType{
	models.SensorType_Sensor_Type_Face:             ConfigTypeImportorLibraF,
	models.SensorType_Sensor_Type_Capture:          ConfigTypeImportorFTP,
	models.SensorType_Sensor_Type_Ipc:              ConfigTypeVsd,
	models.SensorType_Sensor_Type_Video:            ConfigTypeVsd,
	models.SensorType_Sensor_Type_Picture:          ConfigTypeImportorFTP,
	models.SensorType_Sensor_Type_WithID_Device:    ConfigTypeImportorFTP,
	models.SensorType_Sensor_Type_WithoutID_Device: ConfigTypeNeedNotOlympus,
	models.SensorType_Sensor_Type_Car:              ConfigTypeNeedNotOlympus,
	models.SensorType_Sensor_Type_Netposa_PVG:      ConfigTypeVsd,
	models.SensorType_Sensor_Type_GB28181:          ConfigTypeVsd,
	models.SensorType_Sensor_Type_HuiMu:            ConfigTypeSDKHuiMu,
	models.SensorType_Sensor_Type_HaiKang:          ConfigTypeSDKHaiKang,
	models.SensorType_Sensor_Type_DaHua:            ConfigTypeSDKDaHua,
	models.SensorType_Sensor_Type_HuaWei:           ConfigTypeSDKHuaWei,
	models.SensorTypeDC84:                          ConfigTypeDC84,
	models.SensorTypeDC83:                          ConfigTypeDC83,
}

var configTypeToPathMap = map[ConfigType]string{
	ConfigTypeVsd:            VSDCONFIGPATH,
	ConfigTypeImportorLibraF: IMPORTERLIBRAFCONFIGPATH,
	ConfigTypeImportorFTP:    IMPORTERFTPCONFIGPATH,
	ConfigTypeSDKHuiMu:       SDKHUIMUCONFIGPATH,
	ConfigTypeSDKHaiKang:     SDKHAIKANGCONFIGPATH,
	ConfigTypeSDKDaHua:       SDKDAHUACONFIGPATH,
	ConfigTypeSDKHuaWei:      SDKHUAWEICONFIGPATH,
	ConfigTypeHistoryVideo:   VSDCONFIGPATH,
	ConfigTypeDC84:           SDKDC84PATH,
	ConfigTypeDC83:           SDKDC84PATH,
}

type TemplatesParse struct {
	basePath string
}

func NewTemplatesParser() (*TemplatesParse, error) {
	conf := config.GetConfig()
	basePath := conf.GetString("api.templatesBasePath")
	if basePath == "" {
		return nil, errors.New("template base path not set")
	}

	return &TemplatesParse{
		basePath: basePath,
	}, nil
}

func (tc *TemplatesParse) GetConfigJSON(sensorType int) (string, error) {
	configType, ok := sensorTypeToConfigTypeMap[sensorType]
	if !ok {
		err := fmt.Errorf("get config json,but sensor type:%v not found", sensorType)
		return "", err
	}

	templatePath, ok := configTypeToPathMap[configType]
	if !ok {
		// err := fmt.Errorf("get config json, but template path for config type:%v not found", configType)
		return "", nil
	}

	path := tc.basePath + "/" + templatePath
	content, err := tc.loadConfig(path)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return content, nil
}

func (tc *TemplatesParse) loadConfig(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.WithStack(err)
	}

	// delete space
	str := strings.Replace(string(content), " ", "", -1)

	// delete \n \r
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)

	return str, nil
}
