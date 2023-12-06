package controller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	erpservice "erp/service"
	"erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RevenueController struct {
	api.BaseController
	revenueService erpservice.RevenueService
}

func NewRevenueController(revenueService erpservice.RevenueService) *RevenueController {
	return &RevenueController{
		revenueService: revenueService,
	}
}

func (p *RevenueController) Create(c *gin.Context) {
	var req erpdto.CreateRevenueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	revenue, err := p.revenueService.Create(nil, c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", revenue.ID)
}

func (p *RevenueController) Update(c *gin.Context) {
	var req erpdto.UpdateRevenueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	_, err := p.revenueService.Update(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *RevenueController) List(c *gin.Context) {
	var req erpdto.ListRevenueRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	revenues, total, err := p.revenueService.GetList(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, revenues)
}

func (p *RevenueController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	err := p.revenueService.Delete(c.Request.Context(), id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}
