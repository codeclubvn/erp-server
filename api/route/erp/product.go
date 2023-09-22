package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type ProductRoutes struct {
	handler *ERPHandler
}

func NewProductRoutes(handler *ERPHandler, controller *erpcontroller.ERPProductController, middleware *middlewares.GinMiddleware) *ProductRoutes {
	g := handler.Group("/products")

	g.POST("/", middleware.Auth(true), controller.Create)
	g.PUT("/", middleware.Auth(true), controller.Update)
	g.DELETE("/:id", middleware.Auth(true), controller.Delete)
	g.GET("/:id", middleware.Auth(false), controller.GetOne)
	g.GET("/", middleware.Auth(false), controller.GetList)

	return &ProductRoutes{
		handler: handler,
	}
}
