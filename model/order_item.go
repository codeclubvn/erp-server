package model

type OrderItem struct {
	BaseModel
	OrderId     string  `json:"order_id"`
	ProductId   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

type OrderItemRequest struct {
	ProductId *string  `json:"product_id"`
	Quantity  *int     `json:"quantity"`
	Price     *float64 `json:"price"`
}
