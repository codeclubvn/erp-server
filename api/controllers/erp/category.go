package erpcontroller

import (
	"erp/api"
	"erp/dto"
	service "erp/service"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ERPCategoryController struct {
	categoryService service.ERPCategoryService
	api.BaseController
}

func NewERPCategoryController(categoryService service.ERPCategoryService) *ERPCategoryController {
	return &ERPCategoryController{
		categoryService: categoryService,
	}
}

func (b *ERPCategoryController) Create(c *gin.Context) {
	userId, err := utils.CurrentUser(c.Request)
	if err != nil {
		b.ResponseError(c, err)
		return
	}

	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	req.UserId = userId

	res, err := b.categoryService.Create(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *ERPCategoryController) Update(c *gin.Context) {
	var req dto.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, err := b.categoryService.Update(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *ERPCategoryController) Delete(c *gin.Context) {

	id := c.Query("id")

	// todo: doi thanh function cua mnnguyen
	userId, err := utils.CurrentUser(c.Request)
	if err != nil {
		b.ResponseError(c, err)
		return
	}

	if err := b.categoryService.Delete(c, id, userId); err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "delete success", nil)
}

func (b *ERPCategoryController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := b.categoryService.GetOne(c, id)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *ERPCategoryController) GetList(c *gin.Context) {
	var req dto.GetListCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, err := b.categoryService.GetList(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}
