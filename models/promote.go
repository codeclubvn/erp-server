package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Promote struct {
	BaseModel
	PromoteType   string     `json:"promote_type"`                   // amount, percent
	DiscountValue float64    `json:"discount_value" gorm:"not null"` //
	Quantity      *int       `json:"quantity,omitempty" gorm:"default:null"`
	QuantityUse   *int       `json:"quantity_use,omitempty" gorm:"default:null"`
	Code          string     `json:"code" gorm:"unique;type:varchar(20);"` // code
	Note          string     `json:"note"`
	StartDate     *time.Time `json:"start_date" gorm:"default:null"`
	EndDate       *time.Time `json:"end_date" gorm:"default:null"`
	Status        string     `json:"status"` // active, inactive
	MaxAmount     *float64   `json:"max_amount" gorm:"default:null"`
	MaxUse        int        `json:"max_use" gorm:"max_use;type:int;default:1"` // max use per user
}

type PromoteUse struct {
	BaseModel
	CustomerId  uuid.UUID `json:"customer_id"`
	PromoteCode string    `json:"promote_code" gorm:"type:varchar(20);"`
	Customer    Customer  `json:"customer" gorm:"foreignKey:CustomerId"`
	Promote     Promote   `json:"promote" gorm:"foreignKey:PromoteCode;references:Code"`
}

func (PromoteUse) TableName() string {
	return "promote_uses"
}
