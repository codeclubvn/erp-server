package route

import (
	controller "erp/api/controllers"
	"erp/lib"
)

type AuthRoutes struct {
	handler *lib.Handler
}

func NewAuthRoutes(handler *lib.Handler, controller *controller.AuthController) *AuthRoutes {
	g := handler.Group("/auth")
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
	return &AuthRoutes{
		handler: handler,
	}
}
