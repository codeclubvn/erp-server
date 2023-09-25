package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type CustomerRoutes struct {
	handler *ERPHandler
}

func NewCustomerRoutes(handler *ERPHandler, controller *erpcontroller.ERPCustomerController, middleware *middlewares.GinMiddleware) *CustomerRoutes {
	g := handler.Group("/customer")

	g.GET("/", middleware.Auth(false), controller.ListCustomer)
	g.GET("/:id", middleware.Auth(false), controller.CustomerDetail)
	g.POST("/", middleware.Auth(false), controller.CreateCustomer)
	g.PUT("/:id", middleware.Auth(false), controller.UpdateCustomer)
	g.DELETE("/:id", middleware.Auth(false), controller.DeleteCustomer)

	return &CustomerRoutes{
		handler: handler,
	}
}
