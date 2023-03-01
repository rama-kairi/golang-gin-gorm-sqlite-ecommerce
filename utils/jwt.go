package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rama-kairi/blog-api-golang-gin/schema"
	"github.com/spf13/viper"
)

var secret = []byte(viper.GetString("JWT_SECRET"))

var expireMap = map[schema.TokenType]time.Time{
	schema.TokenTypeAccess:  time.Now().Add(15 * time.Minute),
	schema.TokenTypeRefresh: time.Now().Add(24 * time.Hour),
	schema.TokenTypeReset:   time.Now().Add(1 * time.Hour),
	schema.TokenTypeVerify:  time.Now().Add(1 * time.Hour),
}

type TokenClaims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

// GenerateJWTToken - generates a JWT token
func GenerateJWTToken(
	uuid string, tokenType schema.TokenType,
) (schema.SingleTokenResponse, error) {
	// Create the Claims
	claims := &TokenClaims{
		Type: string(tokenType),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "blog-api-golang-gin",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireMap[tokenType]),
			Subject:   uuid,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return schema.SingleTokenResponse{}, err
	}

	return schema.SingleTokenResponse{
		Token:     tokenString,
		ExpiresAt: expireMap[tokenType],
		Type:      tokenType,
	}, nil
}

// VerifyJWTToken - verifies a JWT token
func VerifyJWTToken(tokenString string, tokenType schema.TokenType) (tokenStr string, err error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		},
	)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if CheckTokenType(claims["type"].(string), tokenType) != nil {
		return "", fmt.Errorf("invalid token type")
	}

	return claims["sub"].(string), nil
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
func CheckTokenType(tokenTypeStr string, tokenType schema.TokenType) error {
	if tokenTypeStr != string(tokenType) {
		return fmt.Errorf("invalid token type")
	}
	return nil
}
