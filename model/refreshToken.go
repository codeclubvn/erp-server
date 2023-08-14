package model

import "time"

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	ID        int `json:"id"`
	UserID    int `json:"userID"`
	Token     int
	JwtID     int
	IsUsed    bool
	IsRevoked bool
	IssueAt   bool
	ExpireAt  time.Time
}
