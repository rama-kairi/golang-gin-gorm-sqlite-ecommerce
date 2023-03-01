package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
)

func subCategoryRoutes(e *gin.Engine, client *ent.Client) {
	subCategoryApi := controllers.NewSubCategoryController(client)

	subCategoryGroup := e.Group("/sub-category")
	{
		subCategoryGroup.GET("", subCategoryApi.GetAll)
		subCategoryGroup.POST("", subCategoryApi.Create)
		subCategoryGroup.GET("/:id", subCategoryApi.Get)
		subCategoryGroup.DELETE("/:id", subCategoryApi.Delete)
		subCategoryGroup.PATCH("/:id", subCategoryApi.Update)
	}
}
