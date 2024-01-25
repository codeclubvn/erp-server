package erpdto

import (
	"erp/handler/dto/request"
)

type CustomerUriRequest struct {
	ID string `uri:"id" json:"id"`
}

type ListCustomerRequest struct {
	request.PageOptions
}

type CreateCustomerRequest struct {
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
}

type UpdateCustomerRequest struct {
	ID string `uri:"id" json:"id"`
	CreateCustomerRequest
}
