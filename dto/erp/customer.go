package erpdto

import (
	"erp/api/request"
	"github.com/lib/pq"
)

type CustomerUriRequest struct {
	ID string `uri:"id" json:"id"`
}

type ListCustomerRequest struct {
	request.PageOptions
}

type CreateCustomerRequest struct {
	FirstName      string         `json:"first_name" validate:"required"`
	LastName       string         `json:"last_name" validate:"required"`
	Gender         string         `json:"gender" validate:"required"`
	Age            int            `json:"age" validate:"required"`
	AddressStrings pq.StringArray `json:"address_strings" validate:"required"`
	PhoneNumber    string         `json:"phone_number" validate:"required"`
	Email          string         `json:"email" validate:"required"`
}

type UpdateCustomerRequest struct {
	ID string `uri:"id" json:"id"`
	CreateCustomerRequest
}
