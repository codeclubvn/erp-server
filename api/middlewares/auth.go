package middlewares

import (
	"context"
	"erp/api/response"
	"erp/api_errors"
	dto "erp/dto/auth"
	"erp/models"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (e *GinMiddleware) Auth(authorization bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: "Unauthorized",
				Code:    api_errors.ErrUnauthorizedAccess,
			})
			return
		}
		jwtToken := strings.Split(auth, " ")[1]

		if jwtToken == "" {
			c.Errors = append(c.Errors, &gin.Error{
				Err: errors.New(api_errors.ErrTokenMissing),
			})

			mas := api_errors.MapErrorCodeMessage[api_errors.ErrTokenMissing]
			c.AbortWithStatusJSON(mas.Status, response.ResponseError{
				Message: mas.Message,
				Code:    api_errors.ErrTokenMissing,
			})
			return
		}

		claims, err := parseToken(jwtToken, e.config.Jwt.Secret)
		if err != nil {
			c.Errors = append(c.Errors, &gin.Error{
				Err: errors.WithStack(err),
			})
			mas := api_errors.MapErrorCodeMessage[err.Error()]
			c.AbortWithStatusJSON(mas.Status, response.ResponseError{
				Message: mas.Message,
				Code:    api_errors.ErrTokenInvalid,
			})
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "x-user-id", claims.Subject))
		if !authorization {
			c.Next()
			return
		}

		storeID := c.Request.Header.Get("x-store-id")
		if storeID == "" {
			c.Errors = append(c.Errors, &gin.Error{
				Err: errors.New(api_errors.ErrMissingXStoreID),
			})

			mas := api_errors.MapErrorCodeMessage[api_errors.ErrMissingXStoreID]

			c.AbortWithStatusJSON(mas.Status, response.ResponseError{
				Message: mas.Message,
				Code:    api_errors.ErrMissingXStoreID,
			})
			return
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "x-store-id", storeID))

		ur := new(models.UserRole)
		if err = e.db.Model(models.UserRole{}).Where("user_id = ? AND store_id = ?", claims.Subject, storeID).First(ur).Error; err != nil {
			c.Errors = append(c.Errors, &gin.Error{
				Err: errors.Wrap(err, "cannot find user role"),
			})

			mas := api_errors.MapErrorCodeMessage[api_errors.ErrUnauthorizedAccess]

			c.AbortWithStatusJSON(mas.Status, response.ResponseError{
				Message: mas.Message,
				Code:    api_errors.ErrUnauthorizedAccess,
			})
			return
		}
		if ur.IsStoreOwner {
			c.Next()
			return
		}

		role := new(models.Role)
		if err = e.db.Model(models.Role{}).Where("id = ?", ur.RoleID).First(role).Error; err != nil {
			c.Errors = append(c.Errors, &gin.Error{
				Err: errors.Wrap(err, "cannot find role"),
			})

			mas := api_errors.MapErrorCodeMessage[api_errors.ErrUnauthorizedAccess]

			c.AbortWithStatusJSON(mas.Status, response.ResponseError{
				Message: mas.Message,
				Code:    api_errors.ErrUnauthorizedAccess,
			})
			return
		}
		up := new(models.Permission)
		if err = e.db.Model(models.Permission{}).Where("role_id = ? AND route_path = ?", ur.RoleID, c.Request.URL.Path).First(up).Error; err != nil {
			c.Errors = append(c.Errors, &gin.Error{
				Err: errors.Wrap(err, "cannot find permission"),
			})

			mas := api_errors.MapErrorCodeMessage[api_errors.ErrUnauthorizedAccess]

			c.AbortWithStatusJSON(mas.Status, response.ResponseError{
				Message: mas.Message,
				Code:    api_errors.ErrUnauthorizedAccess,
			})
			return
		}

		c.Next()
	}
}

func parseToken(jwtToken string, secret string) (*dto.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &dto.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if (err.(*jwt.ValidationError)).Errors == jwt.ValidationErrorExpired {
			return nil, errors.New(api_errors.ErrTokenExpired)
		}
		return nil, errors.Wrap(err, "cannot parse token")
	}

	if claims, OK := token.Claims.(*dto.JwtClaims); OK && token.Valid {
		return claims, nil
	}

	return nil, errors.New(api_errors.ErrTokenInvalid)
}
