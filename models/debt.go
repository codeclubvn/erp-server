package models

import uuid "github.com/satori/go.uuid"

type Debt struct {
	BaseModel
	OrderId    uuid.UUID `gorm:"column:order_id;type:varchar(255);not null"`
	Amount     float64   `gorm:"column:amount;type:float;not null"`
	Status     string    `gorm:"column:status;type:varchar(255);not null"` // in, out
	CustomerId uuid.UUID `gorm:"column:customer_id;type:varchar(255);not null"`
	IsPay      bool      `gorm:"column:is_pay;type:boolean;"`
}
