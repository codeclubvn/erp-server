package dto

type RegisterRequest struct {
	Email       string `json:"email" binding:"required" validate:"email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	FirstName   string `json:"first_name" binding:"required" validate:"min=1,max=50"`
	LastName    string `json:"last_name" binding:"required" validate:"min=1,max=50"`
	RequestFrom string `json:"request_from" binding:"required" enums:"erp/,web,app"`
}

type UserGoogleRequest struct {
	Email     string `mapstructure:"email" json:"email" binding:"required" validate:"email"`
	GoogleID  string `mapstructure:"id" json:"id" binding:"required"`
	FirstName string `mapstructure:"family_name" json:"family_name"`
	LastName  string `mapstructure:"given_name" json:"given_name"`
}
