package svc

import "net/http"

var (
	ErrInvalidPara    = NewError(http.StatusBadRequest, "invalid para")
	ErrNotFound       = NewError(http.StatusNotFound, "not found")
	ErrInternalServer = NewError(http.StatusInternalServerError, "internal main-server error")
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) StatusCode() int {
	return e.Code
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
