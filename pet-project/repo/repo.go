package repo

import (
	"gorm.io/gorm"

	"pet-project/model"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

type IRepo interface {
	GetUserByEmail(email string) (model.User, error)
}
