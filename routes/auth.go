package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
)

func authRoutes(e *gin.Engine, client *ent.Client) {
	authApi := controllers.NewAuthController(client)
	auth := e.Group("/auth")
	{
		auth.POST("/signup", authApi.Signup)
		auth.POST("/login", authApi.Login)
		auth.PATCH("/verify/:token", authApi.Verify)
		auth.PATCH("/forgot-password/:email", authApi.ForgotPassword)
		auth.PATCH("/reset-password/:token", authApi.ResetPassword)
		auth.PATCH("/change-password", authApi.ChangePassword)
		auth.PATCH("/refresh", authApi.RefreshToken)
	}
}
