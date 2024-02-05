package domain

import uuid "github.com/satori/go.uuid"

type User struct {
	BaseModel
	FullName string     `json:"full_name" gorm:"column:full_name;type:varchar(50);not null"`
	Email    string     `json:"email" gorm:"column:email;type:varchar(100);not null"`
	Password string     `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Social   string     `json:"social"`
	SocialID string     `json:"social_id"`
	Role     Role       `json:"role" gorm:"foreignKey:RoleID"`
	RoleID   *uuid.UUID `json:"role_id" gorm:"column:role_id;type:uuid"`
}
