package utils

import "encoding/base64"

// GenerateBasicAuthToken - This function generates a basic auth token
func GenerateBasicAuthToken(email string) string {
	// BasicAuth
	token := base64.StdEncoding.EncodeToString([]byte(email))
	return token
}

// DecodeBasicAuthToken - This function decodes a basic auth token
func DecodeBasicAuthToken(token string) string {
	// BasicAuth
	decoded, _ := base64.StdEncoding.DecodeString(token)
	return string(decoded)
}
