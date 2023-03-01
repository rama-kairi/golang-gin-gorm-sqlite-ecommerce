package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
	"github.com/rama-kairi/blog-api-golang-gin/ent/user"
	"github.com/rama-kairi/blog-api-golang-gin/schema"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the header
		token, err := utils.ParseToken(c)
		if err != nil {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token 1")
			c.Abort()
			return
		}

		// get the email from the token
		uuidStr, err := utils.VerifyJWTToken(token, schema.TokenTypeAccess)
		if err != nil {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token 2")
			c.Abort()
			return
		}

		// Parse uuidStr to uuid.UUID
		log.Println(uuidStr)
		uuid, err := uuid.Parse(uuidStr)
		log.Println(err)
		if err != nil {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token 3")
			c.Abort()
			return
		}

		// Get the user from the database
		userRes, err := client.User.
			Query().
			Where(user.ID(uuid)).Only(c)
		if err != nil {
			utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token 4")
			c.Abort()
			return
		}

		// Set the user in the context
		c.Set("user", userRes)

		c.Next()
	}
}
