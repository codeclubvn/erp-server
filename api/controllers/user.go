package controller

import (
	"erp/api"
	service "erp/service"

	"go.uber.org/zap"
)

type UserController struct {
	api.BaseController
	userService service.UserService
	logger      *zap.Logger
}

func NewUserController(userService service.UserService, logger *zap.Logger) *UserController {
	controller := &UserController{
		userService: userService,
		logger:      logger,
	}
	return controller
}
