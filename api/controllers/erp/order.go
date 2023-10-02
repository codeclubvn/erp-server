package erpcontroller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	erpservice "erp/service/erp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ERPOrderController struct {
	api.BaseController
	storeService erpservice.ERPOrderService
}

func NewERPOrderController(storeService erpservice.ERPOrderService) *ERPOrderController {
	return &ERPOrderController{
		storeService: storeService,
	}
}

func (p *ERPOrderController) CreateOrder(c *gin.Context) {
	var req erpdto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	store, err := p.storeService.CreateOrder(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", store.ID)
}
