package model

import "strings"

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// Được gọi là empty struct
// Mục đích: sử dụng nó như một kiểu dữ liệu không chứa bất kỳ thông tin nào
// Đơn giản là để thực hiện một số hah động hoặc hỗ tợ trong mã m khng cần lưu tữ bất kỳ dữ liệu nào
type EmptyObj struct{}

func BUildResponse(status bool, message string, data interface{}) ApiResponse {
	res := ApiResponse{
		Success: status,
		Message: message,
		Data:    data,
		Error:   nil,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) ApiResponse {
	splittedError := strings.Split(err, "\n")
	res := ApiResponse{
		Success: false,
		Message: message,
		Data:    data,
		Error:   splittedError,
	}
	return res
}
