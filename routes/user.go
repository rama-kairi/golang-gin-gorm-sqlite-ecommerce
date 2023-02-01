package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
)

func userRoutes(e *gin.Engine) {
	userApi := controllers.NewUserController()

	e.GET("/user", userApi.GetAll)
	e.POST("/user", userApi.Create)
	e.GET("/user/:id", userApi.Get)
	e.DELETE("/user/:id", userApi.Delete)
	e.PATCH("/user/:id", userApi.Update)
}
