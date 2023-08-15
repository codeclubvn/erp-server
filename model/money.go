package model

import (
	"github.com/google/uuid"
)

type Money struct {
	BaseModel
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	UserId      uuid.UUID `json:"user_id"`
}

func (Money) TableName() string {
	return "money"
}

type MoneyRequest struct {
	ID          *uuid.UUID `json:"id"`
	Status      *string    `json:"status"`
	Description *string    `json:"description"`
	Price       *float64   `json:"price"`
	UserId      *uuid.UUID
}

type Moneys []Money

type OneMoneyRequest struct {
	Id     string
	UserId string
}
