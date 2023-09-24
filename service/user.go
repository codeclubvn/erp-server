package service

import (
	"context"
	config "erp/config"
	"erp/dto/dto"
	models "erp/models"
	repository "erp/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine/log"

	"gorm.io/gorm"
)

type (
	UserService interface {
		Create(ctx context.Context, user models.User) (*models.User, error)
		GetByID(ctx context.Context, id string) (*models.User, error)
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		GetBySocialId(ctx context.Context, socialId string) (*models.User, error)
		ChangePassword(ctx context.Context, req dto.ChangePassword, userId string) error
	}
	UserServiceImpl struct {
		userRepo repository.UserRepository
		config   *config.Config
	}
)

func (u *UserServiceImpl) GetBySocialId(ctx context.Context, socialId string) (*models.User, error) {
	//TODO implement me
	user, err := u.userRepo.GetBySocailId(ctx, socialId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

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
func (u *UserServiceImpl) ChangePassword(ctx context.Context, req dto.ChangePassword, userId string) error {
	user, err := u.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
	if err != nil {
		log.Errorf(ctx, fmt.Sprintf("Old password not match, err : %+v", err))
		return err
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	user.Password = string(encryptedPassword)
	_, err = u.userRepo.ChangePassword(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
