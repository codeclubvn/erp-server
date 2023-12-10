package models

type TransactionCategory struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
}

func (p *TransactionCategory) TableName() string {
	return "transaction_categories"
}
