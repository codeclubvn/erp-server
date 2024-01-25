package finance

import (
	"erp/handler/dto/request"
	uuid "github.com/satori/go.uuid"
)

type CreateWalletRequest struct {
	Name      string  `json:"name" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	IsDefault bool    `json:"is_default"`
}

type UpdateWalletRequest struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	IsDefault bool      `json:"is_default"`
}

type ListWalletRequest struct {
	request.PageOptions
}
