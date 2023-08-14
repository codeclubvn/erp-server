package model

import (
	"github.com/google/uuid"
)

type Product struct {
	BaseModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
	UserId      uuid.UUID `json:"user_id"`
}

func (Product) TableName() string {
	return "product"
}

type ProductRequest struct {
	ID          *uuid.UUID `json:"id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Price       *int       `json:"price"`
	Quantity    *int       `json:"quantity"`
	Status      *string    `json:"status"`
	UserId      *uuid.UUID
}

type Products []Product
