package erpdto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateTransactionRequest struct {
	Amount                float64    `json:"amount" binding:"required"`
	Status                string     `json:"status" binding:"required"`
	Note                  string     `json:"note"`
	OrderId               *uuid.UUID `json:"order_id"`
	TransactionCategoryId *uuid.UUID `json:"transaction_category_id"`
	WalletId              *uuid.UUID `json:"wallet_id" binding:"required"`
	DateTime              *time.Time `json:"date_time" binding:"required"`
}

type UpdateTransactionRequest struct {
	Id                    uuid.UUID  `json:"id"`
	OrderId               *uuid.UUID `json:"order_id"`
	Amount                float64    `json:"amount" binding:"required"`
	Status                string     `json:"status" binding:"required"`
	Note                  string     `json:"note"`
	TransactionCategoryId *uuid.UUID `json:"transaction_category_id"`
	WalletId              *uuid.UUID `json:"wallet_id" binding:"required"`
	DateTime              *time.Time `json:"date_time" binding:"required"`
}

type ListTransactionRequest struct {
	request.PageOptions
}

type TotalTransactionByCategoryResponse struct {
	CategoryId uuid.UUID `json:"category_id"`
	Total      float64   `json:"total"`
}
