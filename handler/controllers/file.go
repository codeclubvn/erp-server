package controller

import (
	"erp/handler/dto"
	erpservice "erp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct {
	dto.BaseController
	uploadService erpservice.FileService
}

func NewFileController(uploadService erpservice.FileService) *FileController {
	return &FileController{
		uploadService: uploadService,
	}
}

func (p *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file_request")
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	output, err := p.uploadService.UploadImage(c.Request.Context(), file)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", output)
}
