package controller

import (
	"erp/config"
	"erp/dto/dto"
	service "erp/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	userService service.UserService
	logger      *zap.Logger
	jwtService  service.JwtService
	BaseController
}

func NewUserController(c *gin.RouterGroup, userService service.UserService, logger *zap.Logger, config *config.Config) *UserController {
	controller := &UserController{
		userService: userService,
		logger:      logger,
	}
	g := c.Group("/user")
	g.PATCH("/change_password", controller.ChangePassword)
	return controller
}

func (u *UserController) ChangePassword(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	jwtToken := strings.Split(auth, " ")[1]
	userIdJwt, err := u.jwtService.ValidateToken(jwtToken)

	if err != nil {
		u.ResponseError(c, http.StatusUnauthorized, []error{err})
		return
	}
	var req dto.ChangePassword
	if err = c.ShouldBindJSON(&req); err != nil {
		u.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	err = u.userService.ChangePassword(c, req, *userIdJwt)
	if err != nil {
		u.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	u.Response(c, http.StatusOK, "Change password success", nil)
}
