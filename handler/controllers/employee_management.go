package controller

import (
	"erp/cmd/lib"
	"erp/handler/dto"
	"erp/handler/dto/erp"
	erpservice "erp/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ERPEmployeeManagementController struct {
	dto.BaseController
	handler                   *lib.Handler
	employeeManagementService erpservice.ERPEmployeeManagementService
	logger                    *zap.Logger
}

func NewERPEmployeeManagementController(handler *lib.Handler, logger *zap.Logger, employeeManagementService erpservice.ERPEmployeeManagementService) *ERPEmployeeManagementController {
	return &ERPEmployeeManagementController{
		handler:                   handler,
		logger:                    logger,
		employeeManagementService: employeeManagementService,
	}
}

func (p *ERPEmployeeManagementController) GetList(c *gin.Context) {
	data, total, err := p.employeeManagementService.ListPermission()
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", total, data)
}

func (p *ERPEmployeeManagementController) CreateRole(c *gin.Context) {
	var req erpdto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	role, err := p.employeeManagementService.CreateRole(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", role)
}

func (p *ERPEmployeeManagementController) CreateEmployee(c *gin.Context) {
	var req erpdto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	employee, err := p.employeeManagementService.CreateEmployee(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", employee)
}

// func (p *ERPPermissionController) Manage(c *gin.Context) {
// 	var req erpdto.ManagePermissionRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		p.ResponseValidationError(c, err)
// 		return
// 	}

// 	if err := p.permissonService.Manage(req); err != nil {
// 		p.ResponseError(c, err)
// 		return
// 	}

// 	p.ResponseSuccess(c, "Success")
// }
