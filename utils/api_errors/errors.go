package api_errors

import "net/http"

var (
	ErrInternalServerError         = "10000"
	ErrUnauthorizedAccess          = "10001"
	ErrTokenExpired                = "10003"
	ErrTokenInvalid                = "10004"
	ErrUserNotFound                = "10006"
	ErrRequestTimeout              = "10008"
	ErrTokenMissing                = "10009"
	ErrValidation                  = "10010"
	ErrInvalidUserID               = "10011"
	ErrMissingXStoreID             = "10012"
	ErrPermissionDenied            = "10013"
	ErrInvalidPassword             = "10014"
	ErrStoreNotFound               = "10015"
	ErrOrderItemRequired           = "10016"
	ErrTypeInvalid                 = "10017"
	ErrNotFound                    = "10018"
	ErrDateNotBetween              = "10019"
	ErrTotalInvalid                = "10020"
	ErrPaymentInvalid              = "10021"
	ErrPromoteCodeExist            = "10022"
	ErrDiscountPercentInvalid      = "10023"
	ErrDiscountAmountInvalid       = "10024"
	ErrDeliveryFeeInvalid          = "10025"
	ErrOrderItemInvalid            = "10026"
	ErrQuantityIsNotEnough         = "10029"
	ErrProductInvalid              = "10030"
	ErrPromoteCodeMaxUse           = "10031"
	ErrPromoteCodeRequiredCustomer = "10032"
	ErrOrderStatus                 = "10033"
	ErrWalletNameAlreadyExist      = "10034"
	ErrFileNotFound                = "10035"
)

type MessageAndStatus struct {
	Message string
	Status  int
}

var MapErrorCodeMessage = map[string]MessageAndStatus{
	ErrRequestTimeout:         {"Request Timeout", http.StatusRequestTimeout},
	ErrTokenMissing:           {"Token Missing", http.StatusUnauthorized},
	ErrValidation:             {"Validation Error", http.StatusBadRequest},
	ErrInvalidUserID:          {"Invalid User ID", http.StatusBadRequest},
	ErrMissingXStoreID:        {"Missing x-store-id", http.StatusBadRequest},
	ErrPermissionDenied:       {"Permission Denied", http.StatusForbidden},
	ErrInvalidPassword:        {"Invalid Password", http.StatusBadRequest},
	ErrStoreNotFound:          {"Store Not Found", http.StatusNotFound},
	ErrOrderItemRequired:      {"order_items Required", http.StatusBadRequest},
	ErrTypeInvalid:            {"Only accept type 'percent' or 'amount'", http.StatusBadRequest},
	ErrNotFound:               {"Status Not Found", http.StatusNotFound},
	ErrDateNotBetween:         {"Date Not Between", http.StatusBadRequest},
	ErrWalletNameAlreadyExist: {"Wallet name already exists", http.StatusBadRequest},
	ErrQuantityIsNotEnough:    {"Quantity is not enough", http.StatusBadRequest},
}
