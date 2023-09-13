package erproute

import (
	"erp/lib"

	"github.com/gin-gonic/gin"
)

type ERPHandler struct {
	*gin.RouterGroup
}

func NewERPHandler(handler *lib.Handler) *ERPHandler {
	return &ERPHandler{
		handler.Group("/erp"),
	}
}
