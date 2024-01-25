package controller

import (
	"erp/handler/dto"
	erpdto "erp/handler/dto/finance"
	erpservice "erp/service"
	"erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	dto.BaseController
	WalletService erpservice.WalletService
}

func NewWalletController(walletService erpservice.WalletService) *WalletController {
	return &WalletController{
		WalletService: walletService,
	}
}

func (p *WalletController) Create(c *gin.Context) {
	var req erpdto.CreateWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	wallet, err := p.WalletService.Create(nil, c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", wallet.ID)
}

func (p *WalletController) Update(c *gin.Context) {
	var req erpdto.UpdateWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	_, err := p.WalletService.Update(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *WalletController) List(c *gin.Context) {
	var req erpdto.ListWalletRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	wallets, total, err := p.WalletService.GetList(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", &total, wallets)
}

func (p *WalletController) Delete(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	err := p.WalletService.Delete(c.Request.Context(), id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *WalletController) GetOne(c *gin.Context) {
	id := utils.ParseStringIDFromUri(c)
	res, err := p.WalletService.GetOne(c, id)
	if err != nil {
		p.ResponseError(c, err)
		return
	}
	p.Response(c, http.StatusOK, "success", res)
}
