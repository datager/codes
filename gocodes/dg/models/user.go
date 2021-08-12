package models

import (
	"errors"
	"strings"
)

type User struct {
	Ts             int64
	UserId         string
	UserName       string
	UserPassword   string
	OrgId          string
	OrgName        string
	UserFuncRole   *FuncRole
	UserDeviceRole *DeviceRole
	IsValid        bool
	RealName       string
	Comment        string
	Status         int
	SecurityToken  string `json:"-"`
}

func (user *User) CheckForSave() error {
	if strings.EqualFold(user.UserName, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserNameShouldNotEmpty)
	}
	if strings.EqualFold(user.OrgId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelOrgIDShouldNotEmpty)
	}
	if user.UserFuncRole == nil {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserFuncRoleShouldNotEmpty)
	}
	if strings.EqualFold(user.UserFuncRole.FuncRoleId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelFuncRoleIDShouldNotEmpty)
	}
	if user.UserDeviceRole == nil {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserDeviceRoleShouldNotEmpty)
	}
	if strings.EqualFold(user.UserDeviceRole.DeviceRoleId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelDeviceRoleIDShouldNotEmpty)
	}
	if strings.EqualFold(user.UserPassword, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserPasswordShouldNotEmpty)
	}
	return nil
}

func (user *User) CheckForUpdate() error {
	if strings.EqualFold(user.UserName, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserNameShouldNotEmpty)
	}
	if strings.EqualFold(user.OrgId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelOrgIDShouldNotEmpty)
	}
	if user.UserFuncRole == nil {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserFuncRoleShouldNotEmpty)
	}
	if strings.EqualFold(user.UserFuncRole.FuncRoleId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelFuncRoleIDShouldNotEmpty)
	}
	if user.UserDeviceRole == nil {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserDeviceRoleShouldNotEmpty)
	}
	if strings.EqualFold(user.UserDeviceRole.DeviceRoleId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelDeviceRoleIDShouldNotEmpty)
	}
	if strings.EqualFold(user.UserId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelUserIDShouldNotEmpty)
	}
	return nil
}

func (user *User) Normalized() {
	user.UserName = strings.TrimSpace(user.UserName)
}

type AppUser struct {
	UserName string `json:"username"`
}

type UserRequest struct {
	UserName string
	OrgId    string
	Pagination
}

func (userRequest *UserRequest) Check() error {
	if strings.EqualFold(userRequest.OrgId, "") {
		return NewErrorVerificationFailed(ErrorCodeUserModelOrgIDShouldNotEmpty)
	}
	return nil
}

func (userRequest *UserRequest) CheckFuzzyMatching() bool {
	return !strings.EqualFold(userRequest.UserName, "")
}

func (userRequest *UserRequest) Normalized() {
	userRequest.UserName = strings.TrimSpace(userRequest.UserName)
}

type UserPatch struct {
	IsValid bool
	IDs     []string
}

func (userPath *UserPatch) Check() error {
	if len(userPath.IDs) == 0 {
		return errors.New("IDs shouldn't be empty")
	}
	return nil
}
