package api_errors

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError  = errors.New("internal server error")
	ErrUnauthorizedAccess   = errors.New("unauthorized access")
	ErrTokenBadSignedMethod = errors.New("bad signed method received")
	ErrTokenExpired         = errors.New("token expired")
	ErrTokenInvalid         = errors.New("token invalid")
	ErrTokenMalformed       = errors.New("token malformed")
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
