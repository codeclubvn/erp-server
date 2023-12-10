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

func NewTransactionController(transactionService erpservice.TransactionService) *TransactionController {
	return &TransactionController{
		TransactionService: transactionService,
	}
}

func (p *TransactionController) Create(c *gin.Context) {
	var req erpdto.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	transaction, err := p.TransactionService.Create(nil, c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", transaction.ID)
}

func (p *TransactionController) Update(c *gin.Context) {
	var req erpdto.UpdateTransactionRequest
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
	var req erpdto.ListTransactionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	transactions, total, err := p.TransactionService.GetList(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, transactions)
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
