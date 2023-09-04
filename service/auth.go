package service

import (
	"context"
	"erp/api_errors"
	config "erp/config"
	"erp/constants"
	dto "erp/dto/auth"
	models "erp/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type (
	AuthService interface {
		Register(ctx context.Context, req dto.RegisterRequest) (user *models.User, err error)
		RegisterByGoogle(ctx context.Context, req dto.UserGoogleRequest) (user *models.User, err error)
		Login(ctx context.Context, req dto.LoginRequest) (res *dto.LoginResponse, err error)
		LoginByGoogle(ctx context.Context, req dto.LoginByGoogleRequest) (res *dto.LoginResponse, err error)
	}
	AuthServiceImpl struct {
		userService UserService
		jwtService  JwtService
		config      *config.Config
	}
)

func NewAuthService(userService UserService, config *config.Config, jwtService JwtService) AuthService {
	return &AuthServiceImpl{
		userService: userService,
		jwtService:  jwtService,
		config:      config,
	}
}

func (a *AuthServiceImpl) Register(ctx context.Context, req dto.RegisterRequest) (user *models.User, err error) {
	roleKey := constants.RoleCustomer
	log.Println(fmt.Sprintf("Request info %+v", req))
	if req.RequestFrom != string(constants.Web) {
		roleKey = constants.RoleStoreOwner
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
		RoleKey:   roleKey,
	})

	return user, err
}
func (a *AuthServiceImpl) RegisterByGoogle(ctx context.Context, req dto.UserGoogleRequest) (user *models.User, err error) {
	roleKey := constants.RoleCustomer
	log.Println(fmt.Sprintf("Request info %+v", req))
	//if req.RequestFrom != string(constants.Web) {
	//	roleKey = constants.RoleStoreOwner
	//}
	user, err = a.userService.Create(ctx, models.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Social:    "Google",
		SocialID:  req.GoogleID,
		RoleKey:   roleKey,
	})
	return user, err
}
func (a *AuthServiceImpl) Login(ctx context.Context, req dto.LoginRequest) (res *dto.LoginResponse, err error) {
	user, err := a.userService.GetByEmail(ctx, req.Email)

	if req.RequestFrom != string(constants.Web) {
		// TODO: Sẽ có thêm role của nhân viên
		if user.RoleKey != constants.RoleStoreOwner {
			return nil, api_errors.ErrUnauthorizedAccess
		}
	}

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := a.jwtService.GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, err
	}

	res = &dto.LoginResponse{
		User: dto.UserResponse{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			RoleKey:   user.RoleKey,
		},
		Token: dto.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    a.config.Jwt.AccessTokenExpiresIn,
		},
	}

	return res, nil
}
func (a *AuthServiceImpl) LoginByGoogle(ctx context.Context, req dto.LoginByGoogleRequest) (res *dto.LoginResponse, err error) {
	user, err := a.userService.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	_, err = a.userService.GetBySocialId(ctx, req.GoogleId)
	if err != nil {
		return nil, err
	}
	accessToken, refreshToken, err := a.jwtService.GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, err
	}
	res = &dto.LoginResponse{
		User: dto.UserResponse{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			RoleKey:   user.RoleKey,
		},
		Token: dto.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    a.config.Jwt.AccessTokenExpiresIn,
		},
	}

	return res, nil
}
