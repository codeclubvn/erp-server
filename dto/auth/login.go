package dto

type LoginRequest struct {
	Email       string `json:"email" binding:"required" validate:"email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	RequestFrom string `json:"request_from" binding:"required" enums:"erp/,web,app"`
}

type LoginResponse struct {
	User  UserResponse  `json:"user"`
	Token TokenResponse `json:"token"`
}

type LoginByGoogleRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	GoogleId string `json:"google_id" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	RoleKey   string `json:"role_key"`
}
