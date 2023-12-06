package erpdto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
)

type CreateRevenueRequest struct {
	OrderId uuid.UUID `json:"order_id"`
	Amount  float64   `json:"amount" binding:"required"`
	Status  string    `json:"status" binding:"required"`
	Note    string    `json:"note" binding:"required"`
}

type UpdateRevenueRequest struct {
	Id      uuid.UUID `json:"id"`
	OrderId uuid.UUID `json:"order_id"`
	Amount  float64   `json:"amount" binding:"required"`
	Status  string    `json:"status" binding:"required"`
	Note    string    `json:"note" binding:"required"`
}

type ListRevenueRequest struct {
	request.PageOptions
}
