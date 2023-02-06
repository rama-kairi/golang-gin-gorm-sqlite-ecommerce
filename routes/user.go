package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/middleware"
)

func userRoutes(e *gin.Engine) {
	userApi := controllers.NewUserController()
	userGroup := e.Group("/user", middleware.AuthMiddleware())
	{
		userGroup.GET("", userApi.GetAll)
		userGroup.POST("", userApi.Create)
		userGroup.GET("/:id", userApi.Get)
		userGroup.DELETE("/:id", userApi.Delete)
		userGroup.PATCH("/:id", userApi.Update)
	}
}
