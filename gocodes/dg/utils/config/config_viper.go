package config

import (
	"fmt"
	"time"

	"codes/gocodes/dg/utils/log"

	"github.com/spf13/viper"
)

type viperConfig struct{}

func (config *viperConfig) Get(key string) interface{} {
	if key == "" {
		return viper.AllSettings()
	}
	return viper.Get(key)
}

func (config *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (config *viperConfig) GetStringOrDefault(key string, defaultVal string) string {
	if viper.IsSet(key) {
		return config.GetString(key)
	}
	return defaultVal
}

func (config *viperConfig) MustGetString(key string) string {
	val := viper.GetString(key)
	if val == "" {
		panic(fmt.Errorf("Can not get a string from configuration with key %v", key))
	}
	return val
}

func (config *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (config *viperConfig) GetBoolOrDefault(key string, defaultVal bool) bool {
	if viper.IsSet(key) {
		return config.GetBool(key)
	}
	return defaultVal
}

func (config *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (config *viperConfig) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func (config *viperConfig) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func (config *viperConfig) GetIntOrDefault(key string, defaultVal int) int {
	if viper.IsSet(key) {
		return config.GetInt(key)
	}
	return defaultVal
}

func (config *viperConfig) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func (config *viperConfig) GetDurationOrDefault(key string, defaultVal time.Duration) time.Duration {
	if viper.IsSet(key) {
		return config.GetDuration(key)
	}
	return defaultVal
}

func (config *viperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (config *viperConfig) Unmarshal(key string, val interface{}) error {
	return viper.UnmarshalKey(key, val)
}

func (config *viperConfig) TryUnmarshal(key string, val interface{}) bool {
	if viper.IsSet(key) {
		if err := config.Unmarshal(key, val); err != nil {
			log.Warnf("Failed to unmarshal by key %v, err = %v", key, err)
			return false
		}
		return true
	}
	return false
}

func (config *viperConfig) MustUnmarshal(key string, val interface{}) {
	if err := viper.UnmarshalKey(key, val); err != nil {
		panic(err)
	}
}
func (config *viperConfig) GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
