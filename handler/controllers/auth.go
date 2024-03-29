package controller

import (
	"erp/handler/dto"
	dto2 "erp/handler/dto/auth"
	service "erp/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	dto.BaseController
	authService service.AuthService
	logger      *zap.Logger
}

func NewAuthController(authService service.AuthService, logger *zap.Logger) *AuthController {
	controller := &AuthController{
		authService: authService,
		logger:      logger,
	}
	return controller
}

func (b *AuthController) Register(c *gin.Context) {
	var req dto2.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	user, err := b.authService.Register(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", user)
}

func (b *AuthController) Login(c *gin.Context) {
	var req dto2.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	res, err := b.authService.Login(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *AuthController) ForgotPassword(c *gin.Context) {
	var req dto2.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	if err := b.authService.ForgotPassword(c.Request.Context(), req); err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *AuthController) ResetPassword(c *gin.Context) {
	var req dto2.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	if err := b.authService.ResetPassword(c.Request.Context(), req); err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}
