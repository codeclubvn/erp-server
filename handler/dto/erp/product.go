package erpdto

import (
	"erp/handler/dto/request"
	uuid "github.com/satori/go.uuid"
)

type CreateProductRequest struct {
	UserId      string
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"required,numeric,gte=0"` // gte: greater than or equal
	Status      bool    `json:"status"`                                 // true: active, false: inactive
	Quantity    *int    `json:"quantity"`
	StoreId     string
	Cost        float64 `json:"cost" binding:"numeric,gte=0"` // gte: greater than or equal
}

type UpdateProductRequest struct {
	ID string `json:"id"`
	CreateProductRequest
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Status      bool      `json:"status"`
	Quantity    int       `json:"quantity"`
	Cost        float64   `json:"cost"`
}

type GetListProductRequest struct {
	request.PageOptions
}
