package models

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cashbook struct {
	BaseModel
	Amount   float64        `json:"amount" gorm:"column:amount;type:float;not null"`
	Status   string         `json:"status" gorm:"column:status;type:varchar(255);not null"` // in, out
	Note     string         `json:"note" gorm:"column:note;type:varchar(255);"`
	Images   pq.StringArray `json:"images" gorm:"column:images;type:varchar(500)[];"`
	DateTime *time.Time     `json:"date_time" gorm:"column:date_time;not null"`

	OrderId *uuid.UUID `json:"order_id"  gorm:"column:order_id;type:uuid;"`
	Order   *Order     `json:"order" gorm:"foreignKey:OrderId"`

	WalletId *uuid.UUID `json:"wallet_id" gorm:"column:wallet_id;type:uuid;"`
	Wallet   *Wallet    `json:"wallet" gorm:"foreignKey:WalletId"`

	CashbookCategoryId *uuid.UUID        `json:"cashbook_category_id" gorm:"column:cashbook_category_id;type:uuid;"`
	CashbookCategory   *CashbookCategory `json:"cashbook_category" gorm:"foreignKey:CashbookCategoryId"`

	CustomerId *uuid.UUID `json:"customer_id" gorm:"column:customer_id;type:uuid;"`
	Customer   *Customer  `json:"customer" gorm:"foreignKey:CustomerId"`

	IsPay bool `json:"is_pay" gorm:"column:is_pay;type:boolean;default:false"`
}

func (p *Cashbook) TableName() string {
	return "cashbooks"
}
