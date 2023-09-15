package models

import uuid "github.com/satori/go.uuid"

type User struct {
	BaseModel
	FirstName string     `json:"first_name" gorm:"column:first_name;type:varchar(50);not null"`
	LastName  string     `json:"last_name" gorm:"column:last_name;type:varchar(50);not null"`
	Email     string     `json:"email" gorm:"column:email;type:varchar(100);not null"`
	Password  string     `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Social    string     `json:"social"`
	SocialID  string     `json:"social_id"`
	Role      Role       `json:"role" gorm:"foreignKey:RoleID"`
	RoleID    *uuid.UUID `json:"role_id" gorm:"column:role_id;type:uuid"`
}
