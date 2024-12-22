package code

import (
	"github.com/golangpoke/nux/nux"
)

type errorCode int

const (
	ErrInternalServerError errorCode = nux.Success + iota + 11
	ErrBadRequest
	ErrUnauthorized
	ErrForbidden
	ErrNotFound
	ErrConflict
)

var MapErrorMsg = map[errorCode]string{
	ErrInternalServerError: "internal server error",
	ErrBadRequest:          "bad request",
	ErrUnauthorized:        "unauthorized",
	ErrForbidden:           "forbidden",
	ErrNotFound:            "not found",
	ErrConflict:            "conflict",
}

func (e errorCode) Data() any {
	return nil
}

func (e errorCode) Code() int {
	return int(e)
}

func (e errorCode) Message() string {
	return MapErrorMsg[e]
}

func (e errorCode) Error() error {
	return nil
}

func (e errorCode) With(err error) nux.Response {
	return &response{
		data:    e.Data(),
		code:    e.Code(),
		message: e.Message(),
		err:     err,
	}
}
