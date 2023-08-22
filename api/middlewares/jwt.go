package middlewares

import (
	config "erp/config"
	constants "erp/constants"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (e *GinMiddleware) JWT(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Server.Env != constants.Dev && constants.Local != config.Server.Env {
			auth := c.Request.Header.Get("Authorization")
			if auth == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				return
			}
			token := strings.Split(auth, " ")[1]
			if token == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				return
			}
		}
		c.Next()
	}
}
