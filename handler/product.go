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

	productReq := model.ProductRequest{}
	if err := ctx.ShouldBindJSON(&productReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	productReq.UserId = &userId

	// Check valid request

	// Create product
	product, err := h.service.CreateProduct(ctx, productReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, product)
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

	productReq := model.ProductRequest{}
	if err := ctx.ShouldBindJSON(&productReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	productReq.UserId = &userId

	// Check valid request

	// Create product
	product, err := h.service.UpdateProduct(ctx, productReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (h *Product) GetProducts(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create product
	product, err := h.service.GetProducts(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *Product) GetProduct(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// get id
	id := ctx.Param("id")

	oneProductReq := model.OneProductRequest{
		UserId: userId,
		Id:     id,
	}

	// Create product
	product, err := h.service.GetProduct(ctx, oneProductReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, product)
}
