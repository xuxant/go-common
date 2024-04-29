package errors

import (
	"encoding/json"
	"net/http"
)

const (
	CodeInternal = iota + 100
	CodeNotFound
	CodeBadRequest
	CodeUnauthorized
	CodeConflict
)

var (
	Internal = Error{
		Message:    "Internal Server Error",
		Code:       CodeInternal,
		HttpStatus: http.StatusInternalServerError,
	}
	NotFound = Error{
		Message:    "Not found",
		HttpStatus: CodeNotFound,
		Code:       http.StatusNotFound,
	}
	BadRequest = Error{
		Message:    "Bad request",
		HttpStatus: CodeBadRequest,
		Code:       http.StatusBadRequest,
	}
	Unauthorized = Error{
		Message:    "Unauthorized",
		HttpStatus: http.StatusUnauthorized,
		Code:       CodeUnauthorized,
	}
	Conflict = Error{
		Message:    "Conflict",
		HttpStatus: http.StatusConflict,
		Code:       CodeConflict,
	}
)

type Error struct {
	Code       int    `json:"code,omitempty"`
	Message    string `json:"message"`
	Detail     error  `json:"-"`
	HttpStatus int    `json:"-"`
}

func (e Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e Error) WithCode(code int) *Error {
	e.Code = code
	return &e
}

func (e Error) WithDetail(err error) *Error {
	e.Detail = err
	return &e
}

func (e Error) WithMessage(msg string) *Error {
	e.Message = msg
	return &e
}

func (e Error) WithStatus(status int) *Error {
	e.HttpStatus = status
	return &e
}
