package middlewares

import (
	"erp/api/response"
	"erp/api_errors"
	config "erp/config"
	constants "erp/constants"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (e *GinMiddleware) JWT(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Server.Env != constants.Dev && constants.Local != config.Server.Env {
			auth := c.Request.Header.Get("Authorization")

			resError := func() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
					Message: "Unauthorized",
					Errors:  []error{api_errors.ErrUnauthorizedAccess},
				})
			}

			if auth == "" {
				resError()
				return
			}
			jwtToken := strings.Split(auth, " ")[1]
			if jwtToken == "" {
				resError()
				return
			}

			token, err := parseToken(jwtToken, config.Jwt.Secret)
			if err != nil {
				resError()
				return
			}

			claims, OK := token.Claims.(jwt.MapClaims)
			if !OK {
				resError()
				return
			}

			_, OK = claims["sub"].(string)
			if !OK {
				resError()
				return
			}

			// user, err := userRepo.GetByID(c, claimedUID)
			// if err != nil {
			// 	resError()
			// 	return
			// }

			// c.Set("user", user)
		}
		c.Next()
	}
}

func parseToken(jwtToken string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, api_errors.ErrTokenBadSignedMethod
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, api_errors.ErrUnauthorizedAccess
	}

	return token, nil
}
