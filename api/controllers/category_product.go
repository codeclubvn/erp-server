package controller

import (
	"erp/api"
	"erp/dto"
	service "erp/service"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type CategoryProductController struct {
	categoryProductService service.CategoryProductService
	logger                 *zap.Logger
	api.BaseController
}

func NewCategoryProductController(c *gin.RouterGroup, categoryProductService service.CategoryProductService, logger *zap.Logger) *CategoryProductController {
	controller := &CategoryProductController{
		categoryProductService: categoryProductService,
		logger:                 logger,
	}
	g := c.Group("/categories_products")

	g.POST("", controller.Create)
	g.PUT("", controller.Update)
	g.DELETE("", controller.Delete)
	g.GET("", controller.GetList)

	return controller
}

func (b *CategoryProductController) Create(c *gin.Context) {
	userId, err := utils.CurrentUser(c.Request)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}

	var req dto.CategoryProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	req.UserId = userId

	res, err := b.categoryProductService.Create(c, req)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *CategoryProductController) Update(c *gin.Context) {
	var req dto.CategoryProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	res, err := b.categoryProductService.Update(c, req)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *CategoryProductController) Delete(c *gin.Context) {
	var req dto.DeleteCatagoryProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	if err := b.categoryProductService.Delete(c, req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "delete success", nil)
}

func (b *CategoryProductController) GetList(c *gin.Context) {
	var req dto.CategoryProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	res, err := b.categoryProductService.GetList(c, req)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}
