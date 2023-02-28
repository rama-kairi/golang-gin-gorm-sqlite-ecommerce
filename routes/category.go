package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/db"
)

func categoryRoutes(e *gin.Engine) {
	entClient := db.InitEntDb()
	categoryApi := controllers.NewCategoryController(entClient)

	categoryGroup := e.Group("/category")
	{
		categoryGroup.GET("", categoryApi.GetAll)
		categoryGroup.POST("", categoryApi.Create)
		categoryGroup.GET("/:id", categoryApi.Get)
		categoryGroup.DELETE("/:id", categoryApi.Delete)
		categoryGroup.PATCH("/:id", categoryApi.Update)
	}
}
