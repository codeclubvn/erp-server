package erpdto

import (
	"erp/api/request"
	"erp/models"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateBudgetRequest struct {
	Amount                float64    `json:"amount" binding:"required"`
	Note                  string     `json:"note"`
	TransactionCategoryId uuid.UUID  `json:"transaction_category_id" binding:"required"`
	StartTime             *time.Time `json:"start_time"`
	EndTime               *time.Time `json:"end_time"`
	Repeat                bool       `json:"repeat"`
}

type UpdateBudgetRequest struct {
	Id                    uuid.UUID  `json:"id" binding:"required"`
	Amount                float64    `json:"amount" binding:"required"`
	Note                  string     `json:"note"`
	TransactionCategoryId uuid.UUID  `json:"transaction_category_id" binding:"required"`
	StartTime             *time.Time `json:"start_time"`
	EndTime               *time.Time `json:"end_time"`
	Repeat                bool       `json:"repeat"`
}

type ListBudgetRequest struct {
	request.PageOptions
}

type BudgetResponse struct {
	models.Budget
	Spent float64 `json:"spent"`
}
