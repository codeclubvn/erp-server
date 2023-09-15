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

type CategoryController struct {
	categoryService service.CategoryService
	logger          *zap.Logger
	api.BaseController
}

func NewCategoryController(c *gin.RouterGroup, categoryService service.CategoryService, logger *zap.Logger) *CategoryController {
	controller := &CategoryController{
		categoryService: categoryService,
		logger:          logger,
	}
	g := c.Group("/categories")

	g.POST("", controller.Create)
	g.PUT("", controller.Update)
	g.DELETE("", controller.Delete)
	g.GET("/:id", controller.GetOne)
	g.GET("", controller.GetList)

	return controller
}

func (b *CategoryController) Create(c *gin.Context) {
	userId, err := utils.CurrentUser(c.Request)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}

	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	req.UserId = userId

	res, err := b.categoryService.Create(c, req)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *CategoryController) Update(c *gin.Context) {
	var req dto.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	res, err := b.categoryService.Update(c, req)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *CategoryController) Delete(c *gin.Context) {
	var req dto.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	if err := b.categoryService.Delete(c, req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "delete success", nil)
}

func (b *CategoryController) GetOne(c *gin.Context) {
	var req dto.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	res, err := b.categoryService.GetOne(c, req)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}

func (b *CategoryController) GetList(c *gin.Context) {
	var req dto.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	res, err := b.categoryService.GetList(c, req)
	if err != nil {
		// todo: change error to error_code
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "", res)
}
