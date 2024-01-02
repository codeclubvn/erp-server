package api

import (
	"erp/api/request"
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
	var o request.PageOptions
	if err := c.ShouldBindQuery(&o); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	if o.Limit == 0 {
		o.Limit = 10
	}

	if o.Page == 0 {
		o.Page = 1
	}

	pageCount := utils.GetPageCount(*total, o.Limit)
	c.JSON(http.StatusOK, response.SimpleResponseList{
		Message: message,
		Data:    data,
		Meta: response.Meta{
			Total:       total,
			Page:        o.Page,
			Limit:       o.Limit,
			Sort:        o.Sort,
			PageCount:   pageCount,
			HasPrevPage: o.Page > 1,
			HasNextPage: o.Page < pageCount,
		},
	})
}

func (b *BaseController) ResponseError(c *gin.Context, err error) {
	mas, ok := api_errors.MapErrorCodeMessage[err.Error()]
	var status int
	ginType := gin.ErrorTypePublic
	errp := err
	if !ok {
		status = http.StatusInternalServerError
		ginType = gin.ErrorTypePrivate
		mas = api_errors.MapErrorCodeMessage[api_errors.ErrInternalServerError]
		errp = errors.New(api_errors.ErrInternalServerError)
	}

	c.Errors = append(c.Errors, &gin.Error{
		Err:  err,
		Type: ginType,
		Meta: status,
	})

	c.AbortWithStatusJSON(mas.Status, response.ResponseError{
		Code:    errp.Error(),
		Message: mas.Message,
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
