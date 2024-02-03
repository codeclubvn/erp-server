package service

import (
	"context"
	config "erp/config"
	models "erp/domain"
	repository "erp/repository"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type (
	UserService interface {
		Create(ctx context.Context, user models.User) (*models.User, error)
		GetByID(ctx context.Context, id string) (*models.User, error)
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		UpdatePassword(ctx context.Context, userId, password string) (err error)
	}
	userService struct {
		userRepo repository.UserRepository
		config   *config.Config
	}
)

func (u *userService) Create(ctx context.Context, user models.User) (*models.User, error) {
	r, err := u.userRepo.Create(nil, ctx, user)
	return r, err
}

func NewUserService(userRepo repository.UserRepository, config *config.Config) UserService {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}

func (u *userService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		}
		return nil, err
	}
	return user, err
}

func (u *userService) GetByID(ctx context.Context, id string) (user *models.User, err error) {
	user, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		}
		return nil, err
	}
	return
}

func (u *userService) UpdatePassword(ctx context.Context, userId, password string) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return
	}
	return u.userRepo.UpdatePassword(nil, ctx, userId, encryptedPassword)
}
