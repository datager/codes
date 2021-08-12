package utils

import "time"

const (
	CONFIG_DIR       = "config"
	HTTP_ADDR        = ":3000"
	AUTH_TIMEOUT     = 24 * time.Hour
	AUTH_MAX_REFRESH = 24 * time.Hour
	ORG_ROOT_ID      = "0000"
	UPLOAD_PATH      = "upload"
	TMP_PATH         = "tmp"
	CLIENT_TIMEOUT   = 1 * time.Minute
	TZ_APP_ACCOUNT   = "tzApp"
)
