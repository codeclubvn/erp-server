package controller

import (
	"erp/handler/dto"
	erpdto "erp/handler/dto/finance"
	erpservice "erp/service"
	"erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionCategoryController struct {
	dto.BaseController
	TransactionCategoryService erpservice.TransactionCategoryService
}

func NewTransactionCategoryController(transactionCategoryService erpservice.TransactionCategoryService) *TransactionCategoryController {
	return &TransactionCategoryController{
		TransactionCategoryService: transactionCategoryService,
	}
}

func (p *TransactionCategoryController) Create(c *gin.Context) {
	var req erpdto.CreateTransactionCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	transactionCategory, err := p.TransactionCategoryService.Create(nil, c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", transactionCategory.ID)
}

func (p *TransactionCategoryController) Update(c *gin.Context) {
	var req erpdto.UpdateTransactionCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	_, err := p.TransactionCategoryService.Update(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *TransactionCategoryController) List(c *gin.Context) {
	var req erpdto.ListTransactionCategoryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	transactionCategorys, total, err := p.TransactionCategoryService.GetList(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, transactionCategorys)
}

func (p *TransactionCategoryController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	err := p.TransactionCategoryService.Delete(c.Request.Context(), id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *TransactionCategoryController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := p.TransactionCategoryService.GetOne(c, id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}
	p.Response(c, http.StatusOK, "success", res)
}
