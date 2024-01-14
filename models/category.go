package models

import uuid "github.com/satori/go.uuid"

type Category struct {
	BaseModel
	StoreId      uuid.UUID `json:"store_id" gorm:"type:uuid;default:uuid_generate_v4();index"`
	Name         string    `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Image        string    `json:"image" gorm:"column:image;type:varchar(250);null"`
	TotalProduct int       `json:"total_product" gorm:"column:total_product;type:int;default:0;"`
}

type CategoryResult struct {
	BaseModel
	StoreId      uuid.UUID `json:"store_id" gorm:"type:uuid;default:uuid_generate_v4();index"`
	Name         string    `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Image        string    `json:"image" gorm:"column:image;type:varchar(250);null"`
	TotalProduct int       `json:"total_product" gorm:"column:total_product;type:int;default:0;"`
}

func (Category) TableName() string {
	return "categories"
}
