package domain

type ResetPasswordToken struct {
	BaseModel
	UserID string `json:"user_id"`
}

func (ResetPasswordToken) TableName() string {
	return "reset_password_token"
}
