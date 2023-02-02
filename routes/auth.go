package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
)

func authRoutes(e *gin.Engine) {
	authApi := controllers.NewAuthController()

	e.POST("/auth/signup", authApi.Signup)
	e.POST("/auth/login", authApi.Login)
	e.PATCH("/auth/verify/:email", authApi.Verify)
}
