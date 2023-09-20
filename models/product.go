package models

type Product struct {
	BaseModel
	Name          string  `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Description   string  `json:"description" gorm:"column:description;type:varchar(250);null"`
	Image         string  `json:"image" gorm:"column:image;type:varchar(250);null"`
	Price         float64 `json:"price" gorm:"column:price;type:float;default:0;"`
	Status        bool    `json:"status" gorm:"column:status;type:boolean;default:true;"`
	NumberProduct int     `json:"number_product" gorm:"column:number_product;type:int;default:0;"`
}

func (Product) TableName() string {
	return "products"
}
