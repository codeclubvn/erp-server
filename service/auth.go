package service

import (
	"context"
	"erp/cmd/lib"
	config "erp/config"
	models "erp/domain"
	dto2 "erp/handler/dto/auth"
	"erp/repository"
	"erp/utils/api_errors"
	"erp/utils/constants"
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
		email              lib.Sendinblue
	}
)

func NewAuthService(userService UserService, config *config.Config, jwtService JwtService, resetPasswordToken repository.ResetPasswordToken, email lib.Sendinblue) AuthService {
	return &authService{
		userService:        userService,
		jwtService:         jwtService,
		config:             config,
		resetPasswordToken: resetPasswordToken,
		email:              email,
	}
}

func (s *authService) Register(ctx context.Context, req dto2.RegisterRequest) (user *models.User, err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return user, err
	}

	req.Password = string(encryptedPassword)

	user, err = s.userService.Create(ctx, models.User{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	})

	return user, err
}

func (s *authService) Login(ctx context.Context, req dto2.LoginRequest) (res *dto2.LoginResponse, err error) {
	user, err := s.userService.GetByEmail(ctx, req.Email)

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

	accessToken, refreshToken, err := s.jwtService.GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, errors.Wrap(err, "cannot generate auth tokens")
	}

	return &dto2.LoginResponse{
		User: dto2.UserResponse{
			ID:       user.ID.String(),
			FullName: user.FullName,
			Email:    user.Email,
			RoleID:   user.RoleID,
		},
		Token: dto2.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    s.config.Jwt.AccessTokenExpiresIn,
		},
	}, nil
}

func (s *authService) ForgotPassword(ctx context.Context, req dto2.ForgotPasswordRequest) (err error) {
	user, err := s.userService.GetByEmail(ctx, req.Email)
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
	if err = s.resetPasswordToken.Create(nil, ctx, output); err != nil {
		return errors.Wrap(err, "cannot create reset password token")
	}

	go func() {
		// {{params.RESET_PASSWORD_LINK}} will be replaced by the actual link
		params := map[string]interface{}{
			constants.ResetPasswordLink: fmt.Sprintf("%s/auth/reset-password/%s", s.config.Server.WebsiteURL, output.ID),
		}
		s.email.SendMail(context.Background(), user.Email, user.FullName, constants.TemplateIdOfEmailResetPassword, params)
	}()
	return
}

func (s *authService) ResetPassword(ctx context.Context, req dto2.ResetPasswordRequest) (err error) {
	// validate token
	resetPasswordToken, err := s.resetPasswordToken.GetOneById(ctx, req.Token)
	if err != nil {
		return errors.Wrap(err, "cannot get reset password token")
	}

	if resetPasswordToken == nil {
		return errors.New("invalid token")
	}

	// reset password
	if err = s.userService.UpdatePassword(ctx, resetPasswordToken.UserID, req.Password); err != nil {
		return errors.Wrap(err, "cannot update password")
	}

	return
}
