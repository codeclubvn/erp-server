package finance

import (
	"erp/handler/dto/request"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateCashbookRequest struct {
	Amount             float64    `json:"amount" binding:"required"`
	Status             string     `json:"status" binding:"required"`
	Note               string     `json:"note"`
	OrderId            *uuid.UUID `json:"order_id"`
	CashbookCategoryId *uuid.UUID `json:"cashbook_category_id"`
	WalletId           *uuid.UUID `json:"wallet_id" binding:"required"`
	DateTime           *time.Time `json:"date_time" binding:"required"`
	IsSaveInReport     bool       `json:"is_save_in_report"`
	CustomerId         *uuid.UUID `json:"customer_id"`
}

type UpdateCashbookRequest struct {
	Id                 uuid.UUID  `json:"id"`
	OrderId            *uuid.UUID `json:"order_id"`
	Amount             float64    `json:"amount" binding:"required"`
	Status             string     `json:"status" binding:"required"`
	Note               string     `json:"note"`
	CashbookCategoryId *uuid.UUID `json:"cashbook_category_id"`
	WalletId           *uuid.UUID `json:"wallet_id" binding:"required"`
	DateTime           *time.Time `json:"date_time"`
	IsSaveInReport     bool       `json:"is_save_in_report"`
	CustomerId         *uuid.UUID `json:"customer_id"`
}

type ListCashbookRequest struct {
	request.PageOptions
}

type TotalTransactionByCategoryResponse struct {
	CategoryId uuid.UUID `json:"category_id"`
	Total      float64   `json:"total"`
}
