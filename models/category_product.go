package models

import uuid "github.com/satori/go.uuid"

type CategoryProduct struct {
	BaseModel
	CategoryId uuid.UUID `json:"category_id" gorm:"type:uuid;default:uuid_generate_v4();index"`
	ProductId  uuid.UUID `json:"product_id" gorm:"type:uuid;default:uuid_generate_v4();index"`
	Category   Category  `gorm:"foreignKey:category_id;"`
	Product    Product   `gorm:"foreignKey:product_id;"`
}

func (CategoryProduct) TableName() string {
	return "categories_products"
}
