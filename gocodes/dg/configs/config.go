package configs

import (
	"io/ioutil"
	"sync"

	"codes/gocodes/dg/utils/json"

	"github.com/golang/glog"
)

const (
	ConfigPath = "./config.json" // "/Users/sunyuchuan/language/src/codes/gocodes/dg/1.json"
)

func init() {
	GetInstance()
}

func GetInstance() *Config {
	configOnce.Do(func() {
		data, err := ioutil.ReadFile(ConfigPath)
		if err != nil {
			glog.Fatal(err)
		}

		config = &Config{}
		err = json.Unmarshal(data, config)
		if err != nil {
			glog.Fatal(err)
		}
	})
	return config
}

var (
	config     *Config
	configOnce sync.Once
)

type Config struct {
	Debug                bool
	LokiConfig           *LokiConfig
	FaceRepoInsertConfig *FaceRepoInsertConfig
	DBServiceConfig      *DBServiceConfig
	RedisServiceConfig   *RedisServiceConfig
	CaptureConfig        *CaptureConfig
}

type RedisServiceConfig struct {
	Addr     string
	Password string
	DB       int
}

type LokiConfig struct {
	Addr string
}

type FaceRepoInsertConfig struct {
	ThreadNum int64
}

type DBServiceConfig struct {
	V5ConnStr             string
	StatisticsConnStr     string
	SpecialConnStr        string
	GaussTestTableName    string
	GaussTestTs           int64
	IsShowSQL             bool
	MaxOpenConnsPerClient int
	MaxIdleConnsPerClient int
	MaxLifeTimeSeconds    int
}

type CaptureConfig struct {
	NumInsideOneInsertBatch int64
	Concurrency             int64
	Ts                      struct {
		TsStart int64 // 1580659200000,
		TsEnd   int64 // 1580659200000,
		//TsDuration int64 // 15724700000 4周=28天的ms值
	}
	Sensor struct {
		MinSensorIntID int64 // 1
		MaxSensorIntID int64 // 11000 		// 需要保证灌sensor的时候不要操作sensor, 这样sensor_int_id就是连续的
	}
	CaptureNumToInsertOfOneSensor int64 // 9091 需要灌1亿抓拍, 则每个sensor=1亿/11000个sensor=9091
	URL                           struct {
		VehicleURL    string // "http://192.168.2.222:8501/api/file/165,02a69376f875afd2" 方便前端能看个图
		PedestrianURL string
		NonmotorURL   string
		FaceURL       string
	}
}
