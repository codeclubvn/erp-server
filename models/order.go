package models

import uuid "github.com/satori/go.uuid"

type Order struct {
	BaseModel

	Code   string `json:"code" gorm:"unique_index:idx_order_code"`
	Status string `json:"status"`
	Note   string `json:"note"`

	Amount        float64 `json:"amount" gorm:"default:0.0"`
	Total         float64 `json:"total" gorm:"default:0.0"`
	Payment       float64 `json:"payment" gorm:"default:0.0"`
	PaymentMethod string  `json:"payment_method" gorm:"not null"`

	CustomerId  *uuid.UUID `json:"customer_id,omitempty" gorm:"default:null;"`
	Customer    *Customer  `json:"customer,omitempty" gorm:"foreignkey:CustomerId;association_foreignkey:ID"`
	DeliveryFee *float64   `json:"delivery_fee,omitempty"`

	Discount     *float64 `json:"discount,omitempty"`
	DiscountType string   `json:"discount_type" gorm:"default:null;"`

	PromoteFee  *float64 `json:"promote_fee,omitempty"`
	PromoteCode *string  `json:"promote_code,omitempty"`
	Cost        float64  `json:"cost"`
	//StoreId     uuid.UUID `json:"store_id" gorm:"not null"`
	OrderItems []*OrderItem `json:"order_item,omitempty" gorm:"foreignkey:OrderId;association_foreignkey:ID"`
}

type OrderOverview struct {
	Revenue  float64 `json:"revenue"`
	Income   float64 `json:"income"`
	Confirm  int     `json:"confirm"`
	Delivery int     `json:"delivery"`
	Complete int     `json:"complete"`
	Cancel   int     `json:"cancel"`
}
