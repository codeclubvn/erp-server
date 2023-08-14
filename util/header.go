package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func GetXUserId(ctx *gin.Context) (uuid.UUID, error) {
	userIdStr := ctx.GetHeader("x-user-id")
	if strings.Contains(userIdStr, "|") {
		userIdStr = strings.Split(userIdStr, "|")[0]
	}
	res, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, err
	}
	return res, nil
}
