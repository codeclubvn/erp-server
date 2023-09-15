package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type employeeManagementRoutes struct {
	handler *ERPHandler
}

func NewEmployeeManagementRoutes(handler *ERPHandler, controller *erpcontroller.ERPEmployeeManagementController, middleware *middlewares.GinMiddleware) *employeeManagementRoutes {
	g := handler.Group("/employee-management")

	p := g.Group("/permission")
	p.GET("/", middleware.Auth(false), controller.ListPermission)

	r := g.Group("/role")
	r.POST("/", middleware.Auth(true), controller.CreateRole)

	e := g.Group("/employee")
	e.POST("/", middleware.Auth(true), controller.CreateEmployee)

	return &employeeManagementRoutes{
		handler: handler,
	}
}
