package controller

import (
	"erp/api"
	erpdto "erp/api/dto/finance"
	erpservice "erp/service"
	"erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CashbookController struct {
	api.BaseController
	cashbookService erpservice.CashbookService
}

func NewCashbookController(outputervice erpservice.CashbookService) *CashbookController {
	return &CashbookController{
		cashbookService: outputervice,
	}
}

func (p *CashbookController) Create(c *gin.Context) {
	var req erpdto.CreateCashbookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	output, err := p.cashbookService.Create(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", output.ID)
}

func (p *CashbookController) Update(c *gin.Context) {
	var req erpdto.UpdateCashbookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	_, err := p.cashbookService.Update(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *CashbookController) List(c *gin.Context) {
	var req erpdto.ListCashbookRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	output, total, err := p.cashbookService.GetList(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, output)
}

func (p *CashbookController) ListDebt(c *gin.Context) {
	var req erpdto.ListCashbookRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	output, total, err := p.cashbookService.GetListDebt(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, output)
}

func (p *CashbookController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	err := p.cashbookService.Delete(c.Request.Context(), id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *CashbookController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := p.cashbookService.GetOne(c, id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}
	p.Response(c, http.StatusOK, "success", res)
}
