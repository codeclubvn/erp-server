package erpdto

import uuid "github.com/satori/go.uuid"

type CreateTransactionRequest struct {
	OrderId uuid.UUID `json:"order_id"`
	Amount  float64   `json:"amount"`
	Status  string    `json:"status"`
}
