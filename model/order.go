package model

import (
	"github.com/google/uuid"
)

type Order struct {
	BaseModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
	UserId      uuid.UUID `json:"user_id"`
}

func (Order) TableName() string {
	return "order"
}

type OrderRequest struct {
	ID          *uuid.UUID `json:"id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Price       *int       `json:"price"`
	Quantity    *int       `json:"quantity"`
	Status      *string    `json:"status"`
	UserId      *uuid.UUID
}

type Orders []Order
