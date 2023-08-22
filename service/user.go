package service

import (
	"context"
	config "erp/config"
	models "erp/models"
	repository "erp/repository"
)

type (
	UserService interface {
		Create(ctx context.Context, user models.User) (models.User, error)
		GetByID(ctx context.Context, id string) (models.User, error)
	}
	UserServiceImpl struct {
		userRepo repository.UserRepository
		config   config.Config
	}
)

// Create implements UserService.
func (u *UserServiceImpl) Create(ctx context.Context, user models.User) (models.User, error) {
	user, err := u.userRepo.Create(ctx, user)
	return user, err
}

func NewUserService(itemRepo repository.UserRepository, config config.Config) UserService {
	return &UserServiceImpl{
		userRepo: itemRepo,
		config:   config,
	}
}

func (s *UserServiceImpl) GetByID(ctx context.Context, id string) (item models.User, err error) {
	item, err = s.userRepo.GetByID(ctx, id)
	return
}
