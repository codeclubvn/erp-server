package model

import (
	"github.com/google/uuid"
	"time"
)

type Business struct {
	BaseModel
	Name         string    `json:"name" gorm:"not null"`
	Avatar       string    `json:"avatar"`
	Background   string    `json:"background"`
	Domain       string    `json:"domain" gorm:"not null" sql:"index"`
	PhoneNumber  string    `json:"phone_number" sql:"index"`
	Bio          string    `json:"bio" gorm:"null"`
	Address      string    `json:"address"`
	OpenTime     time.Time `json:"open_time"`
	CloseTime    time.Time `json:"close_time"`
	BusinessType string    `json:"business_type,omitempty"`
	IsClose      bool      `json:"is_close"`
}

func (Business) TableName() string {
	return "business"
}

type BusinessRequest struct {
	ID          *uuid.UUID `json:"id"`
	Name        *string    `json:"name" valid:"Required"`
	Domain      *string    `json:"domain"`
	Bio         *string    `json:"bio"`
	Avatar      *string    `json:"avatar"`
	Background  *string    `json:"background"`
	Address     *string    `json:"address"`
	PhoneNumber *string    `json:"phone_number"`
	IsClose     *bool      `json:"is_close,omitempty"`
	OpenTime    *time.Time `json:"open_time"`
	CloseTime   *time.Time `json:"close_time"`
	UserId      *uuid.UUID
}
