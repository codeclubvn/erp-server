package repository

import (
	"context"
	api_errors "erp/api_errors"
	infrastructure "erp/infrastructure"
	models "erp/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (res models.User, err error)
	IsExistEmail(ctx context.Context, email string) (res models.User, err error)
	Create(ctx context.Context, user models.User) (res models.User, err error)
}

type UserRepositoryImpl struct {
	*infrastructure.Database
}

func NewUserRepository(db *infrastructure.Database) UserRepository {
	if db == nil {
		panic("Database engine is null")
	}
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, user models.User) (res models.User, err error) {
	err = u.DB.Create(&user).Error

	return user, err
}

func (u *UserRepositoryImpl) GetByID(ctx context.Context, id string) (res models.User, err error) {
	err = u.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(api_errors.UserNotFound)
		}
		return models.User{}, err
	}
	return
}

func (u *UserRepositoryImpl) IsExistEmail(ctx context.Context, email string) (res models.User, err error) {
	err = u.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(api_errors.UserNotFound)
		}
		return models.User{}, err
	}
	return
}
