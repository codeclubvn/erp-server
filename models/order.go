package models

import uuid "github.com/satori/go.uuid"

type Order struct {
	BaseModel

	Code   string `json:"code" gorm:"unique_index:idx_order_code"`
	Status string `json:"status"`
	Note   string `json:"note"`

	Amount        float64 `json:"amount" gorm:"default:0.0"`
	Total         float64 `json:"total" gorm:"default:0.0"`
	Payment       float64 `json:"payment"`
	PaymentMethod string  `json:"payment_method"`

	ContactId   uuid.UUID `json:"contact_id"`
	DeliveryFee float64   `json:"delivery_fee"`

	Discount     float64 `json:"discount"`
	DiscountType string  `json:"discount_type"`

	PromoteFee float64   `json:"promote_fee"`
	PromoteId  uuid.UUID `json:"promote_id"`
}
