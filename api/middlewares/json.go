package middlewares

import "github.com/gin-gonic/gin"

func (e *GinMiddleware) JSONMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()
}
