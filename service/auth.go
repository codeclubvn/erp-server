package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("super_secret_key")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type IAuth interface {
	GenerateJWT(email string, username string) (tokenString string, err error)
	ValidateToken(signedToken string) (err error)
}

// GenerateJWT generates a JWT token and assign a username to it's claims and return it
func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	// Chuyển chuỗi token thành token object
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	// Kiểm tra xem token có hợp lệ không
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.New("token is expired")
		return
	}
	return
}
