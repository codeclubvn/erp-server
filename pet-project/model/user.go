package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model // Tạo ID tự tăng "identity"

	Username string `json:"username" gorm:"unique"`

	// Thực ra là hashpassword nhưng sử dụng bằng tên password để lưu vào databaseư
	Password  string    `json:"password" gorm:"password"`
	Hoten     string    `json:"name"`
	Address   string    `json:"address"`
	NgaySinh  time.Time `json:"ngaysinh"`
	Email     string    `json:"email" gorm:"unique"`
	Role      string    `json:"role,omitempty"`
	SDT       string    `json:"sdt"`
	Create_id int       `json:"create_id"`
	Update_id int       `json:"update_id"`
	DelatedAt time.Time `json:"delated_at" gorm:"datetime"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
