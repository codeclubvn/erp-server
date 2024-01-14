package erpdto

import (
	"erp/api/request"
	uuid "github.com/satori/go.uuid"
)

type CreateOrderRequest struct {
	OrderId uuid.UUID
	Code    string

	Status StatusOrder `json:"status" binding:"required"` // confirm, delivery, complete, cancel

	Note    *string `json:"note"` // note for order
	Amount  float64 // tổng tiền của toàn bộ sản phẩm
	Total   float64 `json:"total"`   // grand total
	Payment float64 `json:"payment"` // COD | Online

	PaymentMethod string     `json:"payment_method" binding:"required"`
	CustomerId    *uuid.UUID `json:"customer_id"`

	DeliveryFee *float64 `json:"delivery_fee"`

	Discount     *float64      `json:"discount"` // chiết khấu
	DiscountType *DiscountType `json:"discount_type"`

	PromoteCode *string `json:"promote_code"` // mã giảm giá
	PromoteFee  *float64

	OrderItems []OrderItemRequest `json:"order_items"`
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
type StatusOrder string

const (
	OrderConfirm  StatusOrder = "confirm"
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
	tCheck := []StatusOrder{OrderConfirm, OrderDelivery, OrderComplete, OrderCancel}
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

type GetListOrderRequest struct {
	request.PageOptions
}
