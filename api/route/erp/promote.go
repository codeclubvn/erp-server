package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type PromoteRoutes struct {
	handler *ERPHandler
}

func NewPromoteRoutes(handler *ERPHandler, controller *erpcontroller.PromoteController, middleware *middlewares.GinMiddleware) *PromoteRoutes {
	g := handler.Group("/promote")

	g.POST("/", middleware.Auth(true), controller.Create)

	return &PromoteRoutes{
		handler: handler,
	}
}
