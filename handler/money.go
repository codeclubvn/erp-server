package handler

import (
	"erp-server/model"
	"erp-server/service"
	"erp-server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Money struct {
	service service.IMoney
}

func NewMoney(service service.IMoney) *Money {
	return &Money{service: service}
}

func (h *Money) CreateMoney(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.MoneyRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.CreateMoney(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Money) UpdateMoney(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.MoneyRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.UpdateMoney(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Money) GetMoneys(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	business, err := h.service.GetMoneys(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, business)
}

func (h *Money) GetMoney(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	business, err := h.service.GetMoney(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, business)
}

func (h *Money) DeleteMoney(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	if err := h.service.DeleteMoney(ctx, userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
