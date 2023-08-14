package handler

import (
	"erp-server/model"
	"erp-server/service"
	"erp-server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	service service.IProduct
}

func NewProduct(service service.IProduct) *Product {
	return &Product{service: service}
}

func (h *Product) CreateProduct(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.ProductRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.CreateProduct(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Product) UpdateProduct(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.ProductRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.UpdateProduct(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Product) GetProducts(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	business, err := h.service.GetProducts(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, business)
}

func (h *Product) GetProduct(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	business, err := h.service.GetProduct(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, business)
}
