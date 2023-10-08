package erpdto

import uuid "github.com/satori/go.uuid"

type CreateDebtRequest struct {
	OrderId    uuid.UUID `json:"order_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	CustomerId uuid.UUID `json:"customer_id"`
}

type UpdateDebtRequest struct {
	ID         uuid.UUID `json:"id"`
	OrderId    uuid.UUID `json:"order_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	CustomerId uuid.UUID `json:"customer_id"`
}
