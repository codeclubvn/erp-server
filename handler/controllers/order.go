package controller

import (
	"erp/handler/dto"
	"erp/handler/dto/erp"
	erpservice "erp/service"
	"erp/utils"
	"erp/utils/api_errors"
	"errors"

	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	dto.BaseController
	orderService erpservice.OrderService
}

func NewOrderController(orderService erpservice.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (h *OrderController) Create(c *gin.Context) {
	req, err := utils.GetRequest(c, erpdto.CreateOrderRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//req.StoreId = utils.GetStoreIDFromContext(c.Request.Context())

	if req.DiscountType != nil {
		if ok := req.DiscountType.CheckValid(); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": req.DiscountType.ErrorMessage(),
			})
			return
		}
	}

	if ok := req.Status.CheckValid(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": req.Status.ErrorCreateMessage(),
		})
		return
	}

	if err := h.validateOrderItem(req.OrderItems); err != nil {
		h.ResponseError(c, err)
		return
	}

	order, err := h.orderService.CreateFlow(c.Request.Context(), req)
	if err != nil {
		h.ResponseError(c, err)
		return
	}

	h.Response(c, http.StatusCreated, "Success", order)
}

func (h *OrderController) Update(c *gin.Context) {
	req, err := utils.GetRequest(c, erpdto.UpdateOrderRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//req.StoreId = utils.GetStoreIDFromContext(c.Request.Context())

	if ok := req.Status.CheckValid(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": req.Status.ErrorUpdateMessage(),
		})
		return
	}

	order, err := h.orderService.UpdateFlow(c.Request.Context(), req)
	if err != nil {
		h.ResponseError(c, err)
		return
	}

	h.Response(c, http.StatusOK, "Success", order)
}

func (h *OrderController) GetList(c *gin.Context) {
	var req erpdto.GetListOrderRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		h.ResponseValidationError(c, err)
		return
	}

	res, total, err := h.orderService.GetList(c.Request.Context(), req)
	if err != nil {
		h.ResponseError(c, err)
		return
	}

	h.ResponseList(c, "success", &total, res)
}

func (h *OrderController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := h.orderService.GetOne(c, id)
	if err != nil {
		h.ResponseError(c, err)
		return
	}
	h.Response(c, http.StatusOK, "success", res)
}

func (h *OrderController) GetOverview(c *gin.Context) {
	var req erpdto.GetListOrderRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		h.ResponseValidationError(c, err)
		return
	}
	res, err := h.orderService.GetOverview(c, req)
	if err != nil {
		h.ResponseError(c, err)
		return
	}
	h.Response(c, http.StatusOK, "success", res)
}

func (h *OrderController) GetBestSeller(c *gin.Context) {
	var req erpdto.GetListOrderRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		h.ResponseValidationError(c, err)
		return
	}
	res, err := h.orderService.GetBestSeller(c, req)
	if err != nil {
		h.ResponseError(c, err)
		return
	}
	h.Response(c, http.StatusOK, "success", res)
}

func (h *OrderController) validateOrderItem(req []erpdto.OrderItemRequest) (err error) {
	if len(req) == 0 {
		return errors.New(api_errors.ErrOrderItemRequired)
	}
	return nil
}
