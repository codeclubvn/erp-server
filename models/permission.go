package models

import uuid "github.com/satori/go.uuid"

type Permission struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Method    string    `json:"method" gorm:"column:method;type:varchar(50);not null"`
	RoutePath string    `json:"route_path" gorm:"column:route_path;type:varchar(255);not null"`
}
