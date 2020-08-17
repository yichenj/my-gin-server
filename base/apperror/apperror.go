package apperror

import (
	"fmt"
)

var (
	ErrInvalidResource = New(ERR_INVALID_RESOURCE, "invalid operating resource")
	ErrNotFound        = New(ERR_NOT_FOUND, "resource not found")
	ErrDatabase        = New(ERR_DATABASE, "database error")
	ErrUnknown         = New(ERR_UNKNOWN, "unknown error")
)

type AppError struct {
	Code ErrCode
	Msg  string
}

func (appError *AppError) Error() string {
	return fmt.Sprintf("App Error %d: %s", appError.Code, appError.Msg)
}

func New(code ErrCode, msg string) error {
	return &AppError{Code: code, Msg: msg}
}
