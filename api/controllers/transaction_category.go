package controller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	erpservice "erp/service"
	"erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	api.BaseController
	TransactionService erpservice.TransactionService
}

func NewTransactionController(revenueService erpservice.TransactionService) *TransactionController {
	return &TransactionController{
		TransactionService: revenueService,
	}
}

func (p *TransactionController) Create(c *gin.Context) {
	var req erpdto.CreateRevenueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	revenue, err := p.TransactionService.Create(nil, c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", revenue.ID)
}

func (p *TransactionController) Update(c *gin.Context) {
	var req erpdto.UpdateRevenueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	_, err := p.TransactionService.Update(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *TransactionController) List(c *gin.Context) {
	var req erpdto.ListRevenueRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	revenues, total, err := p.TransactionService.GetList(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, revenues)
}

func (p *TransactionController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	err := p.TransactionService.Delete(c.Request.Context(), id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *TransactionController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := p.TransactionService.GetOne(c, id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}
	p.Response(c, http.StatusOK, "success", res)
}
