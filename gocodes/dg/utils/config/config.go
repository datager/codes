package config

import (
	"time"

	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config interface {
	Get(key string) interface{}
	GetString(key string) string
	GetStringOrDefault(key string, defaultVal string) string
	MustGetString(key string) string
	GetBool(key string) bool
	GetBoolOrDefault(key string, defaultVal bool) bool
	GetInt(key string) int
	GetInt64(key string) int64
	GetIntOrDefault(key string, defaultVal int) int
	GetDuration(key string) time.Duration
	GetDurationOrDefault(key string, defaultVal time.Duration) time.Duration
	GetStringSlice(key string) []string
	Unmarshal(key string, val interface{}) error
	TryUnmarshal(key string, val interface{}) bool
	MustUnmarshal(key string, val interface{})
	GetStringMapString(key string) map[string]string
	GetFloat64(key string) float64
}

var (
	config Config
)

func InitConfig(configDir string) (Config, error) {
	if config != nil {
		return config, nil
	}
	viper.AddConfigPath(configDir)
	viper.SetConfigName("default")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.WithStack(err)
	}
	env := utils.GetEnv()
	if env != "" {
		viper.AddConfigPath("config")
		viper.SetConfigName(env)
		viper.SetConfigType("json")
		if err := viper.MergeInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Warnf("Failed to load env config file, ignore\n%+v", err)
			} else {
				return nil, errors.WithStack(err)
			}
		}
	}
	config = &viperConfig{}
	return config, nil
}

func GetConfig() Config {
	return config
}
