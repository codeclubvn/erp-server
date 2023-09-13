package dto

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	jwt.StandardClaims
	// RoleID    string              `json:"role_id"`
	TokenType string `json:"token_type"`
}
