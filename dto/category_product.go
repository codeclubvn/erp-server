package dto

type CategoryProductRequest struct {
	UserId     string `json:"user_id" binding:"required"`
	CategoryId string `json:"category_id" binding:"required"`
	ProductId  string `json:"product_id" binding:"required"`
}

type DeleteCatagoryProductRequest struct {
	UserId string `json:"user_id" binding:"required"`
	ID     string `json:"id" binding:"required"`
}
