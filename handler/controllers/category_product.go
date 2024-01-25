package controller

import (
	"erp/handler/dto"
	"erp/handler/dto/erp"
	"erp/service"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryProductController struct {
	categoryProductService service.CategoryProductService
	dto.BaseController
}

func NewCategoryProductController(categoryProductService service.CategoryProductService) *CategoryProductController {
	return &CategoryProductController{
		categoryProductService: categoryProductService,
	}
}

func (b *CategoryProductController) Create(c *gin.Context) {
	var req erpdto.CategoryProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	res, err := b.categoryProductService.Create(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *CategoryProductController) Update(c *gin.Context) {
	var req erpdto.CategoryProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, err := b.categoryProductService.Update(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *CategoryProductController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	if err := b.categoryProductService.Delete(c, id); err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *CategoryProductController) GetList(c *gin.Context) {
	var req erpdto.GetListCatProRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, total, err := b.categoryProductService.GetList(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.ResponseList(c, "success", total, res)
}
