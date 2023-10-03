package models

import uuid "github.com/satori/go.uuid"

type OrderItem struct {
	BaseModel
	OrderId   uuid.UUID `json:"order_id"`
	ProductId uuid.UUID `json:"product_id"`
	Quantity  float64   `json:"quantity"`
}
