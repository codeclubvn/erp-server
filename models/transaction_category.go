package models

type CashbookCategory struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Type string `json:"type" gorm:"column:type;type:varchar(255);not null"` // expense, income, debt
}

func (p *CashbookCategory) TableName() string {
	return "cashbook_categories"
}
