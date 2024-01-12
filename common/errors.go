package common

import (
	"fmt"
	"net/http"
)

var (
	ErrNotFound            = NewError(http.StatusNotFound, "Not Found")
	ErrInternalServerError = NewError(http.StatusInternalServerError, "Internal Server Error")
	ErrBadRequest          = NewError(http.StatusBadRequest, "Bad Request")
	ErrUnauthorized        = NewError(http.StatusUnauthorized, "Unauthorized")
	ErrForbidden           = NewError(http.StatusForbidden, "Forbidden")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) *Error {
	return &Error{code, message}
}

func (e *Error) Error() error {
	return fmt.Errorf("code: %d, message: %s", e.Code, e.Message)
}
