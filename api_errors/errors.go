package api_errors

import (
	"errors"
	"net/http"
)

var (
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)

func GetStatusCode(err error) (int, bool) {
	if v, ok := MapErrorStatusCode[err.Error()]; !ok {
		return http.StatusInternalServerError, false
	} else {
		return v, true
	}
}

const (
	InternalServerError = "internal Server Error"
	UserNotFound        = "user not found"
)

var MapErrorStatusCode = map[string]int{
	InternalServerError: http.StatusInternalServerError,
	UserNotFound:        http.StatusNotFound,
}
