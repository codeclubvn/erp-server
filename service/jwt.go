package service

import (
	"erp/api_errors"
	config "erp/config"
	"erp/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService interface {
	GenerateToken(userID string, tokenType constants.TokenType, expiresIn int64) (string, error)
	ValidateToken(token string, tokenType constants.TokenType) (*string, error)
	GenerateAuthTokens(userID string) (string, string, error)
}

type jwtService struct {
	config *config.Config
}

func NewJwtService(config *config.Config) JwtService {
	return &jwtService{
		config: config,
	}
}

func (j *jwtService) GenerateToken(userID string, tokenType constants.TokenType, expiresIn int64) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(time.Duration(expiresIn)).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "erp",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtService) GenerateAuthTokens(userID string) (string, string, error) {
	accessToken, err := j.GenerateToken(userID, constants.AccessToken, j.config.Jwt.AccessTokenExpiresIn)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := j.GenerateToken(userID, constants.RefreshToken, j.config.Jwt.RefreshTokenExpiresIn)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (j *jwtService) ValidateToken(token string, tokenType constants.TokenType) (*string, error) {
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, api_errors.ErrTokenExpired
	}

	if claims.Issuer != "erp" {
		return nil, api_errors.ErrTokenInvalid
	}

	if claims.Subject == "" {
		return nil, api_errors.ErrTokenInvalid
	}

	return &claims.Subject, nil
}
