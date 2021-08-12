package models

type UnionApiErrorType int

const (
	UnionApi_Error_Type_InputInvalid     UnionApiErrorType = 400
	UnionApi_Error_Type_InnerError       UnionApiErrorType = 500
	UnionApi_Error_Type_BackServiceError UnionApiErrorType = 503
)

type UnionApiError struct {
	ErrorCode string `json:"error_code"`
	Error     string `json:"error"`
}
