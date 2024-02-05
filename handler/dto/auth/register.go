package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
	FullName string `json:"full_name" binding:"required" validate:"min=1,max=100"`
	//RequestFrom string `json:"request_from" binding:"required" enums:"erp/,web,app"`
}
