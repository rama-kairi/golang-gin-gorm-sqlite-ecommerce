package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization"), " ")
		if len(token) != 2 {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Please provide the Bearer token")
			c.Abort()
			return
		}
		tokenStr := token[1]

		// get the email from the token
		email := utils.DecodeBasicAuthToken(tokenStr)

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
