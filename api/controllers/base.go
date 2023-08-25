package controller

import (
	"erp/api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (b *BaseController) Response(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, response.SimpleResponse{
		Data:    data,
		Message: message,
	})
}

func (b *BaseController) ResponseList(c *gin.Context, message string, total *int64, data interface{}) {
	c.JSON(http.StatusOK, response.SimpleResponseList{
		Message: message,
		Data:    data,
		Total:   total,
	})
}

func (b *BaseController) ResponseError(c *gin.Context, statusCode int, errs []error) {

	errorStrings := make([]error, len(errs))
	for i, err := range errs {
		c.Error(err)
		errorStrings[i] = err
	}

	c.AbortWithStatusJSON(statusCode, response.ResponseError{
		Message: errs[0].Error(),
		Errors:  errorStrings,
	})
}
