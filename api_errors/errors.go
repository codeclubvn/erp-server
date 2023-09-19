package api_errors

import "net/http"

var (
	ErrInternalServerError  = "10000"
	ErrUnauthorizedAccess   = "10001"
	ErrTokenBadSignedMethod = "10002"
	ErrTokenExpired         = "10003"
	ErrTokenInvalid         = "10004"
	ErrTokenMalformed       = "10005"
	ErrUserNotFound         = "10006"
	ErrProductNotFound      = "10007"
	ErrRequestTimeout       = "10008"
	ErrTokenMissing         = "10009"
	ErrValidation           = "10010"
	ErrInvalidUserID        = "10011"
	ErrMissingXStoreID      = "10012"
	ErrPermissionDenied     = "10013"
	ErrInvalidPassword      = "10014"
)

type MessageAndStatus struct {
	Message string
	Status  int
}

var MapErrorCodeMessage = map[string]MessageAndStatus{
	ErrInternalServerError:  {"Internal Server Error", http.StatusInternalServerError},
	ErrUnauthorizedAccess:   {"Unauthorized Access", http.StatusUnauthorized},
	ErrTokenBadSignedMethod: {"Token Bad Signed Method", http.StatusUnauthorized},
	ErrTokenExpired:         {"Token Expired", http.StatusUnauthorized},
	ErrTokenInvalid:         {"Token Invalid", http.StatusUnauthorized},
	ErrTokenMalformed:       {"Token Malformed", http.StatusUnauthorized},
	ErrUserNotFound:         {"User Not Found", http.StatusNotFound},
	ErrProductNotFound:      {"Product Not Found", http.StatusNotFound},
	ErrRequestTimeout:       {"Request Timeout", http.StatusRequestTimeout},
	ErrTokenMissing:         {"Token Missing", http.StatusUnauthorized},
	ErrValidation:           {"Validation Error", http.StatusBadRequest},
	ErrInvalidUserID:        {"Invalid User ID", http.StatusBadRequest},
	ErrMissingXStoreID:      {"Missing x-store-id", http.StatusBadRequest},
	ErrPermissionDenied:     {"Permission Denied", http.StatusForbidden},
	ErrInvalidPassword:      {"Invalid Password", http.StatusBadRequest},
}
