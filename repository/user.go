package repository

import (
	"context"
	api_errors "erp/api_errors"
	"erp/infrastructure"
	models "erp/models"
	"erp/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (res *models.User, err error)
	GetByEmail(ctx context.Context, email string) (res *models.User, err error)
	Create(tx *TX, ctx context.Context, user models.User) (res *models.User, err error)
}

type userRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewUserRepository(db *infrastructure.Database, logger *zap.Logger) UserRepository {
	utils.MustHaveDb(db)
	return &userRepository{db, logger}
}

func (u *userRepository) Create(tx *TX, ctx context.Context, user models.User) (res *models.User, err error) {
	if tx != nil {
		tx = &TX{db: *u.db}
	}
	err = u.db.Create(&user).Error

	return &user, errors.Wrap(err, "create user error")
}

func (u *userRepository) GetByID(ctx context.Context, id string) (res *models.User, err error) {
	err = u.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return res, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, errors.Wrap(err, "get user by id error")
	}
	return
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (res *models.User, err error) {
	err = u.db.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, err
	}
	return
}
