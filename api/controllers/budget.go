package controller

import (
	"erp/api"
	erpdto "erp/api/dto/finance"
	erpservice "erp/service"
	"erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BudgetController struct {
	api.BaseController
	BudgetService erpservice.BudgetService
}

func NewBudgetController(budgetService erpservice.BudgetService) *BudgetController {
	return &BudgetController{
		BudgetService: budgetService,
	}
}

func (p *BudgetController) Create(c *gin.Context) {
	var req erpdto.CreateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	budget, err := p.BudgetService.Create(nil, c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", budget.ID)
}

func (p *BudgetController) Update(c *gin.Context) {
	var req erpdto.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	_, err := p.BudgetService.Update(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *BudgetController) List(c *gin.Context) {
	var req erpdto.ListBudgetRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	budgets, total, err := p.BudgetService.GetList(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, budgets)
}

func (p *BudgetController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	err := p.BudgetService.Delete(c.Request.Context(), id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *BudgetController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := p.BudgetService.GetOne(c, id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}
	p.Response(c, http.StatusOK, "success", res)
}
