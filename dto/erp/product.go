package erpdto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
)

type CreateProductRequest struct {
	UserId        string  `json:"user_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Image         string  `json:"image"`
	Price         float64 `json:"price"`
	Status        bool    `json:"status"`
	NumberProduct int     `json:"number_product"`
}

type UpdateProductRequest struct {
	ID string `json:"id"`
	CreateProductRequest
}

type ProductResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	Price         float64   `json:"price"`
	Status        bool      `json:"status"`
	NumberProduct int       `json:"number_product"`
}

type ProductsResponse struct {
	Data []*ProductResponse     `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type GetListProductRequest struct {
	request.PageOptions
}
