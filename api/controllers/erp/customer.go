package erpcontroller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	erpservice "erp/service/erp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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

func (p *ERPCustomerController) ListCustomer(c *gin.Context) {
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

func (p *ERPCustomerController) CustomerDetail(c *gin.Context) {
	var req erpdto.CustomerUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	customer, err := p.customerService.CustomerDetail(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.Response(c, http.StatusOK, "Customer not found", nil)
		} else {
			p.ResponseError(c, err)
		}
		return
	}

	p.Response(c, http.StatusOK, "Success", customer)
}

func (p *ERPCustomerController) CreateCustomer(c *gin.Context) {
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

func (p *ERPCustomerController) UpdateCustomer(c *gin.Context) {
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

func (p *ERPCustomerController) DeleteCustomer(c *gin.Context) {
	var req erpdto.CustomerUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	fmt.Println("=======================================")
	fmt.Println(req.ID)
	fmt.Println("=======================================")

	err := p.customerService.DeleteCustomer(c.Request.Context(), req.ID)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", nil)
}