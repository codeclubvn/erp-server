package controller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	erpservice "erp/service"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerController struct {
	api.BaseController
	customerService erpservice.ERPCustomerService
}

func NewERPCustomerController(customerService erpservice.ERPCustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (p *CustomerController) GetList(c *gin.Context) {
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

func (p *CustomerController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	customer, err := p.customerService.GetOneById(c.Request.Context(), id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", customer)
}

func (p *CustomerController) Create(c *gin.Context) {
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

func (p *CustomerController) Update(c *gin.Context) {
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

func (p *CustomerController) Delete(c *gin.Context) {
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
