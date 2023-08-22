package controller

import (
	dto "erp/dto/auth"
	service "erp/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthService service.AuthService
	logger      *zap.Logger
	BaseController
}

func NewAuthController(c *gin.RouterGroup, authService service.AuthService, logger *zap.Logger) *AuthController {
	controller := &AuthController{
		AuthService: authService,
		logger:      logger,
	}
	g := c.Group("/auth")
	g.POST("/register", controller.Register)

	return controller
}

func (b *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}

	_, err := b.AuthService.Register(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}
