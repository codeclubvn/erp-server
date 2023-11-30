package controller

import (
	"erp/api"
	erpdto "erp/dto/erp"
	"erp/service"
	"erp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PromoteController struct {
	promoteService service.IPromoteService
	api.BaseController
}

func NewPromoteController(promoteService service.IPromoteService) *PromoteController {
	return &PromoteController{
		promoteService: promoteService,
	}
}

func (b *PromoteController) Create(c *gin.Context) {
	var req erpdto.CreatePromoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if ok := req.PromoteType.CheckValid(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": req.PromoteType.ErrorMessage(),
		})
		return
	}

	req.StoreId = utils.GetStoreIDFromContext(c.Request.Context())

	res, err := b.promoteService.CreateFlow(c, req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}
