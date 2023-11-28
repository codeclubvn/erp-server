package models

type Product struct {
	BaseModel
	Name        string  `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Description string  `json:"description" gorm:"column:description;type:varchar(250);null"`
	Image       string  `json:"image" gorm:"column:image;type:varchar(250);null"`
	Price       float64 `json:"price" gorm:"column:price;type:float;default:0;"`
	Status      bool    `json:"status" gorm:"column:status;type:boolean;default:true;"`
	Quantity    *int    `json:"quantity" gorm:"column:quantity;type:int;default:null;"`
	//StoreId      uuid.UUID `json:"store_id" gorm:"column:store_id;type:uuid;not null"`
	Sold         int     `json:"sold" gorm:"column:sold;type:int;default:0;"`
	PromotePrice float64 `json:"promote_price" gorm:"column:promote_price;type:float;default:0;"`
}

func (Product) TableName() string {
	return "products"
}
