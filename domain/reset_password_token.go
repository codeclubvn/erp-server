package domain

import (
	uuid "github.com/satori/go.uuid"
)

type ResetPasswordToken struct {
	BaseModel
	UserID uuid.UUID `json:"user_id"`
}

func (ResetPasswordToken) TableName() string {
	return "reset_password_token"
}
