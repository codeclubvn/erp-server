package utils

import (
	"net/http"
	"strings"
)

func CurrentUser(c *http.Request) (string, error) {
	userIdStr := c.Header.Get("x-user-id")
	if strings.Contains(userIdStr, "|") {
		userIdStr = strings.Split(userIdStr, "|")[0]
	}

	return userIdStr, nil
}
