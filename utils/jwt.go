package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
	TokenTypeReset   TokenType = "reset"
	TokenTypeVerify  TokenType = "verification"
)

var secret = []byte(viper.GetString("JWT_SECRET"))

type tokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	Type      TokenType `json:"type"`
}

var expireMap = map[TokenType]time.Time{
	TokenTypeAccess:  time.Now().Add(15 * time.Minute),
	TokenTypeRefresh: time.Now().Add(24 * time.Hour),
	TokenTypeReset:   time.Now().Add(1 * time.Hour),
	TokenTypeVerify:  time.Now().Add(1 * time.Hour),
}

type TokenClaims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

// GenerateJWTToken - generates a JWT token
func GenerateJWTToken(email string, tokenType TokenType) (tokenResponse, error) {
	// Create the Claims
	claims := &TokenClaims{
		Type: string(tokenType),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "blog-api-golang-gin",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireMap[tokenType]),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return tokenResponse{}, err
	}

	return tokenResponse{
		Token:     tokenString,
		ExpiresAt: expireMap[tokenType],
		Type:      tokenType,
	}, nil
}

// VerifyJWTToken - verifies a JWT token
func VerifyJWTToken(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		},
	)
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	return claims["sub"].(string), claims["type"].(string), nil
}

// ParseToken - parses a JWT token from Gin context
func ParseToken(c *gin.Context) (string, error) {
	tokenArr := strings.Split(c.GetHeader("Authorization"), " ")
	// Check if the length of the tokenArr is 2
	if len(tokenArr) != 2 && tokenArr[0] != "Bearer" {
		return "", fmt.Errorf("invalid token")
	}

	return tokenArr[1], nil
}

// CheckTokenType - checks if the token type is valid
func CheckTokenType(tokenTypeStr string, tokenType TokenType) error {
	if tokenTypeStr != string(tokenType) {
		return fmt.Errorf("invalid token type")
	}
	return nil
}
