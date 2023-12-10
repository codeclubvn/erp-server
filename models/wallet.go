package models

type Wallet struct {
	BaseModel
	Name      string  `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Amount    float64 `json:"amount" gorm:"column:amount;type:float;not null"`
	IsDefault bool    `json:"is_default" gorm:"column:is_default;type:boolean;not null"`
}

func (p *Wallet) TableName() string {
	return "wallets"
}
