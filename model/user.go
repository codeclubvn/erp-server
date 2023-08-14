package model

import (
	"time"
)

type User struct {
	BaseModel
	Username    string    `json:"username" gorm:"unique"`
	Password    string    `json:"password"`
	FullName    string    `json:"full_name"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email" gorm:"unique"`
	Role        string    `json:"role,omitempty"`
	Phone       string    `json:"phone"`
}

func (User) TableName() string {
	return "users"
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
