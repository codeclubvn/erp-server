package erpdto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
)

type CreateTransactionCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"` // expense, income, debt
}

type UpdateTransactionCategoryRequest struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name" binding:"required"`
	Type string    `json:"type" binding:"required"` // expense, income, debt

}

type ListTransactionCategoryRequest struct {
	request.PageOptions
}
