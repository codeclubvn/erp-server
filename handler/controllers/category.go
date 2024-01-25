package controller

import (
	"erp/handler/dto"
	"erp/handler/dto/erp"
	"erp/service"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ERPCategoryController struct {
	categoryService service.ERPCategoryService
	dto.BaseController
}

func NewERPCategoryController(categoryService service.ERPCategoryService) *ERPCategoryController {
	return &ERPCategoryController{
		categoryService: categoryService,
	}
}

func (b *ERPCategoryController) Create(c *gin.Context) {
	var req erpdto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	res, err := b.categoryService.Create(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *ERPCategoryController) Update(c *gin.Context) {
	var req erpdto.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, err := b.categoryService.Update(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *ERPCategoryController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	if err := b.categoryService.Delete(c, id); err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *ERPCategoryController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)

	res, err := b.categoryService.GetOne(c, id)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *ERPCategoryController) GetList(c *gin.Context) {
	var req erpdto.GetListCategoryRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, total, err := b.categoryService.GetList(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.ResponseList(c, "success", &total, res)
}
