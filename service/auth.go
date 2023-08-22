package service

import (
	"context"
	config "erp/config"
	constants "erp/constants"
	dto "erp/dto/auth"
	models "erp/models"

	"golang.org/x/crypto/bcrypt"
)

type (
	AuthService interface {
		Register(ctx context.Context, req dto.RegisterRequest) (item models.User, err error)
	}
	AuthServiceImpl struct {
		userService UserService
		config      config.Config
	}
)

func NewAuthService(userService UserService, config config.Config) AuthService {
	return &AuthServiceImpl{
		userService: userService,
		config:      config,
	}
}

func (a *AuthServiceImpl) Register(ctx context.Context, req dto.RegisterRequest) (user models.User, err error) {
	role := constants.RoleCustomer

	if req.RequestFrom != string(constants.Web) {
		role = constants.RoleSeller
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return user, err
	}

	req.Password = string(encryptedPassword)

	user, err = a.userService.Create(ctx, models.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      role,
	})

	return user, err
}
