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
	g.PUT("/", middleware.Auth(true), controller.Update)

	return &OrderRoutes{
		handler: handler,
	}
}
