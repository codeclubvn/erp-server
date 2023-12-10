package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Budget struct {
	BaseModel
	Amount             float64           `json:"amount"`
	Note               string            `json:"note"`
	StartTime          *time.Time        `json:"start_time"`
	EndTime            *time.Time        `json:"end_time"`
	Repeat             bool              `json:"repeat"`
	Spent              float64           `json:"spent" migration:"-"`
	CashbookCategoryId *uuid.UUID        `json:"cashbook_category_id"`
	CashbookCategory   *CashbookCategory `json:"cashbook_category" gorm:"foreignKey:CashbookCategoryId"`
}

func (p *Budget) TableName() string {
	return "budgets"
}
