package erpcontroller

import (
	"erp/api"
	"erp/api_errors"
	erpdto "erp/dto/erp"
	erpservice "erp/service/erp"
	"erp/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	api.BaseController
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

	req.StoreId = utils.GetStoreIDFromContext(c.Request.Context())

	if req.DiscountType != "" {
		if ok := req.DiscountType.CheckValid(); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": req.DiscountType.ErrorMessage(),
			})
			return
		}
	}

	if ok := req.Status.CheckValid(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": req.Status.ErrorMessage(),
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

func (h *OrderController) validateOrderItem(req []erpdto.OrderItemRequest) (err error) {
	if len(req) == 0 {
		return errors.New(api_errors.ErrOrderItemRequired)
	}
	return nil
}
