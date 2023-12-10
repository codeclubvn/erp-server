package controller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	erpservice "erp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ERPCustomerController struct {
	api.BaseController
	customerService erpservice.ERPCustomerService
}

func NewERPCustomerController(customerService erpservice.ERPCustomerService) *ERPCustomerController {
	return &ERPCustomerController{
		customerService: customerService,
	}
}

func (p *ERPCustomerController) GetList(c *gin.Context) {
	var req erpdto.ListCustomerRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	customer, total, err := p.customerService.ListCustomer(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", total, customer)
}

func (p *ERPCustomerController) GetOne(c *gin.Context) {
	var req erpdto.CustomerUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	customer, err := p.customerService.CustomerDetail(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", customer)
}

func (p *ERPCustomerController) Create(c *gin.Context) {
	var req erpdto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	customer, err := p.customerService.CreateCustomer(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", customer.ID)
}

func (p *ERPCustomerController) Update(c *gin.Context) {
	var req erpdto.UpdateCustomerRequest

	if err := c.ShouldBindUri(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	_, err := p.customerService.UpdateCustomer(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", nil)
}

func (p *ERPCustomerController) Delete(c *gin.Context) {
	var req erpdto.CustomerUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	err := p.customerService.DeleteCustomer(c.Request.Context(), req.ID)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", nil)
}
