package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type OrderRoutes struct {
	handler *ERPHandler
}

func NewOrderRoutes(handler *ERPHandler, controller *erpcontroller.OrderController, middleware *middlewares.GinMiddleware) *OrderRoutes {
	g := handler.Group("/orders")

	g.POST("/", middleware.Auth(true), controller.Create)

	return &OrderRoutes{
		handler: handler,
	}
}
