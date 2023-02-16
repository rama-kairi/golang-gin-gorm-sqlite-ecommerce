package schema

import "time"

type SignupSchema struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginSchema struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResetPasswordSchema struct {
	Password string `json:"password" binding:"required"`
}

type ChangePasswordSchema struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
	TokenTypeReset   TokenType = "reset"
	TokenTypeVerify  TokenType = "verification"
)

type SingleTokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	Type      TokenType `json:"type"`
}
type TokenResponse struct {
	AccessToken  SingleTokenResponse `json:"access_token,omitempty"`
	RefreshToken SingleTokenResponse `json:"refresh_token,omitempty"`
}
