package dto

import uuid "github.com/satori/go.uuid"

type CategoryRequest struct {
	UserId       string `json:"user_id"`
	ID           string `json:"id"`
	Name         string `json:"name" binding:"required"`
	Image        string `json:"image"`
	TotalProduct int    `json:"total_product"`
}

type CategoryResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	TotalProduct int       `json:"total_product"`
}

type CategoriesResponse struct {
	Data []CategoryResponse     `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
