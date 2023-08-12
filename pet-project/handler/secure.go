package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ISecure interface {
	Ping(context *gin.Context)
}

// an endpoint that will hold some super-secret information, which will be a “pong”, obviously
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
