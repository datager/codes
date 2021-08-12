package models

import (
	"fmt"
)

var (
	// base
	LokiPrefix = "123test" // ip:3000
	//LokiPrefix = configs.GetInstance().LokiConfig.Addr // ip:3000

	// auth
	LoginURL = fmt.Sprintf("%s%s", LokiPrefix, "/login")

	// api
	OrgCreateURL      = fmt.Sprintf("%s%s", LokiPrefix, "/api/org")
	FaceRepoCreateURL = fmt.Sprintf("%s%s", LokiPrefix, "/api/face/repo")
	FaceRepoDeleteURL = fmt.Sprintf("%s%s", LokiPrefix, "/api/face/repo")
	FaceRepoGetURL    = fmt.Sprintf("%s%s", LokiPrefix, "/api/face/repos")

	SensorCreateURL = fmt.Sprintf("%s%s", LokiPrefix, "/api/sensor")
	SensorDelURL    = fmt.Sprintf("%s%s", LokiPrefix, "/api/sensors")
)
