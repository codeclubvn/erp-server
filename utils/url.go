package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UriParse struct {
	ID []string `json:"id" uri:"id"`
}

func ParseStringIDFromUri(c *gin.Context) string {
	tID := UriParse{}
	if err := c.ShouldBindUri(&tID); err != nil {
		_ = c.Error(err)
		return ""
	}
	if len(tID.ID) == 0 {
		_ = c.Error(fmt.Errorf("error: Empty when parse ID from URI"))
		return ""
	}
	return tID.ID[0]
}
