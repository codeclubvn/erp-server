package model

import (
	"github.com/google/uuid"
)

type Order struct {
	BaseModel
	Price       float64     `json:"price"` // total price of all items
	DeliveryFee float64     `json:"delivery_fee"`
	State       string      `json:"state"` // pending, processing, completed, cancelled
	Total       float64     `json:"total"` // total price + delivery fee
	UserId      uuid.UUID   `json:"user_id"`
	OrderItems  []OrderItem `json:"order_items"`
}

func (Order) TableName() string {
	return "order"
}

type OrderRequest struct {
	ID          *uuid.UUID `json:"id"`
	Price       *float64   `json:"price"` // total price of all items
	DeliveryFee *float64   `json:"delivery_fee"`
	State       *string    `json:"state"` // pending, processing, completed, cancelled
	Total       *float64   `json:"total"` // total price + delivery fee
	UserId      *uuid.UUID
	OrderItems  []OrderItemRequest `json:"order_items"`
}

type Orders []Order

type OneOrderRequest struct {
	Id     string
	UserId string
}
