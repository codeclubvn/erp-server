package erpdto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
)

type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image"`
}

type CategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image"`
}

type UpdateCategoryRequest struct {
	ID string `json:"id"`
	CreateCategoryRequest
}

type CategoryResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	TotalProduct int       `json:"total_product"`
}

type GetListCategoryRequest struct {
	request.PageOptions
}
