package handler

import (
	"erp-server/model"
	"erp-server/service"
	"erp-server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Order struct {
	service service.IOrder
}

func NewOrder(service service.IOrder) *Order {
	return &Order{service: service}
}

func (h *Order) CreateOrder(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.OrderRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.CreateOrder(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Order) UpdateOrder(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.OrderRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.UpdateOrder(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Order) GetOrders(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	business, err := h.service.GetOrders(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, business)
}

func (h *Order) GetOrder(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	business, err := h.service.GetOrder(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, business)
}
