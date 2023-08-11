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
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
	GetUserById(id int) (model.User, error)
	GetAllUser() ([]model.User, error)
	GetRoleUser(role string) ([]model.User, error)
}
