package erpdto

import uuid "github.com/satori/go.uuid"

type CreateOrderRequest struct {
	OrderId uuid.UUID

	Status StatusCreateOrder `json:"status" binding:"required"` // confirmed, delivered, completed, canceled
	Note   *string           `json:"note"`                      // note for order

	Amount        float64 `json:"amount"` // total amount of order items
	Total         float64 `json:"total"`  // grand total
	Payment       float64 `json:"payment"`
	PaymentMethod string  `json:"payment_method" binding:"required"`

	CustomerId *string `json:"customer_id"`
	Code       string

	DeliveryFee *float64 `json:"delivery_fee"`

	Discount     *float64     `json:"discount"`
	DiscountType DiscountType `json:"discount_type"`

	PromoteCode *string `json:"promote_code"`
	PromoteFee  *float64

	OrderItems []OrderItemRequest `json:"order_items"`

	StoreId string
}

type OrderItemRequest struct {
	ProductId uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
}

// enum promote method
type DiscountType string

const (
	DiscountPercent DiscountType = "percent"
	DiscountAmount  DiscountType = "amount"
)

func (u DiscountType) ErrorMessage() string {
	return "Error: DiscountType expected [percent,amount]"
}

func (d DiscountType) CheckValid() bool {
	tCheck := []DiscountType{DiscountPercent, DiscountAmount}
	for _, v := range tCheck {
		if v == d {
			return true
		}
	}
	return false
}

// enum status order
type StatusCreateOrder string

const (
	Delivery StatusCreateOrder = "delivered"
	Complete StatusCreateOrder = "completed"
)

func (u StatusCreateOrder) ErrorMessage() string {
	return "Error: Status expected [delivered,completed]"
}

func (d StatusCreateOrder) CheckValid() bool {
	tCheck := []StatusCreateOrder{Delivery, Complete}
	for _, v := range tCheck {
		if v == d {
			return true
		}
	}
	return false
}
