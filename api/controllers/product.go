package controller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	"erp/service"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ERPProductController struct {
	productService service.IProductService
	api.BaseController
}

func NewERPProductController(productService service.IProductService) *ERPProductController {
	return &ERPProductController{
		productService: productService,
	}
}

func (b *ERPProductController) Create(c *gin.Context) {
	var req erpdto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	//req.StoreId = utils.GetStoreIDFromContext(c.Request.Context())

	res, err := b.productService.Create(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *ERPProductController) Update(c *gin.Context) {
	var req erpdto.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, err := b.productService.Update(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *ERPProductController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	if err := b.productService.Delete(c, id); err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "delete success", nil)
}

func (b *ERPProductController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := b.productService.GetOne(c, id)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *ERPProductController) GetList(c *gin.Context) {
	var req erpdto.GetListProductRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}
	res, total, err := b.productService.GetList(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.ResponseList(c, "success", &total, res)
}
