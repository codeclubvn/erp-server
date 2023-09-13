package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserRole struct {
	UpdaterID    uuid.UUID `json:"updater_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:now();not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:now();not null" json:"updated_at"`
	UserID       uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;not null;uniqueIndex:idx_user_store;primary_key"`
	User         User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	RoleID       uuid.UUID `json:"role_id" gorm:"column:role_id;type:uuid;not null;primary_key"`
	Role         Role      `json:"role" gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE"`
	StoreID      uuid.UUID `json:"store_id" gorm:"column:store_id;type:uuid;not null;uniqueIndex:idx_user_store;primary_key"`
	Store        Store     `json:"store" gorm:"foreignKey:StoreID;references:ID;constraint:OnDelete:CASCADE"`
	IsStoreOwner bool      `json:"is_store_owner" gorm:"column:is_store_owner;default:false"`
}
