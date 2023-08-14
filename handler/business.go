package handler

import (
	"erp-server/model"
	"erp-server/service"
	"erp-server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Business struct {
	service service.IBusiness
}

func NewBusiness(service service.IBusiness) *Business {
	return &Business{service: service}
}

func (h *Business) CreateBusiness(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.BusinessRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.CreateBusiness(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Business) UpdateBusiness(ctx *gin.Context) {

	// get x-user-id
	userId, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	businessReq := model.BusinessRequest{}
	if err := ctx.ShouldBindJSON(&businessReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	businessReq.UserId = &userId

	// Check valid request

	// Create business
	business, err := h.service.UpdateBusiness(ctx, businessReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (h *Business) GetBusiness(ctx *gin.Context) {

	// check x-user-id
	userId := ctx.GetHeader("x-user-id")

	// Create business
	business, err := h.service.GetBusiness(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, business)
}
