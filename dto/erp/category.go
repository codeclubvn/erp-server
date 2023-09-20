package dto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
)

type CreateCategoryRequest struct {
	UserId string `json:"user_id"`
	Name   string `json:"name" binding:"required" copier:"Name"`
	Image  string `json:"image"`
}

type CategoryRequest struct {
	UserId string `json:"user_id"`
	Name   string `json:"name" binding:"required" copier:"Name"`
	Image  string `json:"image"`
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

type CategoriesResponse struct {
	Data []*CategoryResponse    `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type GetListCategoryRequest struct {
	request.PageOptions
}
