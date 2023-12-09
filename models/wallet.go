package models

type Wallet struct {
	BaseModel
	Name   string  `gorm:"column:name;type:varchar(255);not null"`
	Amount float64 `gorm:"column:amount;type:float;not null"`
}
