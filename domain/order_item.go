package domain

import uuid "github.com/satori/go.uuid"

type OrderItem struct {
	BaseModel
	OrderId   uuid.UUID `json:"order_id" gorm:"type:uuid;not null"`
	ProductId uuid.UUID `json:"product_id" gorm:"type:uuid;not null"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}
