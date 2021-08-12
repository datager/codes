package utils

import (
	"os"
	"strings"
)

const (
	ENV_DEV      = "dev"
	ENV_PROD     = "production"
	ENV_TEST     = "test"
)

func GetRawEnv() string {
	return os.Getenv("ENVIRONMENT")
}

func GetEnv() string {
	if IsDevEnv() {
		return ENV_DEV
	}

	if IsTestEnv() {
		return ENV_TEST
	}

	return ENV_PROD
}

func IsTestEnv() bool {
	env := GetRawEnv()
	return strings.ToLower(env) == ENV_TEST
}

func IsDevEnv() bool {
	env := GetRawEnv()
	return strings.ToLower(env) != ENV_PROD && strings.ToLower(env) != ENV_TEST
}

func GetOSEnv(key string) string {
	return os.Getenv(key)
}