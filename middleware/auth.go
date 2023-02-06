package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the header
		token, err := utils.ParseToken(c)
		if err != nil {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
			c.Abort()
			return
		}

		// get the email from the token
		email, tokenType, err := utils.VerifyJWTToken(token)
		if err != nil {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
			c.Abort()
			return
		}

		// Check if the token is an access token
		if tokenType != "access" {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
			c.Abort()
			return
		}

		// Get the user from the database
		var user models.User
		if err := db.Db.Where("email = ?", email).First(&user).Error; err != nil {
			// If the user is not found, return 404
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
			c.Abort()
			return
		}

		c.Next()
	}
}
