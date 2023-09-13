package api

import (
	"erp/api/response"
	"erp/api_errors"
	"erp/utils"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		Meta: response.Meta{
			Total: total,
		},
	})
}

func (b *BaseController) ResponseError(c *gin.Context, err error) {

	message, ok := api_errors.MapErrorCodeMessage[err.Error()]
	var status int
	ginType := gin.ErrorTypePublic
	errp := err
	if !ok {
		status = http.StatusInternalServerError
		ginType = gin.ErrorTypePrivate
		message = api_errors.MapErrorCodeMessage[api_errors.ErrInternalServerError]
		errp = errors.New(api_errors.ErrInternalServerError)
	}

	c.Errors = append(c.Errors, &gin.Error{
		Err:  err,
		Type: ginType,
		Meta: status,
	})

	c.AbortWithStatusJSON(status, response.ResponseError{
		Code:    errp.Error(),
		Message: message,
	})
}

func (b *BaseController) ResponseValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		err = errors.New(utils.StructPascalToSnakeCase(ve[0].Field()) + " is " + ve[0].Tag())
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseError{
		Code:    api_errors.ErrValidation,
		Message: err.Error(),
	})
}
