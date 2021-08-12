package models

import (
	"strings"
)

const (
	FuncRoleType_SystemPreserved      = 1
	FuncRoleType_Normal               = 2
	FuncRoleType_SystemPreservedAdmin = 3
)

type FuncRole struct {
	Ts           int64
	OrgId        string
	OrgName      string
	FuncRoleId   string
	FuncRoleName string
	FuncAuthFlag string
	FuncRoleType int
	UserCount    int
	Comment      string
}

func (funcRole *FuncRole) CheckForSave() error {
	if strings.EqualFold(funcRole.OrgId, "") {
		return NewErrorf(ErrorCodeFuncRoleModelOrgIDShouldNotEmpty, "OrgId shouldn't be empty")
	}
	if strings.EqualFold(funcRole.FuncRoleName, "") {
		return NewErrorf(ErrorCodeFuncRoleModelFuncRoleNameShouldNotEmpty, "FuncRoleName shouldn't be empty")
	}
	if funcRole.FuncRoleType != FuncRoleType_Normal {
		return NewErrorf(ErrorCodeFuncRoleModelFuncRoleTypeNotNormal, "Invalid func role type %d", funcRole.FuncRoleType)
	}
	return nil
}

func (funcRole *FuncRole) CheckForUpdate() error {
	if err := funcRole.CheckForSave(); err != nil {
		return err
	}
	if strings.EqualFold(funcRole.FuncRoleId, "") {
		return NewErrorf(ErrorCodeFuncRoleModelFuncRoleIDShouldNotEmpty, "FuncRoleID shouldn't be empty")
	}
	return nil
}

func (funcRole *FuncRole) Normalized() {
	funcRole.FuncRoleName = strings.TrimSpace(funcRole.FuncRoleName)
	funcRole.Comment = strings.TrimSpace(funcRole.Comment)
}

type FuncRoleRequest struct {
	FuncRoleName string
	Pagination
}

func (funcRoleRequest *FuncRoleRequest) CheckFuzzyMatching() bool {
	return !strings.EqualFold(funcRoleRequest.FuncRoleName, "")
}

func (funcRoleRequest *FuncRoleRequest) Normalized() {
	funcRoleRequest.FuncRoleName = strings.TrimSpace(funcRoleRequest.FuncRoleName)
}
