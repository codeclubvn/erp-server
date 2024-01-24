package models

import (
	"time"
)

type Debt struct {
	BaseModel
	Amount    float64    `json:"amount"`
	Note      string     `json:"note"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Repeat    bool       `json:"repeat"`
	Spent     float64    `json:"spent" migration:"-"`
}

func (p *Debt) TableName() string {
	return "budgets"
}
