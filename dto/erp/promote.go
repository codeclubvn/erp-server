package erpdto

import "time"

type CreatePromoteRequest struct {
	PromoteType      PromoteType `json:"promote_type" binding:"required"`                    // amount, percent
	DiscountValue    float64     `json:"discount_value" binding:"required" validate:"min:1"` //
	Quantity         *int        `json:"quantity"`
	QuantityUse      *int        `json:"quantity_use"`
	Code             string      `json:"code" binding:"required"` // code
	Note             string      `json:"note"`
	StartDate        *time.Time  `json:"start_date"`
	EndDate          *time.Time  `json:"end_date"`
	Status           string      `json:"status"` // active, inactive
	MaxPromoteAmount *float64    `json:"max_amount"`
	MaxUse           int         `json:"max_use" gorm:"max_use;type:int;default:1"` // max use per user
	StoreId          string
}

// enum promote method
type PromoteType string

const (
	PromotePercent PromoteType = "percent"
	PromoteAmount  PromoteType = "amount"
)

func (u PromoteType) ErrorMessage() string {
	return "Error: PromoteType expected [percent,amount]"
}

func (d PromoteType) CheckValid() bool {
	tCheck := []PromoteType{PromotePercent, PromoteAmount}
	for _, v := range tCheck {
		if v == d {
			return true
		}
	}
	return false
}

type CreatePromoteUseRequest struct {
	CustomerId  string `json:"customer_id"`
	PromoteCode string `json:"promote_code"`
}
