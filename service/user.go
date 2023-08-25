package service

import (
	"context"
	config "erp/config"
	models "erp/models"
	repository "erp/repository"

	"gorm.io/gorm"
)

type (
	UserService interface {
		Create(ctx context.Context, user models.User) (*models.User, error)
		GetByID(ctx context.Context, id string) (*models.User, error)
		GetByEmail(ctx context.Context, email string) (*models.User, error)
	}
	UserServiceImpl struct {
		userRepo repository.UserRepository
		config   *config.Config
	}
)

func (u *UserServiceImpl) Create(ctx context.Context, user models.User) (*models.User, error) {
	r, err := u.userRepo.Create(ctx, user)
	return r, err
}

func NewUserService(itemRepo repository.UserRepository, config *config.Config) UserService {
	return &UserServiceImpl{
		userRepo: itemRepo,
		config:   config,
	}
}

func (u *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		}
		return nil, err
	}
	return user, err
}

func (u *UserServiceImpl) GetByID(ctx context.Context, id string) (user *models.User, err error) {
	user, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		}
		return nil, err
	}
	return
}
