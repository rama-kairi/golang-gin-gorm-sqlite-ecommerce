package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

var secret = "5eb65440d4d739befe9a0c832c4b39aa10d74748aad00789002e224d45d980b3"

// type tokenResponse struct {
// 	Token     string    `json:"token"`
// 	ExpiresAt time.Time `json:"expires_at"`
// 	Type      TokenType `json:"type"`
// }

// var tokenDuration map[TokenType]time.Time = map[TokenType]time.Time{
// 	TokenTypeAccess:  time.Now().Add(15 * time.Minute),
// 	TokenTypeRefresh: time.Now().Add(24 * time.Hour),
// }

// GenerateJWTToken - generates a JWT token
func GenerateJWTToken(email string, tokenType TokenType) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// tokenString, err := token.SignedString("secret")
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

// VerifyJWTToken - verifies a JWT token
func VerifyJWTToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.RegisteredClaims)
	if !ok {
		return nil, err
	}

	return &claims, nil
}

// ParseToken - parses a JWT token from Gin context
func ParseToken(c *gin.Context) string {
	token := strings.Split(c.GetHeader("Authorization"), " ")
	if len(token) != 2 {
		Response(c, http.StatusUnauthorized, nil, "Unauthorized, Please provide the Bearer token")
		c.Abort()
		return ""
	}
	return token[1]
}
