package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type CategoryRoutes struct {
	handler *ERPHandler
}

func NewCategoryRoutes(handler *ERPHandler, controller *erpcontroller.ERPCategoryController, middleware *middlewares.GinMiddleware) *CategoryRoutes {
	g := handler.Group("/category")

	g.POST("/", middleware.Auth(true), controller.Create)
	g.PUT("/", middleware.Auth(true), controller.Update)
	g.GET("/", middleware.Auth(false), controller.GetList)
	g.GET("/:id", middleware.Auth(false), controller.GetOne)
	g.DELETE("/:id", middleware.Auth(true), controller.Delete)

	return &CategoryRoutes{
		handler: handler,
	}
}
