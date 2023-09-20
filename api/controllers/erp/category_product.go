package erpcontroller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	service "erp/service/erp"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ERPCategoryProductController struct {
	api.BaseController
	categoryProductService service.ERPCategoryProductService
}

func NewERPCategoryProductController(categoryProductService service.ERPCategoryProductService) *ERPCategoryProductController {
	return &ERPCategoryProductController{
		categoryProductService: categoryProductService,
	}
}

func (b *ERPCategoryProductController) Create(c *gin.Context) {
	userId, err := utils.CurrentUser(c.Request)
	if err != nil {
		b.ResponseError(c, err)
		return
	}

	var req erpdto.CategoryProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	req.UserId = userId

	res, err := b.categoryProductService.Create(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *ERPCategoryProductController) Update(c *gin.Context) {
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
	b.Response(c, http.StatusOK, "", res)
}

func (b *ERPCategoryProductController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	userId := utils.GetUserStringIDFromContext(c)

	if err := b.categoryProductService.Delete(c, id, userId); err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "delete success", nil)
}

func (b *ERPCategoryProductController) GetList(c *gin.Context) {
	var req erpdto.GetListCatProRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, err := b.categoryProductService.GetList(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}
