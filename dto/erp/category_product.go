package erpdto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
)

type CategoryProductRequest struct {
	UserId     string `json:"user_id" binding:"required"`
	CategoryId string `json:"category_id" binding:"required"`
	ProductId  string `json:"product_id" binding:"required"`
}

type GetListCatProRequest struct {
	request.PageOptions
}

type CatProductResponse struct {
	CategoryId uuid.UUID `json:"category_id"`
	ProductId  uuid.UUID `json:"product_id"`
}

type DeleteCatagoryProductRequest struct {
	UserId string `json:"user_id" binding:"required"`
	ID     string `json:"id" binding:"required"`
}
