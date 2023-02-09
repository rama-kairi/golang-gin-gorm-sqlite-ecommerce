package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/db"
)

func userRoutes(e *gin.Engine) {
	entClient := db.InitEntDb()
	userApi := controllers.NewUserController(entClient)

	userGroup := e.Group("/user")
	{
		userGroup.GET("", userApi.GetAll)
		userGroup.POST("", userApi.Create)
		userGroup.GET("/:id", userApi.Get)
		userGroup.DELETE("/:id", userApi.Delete)
		userGroup.PATCH("/:id", userApi.Update)
	}
}
