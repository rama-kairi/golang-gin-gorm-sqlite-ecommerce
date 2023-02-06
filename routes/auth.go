package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
)

func authRoutes(e *gin.Engine) {
	authApi := controllers.NewAuthController()

	auth := e.Group("/auth")
	{
		auth.POST("/signup", authApi.Signup)
		auth.POST("/login", authApi.Login)
		auth.PATCH("/verify/:email", authApi.Verify)
		auth.PATCH("/forgot-password/:email", authApi.ForgotPassword)
		auth.PATCH("/reset-password/:token", authApi.ResetPassword)
	}
}
