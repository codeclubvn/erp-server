package erpdto

import uuid "github.com/satori/go.uuid"

type CreateOrderRequest struct {
	Status string `json:"status"`
	Note   string `json:"note"`

	Amount        float64 `json:"amount" validate:"required"` // total amount of order items
	Total         float64 `json:"total" validate:"required"`  // grand total
	Payment       float64 `json:"payment" validate:"required"`
	PaymentMethod string  `json:"payment_method" validate:"required"`

	CustomerID uuid.UUID `json:"customer_id"`
	Code       string    `json:"code"`

	ContactId   uuid.UUID `json:"contact_id"`
	DeliveryFee float64   `json:"delivery_fee"`

	Discount     float64 `json:"discount"`
	DiscountType string  `json:"discount_type"`

	PromoteFee  float64   `json:"promote_fee"`
	PromoteType string    `json:"promote_type"`
	PromoteId   uuid.UUID `json:"promote_id"`

	OrderItems []OrderItemRequest `json:"order_items"`

	StoreId string
}

type OrderItemRequest struct {
	ProductId uuid.UUID `json:"product_id" validate:"required"`
	Quantity  float64   `json:"quantity" validate:"required"`
	Price     float64   `json:"price" validate:"required"`

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
