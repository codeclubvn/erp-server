package api_errors

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

var MapErrorCodeMessage = map[string]string{
	ErrInternalServerError:  "Internal Server Error",
	ErrUnauthorizedAccess:   "Unauthorized Access",
	ErrTokenBadSignedMethod: "Token Bad Signed Method",
	ErrTokenExpired:         "Token Expired",
	ErrTokenInvalid:         "Token Invalid",
	ErrTokenMalformed:       "Token Malformed",
	ErrUserNotFound:         "User Not Found",
	ErrProductNotFound:      "Product Not Found",
	ErrRequestTimeout:       "Request Timeout",
	ErrTokenMissing:         "Token Missing",
	ErrValidation:           "Validation Error",
	ErrInvalidUserID:        "Invalid User ID",
	ErrMissingXStoreID:      "Missing x-store-id",
	ErrPermissionDenied:     "Permission Denied",
	ErrInvalidPassword:      "Invalid Password",
}
