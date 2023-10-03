package erpcontroller

import (
	"erp/api"
	"erp/api_errors"
	"erp/constants"
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
		h.ResponseValidationError(c, err)
		return
	}

	req.StoreId = utils.GetStoreIDFromContext(c)

	if err := h.validateOrderItem(req.OrderItems); err != nil {
		h.ResponseError(c, err)
		return
	}

	if err := validateType(req.DiscountType); err != nil {
		h.ResponseError(c, err)
		return
	}

	if err := validateType(req.PromoteType); err != nil {
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

func validateType[E string](t E) (err error) {
	if t != "" {
		if t != constants.TypePercent && t != constants.TypeAmount {
			return errors.New(api_errors.ErrTypeInvalid)
		}
	}
	return nil
}
