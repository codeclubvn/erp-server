package erpdto

import uuid "github.com/satori/go.uuid"

type CreateOrderRequest struct {
	OrderId uuid.UUID

	Status StatusOrder `json:"status" binding:"required"` // confirm, delivery, complete, cancel
	Note   *string     `json:"note"`                      // note for order

	Amount        float64 // total amount of order items
	Total         float64 `json:"total"` // grand total
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
type StatusOrder string

const (
	OrderDelivery StatusOrder = "delivery"
	OrderComplete StatusOrder = "complete"
	OrderCancel   StatusOrder = "cancel"
)

func (u StatusOrder) ErrorCreateMessage() string {
	return "Error: Status expected [delivery,complete]"
}

func (u StatusOrder) ErrorUpdateMessage() string {
	return "Error: Status expected [delivery,complete, cancel]"
}

func (d StatusOrder) CheckValid() bool {
	tCheck := []StatusOrder{OrderDelivery, OrderComplete, OrderCancel}
	for _, v := range tCheck {
		if v == d {
			return true
		}
	}
	return false
}

type UpdateOrderRequest struct {
	OrderId uuid.UUID `json:"order_id" binding:"required"`

	Status  StatusOrder `json:"status" binding:"required"` // confirm, delivery, complete, cancel
	Payment float64     `json:"payment"`
	StoreId string
}
