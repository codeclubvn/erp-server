package domain

import (
	uuid "github.com/satori/go.uuid"
)

type Role struct {
	BaseModel
	Name        string       `json:"name" gorm:"column:name;type:varchar(200);not null"`
	Extends     []Role       `gorm:"many2many:role_extends"`
	StoreID     uuid.UUID    `json:"store_id" gorm:"column:store_id;type:uuid"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
	Users       []User       `gorm:"foreignKey:RoleID"`
}
