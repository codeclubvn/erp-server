package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (e *GinMiddleware) ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		if len(c.Errors) > 0 {
			logger.Error("Error", zap.String("Error", c.Errors.String()))
			//var message string
			//statusCode := http.StatusInternalServerError
			//err := c.Errors.Last()
			//
			//if err.IsType(gin.ErrorTypePrivate) {
			//	message = utils.ErrInternalServerError.Error()
			//} else {
			//	message = err.Error()
			//	statusCode = err.Meta.(int)
			//}
			//c.JSON(statusCode, dto.ResponseError{
			//	Message: message,
			//	Errors:  c.Errors.Errors(),
			//})
			//return
		}
	}
}
