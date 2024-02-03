package service

import (
	"context"
	config "erp/config"
	models "erp/domain"
	dto2 "erp/handler/dto/auth"
	"erp/repository"
	"erp/utils/api_errors"
	"fmt"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type (
	AuthService interface {
		Register(ctx context.Context, req dto2.RegisterRequest) (user *models.User, err error)
		Login(ctx context.Context, req dto2.LoginRequest) (res *dto2.LoginResponse, err error)
		ForgotPassword(ctx context.Context, req dto2.ForgotPasswordRequest) (err error)
		ResetPassword(ctx context.Context, req dto2.ResetPasswordRequest) (err error)
	}
	authService struct {
		userService        UserService
		jwtService         JwtService
		config             *config.Config
		resetPasswordToken repository.ResetPasswordToken
	}
)

func NewAuthService(userService UserService, config *config.Config, jwtService JwtService, resetPasswordToken repository.ResetPasswordToken) AuthService {
	return &authService{
		userService:        userService,
		jwtService:         jwtService,
		config:             config,
		resetPasswordToken: resetPasswordToken,
	}
}

func (a *authService) Register(ctx context.Context, req dto2.RegisterRequest) (user *models.User, err error) {
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
	})

	return user, err
}

func (a *authService) Login(ctx context.Context, req dto2.LoginRequest) (res *dto2.LoginResponse, err error) {
	user, err := a.userService.GetByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	//if req.RequestFrom == string(constants.Erp) {
	//	// account is not for erp will not have role id
	//	if user.RoleID == nil {
	//		return nil, errors.New(api_errors.ErrUnauthorizedAccess)
	//	}
	//}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, errors.New(api_errors.ErrInvalidPassword)
	}

	accessToken, refreshToken, err := a.jwtService.GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, errors.Wrap(err, "cannot generate auth tokens")
	}

	return &dto2.LoginResponse{
		User: dto2.UserResponse{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			RoleID:    user.RoleID,
		},
		Token: dto2.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    a.config.Jwt.AccessTokenExpiresIn,
		},
	}, nil
}

func (a *authService) ForgotPassword(ctx context.Context, req dto2.ForgotPasswordRequest) (err error) {
	user, err := a.userService.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New(api_errors.ErrUserNotFound)
	}

	// save token to db
	output := &models.ResetPasswordToken{
		UserID: user.ID.String(),
	}
	if err = a.resetPasswordToken.Create(nil, ctx, output); err != nil {
		return errors.Wrap(err, "cannot create reset password token")
	}
	// send email with reset password link
	fmt.Println("resetPasswordToken", output.ID)

	return
}

func (a *authService) ResetPassword(ctx context.Context, req dto2.ResetPasswordRequest) (err error) {
	// validate token
	resetPasswordToken, err := a.resetPasswordToken.GetOneById(ctx, req.Token)
	if err != nil {
		return errors.Wrap(err, "cannot get reset password token")
	}

	if resetPasswordToken == nil {
		return errors.New("invalid token")
	}

	// reset password
	if err = a.userService.UpdatePassword(ctx, resetPasswordToken.UserID, req.Password); err != nil {
		return errors.Wrap(err, "cannot update password")
	}

	return
}
