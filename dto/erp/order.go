package erpdto

import uuid "github.com/satori/go.uuid"

type CreateOrderRequest struct {
	Code   string `json:"code"`
	Status string `json:"status"`
	Note   string `json:"note"`

	Amount        float64 `json:"amount"`
	Total         float64 `json:"total"`
	Payment       float64 `json:"payment"`
	PaymentMethod string  `json:"payment_method"`

	ContactId   uuid.UUID `json:"contact_id"`
	DeliveryFee float64   `json:"delivery_fee"`

	Discount     float64 `json:"discount"`
	DiscountType string  `json:"discount_type"`

	PromoteFee float64   `json:"promote_fee"`
	PromoteId  uuid.UUID `json:"promote_id"`

	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  float64   `json:"quantity"`
	Price     float64   `json:"price"`
	Amount    float64   `json:"amount"`
}
