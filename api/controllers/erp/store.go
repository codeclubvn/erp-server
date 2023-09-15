package erpcontroller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	erpservice "erp/service/erp"
	"erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ERPStoreController struct {
	api.BaseController
	storeService erpservice.ERPStoreService
}

func NewERPStoreController(storeService erpservice.ERPStoreService) *ERPStoreController {
	return &ERPStoreController{
		storeService: storeService,
	}
}

func (p *ERPStoreController) CreateStore(c *gin.Context) {
	var req erpdto.CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	store, err := p.storeService.CreateStoreAndAssignOwner(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", store.ID)
}

func (p *ERPStoreController) UpdateStore(c *gin.Context) {
	var req erpdto.UpdateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	storeID := utils.GetStoreIDFromContext(c.Request.Context())
	_, err := p.storeService.UpdateStore(c.Request.Context(), storeID, req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *ERPStoreController) ListStore(c *gin.Context) {
	var req erpdto.ListStoreRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	stores, total, err := p.storeService.ListStore(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", total, stores)
}

func (p *ERPStoreController) DeleteStore(c *gin.Context) {
	storeID := utils.GetStoreIDFromContext(c.Request.Context())
	err := p.storeService.DeleteStore(c.Request.Context(), storeID)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}
