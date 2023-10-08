package api_errors

import "net/http"

var (
	ErrInternalServerError         = "10000"
	ErrUnauthorizedAccess          = "10001"
	ErrTokenBadSignedMethod        = "10002"
	ErrTokenExpired                = "10003"
	ErrTokenInvalid                = "10004"
	ErrTokenMalformed              = "10005"
	ErrUserNotFound                = "10006"
	ErrProductNotFound             = "10007"
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
	ErrPriceOfProductInvalid       = "10027"
	ErrAmountIsNotMatched          = "10028"
	ErrQuantityIsNotEnough         = "10029"
	ErrProductInvalid              = "10030"
	ErrPromoteCodeMaxUse           = "10031"
	ErrPromoteCodeRequiredCustomer = "10032"
	ErrOrderStatus                 = "10033"
)

type MessageAndStatus struct {
	Message string
	Status  int
}

var MapErrorCodeMessage = map[string]MessageAndStatus{
	ErrInternalServerError:         {"Internal Server Error", http.StatusInternalServerError},
	ErrUnauthorizedAccess:          {"Unauthorized Access", http.StatusUnauthorized},
	ErrTokenBadSignedMethod:        {"Token Bad Signed Method", http.StatusUnauthorized},
	ErrTokenExpired:                {"Token Expired", http.StatusUnauthorized},
	ErrTokenInvalid:                {"Token Invalid", http.StatusUnauthorized},
	ErrTokenMalformed:              {"Token Malformed", http.StatusUnauthorized},
	ErrUserNotFound:                {"User Not Found", http.StatusNotFound},
	ErrProductNotFound:             {"Product Not Found", http.StatusNotFound},
	ErrRequestTimeout:              {"Request Timeout", http.StatusRequestTimeout},
	ErrTokenMissing:                {"Token Missing", http.StatusUnauthorized},
	ErrValidation:                  {"Validation Error", http.StatusBadRequest},
	ErrInvalidUserID:               {"Invalid User ID", http.StatusBadRequest},
	ErrMissingXStoreID:             {"Missing x-store-id", http.StatusBadRequest},
	ErrPermissionDenied:            {"Permission Denied", http.StatusForbidden},
	ErrInvalidPassword:             {"Invalid Password", http.StatusBadRequest},
	ErrStoreNotFound:               {"Store Not Found", http.StatusNotFound},
	ErrOrderItemRequired:           {"Order Item Required", http.StatusBadRequest},
	ErrTypeInvalid:                 {"Only accept type 'percent' or 'amount'", http.StatusBadRequest},
	ErrNotFound:                    {"Status Not Found", http.StatusNotFound},
	ErrDateNotBetween:              {"Date Not Between", http.StatusBadRequest},
	ErrTotalInvalid:                {"Total request and calculated total are not matched", http.StatusBadRequest},
	ErrPaymentInvalid:              {"Payment invalid", http.StatusBadRequest},
	ErrPromoteCodeExist:            {"Promote code exist", http.StatusBadRequest},
	ErrDiscountPercentInvalid:      {"Discount only accept 1 - 100 percent", http.StatusBadRequest},
	ErrDiscountAmountInvalid:       {"Discount only accept > 0 and <= 100% value of order", http.StatusBadRequest},
	ErrDeliveryFeeInvalid:          {"DeliveryFee is Invalid", http.StatusBadRequest},
	ErrOrderItemInvalid:            {"OrderItem is invalid", http.StatusBadRequest},
	ErrPriceOfProductInvalid:       {"Price of product is mismatched, please update price again", http.StatusBadRequest},
	ErrAmountIsNotMatched:          {"Amount request and calculated amount are not matched", http.StatusBadRequest},
	ErrQuantityIsNotEnough:         {"Quantity is not enough", http.StatusBadRequest},
	ErrProductInvalid:              {"Product is invalid", http.StatusBadRequest},
	ErrPromoteCodeMaxUse:           {"Promote code max used", http.StatusBadRequest},
	ErrPromoteCodeRequiredCustomer: {"Promote code is required customer", http.StatusBadRequest},
	ErrOrderStatus:                 {"Order status invalid", http.StatusBadRequest},
}
