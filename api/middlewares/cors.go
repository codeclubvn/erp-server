package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (e *GinMiddleware) CORS(c *gin.Context) {
	cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}
