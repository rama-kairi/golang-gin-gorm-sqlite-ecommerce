package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/db"
)

func productRoutes(e *gin.Engine) {
	entClient := db.InitEntDb()
	productApi := controllers.NewProductController(entClient)

	productGroup := e.Group("/product")
	{
		productGroup.GET("", productApi.GetAll)
		productGroup.POST("", productApi.Create)
		productGroup.GET("/:id", productApi.Get)
		productGroup.DELETE("/:id", productApi.Delete)
		productGroup.PATCH("/:id", productApi.Update)
		productGroup.GET("/user/:id", productApi.GetAllByUser)
	}
}
