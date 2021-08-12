package models

import (
	"strings"
)

// 2: means default deviceRole
// 3: means non-default deviceRole
const (
	//DeviceRoleType_SuperAdmin = 1	// already waste in sql trigger
	DeviceRoleType_OrgAdmin = 2
	DeviceRoleType_Normal   = 3
)

func IsDefaultDeviceRoleType(deviceRoleType int) (bool, error) {
	switch deviceRoleType {
	case DeviceRoleType_OrgAdmin: //case DeviceRoleType_SuperAdmin, DeviceRoleType_OrgAdmin:
		return true, nil
	case DeviceRoleType_Normal:
		return false, nil
	default:
		return false, NewErrorf(ErrorCodeDeviceRoleModelDeviceRoleTypeInvalid, "DeviceRoleType: %d", deviceRoleType)
	}
}

type DeviceRoleSensor struct {
	DeviceSensorPairId string
	SensorId           string
	SensorName         string
	AuthFlag           string
	Status             int
}

func (deviceRoleSensor *DeviceRoleSensor) ToSensorID() string {
	if deviceRoleSensor == nil {
		return ""
	}
	return deviceRoleSensor.SensorId
}

func newDeviceRoleSensor(sensorID string) *DeviceRoleSensor {
	return &DeviceRoleSensor{
		SensorId: sensorID,
	}
}

func NewDeviceRoleSensors(sensorIDs []string) []*DeviceRoleSensor {
	deviceRoleSensor := make([]*DeviceRoleSensor, 0)
	for _, sID := range sensorIDs {
		deviceRoleSensor = append(deviceRoleSensor, newDeviceRoleSensor(sID))
	}
	return deviceRoleSensor
}

type DeviceRole struct {
	Ts             int64
	DeviceRoleId   string
	DeviceRoleName string
	OrgId          string
	OrgName        string
	DeviceRoleType int
	UserCount      int
	Sensors        []*DeviceRoleSensor
	SensorCount    int
	Comment        string
	Status         int
}

func (deviceRole *DeviceRole) CheckForSave() error {
	if strings.EqualFold(deviceRole.OrgId, "") {
		return NewErrorf(ErrorCodeDeviceRoleModelOrgIDShouldNotEmpty, "OrgId shouldn't be empty")
	}
	if strings.EqualFold(deviceRole.DeviceRoleName, "") {
		return NewErrorf(ErrorCodeDeviceRoleModelDeviceRoleNameShouldNotEmpty, "DeviceRoleName shouldn't be empty")
	}
	if isDefault, err := IsDefaultDeviceRoleType(deviceRole.DeviceRoleType); err != nil || isDefault {
		return NewErrorf(ErrorCodeDeviceRoleModelDeviceRoleTypeNotNormal, "Invalid device role type %d", deviceRole.DeviceRoleType)
	}
	return nil
}

func (deviceRole *DeviceRole) CheckForEdit() error {
	if err := deviceRole.CheckForSave(); err != nil {
		return err
	}
	if strings.EqualFold(deviceRole.DeviceRoleId, "") {
		return NewErrorf(ErrorCodeDeviceRoleModelDeviceRoleIDShouldNotEmpty, "DeviceRoleID shouldn't be empty")
	}
	return nil
}

func (deviceRole *DeviceRole) Normalized() {
	deviceRole.DeviceRoleName = strings.TrimSpace(deviceRole.DeviceRoleName)
	deviceRole.Comment = strings.TrimSpace(deviceRole.Comment)
}

type DeviceRoleRequest struct {
	//DeviceRoleIDs  []string
	DeviceRoleName string
	Pagination
	OrgIDs      []string
	CurrentUser *User
}

func (deviceRoleRequest *DeviceRoleRequest) CheckFuzzyMatching() bool {
	return !strings.EqualFold(deviceRoleRequest.DeviceRoleName, "")
}

func (deviceRoleRequest *DeviceRoleRequest) Normalized() {
	deviceRoleRequest.DeviceRoleName = strings.TrimSpace(deviceRoleRequest.DeviceRoleName)
}

func DeviceRoleSensorModelToSensorIDs(deviceRoleSensors []*DeviceRoleSensor) (sensorIDs []string) {
	sensorIDs = make([]string, 0)
	for _, deviceRoleSensor := range deviceRoleSensors {
		sensorIDs = append(sensorIDs, deviceRoleSensor.SensorId)
	}
	return
}
