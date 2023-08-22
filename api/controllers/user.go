package controller

import (
	service "erp/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	userService service.UserService
	logger      *zap.Logger
	BaseController
}

func NewUserController(c *gin.RouterGroup, userService service.UserService, logger *zap.Logger) *UserController {
	controller := &UserController{
		userService: userService,
		logger:      logger,
	}
	_ = c.Group("/user")

	return controller
}
