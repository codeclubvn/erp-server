package repository

import (
	"context"
	api_errors "erp/api_errors"
	"erp/infrastructure"
	models "erp/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (res *models.User, err error)
	GetByEmail(ctx context.Context, email string) (res *models.User, err error)
	Create(ctx context.Context, user models.User) (res *models.User, err error)
	GetBySocailId(ctx context.Context, socialId string) (res *models.User, err error)
}

type userRepositoryImpl struct {
	*infrastructure.Database
}

func NewUserRepository(db *infrastructure.Database) UserRepository {
	if db == nil {
		panic("Database engine is null")
	}
	return &userRepositoryImpl{db}
}

func (u *userRepositoryImpl) Create(ctx context.Context, user models.User) (res *models.User, err error) {
	err = u.DB.Create(&user).Error

	return &user, err
}

func (u *userRepositoryImpl) GetByID(ctx context.Context, id string) (res *models.User, err error) {
	err = u.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(api_errors.UserNotFound)
		}
		return nil, err
	}
	return
}

func (u *userRepositoryImpl) GetByEmail(ctx context.Context, email string) (res *models.User, err error) {
	err = u.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(api_errors.UserNotFound)
		}
		return nil, err
	}
	return
}

func (u *userRepositoryImpl) GetBySocailId(ctx context.Context, socialId string) (res *models.User, err error) {
	err = u.WithContext(ctx).Where("social_id = ?", socialId).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(api_errors.UserNotFound)
		}
		return nil, err
	}
	return
}
