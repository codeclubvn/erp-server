package models

type RevenueCategory struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (p *RevenueCategory) TableName() string {
	return "revenue_category"
}
