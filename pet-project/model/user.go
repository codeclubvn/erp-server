package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model              // Tạo ID tự tăng "identity"
	Username      string    `json:"username" gorm:"unique"`
	Password      string    `json:"password"`
	Hoten         string    `json:"name"`
	Address       string    `json:"address"`
	Date_Of_Birth time.Time `json:"date_of_birth"`
	Email         string    `json:"email" gorm:"unique"`
	Role          string    `json:"role,omitempty"`
	Phone         string    `json:"phone"`
	Create_id     int       `json:"create_id"`
	Update_id     int       `json:"update_id"`
	DelatedTime   time.Time `json:"delated_time" gorm:"datetime"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
