package models

type SensorResource struct {
	OrgID     string
	SensorIDs []string
}

type UserResrouce struct {
	UserOrgID       string
	UserOrgComment  string
	SensorResources []*SensorResource
}
