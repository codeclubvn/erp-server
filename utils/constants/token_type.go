package constants

type TokenType string

const (
	AccessToken        TokenType = "access_token"
	RefreshToken       TokenType = "refresh_token"
	VerifyToken        TokenType = "verify_token"
	ResetPasswordToken TokenType = "reset_password_token"
)
